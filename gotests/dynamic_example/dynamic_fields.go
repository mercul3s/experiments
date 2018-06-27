package main

import (
	"encoding/json"
	"fmt"

	"github.com/sensu/sensu-go/types/dynamic"
)

// StaticType is a basic struct with a known field and one for arbitrary json
// values.
type StaticType struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Attrs []byte `json:",omitempty"` // note that this will not be marshalled directly, despite missing the `json"-"`!
}

// GetExtendedAttributes returns the attributes field
func (s *StaticType) GetExtendedAttributes() []byte {
	return s.Attrs
}

// SetExtendedAttributes sets the attributes field to the passed in value
func (s *StaticType) SetExtendedAttributes(a []byte) {
	s.Attrs = a
}

// Get returns an interface or error
func (s *StaticType) Get(name string) (interface{}, error) {
	return dynamic.GetField(s, name)
}

// MarshalJSON marshals the struct to json format
func (s *StaticType) MarshalJSON() ([]byte, error) {
	return dynamic.Marshal(s)
}

// UnmarshalJSON returns json into a struct
func (s *StaticType) UnmarshalJSON(p []byte) error {
	return dynamic.Unmarshal(p, s)
}

func main() {
	// predefined object with already filled in fields
	sType := StaticType{
		ID:   "testID",
		Name: "Michaelangelo",
	}

	// create a map to hold the values of the message we get
	var msgType map[string]interface{}
	msg := `{"field":"I like turtles", "another":{"key": "value"}}`
	json.Unmarshal([]byte(msg), &msgType)
	fmt.Println(msgType)
	// iterate over the keys and values in the map, and set them as keys and values
	// in sType
	for k, v := range msgType {

		dynamic.SetField(&sType, k, v)
	}
	// print out the struct - the values will be in Attrs
	fmt.Println(sType)
	// string version of Attrs
	fmt.Println(string(sType.Attrs))
	//  the marshalled representation of the object shows the keys and values as
	//  intended.
	s, _ := dynamic.Marshal(&sType)

	fmt.Println(string(s))
}
