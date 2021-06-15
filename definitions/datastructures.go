package definitions
var M,N uint32
var Dir byte			//can have 0,1,2,3 - for up, down, left, right directions respectively
var OldDir byte			//can have 0,1,2,3 - for up, down, left, right directions respectively
var FoodLoc Coords
var Score int
var UserDir = map[string]byte {
	"w":0,
	"s":1,
	"a":2,
	"d":3,
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


