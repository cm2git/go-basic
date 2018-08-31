package impl

//interface implement
type Man struct {
	Name string
	age  int
}

func (m Man) DisplaySalary() {
	panic("implement me")
}

func (m Man) CalculateLeavesLeft() int {
	panic("implement me")
}

type Woman struct {
	Name string
	age  int
}

func (m Man) QueryName() string {
	return m.Name
}

func (w *Woman) QueryName() string {
	w.Name = "rose"
	return w.Name
}
