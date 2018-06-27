package example

import "fmt"

// Example ...
type Example interface {
	Get() string
}

// TwoAttrs is a struct to hold two string attributes.
type TwoAttrs struct {
	attrOne string
	attrTwo string
}

// Get prints out the two attributes.
func (ta *TwoAttrs) Get() string {
	return fmt.Sprintf(ta.attrOne, ta.attrTwo)
}
