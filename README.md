# PlayWordle

To play the game

```bash
go run *.go
```

To play, you will get the first word, then you can write the colors in command line (for eg. YYGGB ), then it will provide you with the next word. Write "WIN" in command line to exit after victory ;)

## The game

Yet to be build

## The helper

[helper.go](helper.go) - has functions to initialize the game and to run on every new word tried

## The Player

[player.go](player.go) - has functions to be used by the player to get the most probable word

Hopefully ready for use

### TC#1 - BRINE - 05/03

Wordle 259 5/6

⬛🟩⬛⬛🟩

⬛🟩🟩🟩🟩

⬛🟩🟩🟩🟩

⬛🟩🟩🟩🟩

🟩🟩🟩🟩🟩

For such situations one possible solution, once we have multiple final letters - we can try to eliminate the most common letters by using an entirely different word made with non selected characters.

Has failed for similar situation, so solution is needed. Try to fix 2 starting words maybe ?
Will it effect success probability ? will it increase/decrease the number of attempts ?

Playing archives from [here](https://metzger.media/games/wordle-archive/)

### TC#2 - MOURN - 03/03

⬜🟨🟨⬜⬜

⬜⬜🟨🟨⬜

🟨🟩🟩🟨⬜

⬜🟩🟩🟩🟩

🟩🟩🟩🟩🟩

### TC#3 - NASTY - 02/03

🟨⬜⬜🟨⬜

⬜⬜🟨🟨🟨

🟨🟩⬜🟨🟨

🟩🟩🟩🟩🟩

### TC#4 - RUPEE - 01/03

⬜🟨⬜⬜🟩

⬜⬜⬜🟨🟩

🟩⬜⬜⬜🟩

🟩🟩🟩🟩🟩

### TC#5 - CHOKE - 28/02

⬜⬜🟩⬜🟩

🟩⬜🟩⬜🟩

🟩🟩🟩⬜🟩

🟩🟩🟩🟩🟩

### TC#6 - CHANT - 27/02

🟨⬜⬜⬜⬜

⬜⬜🟨🟨⬜

⬜🟨🟩⬜🟩

🟩🟩🟩🟩🟩

### TC#7 - SPILL - 26/02

⬜⬜⬜⬜🟨

🟩🟨⬜🟨⬜

🟩⬜🟩🟩⬜

🟩🟩🟩🟩🟩

So the algorithm for playing the game is definitely a success, let's see how much optimization can be done.

Let us simulate for different types of algorithms
