package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/noxworld-dev/opennox/v1/internal/vers/git"
)

const (
	BinOpenNox        = "opennox"
	BinOpenNoxHD      = "opennox-hd"
	BinOpenNoxDebug   = "opennox-debug"
	BinOpenNoxHDDebug = "opennox-hd-debug"
	BinServer         = "opennox-server"
	BinLauncher       = "opennox-launcher"
)

const (
	versPackage = "github.com/noxworld-dev/opennox/v1/internal/version"
)

var (
	defTargets = []string{
		BinOpenNox, BinOpenNoxHD, BinLauncher, BinServer,
	}
	fOut     = flag.String("o", "", "output directory")
	fSrc     = flag.String("s", "", "source directory")
	fOS      = flag.String("os", runtime.GOOS, "target OS to build for")
	fSafe    = flag.Bool("safe", false, "build a safe version (will run significantly slower)")
	fGo      = flag.String("go", "go", "go command to use")
	fVerbose = flag.Bool("v", false, "verbose mode")
)

func main() {
	flag.Parse()
	if err := build(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func build(args []string) error {
	if *fOS == "windows" {
		defTargets = append(defTargets, BinOpenNoxDebug, BinOpenNoxHDDebug)
	}
	targets := args
	if len(targets) == 0 {
		targets = defTargets
	}
	start := time.Now()
	// TODO: run in parallel
	for _, target := range targets {
		if err := buildTarget(target); err != nil {
			return err
		}
	}
	fmt.Printf("=== DONE in %v ===\n", time.Since(start))
	vers, sha := git.Version()
	fmt.Printf("version: %s (%s)\n", vers, sha)
	return nil
}

func buildTarget(target string) error {
	switch target {
	default:
		return fmt.Errorf("unsupported target: %q", target)
	case BinServer, "server":
		return goBuild("opennox", BinServer, &buildOpts{
			CGO: true, Tags: []string{"server"},
		})
	case BinOpenNox, "client":
		return goBuild("opennox", BinOpenNox, &buildOpts{
			CGO: true, GUI: true,
		})
	case BinOpenNoxHD, "client-hd", "hd":
		return goBuild("opennox", BinOpenNoxHD, &buildOpts{
			CGO: true, GUI: true, Tags: []string{"highres"},
		})
	case BinOpenNoxDebug, "client-debug":
		return goBuild("opennox", BinOpenNoxDebug, &buildOpts{
			CGO: true,
		})
	case BinOpenNoxHDDebug, "client-hd-debug":
		return goBuild("opennox", BinOpenNoxHDDebug, &buildOpts{
			CGO: true, Tags: []string{"highres"},
		})
	case BinLauncher, "launcher":
		return goBuild(BinLauncher, BinLauncher, &buildOpts{
			Native: true, CGO: true, GUI: true,
		})
	}
}

type buildOpts struct {
	Native bool
	CGO    bool
	GUI    bool
	Tags   []string
}

func goBuild(cmd string, bin string, opts *buildOpts) error {
	goos := *fOS
	switch goos {
	case "windows":
		if filepath.Ext(bin) == "" {
			bin += ".exe"
		}
	}
	bin = filepath.Join(*fOut, bin)
	isCross := goos != runtime.GOOS
	if opts == nil {
		opts = &buildOpts{}
	}
	if *fSafe {
		opts.Tags = append(opts.Tags, "safe")
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	vers, sha := git.Version()
	var (
		LDFLAGS = []string{
			"-X '" + versPackage + ".commit=" + sha + "'",
			"-X '" + versPackage + ".version=" + vers + "'",
		}
		GCFLAGS  = "-trimpath=" + wd
		ASMFLAGS = "-trimpath=" + wd
	)
	if opts.GUI {
		switch goos {
		case "windows":
			opts.Tags = append(opts.Tags, "guiapp")
			LDFLAGS = append(LDFLAGS, "-H windowsgui")
		}
	}
	var (
		envs []string
		args = []string{
			*fGo, "build", "-v",
			`-ldflags=` + strings.Join(LDFLAGS, " "),
			"-gcflags=" + GCFLAGS,
			"-asmflags=" + ASMFLAGS,
		}
	)
	if *fVerbose {
		args = append(args, "-work", "-x")
		work := filepath.Join(os.TempDir(), "go-opennox-build")
		if err := os.RemoveAll(work); err != nil {
			return err
		}
		if err := os.MkdirAll(work, 0755); err != nil {
			return err
		}
		envs = append(envs,
			"GOTMPDIR="+work,
		)
	}
	envs = append(envs,
		"GOOS="+goos,
	)
	if opts.CGO {
		envs = append(envs,
			"CGO_ENABLED=1",
		)
		if !opts.Native {
			envs = append(envs,
				"GOARCH=386",
				`CGO_CFLAGS_ALLOW=(-fshort-wchar)|(-fno-strict-aliasing)|(-fno-strict-overflow)`,
			)
		}
		if isCross {
			switch goos {
			case "windows":
				envs = append(envs,
					"CXX_FOR_TARGET=i686-w64-mingw32-g++",
					"CC_FOR_TARGET=i686-w64-mingw32-gcc",
					"CC=i686-w64-mingw32-gcc",
				)
			}
		}
	} else {
		envs = append(envs, "CGO_ENABLED=0")
	}
	if len(opts.Tags) != 0 {
		args = append(args, "-tags", strings.Join(opts.Tags, ","))
	}
	src := filepath.Join(*fSrc, "cmd", cmd)
	if !filepath.IsAbs(src) {
		src = "." + string(filepath.Separator) + src
	}
	args = append(args,
		"-o", bin,
		src,
	)
	return doEnvs(wd, envs, args...)
}

func do(cmd ...string) error {
	fmt.Println("+ " + strings.Join(cmd, " "))
	c := exec.Command(cmd[0], cmd[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Env = os.Environ()
	return c.Run()
}

func doEnvs(wd string, envs []string, cmd ...string) error {
	fmt.Println("+ " + strings.Join(cmd, " "))
	c := exec.Command(cmd[0], cmd[1:]...)
	c.Dir = wd
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Env = append(os.Environ(), envs...)
	return c.Run()
}
