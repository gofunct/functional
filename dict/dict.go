package dict

import "github.com/imdario/mergo"

func Set(d map[string]interface{}, key string, value interface{}) map[string]interface{} {
	d[key] = value
	return d
}

func Unset(d map[string]interface{}, key string) map[string]interface{} {
	delete(d, key)
	return d
}

func HasKey(d map[string]interface{}, key string) bool {
	_, ok := d[key]
	return ok
}

func Pluck(key string, d ...map[string]interface{}) []interface{} {
	res := []interface{}{}
	for _, dict := range d {
		if val, ok := dict[key]; ok {
			res = append(res, val)
		}
	}
	return res
}

func Keys(dicts ...map[string]interface{}) []string {
	k := []string{}
	for _, dict := range dicts {
		for key := range dict {
			k = append(k, key)
		}
	}
	return k
}

func Pick(dict map[string]interface{}, keys ...string) map[string]interface{} {
	res := map[string]interface{}{}
	for _, k := range keys {
		if v, ok := dict[k]; ok {
			res[k] = v
		}
	}
	return res
}

func Omit(dict map[string]interface{}, keys ...string) map[string]interface{} {
	res := map[string]interface{}{}

	omit := make(map[string]bool, len(keys))
	for _, k := range keys {
		omit[k] = true
	}

	for k, v := range dict {
		if _, ok := omit[k]; !ok {
			res[k] = v
		}
	}
	return res
}

func Dict(v ...interface{}) map[string]interface{} {
	dict := map[string]interface{}{}
	lenv := len(v)
	for i := 0; i < lenv; i += 2 {
		key := StrVal(v[i])
		if i+1 >= lenv {
			dict[key] = ""
			continue
		}
		dict[key] = v[i+1]
	}
	return dict
}

func Merge(dst map[string]interface{}, srcs ...map[string]interface{}) interface{} {
	for _, src := range srcs {
		if err := mergo.Merge(&dst, src); err != nil {
			// Swallow errors inside of a template.
			return ""
		}
	}
	return dst
}

func Values(dict map[string]interface{}) []interface{} {
	values := []interface{}{}
	for _, value := range dict {
		values = append(values, value)
	}

	return values
}
