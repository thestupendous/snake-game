package definitions
var M,N uint32
var Dir byte			//can have 0,1,2,3 - for up, down, left, right directions respectively
/*
func Abra() int {
	return 90
}
*/
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


