package main
import (
	d "github.com/thestupendous/snake-game/definitions"
	//"./definitions"
	"fmt"
 	"math/rand"
	"time"
	"os"
	"os/exec"
	"strconv"
)
type myBoard [][]string
func placeInitialSnake(s *d.Snake, board  myBoard) {
	midX := d.M / 2
	midY := d.N / 2
	_ = midX
	_ = midY
	s.Enqueue(d.Coords{midX,midY-1}) ; board[midX][midY-1]="─"
	s.Enqueue(d.Coords{midX,midY})   ; board[midX][midY]="─"
	s.Enqueue(d.Coords{midX,midY+1}) ; board[midX][midY+1]=">"
	// fmt.Println("len of snake after initialising ",len(s.Tail))
}
func (board myBoard) String() string {
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
func placeFood(board myBoard) {
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	x,y := uint32(r1.Intn(int(d.M))),uint32(r1.Intn(int(d.N)))
	for board[x][y]!=" " {
		x,y = uint32(r1.Intn(int(d.M))),uint32(r1.Intn(int(d.N)))
	}
	board[x][y]= "Ø"
	d.FoodLoc.X = x
	d.FoodLoc.Y = y
}
func gameWon(s d.Snake) bool {
	if uint32(len(s.Tail))>= d.M*d.N-1 {
		return true
	}
	return false
}

func getNextCell(coord d.Coords,dir byte,board myBoard) (d.Coords,bool) {
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
	if res.X >=d.M || res.Y>=d.N || res.X <0 || res.Y <0 {
		d.GameOverReason = "You hit the wall too Hard!!"
		return res,false
	}
	if board[res.X][res.Y]!=" " && board[res.X][res.Y]!="Ø"{
		d.GameOverReason = "Trying to eat yourself!!"
		return res,false
	}
	return res,true
}

func newHead(dir byte) string {
	switch dir {
	case 0: return "^"
	case 1: return "v"
	case 2: return "<"
	default: return ">"
	}
	return "E"
}
func getDashBar(oldDir byte,newDir byte) string {
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
//----------------------important function here----------------------------------------------
func updateBoard(board myBoard, snake *d.Snake,won *bool, lost *bool) {
	// fmt.Println("board updated")
	// fmt.Println("len of Snake ",len(snake.Tail))
	last := int(len(snake.Tail) - 1)
	newCell , cont:= getNextCell(snake.Tail[last],d.Dir, board)
	if !cont{
			// game is lost
			*lost = true
			return
	}
	var newHead=newHead(d.Dir)
	snake.Enqueue(newCell)
	if len(snake.Tail) >= int( int(d.M)*int(d.N) - 5 )  {
		*won = true
		return
	}
	last = len(snake.Tail) - 1
	board[snake.Tail[last].X][snake.Tail[last].Y] = newHead
	board[snake.Tail[last-1].X][snake.Tail[last-1].Y] = getDashBar(d.OldDir,d.Dir)
	
	if newCell==d.FoodLoc /*&& board[newCell.X][newCell.Y]=="Ø"*/ {		//no dequeue
		d.Score +=5
		placeFood(board)
	}else {
		board[snake.Tail[0].X][snake.Tail[0].Y]=" "
		snake.Dequeue()
	}


	d.OldDir = d.Dir
}

func checknCorrectWrongDir(){
	m := map[byte]string {		0: "u", 1:"d", 2:"l", 3:"r"}
	n,o := m[d.Dir],m[d.OldDir]
	if (n=="u" && o=="d" ) || (n=="d" && o=="u") || (n=="r" && o=="l") || (n=="l" && o=="r") {
		d.Dir = d.OldDir
	}
}

func main() {
	won,lost := false, false
	//defining dimensions of board
	d.M,d.N = 20,40
	//defining game speed
	d.TickDelay = 200
	//defining empty board of M,N dimensions
	var board myBoard
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


	// fmt.Println("len of snake before starting game ",len(sn.Tail))

	//game progression steps
	d.Score = 0
	d.Dir = d.UserDir["d"]
	d.OldDir = d.UserDir["d"]
	// ch := ""
	// for i:=0;i>0;i++ {
	// 	fmt.Scanln(&ch)
	// 	d.Dir = userDir[ch]
	// 	// fmt.Println("\033[H\033[2J")				//for clearing screen
	// 	updateBoard(board)
	// }

    ch := make(chan string)
    go func(ch chan string) {
        // disable input buffering
        exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
        // do not display entered characters on the screen
        exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
        var b []byte = make([]byte, 1)
        for {
			// time.Sleep(time.Millisecond * 1000)
            os.Stdin.Read(b)
            ch <- string(b)
        }
		if lost {
			close(ch)
		}
    }(ch)

	counter := 0
    for !lost {
        select {
            case stdin, _ := <-ch:
				counter+=1
                fmt.Println("Moves : ",counter,"\nKey pressed: ", stdin)
					//updateBoard(newDir)  //update snake, placefood, lost, won
					d.Dir = d.UserDir[stdin]
					checknCorrectWrongDir()
					updateBoard(board,&sn,&won,&lost)
            default:
                fmt.Println("Moves : ",counter)
					d.Dir = d.OldDir
					updateBoard(board,&sn,&won,&lost)
				//updateBoard(oldDir)
        }
        time.Sleep(time.Millisecond * time.Duration(d.TickDelay))
		fmt.Println("\033[H\033[2J")			//clear the screen
		fmt.Println(board)						//printing board at every clock tick
		fmt.Println("  Score : " + strconv.Itoa(d.Score))
    }

	if lost {
		fmt.Println("GAME OVER! ",d.GameOverReason)
	}
	fmt.Println("your final Score : ",d.Score)


}
