package main

//Реализовать композицию структуры Action от родительской структуры Human
import "fmt"

type Human struct {
	age       int
	firstName string
	shureName string
}

type Action struct {
	human  Human
	action string
}

func (hum Human) printHuman() {
	fmt.Println("my name is ", hum.firstName, " ", hum.shureName, " my age is ", hum.age)
}
func (act Action) actHum() {
	fmt.Println(act.human.firstName, " likes ", act.action)
}
func main() {
	pers := Human{age: 21, firstName: "Sam", shureName: "Trump"}
	actPres := Action{action: "run", human: pers}
	actPres.actHum()
}
