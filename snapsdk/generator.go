package main

type Generator interface {
	Generate(stubby Stubby) (string, error)
	Name() string
}

func WriteSDK(generator Generator, stubby Stubby) (string, error) {
	return generator.Generate(stubby)
}
