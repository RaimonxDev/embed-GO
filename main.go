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

// Que sucede si la struct de employee tiene un metodo con el mismo nombre de una struct embebida?
// Go oculta el metodo de Person y ejecuta el meotodo de Employee
func (e Employee) Greet() {
	fmt.Println("Hola desde empleado")

}

type Human struct {
	// age generará un conflicto con Person ya que las 2 estructuras tiene el mismo nombre
	// genera el conflicto debido a que embebemos human en employee
	Age      uint
	Children uint
}

func NewHuman(age, children uint) Human {
	return Human{age, children}

}

type Employee struct {
	// embebemos la estructura de Person para que employee pueda acceder a sus valores y metodos
	Person
	Human
	salary float64
}

func NewEmployee(name string, age uint, children uint, salary float64) Employee {
	return Employee{NewPerson(name, age), NewHuman(age, children), salary}

}

func (e Employee) Payroll() {
	fmt.Println(e.salary * 0.90)
}

func main() {
	e := NewEmployee("ramon", 28, 2, 60000)
	// por medio de Employee podemos ejecutar metodos de Person. debimos que lo embebimos en la estructura de Employee
	// Ejucatara el metodo de employee
	e.Greet()
	// Si queremos ejecutar el metodo de Greet de Person accedemos a el por el nombre de su strutc
	// GO NO SOBREESCRIBE LOS METODOS
	e.Person.Greet()
	e.Payroll()
	// tambien a sus campos
	fmt.Println(e.Name)
	// => esta lina genera un error por que age lo comparten Person y Human. Go nos dice que ambiguo fmt.Println(e.Age)
	// Tenemos que especificar a que struct haremos la referencia para age
	fmt.Println("AGE DE PERSON", e.Person.Age)
	fmt.Println("AGE DE HUMAN", e.Human.Age)
	fmt.Println(e)

}
