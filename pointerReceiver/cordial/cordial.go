package cordial

type Person struct {
	name string
}

func (ct Person) SayHi() {
	println("Hi", ct.name)
}

func (ct *Person) IsCalled(name string) {
	ct.name = name
}
