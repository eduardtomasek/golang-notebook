package main

import (
	"fmt"
	"reflect"
)

type rowStruct struct {
	TopicID        string `sql:"topic_id"`
	TitleCanonical string `sql:"title_canonical"`
}

func main() {
	var str rowStruct

	// func ValueOf(i interface{}) Value
	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i. ValueOf(nil) returns the zero Value.
	structVal := reflect.ValueOf(&str)
	fmt.Printf("ValueOf: %s\n", structVal.Type())

	// func (v Value) Type() Type
	// Type returns v's type.
	typ := structVal.Type()
	fmt.Printf("Type: %s\n", typ)

	// func (v Value) Kind() Kind
	// Kind returns v's Kind. If v is the zero Value (IsValid returns false),
	// Kind returns Invalid.
	kind := typ.Kind()
	fmt.Printf("Type -> Kind: %s\n", kind)

	// func (v Value) Elem() Value
	// Elem returns the value that the interface v contains or that
	// the pointer v points to. It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.
	elem := typ.Elem()

	// func (v Value) NumField() int
	// NumField returns the number of fields in the struct v.
	// It panics if v's Kind is not Struct.
	fmt.Printf("Type -> Elem -> NumField: %d\n", elem.NumField())

	// func (v Value) Field(i int) Value
	// Field returns the i'th field of the struct v.
	// It panics if v's Kind is not Struct or i is out of range.
	field := elem.Field(0)

	// func (tag StructTag) Get(key string) string
	// Get returns the value associated with key in the tag string.
	// If there is no such key in the tag, Get returns the empty string.
	// If the tag does not have the conventional format, the value
	// returned by Get is unspecified. To determine whether a tag
	// is explicitly set to the empty string, use Lookup.
	fmt.Printf("Type -> Elem -> Field -> Tag -> Get('sql'): %s\n", field.Tag.Get("sql"))

	fmt.Println(str)
	structValElem := structVal.Elem()
	elementAddres := structValElem.FieldByName("TopicID")
	tmpVal := "TestValue"
	elementAddres.Set(reflect.ValueOf(tmpVal))
	fmt.Println(str)
}
