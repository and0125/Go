package main

import "fmt"

type Student struct {
    name string
    grades []int
    age int
}

func (s Student) getAge() int {
    return s.age
}

// want to use the pointer to the student to permanently change the age value.
func (s *Student) setAge(age int) {
    s.age = age
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func main() {
    s1 := Student{"Aaron", []int{70,100}, 19}
	fmt.Println(s1)
    fmt.Println(s1.getAge())
    //equivalent to s1.age but wanted to use the method for this example

    s1.setAge(20)
    fmt.Println(s1)
	fmt.Println(s1.getAverageGrade())
}