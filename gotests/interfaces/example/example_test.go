package example

import (
	"github.com/stretchr/testify/mock"
)

type mockExample struct {
	mock.Mock
}

func TestGet(t *testing.T) {
	ta := TwoAttrs{
		attrOne: "hello",
		attrTwo: "world",
	}

	example := &mockExample{}

}
