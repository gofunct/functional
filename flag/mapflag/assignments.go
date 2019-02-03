package mapflag

import (
	"fmt"
	"strings"
)

// Maps is a `flag.Value` for `KEY=VALUE` arguments.
// The value of the `Separator` field is used instead  of `"="` when set.
type Maps struct {
	Separator string

	Values []struct {
		Key   string
		Value string
	}
	Texts []string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *Maps) Help() string {
	separator := "="
	if fv.Separator != "" {
		separator = fv.Separator
	}
	return fmt.Sprintf("a key/value pair KEY%sVALUE", separator)
}

// Set is flag.Value.Set
func (fv *Maps) Set(v string) error {
	separator := "="
	if fv.Separator != "" {
		separator = fv.Separator
	}
	i := strings.Index(v, separator)
	if i < 0 {
		return fmt.Errorf(`"%s" must have the form KEY%sVALUE`, v, separator)
	}
	fv.Texts = append(fv.Texts, v)
	fv.Values = append(fv.Values, struct {
		Key   string
		Value string
	}{
		Key:   v[:i],
		Value: v[i+len(separator):],
	})
	return nil
}

func (fv *Maps) String() string {
	return strings.Join(fv.Texts, ", ")
}
