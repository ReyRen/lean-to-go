package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) showInfo() {
	fmt.Printf("studentName=%v age=%v score=%v\n", stu.Name, stu.Age, stu.Score)
}
func (stu *Student) SetScore(score int) {
	stu.Score = score
}

type Pupil struct {
	Student
}
type Graduate struct {
	Student
}

func (p *Graduate) testing() {
	fmt.Println("graduating is testing...")
}

func (p *Pupil) testing() {
	fmt.Println("pupil is testing...")
}

func main() {
	pupil := &Pupil{}
	pupil.Student.Name = "Tom"
	pupil.Age = 8
	pupil.testing()
	pupil.Student.SetScore(100)
	pupil.Student.showInfo()
}
