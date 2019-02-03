package mapflag

import (
	"fmt"
	"strings"
)

// Map is a `flag.Value` for `KEY=VALUE` arguments.
// The value of the `Separator` field is used instead  of `"="` when set.
type Map struct {
	Separator string

	Value struct {
		Key   string
		Value string
	}
	Text string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *Map) Help() string {
	separator := "="
	if fv.Separator != "" {
		separator = fv.Separator
	}
	return fmt.Sprintf("a key/value pair KEY%sVALUE", separator)
}

// Set is flag.Value.Set
func (fv *Map) Set(v string) error {
	separator := "="
	if fv.Separator != "" {
		separator = fv.Separator
	}
	i := strings.Index(v, separator)
	if i < 0 {
		return fmt.Errorf(`"%s" must have the form KEY%sVALUE`, v, separator)
	}
	fv.Text = v
	fv.Value = struct {
		Key   string
		Value string
	}{
		Key:   v[:i],
		Value: v[i+len(separator):],
	}
	return nil
}

func (fv *Map) String() string {
	return fv.Text
}
