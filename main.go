package main

import (
	"fmt"
	"time"
)

// struct는 기본적으로 call by value
type person struct{
	name string
	age int
}

func newPerson(name string, age int) * person{
	p := person{}
	p.name = name
	p.age = age

	return &p
}
func (np person) get() int {

	np.age = 100
	return np.age
}

func testfunc(st *int) {
	*st = 10
}
func main() {

	ch := make(chan int)

	ch <- 10

	fmt.Println("대기전")
	time.Sleep(time.Second*3)
	fmt.Println("대기후")
	go func() {
		var a int
		a = <- ch
		fmt.Println(a)
	}()


}