package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Triangle struct {
	a, b, c float64
}

//NewTriangle return the triangle pointer if is valid
func NewTriangle(a float64, b float64, c float64) (*Triangle, error) {
	if (a+b) > c && (b+c) > a && (c+a) > b {
		return &Triangle{a, b, c}, nil
	}
	return nil, errors.New("This is not a triangle")
}

func (t *Triangle) isEquilatero() bool {
	return (t.a == t.b) && (t.b == t.c)
}

func (t *Triangle) isIsosceles() bool {
	return (t.a == t.b) && (t.c != t.a) || (t.a == t.c) && (t.b != t.a) || (t.b == t.c) && (t.a != t.c)
}

func (t *Triangle) isEscaleno() bool {
	return (t.a != t.b) && (t.a != t.c) && (t.b != t.c)
}

func main() {

	values := os.Args[1:]

	if len(values) != 3 {
		//TODO Add recovery
		panic("You should set three paraments")
	}

	var newValues = make([]float64, 3)
	for k, v := range values {
		if f, err := strconv.ParseFloat(v, 64); err == nil && f > 0 {
			newValues[k] = f
			continue
		}
		//TODO Add recovery
		panic("You must pass a numeric greater than zero floating value")
	}

	t, err := NewTriangle(newValues[0], newValues[1], newValues[2])
	if err != nil {
		//TODO Add recovery
		panic(err.Error())
	}

	fmt.Printf("isEquilatero: %t\n", t.isEquilatero())
	fmt.Printf("isIsosceles: %t\n", t.isIsosceles())
	fmt.Printf("isEscaleno: %t\n", t.isEscaleno())
}
