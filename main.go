package main
import (
	d "github.com/thestupendous/snake-game/definitions"
	//"./definitions"
	"fmt"
 	"math/rand"
	"time"
)
type myboard [][]string
func placeInitialSnake(s *d.Snake, board  myboard) {
	midX := d.M / 2
	midY := d.N / 2
	_ = midX
	_ = midY
//	c := d.Coords{1,2}
//	s.Enqueue(d.Coords{1,2})
//	fmt.Println("after enqueue s ",s)
	s.Enqueue(d.Coords{midX-1,midY}) ; board[midX][midY-1]="-"
	s.Enqueue(d.Coords{midX,midY})   ; board[midX][midY]="-"
	s.Enqueue(d.Coords{midX+1,midY}) ; board[midX][midY+1]=">"
	// fmt.Println("after enqueue s ",s)
}
func (board myboard) String() string {
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
func placeFood(board myboard) {
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	x,y := uint32(r1.Intn(int(d.M))),uint32(r1.Intn(int(d.N)))
	for board[x][y]!=" " {
		x,y = uint32(r1.Intn(int(d.M))),uint32(r1.Intn(int(d.N)))
	}
	board[x][y]= "Ø"
}
func gameWon(s d.Snake) bool {
	if uint32(len(s.Tail))>= d.M*d.N-1 {
		return true
	}
	return false
}
func updateBoard(board myBoard) {
	
}


func main() {
	//defining dimensions of board
	d.M,d.N = 20,20
	//defining empty board of M,N dimensions
	var board myboard
	board = make([][]string,d.M)
	for  i:=0;i<int(d.M);i++  {
		board[i] = make([]string,d.N)
	}

	//initialising board
	for i:=0 ; i< int(d.M) ; i++ {
		for j:=0 ; j < int(d.N) ; j++ {
			board[i][j] = string(" ")
		}
	}

	//initialising snake
	var sn d.Snake
	// board[0][4] = "A"
	// fmt.Println(board)
	placeInitialSnake(&sn,board)
	placeFood(board)				//for the first time

	fmt.Println(board)

	userDir := map[string]byte {
		"U":0,
		"D":1,
		"L":2,
		"R":3,
	}
	ch := ""
	for i:=0;i>0;i++ {
		fmt.Scanln(&ch)
		d.Dir = userDir(ch)
		// fmt.Println("\033[H\033[2J")				//for clearing screen
		updateBoard(board)
	}

}
