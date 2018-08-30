package model

type User struct {
	ID       string `bson:"_id"`
	Name     string
	Age      int
	Username string
	Pwd      string
	Pet      []string
	Toy      []Toy
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
	ID      string `bson:"_id"`
	Name    string
	Age     uint
	Toys    []Toy
	Account int `bson:"account"`
}
