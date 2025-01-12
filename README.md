# Trixie the Truffler
In this cozy puzzle game, you play as the titular Trixie; a truffle pig on the hunt for mushrooms deep within the forest. 

Unfortunately, Trixie is bad at directions. It's your job to help her find the truffles in each level. You can only move Trixie left or down. However, you can give her a "change in perspective" by shifting (rotating) the board to her left or behind her.

This was created for the [Ebitengine Holiday Hack 2024](https://itch.io/jam/ebitengine-holiday-hack-2024) game jam.

### Controls:
* **A** - move Trixie left
* **S** - move Trixie down
* ← - rotate board left
* ↓ - rotate board behind
* **R** - restart level*
* **Esc** - open pause menu

*you can get stuck in some levels so this is your get-out-of-jail-free card.

## itch.io
You can play it in the browser or download the desktop version on itch.io [here](https://milk9111.itch.io/trixie-the-truffler).

## Running the game
You can run the game locally from source using 
```
go run .
```
or you can run a local WASM build with the make command
```
make web-serve
```
and navigate to [localhost:8080](http://localhost:8080)

## Making new levels 
If you want to make your own levels, you can run the `cmd/newlevel` command which will make a new JSON file in the `assets/levels` directory. It requires the flag `-cr [columns],[rows]` and the flag `-n [level name]`. 

The empty level file's `data` field is an array of strings formatted to resemble the board in-game for easier level layout. These are the supported characters for each cell:

* `P` - Player (_can only have one_)
* `G` - Goal (_can only have one_)
* `f` - Rock
* `s` - Bunny

Make sure to add your level to the `level_order` array in the `assets/levels/_config.json` file in order for it to show up with the rest of the levels in-game. You can use the debug flag `-L [level name]` when running the game to skip the main menu and load your level immediately. 

While the game is running, you can press the **F key** to do a hot reload of your level file. _You may need to uncomment the debug system in the `scene/game.go` in order for that key to work._