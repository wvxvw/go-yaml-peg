package goyamlpeg

import "testing"

func TestEnum(t *testing.T) {
	VerifyEnum(One{})
	VerifyEnum(Two{})
	VerifyEnum(Three{})
}
