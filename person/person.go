package person

type person struct {
	name   string
	age    int
	gender string
	Type   string
}

var (
	p = person{name: "Bob", age: 42, gender: "Male"}
)

func Description(name string) string {
	return "The person name is: " + p.name
}
