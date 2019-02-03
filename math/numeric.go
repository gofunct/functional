package math

import (
	"math"
	"reflect"
	"strconv"
)

// ToFloat64 converts 64-bit floats
func ToFloat64(v interface{}) float64 {
	if str, ok := v.(string); ok {
		iv, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0
		}
		return iv
	}

	val := reflect.Indirect(reflect.ValueOf(v))
	switch val.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return float64(val.Int())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return float64(val.Uint())
	case reflect.Uint, reflect.Uint64:
		return float64(val.Uint())
	case reflect.Float32, reflect.Float64:
		return val.Float()
	case reflect.Bool:
		if val.Bool() == true {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func ToInt(v interface{}) int {
	//It's not optimal. Bud I don't want duplicate ToInt64 code.
	return int(ToInt64(v))
}

// ToInt64 converts integer types to 64-bit integers
func ToInt64(v interface{}) int64 {
	if str, ok := v.(string); ok {
		iv, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return 0
		}
		return iv
	}

	val := reflect.Indirect(reflect.ValueOf(v))
	switch val.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return val.Int()
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return int64(val.Uint())
	case reflect.Uint, reflect.Uint64:
		tv := val.Uint()
		if tv <= math.MaxInt64 {
			return int64(tv)
		}
		// TODO: What is the sensible thing to do here?
		return math.MaxInt64
	case reflect.Float32, reflect.Float64:
		return int64(val.Float())
	case reflect.Bool:
		if val.Bool() == true {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func Max(a interface{}, i ...interface{}) int64 {
	aa := ToInt64(a)
	for _, b := range i {
		bb := ToInt64(b)
		if bb > aa {
			aa = bb
		}
	}
	return aa
}

func Min(a interface{}, i ...interface{}) int64 {
	aa := ToInt64(a)
	for _, b := range i {
		bb := ToInt64(b)
		if bb < aa {
			aa = bb
		}
	}
	return aa
}

func Until(count int) []int {
	step := 1
	if count < 0 {
		step = -1
	}
	return UntilStep(0, count, step)
}

func UntilStep(start, stop, step int) []int {
	v := []int{}

	if stop < start {
		if step >= 0 {
			return v
		}
		for i := start; i > stop; i += step {
			v = append(v, i)
		}
		return v
	}

	if step <= 0 {
		return v
	}
	for i := start; i < stop; i += step {
		v = append(v, i)
	}
	return v
}

func Floor(a interface{}) float64 {
	aa := ToFloat64(a)
	return math.Floor(aa)
}

func Ceil(a interface{}) float64 {
	aa := ToFloat64(a)
	return math.Ceil(aa)
}

func Round(a interface{}, p int, r_opt ...float64) float64 {
	roundOn := .5
	if len(r_opt) > 0 {
		roundOn = r_opt[0]
	}
	val := ToFloat64(a)
	places := ToFloat64(p)

	var round float64
	pow := math.Pow(10, places)
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}
