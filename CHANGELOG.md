# Changelog

## v1.8.11

- Preparations to make more spells moddable.
- Hidden Death spell now correctly draws spinning skull on monsters.
- Allow using two more effects from scripts: `ENERGY_BOLT` and `PLASMA`.
- Lots of internal changes related to color handling.
- Allow force-switching maps in more cases.
- Stretch image option is now persisted.
- Fix: Potential crash when loading certain maps.
- Fix: Downloading maps via Nox protocol.
- Fix: Incorrect rendering of the observer effect.
- Fix: Rendering of abilities panel when player dies.
- Fix: Obliteration spell effect.
- Fix: Minimap rendering.
- Fix: Parsing of network filters from the config.
- Fix: Transparent color not rendering correctly in some GUI elements.
- Fix: Correctly override background color in GUI.
- Fix: Crash when reading gamma on the server.
- Fix: Server requiring audio/video libs.
- Fix: Text message when switching maps.
- Fix: `GREEN_BOLT` effect when called from a script.
- Fix: Build now doesn't require Git.

## v1.8.10

- First release as an open source project!
- Significant cleanup of the codebase.
- Fix long-standing bug for rendering lights in high-res.
- Long-awaited support for movie playback.
- Experimental support loading campaign in multiplayer.
- Reworked most of the rendering code.
- Allow console commands in Quest.
- Use shaders for gamma control. Should work more reliably.
- Send server resolution and player list to the lobby.
- Allow overriding and initial modding support for spells.
- Allow getting all spells, including hidden ones via `cheat spells all`.
- Allow removing summon limit via `cheat summon.nolimit`.
- Support XDG paths on Linux.
- Fix: Correctly handle `goto` console command.
- Fix: Minor physics deviation compared to vanilla.
- Fix: Teams creation in Quest.
- Fix: Reading gamma settings from config.

## v1.8.9

- OpenNox will now use the new lobby server and fallback to XWIS if no servers were found.
- Ping will now be shown in servers list for servers that support it.
- Server discovery was optimized and is now faster.
- Implemented a tiny builtin SSH remote console. Telnet support removed.
- Game load/save was optimized and works faster now.
- Fix: Not being able to connect to non-OpenNox servers.
- Fix: Spawned creatures being stuck when damaged.
- Fix: Magic missiles targeting dead entities.
- Fix: Energy bolt not working properly in traps.
- Fix: UI no longer freezes when discovering game servers.
- Fix: Player names not being shown in scores table.
- Fix: Sorting in the character list.

## v1.8.8

- Notify if OpenNox update is available.
- Server: add metrics for monitoring.
- Fix: Jumping no longer prevents from taking damage from lava.
- Fix: Do not announce Solo Quest games online.
- Fix: Team colors (regression from v1.8.7).
- Fix: Correctly set config dir for Snap.

## v1.8.7

- Support loading of TrueType and OpenType fonts.
- Automatically fix Russian translation and font files.
- A ton of internal changes and cleanup.
- Fix: Creatures and NPCs with ranging attack trying to fight melee only.
- Fix: Server API being stuck than there are no players on the server.
- Fix: Crash when falling into a pit.
- Fix: AI debug logging.

## v1.8.6

- Release OpenNox for Linux on [Snapcraft](https://snapcraft.io/opennox).
- Reading game files is now safer and more extendable.
- Preparations for saving and replaying Nox matches.
- Preparations for rendering TTF fonts in game engine.
- E2E mode now allows recording player input directly (to reproduce bugs).
- Safer and faster sprite rendering.
- Initial work on rendering map previews.
- A lot of internal refactorings and improvements.
- Allow configuring manual spell cast timeout.
- Allow binding any keys with `bind` command.
- New `cheat spawn` command to spawn items and monsters.
- Fix: Safer GUI code, preventing a major memory corruption.
- Fix: Races in the audio playback.
- Fix: Potential fix for native map downloads.
- Fix: Crashes when playing on broken maps.

## v1.8.5

- Add more resolution options for both regular and HD version.
- Move screen resolution to `opennox.yml` config.
- Fix: Instant victory if player commits a suicide in Arena.
- Fix: Properly propagate closed and private game flags to XWIS.
- Fix: It was possible to join closed games.
- Fix: Map filtering by game type.
- Fix: `cheat health` and `cheat mana` now work properly without arguments.

## v1.8.4

- Better performance for particle effects.
- Allow using emotes in campaign.
- New `cheat sage` command that gives all spells and scrolls, but doesn't make you invincible.
- New `cheat spells` command that only gives all spells for your class.
- New `cheat scrolls` command that only gives all best scrolls.
- New `cheat god` (as opposed to `set god`) that only makes you invincible, but won't give any spells.
- Fix: Set time limit for new key bindings to prevent the same action from executing to fast.
- Fix: Lightning spells were not hitting multiple targets as they should.
- Fix: Controlled creatures disappearing on map switch.
- Fix: Effects from Protection spells were played incorrectly.
- Fix: Random "you cannot wear this" messages when switching equipment.
- Fix: Allow really large amount of gold in `cheat gold`.
- Fix: Announcing Solo games to online lobby.
- Fix: Using relative paths in config.
- Fix: Better protection from potential memory leaks.
- Fix: Crash when reading maps that use the new Panic's script compiler.
- Fix: Multiple other crashes.

## v1.8.3

- Integrated the new (simple) server control panel. 
- Fix: Saving complex input bindings to the config.
- Fix: Wrong interpretation of XWIS option in the new config.
- Fix: Possible crashes for maps using new Panic's compiler.

## v1.8.2

- Write `opennox.yml` config with all extended options.
- Allow changing game data dir via the `opennox.yml` config.
- Fix: Player names being stuck as online in dedicated server.
- Fix: Mouse movement on low sensitivity settings.
- Fix: Assigning keyboard keys instead of mouse for movement.

## v1.8.1

- Fix: Back walls rendering in high-resolution mode.
- Fix: Center in-game menu in high-resolution mode.
- Fix: Possible crash after running for some time.

## v1.8.0

### New

- High-resolution support (up to 4K)! You'll need to run `opennox-hd.exe` for it.
- A lot faster OpenGL-based rendering (partially offloaded to GPU).
- Add `cheat charm.all` to charm any creature (including humanoids).
- Automated end-to-end testing mode, see [opennox-e2e](https://github.com/opennox/opennox-e2e) for details.
- Simple HTTP-based [Server API](./docs/server-api.md) for controlling game servers.
- Allow disabling image smoothing in the video options.
- Allow enabling image stretching in the video options.
- Other minor optimizations.

### Fixes

- Bug with creatures being "deaf" (not reacting to footsteps, etc).
- Properly show the last character in long server names.
- Error when manually saving player in online games.
- Banish of captured creatures (like vampire bats) leading to scripts being stuck.
- Properly forward ports when server port is changed via a flag.
- Disabling soft shadows via `-soft` flag.
- LUA scripts when reloading the same map from the console.
- Nox Reloaded failing to join OpenNox Quest games.
- OpenNox changing `user.rul` file permissions.
- Potential crash on hosting Quest games.
- Potential crash when leaving the game.
- Crash when opening advanced server setting without admin rights.

## v1.7.1

- Fix: Hide OS mouse cursor in windowed mode as well.
- Fix: Rename "Back" button in input config menu to "Accept".
- Fix: Prevent crash on audio errors.
- Fix: Avoid collision of new health and mana cheats with old ones.

## v1.7.0

- Completely reworked input and present pipelines.
- Manual spell casting is now supported. Bind keys to individual spell gestures and become a real wizard/conjurer!
- Gamma and mouse sensitivity sliders were added to the options menu.
- Wide-screen 16:9 resolutions are now available. Older 4:3 resolutions are still supported via config.
- Fix: Input sensitivity should now work correctly in fullscreen mode.
- Fix: Gamma setting in the config is now respected.
- Fix: Stretched video setting in the config is now respected.
- Added key bindings for:
  - Switching to a specific spell slot.
  - Switching to a specific trap slot.
  - Dropping a trap from your inventory.
  - Controlling spawned creature behavior.
  - Accepting item buy/sell/drop.
- Breaking change: `set allow.all` was renamed to `cheat equip.all`
- Breaking change: `set mana` was renamed to `cheat mana`
- Breaking change: `set health` was renamed to `cheat health`

## v1.6.1

- Fix: Prevent freeze when closing the game on Windows. 
- Fix: Prevent crash in `set mana`.
- Fix: Prevent crash when switching shield and bow with `set allow.all` enabled.
- Fix: Ignore item strength requirement when using `allow.all`.
- Fix: Remove restrictions from `cheat gold`, so it's accepted in Quest games.

## v1.6.0

- Rename project to OpenNox to avoid confusion.
- Automatically set Nox data folder. Supports Origin, GoG and Nox Reloaded.
- Keep 4:3 aspect ratio when resizing game window.
- Toggle fullscreen mode with Alt+Enter.
- Add `set health` and `set mana` cheats.
- Add `set allow.all` cheat to remove item class restrictions.
- Allow customizing `gamedata.bin` values without encoding it. See [gamedata.yml](./docs/gamedata-sample.yml).
- LUA now supports trigger events as well as player join/leave events.
- Fix: Detonate traps spell will no longer freeze the game.
- Fix: Fullscreen mode will now be saved correctly.
- Fix: Warriors should now be able to scroll weapons with mouse wheel in multiplayer as well.

## v1.5.0

- Automatically open/forward ports when hosting a game (no need to configure firewall).
- Fixed the bug that prevented warriors from scrolling weapons with mouse wheel.
- Blazing-fast map downloads. Both client and the server should use this version.
- Experimental support for [LUA map scripts](./docs/maps-lua.md).
- Maps will now be transferred with other related files when possible (`.txt`, `.rul`, `.lua`, `.png`).

## v1.4.0

- First public release!
- Compatible with Nox Reloaded.
- Native SDL and OpenAL support (no workarounds required for Win10, streaming, etc).
- Builtin XWIS integration. No need for account in order to host or join games.
- More [Nox Quest](./docs/game-quest.md) options (skip levels, keep portal forever, etc).
- Native Linux build (no Wine).
- Experimental gamepad support.
- Dedicated "headless" server.
