package goyamlpeg

import "fmt"

type enum struct {}

type Enum interface {
	enum()
}

type One enum
type Two enum
type Three enum

func (e One) enum(){}
func (e Two) enum(){}
func (e Three) enum(){}

func VerifyEnum(e Enum) {
	switch e.(type) {
	case One:
		fmt.Println("One")
	case Two:
		fmt.Println("Two")
	case Three:
		fmt.Println("Three")
	}
}
