# Snake Game
The classic snake game we have played in feature phones. The snake hunting food and growing with every time it eats. The game is implemented in golang

##Supported Platforms
Only linux platforms are supported. Windows support is coming soon :)
  
## How to Launch
Just launch main.go file using -

> go run main.go

The game starts with snake in the middle and starts moving.

## Deploying on Docker
For deploying the game on docker, 
- Make sure docker is installed
- in the home directory run 
  - `docker build -t snake-go .`
  - `docker run -i --tty --name snake-game snake-go`

### Removal 
For deleting game on docker, do
- `docker rm -f snake-game`
- `docker rmi -f snake-go`

## Deploying on Kubernetes
For deploying on k8s, 
- make sure kubernetes is installed and you are able to access kubectl or kubeadm command
- in home directory, run
  - `./k8s-deploy/one-click-deploy.sh`
- list your pod which comes up after a few seconds. Wait till it appears
  - `kubectl get po -n snake-go `
- after pod comes up, note down its name in previous command and run this
  - `kubectl exec -n snake-go -stdin --tty <pod-name> go run main.go`

### Removal from kubernetes
For deleting all k8s resources, in home directory, run
- `./k8s-deploy/delete-resource.sh `

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


## References
 - This game is built from scratch, no external resources are used. Still, some help is taken for inputting keypresses, without pressing enter, using this stackoverflow thread - https://stackoverflow.com/questions/54422309/how-to-catch-keypress-without-enter-in-golang-loop
