package opennox

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/noxworld-dev/opennox-lib/log"

	"github.com/noxworld-dev/opennox/v1/common/config"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/internal/version"
)

var (
	apiLog    = log.New("api")
	apiTokens = make(map[string]struct{})
)

func init() {
	if tok := os.Getenv("NOX_API_TOKEN"); tok != "" {
		apiTokens[tok] = struct{}{}
	}
	const pref = "/api/v0/game"
	var token string
	configStrPtr(config.KeyServerAPIToken, "NOX_API_TOKEN", "", &token)
	registerOnConfigRead(func() {
		s := noxServer
		mux := s.HTTP()
		mux.HandleFunc(pref+"/info", handleServerInfo)
		mux.HandleFunc(pref+"/map", handleChangeMap)
		mux.HandleFunc(pref+"/cmd", handleRunCmd)
		s.registerScriptAPIs(pref)
		if token != "" {
			apiTokens[token] = struct{}{}
		}
	})
}

type gameInfoPlayers struct {
	Cur  int              `json:"cur"`
	Max  int              `json:"max"`
	List []gameInfoPlayer `json:"list"`
}

type gameInfoPlayer struct {
	Name  string `json:"name"`
	Class string `json:"class,omitempty"`
}

type gameInfoQuest struct {
	Stage int `json:"stage"`
}

type gameInfoResp struct {
	Name    string          `json:"name"`
	Map     string          `json:"map"`
	Mode    string          `json:"mode"`
	Vers    string          `json:"vers"`
	Players gameInfoPlayers `json:"players"`
	Quest   *gameInfoQuest  `json:"quest,omitempty"`
}

func getGameInfo(ctx context.Context) (*gameInfoResp, error) {
	ch := make(chan *gameInfoResp, 1)
	noxServer.QueueInLoop(ctx, func() {
		s := noxServer
		v := &gameInfoResp{
			Name: s.getServerName(),
			Map:  strings.ToLower(s.getServerMap()),
			Vers: version.Version(),
			Mode: noxflags.GetGame().ModeString(),
		}
		if noxflags.HasGame(noxflags.GameModeQuest) {
			v.Quest = &gameInfoQuest{Stage: s.nox_game_getQuestStage_4E3CC0()}
		}
		v.Players.Max = s.getServerMaxPlayers()
		for _, p := range s.Players.List() {
			v.Players.List = append(v.Players.List, gameInfoPlayer{
				Name:  p.Name(),
				Class: p.PlayerClass().String(),
			})
		}
		v.Players.Cur = len(v.Players.List)
		ch <- v
	})
	select {
	case v := <-ch:
		sort.Slice(v.Players.List, func(i, j int) bool {
			return v.Players.List[i].Name < v.Players.List[j].Name
		})
		return v, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func queueServerMapLoad(name string) {
	name = strings.TrimSpace(name)
	apiLog.Printf("load map: %q", name)
	noxServer.QueueInLoop(context.Background(), func() {
		serverCmdLoadMap(name)
	})
}

func queueServerCmd(cmd string) {
	cmd = strings.TrimSpace(cmd)
	apiLog.Printf("run command: %q", cmd)
	noxServer.QueueInLoop(context.Background(), func() {
		execServerCmd(cmd)
	})
}

func handleServerInfo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	case http.MethodGet, http.MethodHead:
	}
	resp, err := getGameInfo(r.Context())
	if err != nil {
		// TODO: set correct status code
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func handlePostStr(w http.ResponseWriter, r *http.Request, limit int, fnc func(str string)) {
	switch r.Method {
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	case http.MethodHead:
		return
	case http.MethodPost:
	}
	defer r.Body.Close()
	token := r.Header.Get("X-Token")
	if _, ok := apiTokens[token]; !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	data, err := io.ReadAll(io.LimitReader(r.Body, int64(limit)+1))
	if err != nil || len(data) > limit {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fnc(string(data))
}

func handleChangeMap(w http.ResponseWriter, r *http.Request) {
	handlePostStr(w, r, 256, queueServerMapLoad)
}

func handleRunCmd(w http.ResponseWriter, r *http.Request) {
	handlePostStr(w, r, 256, queueServerCmd)
}
