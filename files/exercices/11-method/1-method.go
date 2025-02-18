package main

type person struct {
	firstName string
	lastName  string
}

// TODO: create the 'greet' method that prints the message: "Hello, I am <firstName> <lastName>\n"

func main() {
	alice := person{"alice", "smith"}
	bob := person{"bob", "williams"}

	alice.greet()
	bob.greet()
}
