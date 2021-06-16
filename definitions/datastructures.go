package definitions

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


