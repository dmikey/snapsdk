package main

type Generator interface {
	Generate(snap Snap) (string, error)
	Name() string
}

func WriteSDK(generator Generator, snap Snap) (string, error) {
	return generator.Generate(snap)
}
