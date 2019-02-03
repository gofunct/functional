package manipulate

import (
	"fmt"
	"reflect"
	"sort"
)

// Reflection is used in these functions so that Slices and arrays of strings,
// ints, and other types not implementing []interface{} can be worked with.
// For example, this is useful if you need to work on the output of regexs.

func list(v ...interface{}) []interface{} {
	return v
}

func Push(list interface{}, v interface{}) []interface{} {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		nl := make([]interface{}, l)
		for i := 0; i < l; i++ {
			nl[i] = l2.Index(i).Interface()
		}

		return append(nl, v)

	default:
		panic(fmt.Sprintf("Cannot push on type %s", tp))
	}
}

func Prepend(list interface{}, v interface{}) []interface{} {
	//return append([]interface{}{v}, list...)

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		nl := make([]interface{}, l)
		for i := 0; i < l; i++ {
			nl[i] = l2.Index(i).Interface()
		}

		return append([]interface{}{v}, nl...)

	default:
		panic(fmt.Sprintf("Cannot Prepend on type %s", tp))
	}
}

func Last(list interface{}) interface{} {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		if l == 0 {
			return nil
		}

		return l2.Index(l - 1).Interface()
	default:
		panic(fmt.Sprintf("Cannot find Last on type %s", tp))
	}
}

func First(list interface{}) interface{} {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		if l == 0 {
			return nil
		}

		return l2.Index(0).Interface()
	default:
		panic(fmt.Sprintf("Cannot find First on type %s", tp))
	}
}

func Rest(list interface{}) []interface{} {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		if l == 0 {
			return nil
		}

		nl := make([]interface{}, l-1)
		for i := 1; i < l; i++ {
			nl[i-1] = l2.Index(i).Interface()
		}

		return nl
	default:
		panic(fmt.Sprintf("Cannot find Rest on type %s", tp))
	}
}

func Initial(list interface{}) []interface{} {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		if l == 0 {
			return nil
		}

		nl := make([]interface{}, l-1)
		for i := 0; i < l-1; i++ {
			nl[i] = l2.Index(i).Interface()
		}

		return nl
	default:
		panic(fmt.Sprintf("Cannot find Initial on type %s", tp))
	}
}

func SortAlpha(list interface{}) []string {
	k := reflect.Indirect(reflect.ValueOf(list)).Kind()
	switch k {
	case reflect.Slice, reflect.Array:
		a := StrSlice(list)
		s := sort.StringSlice(a)
		s.Sort()
		return s
	}
	return []string{StrVal(list)}
}

func Reverse(v interface{}) []interface{} {
	tp := reflect.TypeOf(v).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(v)

		l := l2.Len()
		// We do not sort in place because the incoMing array should not be altered.
		nl := make([]interface{}, l)
		for i := 0; i < l; i++ {
			nl[l-i-1] = l2.Index(i).Interface()
		}

		return nl
	default:
		panic(fmt.Sprintf("Cannot find Reverse on type %s", tp))
	}
}

func Compact(list interface{}) []interface{} {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		nl := []interface{}{}
		var item interface{}
		for i := 0; i < l; i++ {
			item = l2.Index(i).Interface()
			if !IsEmpty(item) {
				nl = append(nl, item)
			}
		}

		return nl
	default:
		panic(fmt.Sprintf("Cannot compact on type %s", tp))
	}
}

func Uniq(list interface{}) []interface{} {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		dest := []interface{}{}
		var item interface{}
		for i := 0; i < l; i++ {
			item = l2.Index(i).Interface()
			if !inList(dest, item) {
				dest = append(dest, item)
			}
		}

		return dest
	default:
		panic(fmt.Sprintf("Cannot find Uniq on type %s", tp))
	}
}

func inList(haystack []interface{}, needle interface{}) bool {
	for _, h := range haystack {
		if reflect.DeepEqual(needle, h) {
			return true
		}
	}
	return false
}

func Without(list interface{}, omit ...interface{}) []interface{} {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		res := []interface{}{}
		var item interface{}
		for i := 0; i < l; i++ {
			item = l2.Index(i).Interface()
			if !inList(omit, item) {
				res = append(res, item)
			}
		}

		return res
	default:
		panic(fmt.Sprintf("Cannot find Without on type %s", tp))
	}
}

func Has(needle interface{}, haystack interface{}) bool {
	tp := reflect.TypeOf(haystack).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(haystack)
		var item interface{}
		l := l2.Len()
		for i := 0; i < l; i++ {
			item = l2.Index(i).Interface()
			if reflect.DeepEqual(needle, item) {
				return true
			}
		}

		return false
	default:
		panic(fmt.Sprintf("Cannot find Has on type %s", tp))
	}
}

// $list := [1, 2, 3, 4, 5]
// Slice $list     -> list[0:5] = list[:]
// Slice $list 0 3 -> list[0:3] = list[:3]
// Slice $list 3 5 -> list[3:5]
// Slice $list 3   -> list[3:5] = list[3:]
func Slice(list interface{}, indices ...interface{}) interface{} {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		if l == 0 {
			return nil
		}

		var start, end int
		if len(indices) > 0 {
			start = ToInt(indices[0])
		}
		if len(indices) < 2 {
			end = l
		} else {
			end = ToInt(indices[1])
		}

		return l2.Slice(start, end).Interface()
	default:
		panic(fmt.Sprintf("list should be type of Slice or array but %s", tp))
	}
}
