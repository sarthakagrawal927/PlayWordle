# PlayWordle

```bash
go run *.go
```

## The game

Yet to be build

## The helper

[helper.go](helper.go) - has functions to initialize the game and to run on every new word tried

## The Player

[player.go](player.go) - has functions to be used by the player to get the most probable word

Hopefully ready for use

### TC#1 - BRINE

Wordle 259 5/6

⬛🟩⬛⬛🟩
⬛🟩🟩🟩🟩
⬛🟩🟩🟩🟩
⬛🟩🟩🟩🟩
🟩🟩🟩🟩🟩

For such situations one possible solution, once we have multiple final letters - we can try to eliminate the most common letters by using an entirely different word made with non selected characters.
