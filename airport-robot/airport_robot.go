package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeting interface {
	Hello(name string) string
}

type Italian struct{}

func (Italian) Hello(name string) string {
	return "I can speak Italian: Ciao " + name + "!"
}

type Portuguese struct{}

func (Portuguese) Hello(name string) string {
	return "I can speak Portuguese: Olá " + name + "!"
}

func SayHello(name string, g Greeting) string {
	return g.Hello(name)
}

func SayHello_Italian(name string) string {
	return SayHello(name, Italian{})
}

func SayHello_Portuguese(name string) string {
	return SayHello(name, Portuguese{})
}


