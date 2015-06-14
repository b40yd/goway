package goway

import (
	"fmt"
	"reflect"
)

type Injector interface {
	Invoke(interface{}) ([]reflect.Value, error)
	Map(interface{}) Injector
	Set(reflect.Type, reflect.Value) Injector
	Get(reflect.Type) reflect.Value
	SetParent(Injector)
	MapTo(interface{}, interface{}) Injector
	All() map[reflect.Type]reflect.Value
}

type injector struct {
	values map[reflect.Type]reflect.Value
	parent Injector
}

func (this *injector) All() map[reflect.Type]reflect.Value {
	for k,v := range this.values {
		fmt.Printf("key: %v  value: %v \n", k,v)
	}
	return this.values
}

func (intject *injector) SetParent(parent Injector) {
	intject.parent = parent
}

func InterfaceOf(value interface{}) reflect.Type {
	t := reflect.TypeOf(value)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Interface {
		panic("Called InterfaceOf with a value that is not a pointer to an interface. (*MyInterface)(nil)")
	}

	return t
}

func (intject *injector) Invoke(f interface{}) ([]reflect.Value, error) {
	t := reflect.TypeOf(f)
	// NumIn returns a function type's input parameter count.
	// It panics if the type's Kind is not Func.
	var params = make([]reflect.Value, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		// In returns the type of a function type's i'th input parameter.
		// It panics if the type's Kind is not Func.
		// It panics if i is not in the range [0, NumIn()).
		argType := t.In(i)
		val := intject.Get(argType)
		if !val.IsValid() {
			return nil, fmt.Errorf("Value not found for type %v", argType)
		}

		params[i] = val
	}

	return reflect.ValueOf(f).Call(params), nil
}

func (intject *injector) Map(val interface{}) Injector {
	intject.values[reflect.TypeOf(val)] = reflect.ValueOf(val)
	return intject
}

func (intject *injector) MapTo(val interface{}, interfacePtr interface{}) Injector {
	intject.values[InterfaceOf(interfacePtr)] = reflect.ValueOf(val)
	return intject
}

// Maps the given reflect.Type to the given reflect.Value and returns
// the Typemapper the mapping has been registered in.
func (intject *injector) Set(typ reflect.Type, val reflect.Value) Injector {
	intject.values[typ] = val
	return intject
}

func (intject *injector) Get(t reflect.Type) reflect.Value {
	// fmt.Println("injectoer get :", intject.values)
	val := intject.values[t]

	if val.IsValid() {
		return val
	}

	// no concrete types found, try to find implementors
	// if t is an interface
	if t.Kind() == reflect.Interface {
		for k, v := range intject.values {
			if k.Implements(t) {
				val = v
				break
			}
		}
	}

	// Still no type found, try to look it up on the parent
	if !val.IsValid() && intject.parent != nil {
		val = intject.parent.Get(t)
	}

	return val

}

// New returns a new Injector.
func NewInjector() Injector {
	return &injector{
		values: make(map[reflect.Type]reflect.Value),
	}
}
