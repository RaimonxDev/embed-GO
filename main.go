package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

func (p Person) Greet() {
	fmt.Println("Hola")
}

type Employee struct {
	// embebemos la estructura de Person para que employee pueda acceder a sus valores y metodos
	Person
	salary float64
}

func NewEmployee(name string, age uint, salary float64) Employee {
	return Employee{NewPerson(name, age), salary}

}

func (e Employee) Payroll() {
	fmt.Println(e.salary * 0.90)
}

func main() {
	e := NewEmployee("ramon", 28, 60000)
	// por medio de Employee podemos ejecutar metodos de Person. debimos que lo embebimos en la estructura de Employee
	e.Greet()
	e.Payroll()
	// tambien a sus campos
	fmt.Println(e.Name)
	fmt.Println(e.Age)
	fmt.Println(e)
}
