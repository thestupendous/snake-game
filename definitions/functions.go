package definitions
import (
	"time"
	"math/rand"
)
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