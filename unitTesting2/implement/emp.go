package implement

import (
	"time"
)

type Emp struct {
	Name string
	DOB  string
	Noye int
}

func (e Emp) Greet() string {
	return "Hello and Welcome " + e.Name
}

func (e Emp) Age() int {
	format := "2006-02-04"
	birthday, _ := time.Parse(format, e.DOB)
	now := time.Now()
	years := now.Year() - birthday.Year()
	
	return years
}

// Calculate  and store salary bonus on experience
// (1 year of experience=20000,2 to 5 = 50000,
// 5 to 10 = 100000, more than 10 200000)
func (e Emp) SalaryBonus() int {
	if e.Noye >= 1 && e.Noye < 2 {
		return 20000
	} else if e.Noye >= 2 && e.Noye < 5 {
		return 50000
	} else if e.Noye >= 5 && e.Noye < 10 {
		return 100000
	} else if e.Noye >= 10 {
		return 2000000
	} else {
		return 0
	}
}
