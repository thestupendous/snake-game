package definitions
import (
	"time"
	"math/rand"
)
var M,N uint32
var Dir byte			//can have 0,1,2,3 - for up, down, left, right directions respectively
var OldDir byte			//can have 0,1,2,3 - for up, down, left, right directions respectively
var FoodLoc Coords
type MyBoard [][]string					//Game board
var Score int
var TickDelay int						//increasing this value in multiples of 100 makes game slower
var GameOverReason string				//throwing out reason for your "game over"
var UserDir = map[string]byte {
	"w":0,
	"s":1,
	"a":2,
	"d":3,
}
func (board MyBoard) String() string {
	out:= "╔"
	for i:=0;i<len(board[0]);i++{
		out += "═"
	}
	out += "╗\n"
	for i:=0;i<len(board);i++{
		out += "║"
		for j:=0;j<len(board[i]);j++{
			out += board[i][j]
		}
		out +="║\n"
	}
	out+= "╚"
	for i:=0;i<len(board[0]);i++{
		out += "═"
	}
	out += "╝\n"
	return out
} 
type Coords struct {
	X uint32
	Y uint32
}

type Snake struct {
	Tail []Coords		//could've used another field for snake's total length, but it can be retrieved from the queue size
}

func (s *Snake) Enqueue(newLoc Coords) Coords {
	s.Tail = append(s.Tail,newLoc)
	return newLoc
}

func (s *Snake) Dequeue() Coords {
	temp := s.Tail[0]
	s.Tail = s.Tail[1:]
	return temp
}


//-------------functions--------------------------
func PlaceFood(board MyBoard) {
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	x,y := uint32(r1.Intn(int(M))),uint32(r1.Intn(int(N)))
	for board[x][y]!=" " {
		x,y = uint32(r1.Intn(int(M))),uint32(r1.Intn(int(N)))
	}
	board[x][y]= "Ø"
	FoodLoc.X = x
	FoodLoc.Y = y
}

func PlaceInitialSnake(s *Snake, board  MyBoard) {
	midX := M / 2
	midY := N / 2
	_ = midX
	_ = midY
	s.Enqueue(Coords{midX,midY-1}) ; board[midX][midY-1]="─"
	s.Enqueue(Coords{midX,midY})   ; board[midX][midY]="─"
	s.Enqueue(Coords{midX,midY+1}) ; board[midX][midY+1]=">"
	// fmt.Println("len of snake after initialising ",len(s.Tail))
}

func NewHead(dir byte) string {
	switch dir {
	case 0: return "^"
	case 1: return "v"
	case 2: return "<"
	default: return ">"
	}
	return "E"
}

func GetDashBar(oldDir byte,newDir byte) string {
	m := map[byte]string {	0: "u",
		1:"d",
		2:"l",
		3:"r",
	}
	if m[oldDir]==m[newDir] {
		if m[oldDir]=="u" || m[oldDir]=="d" {
			return "│"
		}else if m[oldDir]=="l" || m[oldDir]=="r" {
			return "─"
		}
	} 
	if m[oldDir] == "u" {
		if m[newDir] == "l"{ return "┐" }else{  return "┌"}
	} else if m[oldDir] =="d" {
		if m[newDir] == "l" { return "┘"} else { return "└"}
	}else if m[oldDir] == "l" {
		if m[newDir] == "u" { return "└"} else { return "┌"}
	} else {		//old dir is right
		if m[newDir] == "u" { return "┘"} else { return "┐"}
	}
	return "L"
}

func ChecknCorrectWrongDir(){
	m := map[byte]string {		0: "u", 1:"d", 2:"l", 3:"r"}
	n,o := m[Dir],m[OldDir]
	if (n=="u" && o=="d" ) || (n=="d" && o=="u") || (n=="r" && o=="l") || (n=="l" && o=="r") {
		Dir = OldDir
	}
}


func GetNextCell(coord Coords,dir byte,board MyBoard) (Coords,bool) {
	res := coord
	switch dir {
	case 0:
		res.X-=1 ; break
	case 2:
		res.Y-=1 ; break
	case 1:
		res.X+=1 ; break
	case 3:
		res.Y+=1 ; break
	}
	if res.X >=M || res.Y>=N || res.X <0 || res.Y <0 {				//reason for dying - wall
		GameOverReason = "You hit the wall too Hard!!"
		return res,false
	}
	if board[res.X][res.Y]!=" " && board[res.X][res.Y]!="Ø"{		//reason for dying - tail hit
		GameOverReason = "Trying to eat yourself?!"
		return res,false
	}
	return res,true
}

//----------------------important function here----------------------------------------------
func UpdateBoard(board MyBoard, snake *Snake,won *bool, lost *bool) {
	// fmt.Println("board updated")
	// fmt.Println("len of Snake ",len(snake.Tail))
	last := int(len(snake.Tail) - 1)
	newCell , cont:= GetNextCell(snake.Tail[last],Dir, board)
	if !cont{
			// game is lost
			*lost = true
			return
	}
	var newHead=NewHead(Dir)
	snake.Enqueue(newCell)
	if len(snake.Tail) >= int( int(M)*int(N) - 5 )  {
		*won = true
		return
	}
	last = len(snake.Tail) - 1
	board[snake.Tail[last].X][snake.Tail[last].Y] = newHead
	board[snake.Tail[last-1].X][snake.Tail[last-1].Y] = GetDashBar(OldDir,Dir)
	
	if newCell==FoodLoc /*&& board[newCell.X][newCell.Y]=="Ø"*/ {		//no dequeue
		Score +=5
		PlaceFood(board)
	}else {
		board[snake.Tail[0].X][snake.Tail[0].Y]=" "
		snake.Dequeue()
	}


	OldDir = Dir
}