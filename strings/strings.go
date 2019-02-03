package strings

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	util "github.com/aokoli/goutils"
)

func Base64encode(v string) string {
	return base64.StdEncoding.EncodeToString([]byte(v))
}

func Base64decode(v string) string {
	data, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func Base32encode(v string) string {
	return base32.StdEncoding.EncodeToString([]byte(v))
}

func Base32decode(v string) string {
	data, err := base32.StdEncoding.DecodeString(v)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func Abbrev(width int, s string) string {
	if width < 4 {
		return s
	}
	r, _ := util.Abbreviate(s, width)
	return r
}

func Abbrevboth(left, right int, s string) string {
	if right < 4 || left > 0 && right < 7 {
		return s
	}
	r, _ := util.AbbreviateFull(s, left, right)
	return r
}
func Initials(s string) string {
	// Wrap this just to eliMinate the var args, which templates don't do well.
	return util.Initials(s)
}

func RandAlphaNumeric(count int) string {
	// It is not possible, it appears, to actually generate an error here.
	r, _ := util.RandomAlphaNumeric(count)
	return r
}

func RandAlpha(count int) string {
	r, _ := util.RandomAlphabetic(count)
	return r
}

func RandAscii(count int) string {
	r, _ := util.RandomAscii(count)
	return r
}

func RandNumeric(count int) string {
	r, _ := util.RandomNumeric(count)
	return r
}

func Untitle(str string) string {
	return util.Uncapitalize(str)
}

func Quote(str ...interface{}) string {
	out := make([]string, len(str))
	for i, s := range str {
		out[i] = fmt.Sprintf("%q", StrVal(s))
	}
	return strings.Join(out, " ")
}

func SQuote(str ...interface{}) string {
	out := make([]string, len(str))
	for i, s := range str {
		out[i] = fmt.Sprintf("'%v'", s)
	}
	return strings.Join(out, " ")
}

func Cat(v ...interface{}) string {
	r := strings.TrimSpace(strings.Repeat("%v ", len(v)))
	return fmt.Sprintf(r, v...)
}

func Indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

func NIndent(spaces int, v string) string {
	return "\n" + Indent(spaces, v)
}

func Replace(old, new, src string) string {
	return strings.Replace(src, old, new, -1)
}

func Plural(one, many string, count int) string {
	if count == 1 {
		return one
	}
	return many
}

func StrSlice(v interface{}) []string {
	switch v := v.(type) {
	case []string:
		return v
	case []interface{}:
		l := len(v)
		b := make([]string, l)
		for i := 0; i < l; i++ {
			b[i] = StrVal(v[i])
		}
		return b
	default:
		val := reflect.ValueOf(v)
		switch val.Kind() {
		case reflect.Array, reflect.Slice:
			l := val.Len()
			b := make([]string, l)
			for i := 0; i < l; i++ {
				b[i] = StrVal(val.Index(i).Interface())
			}
			return b
		default:
			return []string{StrVal(v)}
		}
	}
}

func StrVal(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func Trunc(c int, s string) string {
	if len(s) <= c {
		return s
	}
	return s[0:c]
}

func Join(sep string, v interface{}) string {
	return strings.Join(StrSlice(v), sep)
}

func Split(sep, orig string) map[string]string {
	parts := strings.Split(orig, sep)
	res := make(map[string]string, len(parts))
	for i, v := range parts {
		res["_"+strconv.Itoa(i)] = v
	}
	return res
}

func Splitn(sep string, n int, orig string) map[string]string {
	parts := strings.SplitN(orig, sep, n)
	res := make(map[string]string, len(parts))
	for i, v := range parts {
		res["_"+strconv.Itoa(i)] = v
	}
	return res
}

// Substring creates a Substring of the given string.
//
// If start is < 0, this calls string[:length].
//
// If start is >= 0 and length < 0, this calls string[start:]
//
// Otherwise, this calls string[start, length].
func Substring(start, length int, s string) string {
	if start < 0 {
		return s[:length]
	}
	if length < 0 {
		return s[start:]
	}
	return s[start:length]
}
