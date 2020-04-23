package generator

import (
	"fmt"
	"strings"
)

type Generator interface {
	Generate(settings, seed string) (patch []byte, spoilerLog string, err error)
}

func NewGenerator(id string) (Generator, error) {
	switch name, version := parseID(id); name {
	case "oot-randomizer":
		return NewOOTRandomizer(version), nil
	case "test":
		return NewTest(), nil
	default:
		return nil, fmt.Errorf("unknown generator: %s", id)
	}
}

func parseID(id string) (name, version string) {
	parts := strings.SplitN(id, ":", 2)
	switch len(parts) {
	case 0:
		return "", ""
	case 1:
		return parts[0], ""
	default:
		return parts[0], parts[1]
	}
}

type Test struct{}

func NewTest() *Test {
	return &Test{}
}

func (*Test) Generate(settings, seed string) ([]byte, string, error) {
	return []byte("generated binary for seed: " + seed),
		"generated spoiler log for seed: " + seed,
		nil
}
