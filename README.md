# yet-another-game

Discover (fictional) space from the comfort of your personal terminal.

## How to play?

[Install Go](https://golang.org/doc/install) if you haven't already.

```
git clone https://github.com/Seklfreak/yet-another-game
cd yet-another-game
go build -o yet-another-game
./yet-another-game
```

### Extendability

The game is modular built using "actions". Each action can be hooked in at the game loop, or in other actions. The interface is defined in `models.Action`.

State is shared between actions using a shared state object. This state object is stored and retrieved from the disk on save or launch. The object is defined in `models.State`.

New fields can be added to the state at any time, allowing features to be added to previous saved games. 
