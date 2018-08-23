package model

type User struct {
	Name     string
	Age      int
	Username string
	Pwd      string
	Pet      []string
}

type Animal struct {
	ID    string
	Name  string
	Order string
}

type Toy struct {
	ID   string
	Name string
}

type People struct {
	ID   string
	Name string
	Age  uint
	Toys []Toy
}
