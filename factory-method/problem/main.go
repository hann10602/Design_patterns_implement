package main

import "fmt"

type Employee struct {
	job string
}

func FirstService(){
	e := Employee{
		job: "Police",
	}

	fmt.Printf("This employee is %s", e.job)
}

func SecondService(){
	e := Employee{
		job: "Docter",
	}

	fmt.Printf("This employee is %s", e.job)
}

func main() {
	FirstService()
	SecondService()
}