package definitions
var M,N uint32
var Dir byte			//can have u,d,l,r - for up, down, left, right directions
/*
func Abra() int {
	return 90
}
*/
type Coords struct {
	x uint32
	y uint32
}

type Snake struct {
	tail []Coords
}

func (s *Snake) Enqueue(newLoc Coords) Coords {
	s.tail = append(s.tail,newLoc)
	return newLoc
}

func (s *Snake) Dequeue() Coords {
	temp := s.tail[0]
	s.tail = s.tail[1:]
	return temp
}


var Board [][]Coords
