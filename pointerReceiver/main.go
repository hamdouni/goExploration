package main

import "pointerReceiver/cordial"

func main() {
	var person cordial.Person
	person.IsCalled("Jo")
	present(&person)
}

func present(it politesse) {
	it.SayHi()
}

type politesse interface {
	IsCalled(string)
	SayHi()
}
