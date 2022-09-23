# PlayWordle

To play the game

```bash
go run *.go
```

To play, you will get the first word, then you can write the colors in command line (for eg. YYGGB ), then it will provide you with the next word. Write "WIN" in command line to exit after victory ;)

## The Metrics (PEAK)

```
Success Percentage 94.56643692899503
Average Win moves: 3.5072033898305084
Average Loss moves: 7.961651917404129
Time taken per word 2.087277505048886 ms
```

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

Can advance the algorithm by adding ways to deal with similar words.

