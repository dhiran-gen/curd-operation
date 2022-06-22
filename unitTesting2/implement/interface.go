// Get employee details such as name, date of birth and no of years of experience

// Write methods to greet the employee
// Calculate and store the age given the date of birth
// Calculate  and store salary bonus on experience
// (1 year of experience=20000,2 to 5 = 50000, 5 to 10 = 100000, more than 10 200000)
// Print all the details of the employee in the end

package implement

type Get interface {
	Greet() string
	Age() int
	SalaryBonus() int
}
