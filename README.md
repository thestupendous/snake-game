# Snake Game
The classic snake game we have played in feature phones. The snake hunting food and growing with every time it eats. The game is implemented in golang

  
## How to Launch
Just launch main.go file using -

> go run main.go

The game starts with snake in the middle and starts moving.

## How to play the game
Use WASD  keys to move the snake. ( W - up, S - down, A - left, D - right)
![Game image](/doc/images/game-image.png)

## Scoring
You get 5 points for each food item eaten. Make sure to beat your older scores!

## Quit the game
Press Ctrl+C to quit the game. (You can also hit the walls to let the game end :) , you'll see the final score)

## Tweaking game settings
For now, the game settings are not provided on the game UI. You can change and tweak some game settings for your convenience by editing certain fields in the code itself. You have to change the values of 
some variables in the main.go file, as shown below
- Snake speed - It can be controlled by editing the d.TickDelay variable. Its default value is 200, the more the value, slower the game becomes. 
- Game Board Dimensions - The board length and breadths can be changed by updating d.M (height) and d.N (width) variables in main function. The default values for width and height are 40 and 20 respectively.

