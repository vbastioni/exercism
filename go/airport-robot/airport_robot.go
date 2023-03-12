package airportrobot

import "fmt"

func SayHello(name string, greeter Greeter) string {
	return greeter.Greet(name)
}

type Greeter interface {
	Greet(name string) string
}

type Italian struct{}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("I can speak Italian: Ciao %s!", name)
}

type Portuguese struct{}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("I can speak Portuguese: Olá %s!", name)
}

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
