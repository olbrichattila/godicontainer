package godicontainer

import (
	"fmt"
	"reflect"
)

type CallbackFunc func() (interface{}, error)
type CallbackDefinitions map[string]CallbackFunc

type ContainerInterface interface {
	Get(string) interface{}
	Has(string) bool
	SetDefinitions([]string)
	Set(string, CallbackFunc)
}

type Container struct {
	definitions CallbackDefinitions
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Get(id string) (interface{}, error) {
	if callbackFunc, found := c.definitions[id]; found {
		resolvedStruct, err := callbackFunc()
		if err != nil {
			return nil, err
		}

		return resolvedStruct, nil
	}

	return nil, fmt.Errorf("Cannot resolve %s, use Set or 'SetDefinitions' to map it", id)
}

func (c *Container) Has(id string) bool {
	_, found := c.definitions[id]
	return found
}

func (c *Container) SetDefinitions(definitions CallbackDefinitions) {
	c.definitions = definitions
}

func (c *Container) Set(id string, callback CallbackFunc) {
	if c.definitions == nil {
		c.definitions = make(CallbackDefinitions)
	}
	c.definitions[id] = callback
}

func (c *Container) ResolvDependencies(s interface{}, sptr interface{}) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		fmt.Println("Not a struct.")
		return
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Kind() == reflect.Interface {
			fieldName := v.Type().Field(i).Name
			fieldType := field.Type().Name()
			diTag := v.Type().Field(i).Tag.Get("di")
			if diTag == "autowire" && c.Has(fieldType) {
				value, _ := c.Get(fieldType)
				c.resolveStructDepencency(sptr, fieldName, value)
			}
		}
	}
}

func (c *Container) resolveStructDepencency(str interface{}, fieldName string, value interface{}) error {
	val := reflect.ValueOf(str)
	if val.Kind() != reflect.Ptr || val.IsNil() || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("Invalid input. Expecting a pointer to a struct.")
	}

	field := val.Elem().FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("Field not found: %s", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("Field is not settable: %s", fieldName)
	}

	field.Set(reflect.ValueOf(value))

	return nil
}
