package main

type person struct{
	age int
}

func newPerson() *person{
	p := person{0}
	return &p
}
func (p person) Error() string{
	return "ERROR AT person"
}

func main() {
	p := newPerson()
	a := p.(person)
	println(a)
}