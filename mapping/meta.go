package mapping

import (
	"errors"
	"fmt"
	"reflect"
)

type IterStructObjectFunc func(ownRV reflect.Value, index int, args ...interface{}) error

func ScanAllFields(obj interface{}, iterFunc IterStructObjectFunc, callArgs ...interface{}) error {
	rt := reflect.TypeOf(obj)
	rv := reflect.ValueOf(obj)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return errors.New("the object is a nil pointer")
		}
		rv = rv.Elem()
		rt = rt.Elem()
	}

	for i := 0; i < rt.NumField(); i++ {
		if err := iterFunc(rv, i, callArgs...); err != nil {
			return err
		}
	}

	return nil
}

func CallFieldMethodByName(field reflect.Value, name string, callArgs ...reflect.Value) error {
	method := field.MethodByName(name)
	if !method.IsValid() || method.IsNil() {
		return fmt.Errorf("the %s function dose not exist", name)
	}

	retValues := method.Call(callArgs)
	retErr := retValues[0]
	if retErr.Interface() != nil {
		return retErr.Interface().(error)
	}
	return nil
}

func ListToValues(args ...interface{}) []reflect.Value {
	retValues := make([]reflect.Value, len(args))
	for n, v := range args {
		retValues[n] = reflect.ValueOf(v)
	}
	return retValues
}

func NewFieldByRV(rv reflect.Value) reflect.Value {
	if rv.Kind() != reflect.Ptr {
		return reflect.New(rv.Type())
	}

	return reflect.New(rv.Type().Elem())
}

func SetNewFieldToOwnRV(ownRV reflect.Value, index int) (retNewVal reflect.Value, retErr error) {
	if ownRV.Kind() == reflect.Ptr {
		if ownRV.IsNil() {
			retErr = errors.New("the ownRV is a nil pointer")
			return
		}

		ownRV = ownRV.Elem()
	}

	field := NewFieldByRV(ownRV.Field(index))
	ownRV.Field(index).Set(field)

	return field, nil
}

func NewFieldByRT(rt reflect.StructField) reflect.Value {
	if rt.Type.Kind() != reflect.Ptr {
		return reflect.New(rt.Type)
	}

	return reflect.New(rt.Type.Elem())
}

func GetReflectValueTypeName(v reflect.Value) string {
	if v.Kind() == reflect.Ptr {
		return v.Elem().Type().Name()
	}

	return v.Type().Name()
}
