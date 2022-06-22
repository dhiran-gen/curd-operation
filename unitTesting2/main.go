package main

import (
	"fmt"

	i "github.com/unitTesting2/implement"
)

func main(){

	e := i.Emp{Name: "Niraj",DOB: "2000-02-04",Noye: 2}

	fmt.Println(e.Greet())
	fmt.Println("Age = ",e.Age())
	fmt.Println("SalaryBonus = ",e.SalaryBonus())
}