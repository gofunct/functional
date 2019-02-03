package reflect

import (
	"reflect"
)

func Dfault(d interface{}, given ...interface{}) interface{} {

	if IsEmpty(given) || IsEmpty(given[0]) {
		return d
	}
	return given[0]
}

// IsEmpty returns true if the given value Has the zero value for its type.
func IsEmpty(given interface{}) bool {
	g := reflect.ValueOf(given)
	if !g.IsValid() {
		return true
	}

	// Basically adapted from text/template.isTrue
	switch g.Kind() {
	default:
		return g.IsNil()
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return g.Len() == 0
	case reflect.Bool:
		return g.Bool() == false
	case reflect.Complex64, reflect.Complex128:
		return g.Complex() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return g.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return g.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return g.Float() == 0
	case reflect.Struct:
		return false
	}
}

// coalesce returns the First non-IsEmpty value.
func Coalesce(v ...interface{}) interface{} {
	for _, val := range v {
		if !IsEmpty(val) {
			return val
		}
	}
	return nil
}

// ternary returns the First value if the Last value is true, otherwise returns the second value.
func Ternary(vt interface{}, vf interface{}, v bool) interface{} {
	if v {
		return vt
	}

	return vf
}
