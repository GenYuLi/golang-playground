package util

import (
	"fmt"
)

// interface{} == any type
func AddBroken(a, b interface{}) interface{} {
	return a.(int) + b.(int)
}

// after Go 1.18, we can use type parameters to create a more flexible Add function
type Addable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | complex64 | complex128
}

func Add[T Addable](a, b T) T {
	return a + b
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func ToUnsigned[S Signed, U Unsigned](v S) U {
	if v < 0 {
		panic("negative value cannot be converted safely")
	}
	return U(v)
}

type Animal interface {
	Speak() string
}

type Dog struct{}

func (Dog) Speak() string { return "Woof" }

func SaySomething(a Animal) {
	fmt.Println(a.Speak())
}
