package definitions
var M,N unit32
var Dir byte			//can have u,d,l,r - for up, down, left, right directions

type Coords struct {
	var x,y uint32
}

type Snake struct {
	var tail []Coords
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


var Board [M][N]Coords

