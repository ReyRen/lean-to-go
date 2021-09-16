package model

/*type Student struct {
	Name string
	Score float64
}*/

type student struct {
	Name  string
	Score float64
}

type student2 struct {
	Name  string
	score float64
}

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		Score: s,
	}
}

func NewStudent2(n string, s float64) *student2 {
	return &student2{
		Name:  n,
		score: s,
	}
}

func (stu *student2) GetScore() float64 {
	return stu.score
}
