package impl

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

//embedding interface
type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}
