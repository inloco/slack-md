package slack_md

import (
	"fmt"
	"testing"
)

func TestMarshalGoDefault(t *testing.T) {
	input := []string{"value_one"}
	marshaled, err := Marshal(input)
	if err != nil {
		t.FailNow()
	}

	expected := fmt.Sprintf("%v", input)
	if expected != string(marshaled) {
		t.Fail()
	}
}

func TestMarshalString(t *testing.T) {
	input := "value_one"
	marshaled, err := Marshal(input)
	if err != nil {
		t.FailNow()
	}

	expected := input
	if expected != string(marshaled) {
		t.Fail()
	}
}

type CustomMarshaling struct{}

func (t CustomMarshaling) MarshalSlackMD() ([]byte, error) {
	return []byte("custom marshaling"), nil
}
func TestMarshalMarshaler(t *testing.T) {
	input := CustomMarshaling{}
	marshaled, err := Marshal(input)
	if err != nil {
		t.FailNow()
	}

	expected := "custom marshaling"
	if expected != string(marshaled) {
		t.Fail()
	}
}

func TestMarshalValidPointer(t *testing.T) {
	input := "value_one"
	marshaled, err := Marshal(&input)
	if err != nil {
		t.FailNow()
	}

	expected := input
	if expected != string(marshaled) {
		t.Fail()
	}
}
func TestMarshalNullPointer(t *testing.T) {
	var input *string
	marshaled, err := Marshal(input)
	if err != nil {
		t.FailNow()
	}

	expected := "null"
	if expected != string(marshaled) {
		t.Fail()
	}
}

type MyStruct struct {
	FirstField  string
	SecondField string

	OmittedField    string `slack_md:"-"`
	ObfuscatedField string `slack_md:"obfuscate"`
}

func TestMarshalStruct(t *testing.T) {
	input := MyStruct{"first_value", "second_value", "omitted", "0123456789"}

	expected := `*FirstField*: first_value
*SecondField*: second_value
*ObfuscatedField*: ******6789`

	marshaled, _ := Marshal(input)

	if expected != string(marshaled) {
		t.Fail()
	}
}

type InnerStruct struct {
	InnerField  string
	InnerStruct *InnerStruct
}

type MultiLevelStruct struct {
	FirstField  string
	InnerStruct *InnerStruct
}

func TestMarshalMultiLevelStruct(t *testing.T) {
	input := MultiLevelStruct{
		"first_value",
		&InnerStruct{
			"first_inner_value",
			&InnerStruct{
				"second_inner_value",
				nil,
			},
		},
	}

	expected := `*FirstField*: first_value
*InnerStruct*: 
	*InnerField*: first_inner_value
	*InnerStruct*: 
		*InnerField*: second_inner_value
		*InnerStruct*: null`

	marshaled, err := Marshal(input)
	if err != nil {
		t.FailNow()
	}

	if expected != string(marshaled) {
		t.Fail()
	}
}
