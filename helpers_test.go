package footballapiv2

import (
	"reflect"
	"testing"
)

type ValueInt struct {
	Field int
}

type ValueFloat32 struct {
	Field float32
}

type ValueFloat64 struct {
	Field float64
}

func TestMakeLowerCase(t *testing.T) {
	str := "a"
	res := makeFirstLowerCase(str)

	if res != str {
		t.Error()
	}
}

func TestIntStructToMap(t *testing.T) {
	value := ValueInt{
		Field: -10,
	}
	res := structToMap(&value)
	if res.Get("field") != "-10" {
		t.Error()
	}
}

func TestFloat32StructToMap(t *testing.T) {
	value := ValueFloat32{
		Field: 3.14,
	}
	res := structToMap(&value)
	if res.Get("field") != "3.1400" {
		t.Error()
	}
}

func TestFloat64StructToMap(t *testing.T) {
	value := ValueFloat64{
		Field: 3.14,
	}
	res := structToMap(&value)
	if res.Get("field") != "3.1400" {
		t.Error()
	}
}

func TestEmptyStructToMap(t *testing.T) {
	res := structToMap(nil)
	if reflect.TypeOf(res).Kind() != reflect.Map {
		t.Error()
	}
}
