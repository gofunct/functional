package constflag

import (
	"fmt"
	"strings"
)

// Enum is a `flag.Value` for one-of-a-fixed-set string arguments.
// The value of the `Choices` field defines the valid choices.
// If `CaseSensitive` is set to `true` (default `false`), the comparison is case-sensitive.
type Enum struct {
	Choices       []string
	CaseSensitive bool
	Default       string
	Value         string
	Text          string
}

func (fv *Enum) HasChanged() bool {
	if fv.Default != fv.Value {
		return true
	}
	return false
}

func (fv *Enum) Name() string {
	panic("implement me")
}

func (fv *Enum) ValueString() string {
	panic("implement me")
}

func (fv *Enum) ValueType() string {
	panic("implement me")
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *Enum) Help() string {
	if fv.CaseSensitive {
		return fmt.Sprintf("one of %v (case-sensitive)", fv.Choices)
	}
	return fmt.Sprintf("one of %v", fv.Choices)
}

// Set is flag.Value.Set
func (fv *Enum) Set(v string) error {
	fv.Text = v
	equal := strings.EqualFold
	if fv.CaseSensitive {
		equal = func(a, b string) bool { return a == b }
	}
	for _, c := range fv.Choices {
		if equal(c, v) {
			fv.Value = c
			return nil
		}
	}
	return fmt.Errorf(`"%s" must be one of [%s]`, v, strings.Join(fv.Choices, " "))
}

func (fv *Enum) String() string {
	return fv.Value
}

// Enums is a `flag.Value` for one-of-a-fixed-set string arguments.
// The value of the `Choices` field defines the valid choices.
// If `CaseSensitive` is set to `true` (default `false`), the comparison is case-sensitive.
type Enums struct {
	Choices       []string
	CaseSensitive bool

	Values []string
	Texts  []string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *Enums) Help() string {
	if fv.CaseSensitive {
		return fmt.Sprintf("one of %v (case-sensitive)", fv.Choices)
	}
	return fmt.Sprintf("one of %v", fv.Choices)
}

// Set is flag.Value.Set
func (fv *Enums) Set(v string) error {
	equal := strings.EqualFold
	if fv.CaseSensitive {
		equal = func(a, b string) bool { return a == b }
	}
	for _, c := range fv.Choices {
		if equal(c, v) {
			fv.Values = append(fv.Values, c)
			fv.Texts = append(fv.Texts, v)
			return nil
		}
	}
	return fmt.Errorf(`"%s" must be one of [%s]`, v, strings.Join(fv.Choices, " "))
}

func (fv *Enums) String() string {
	return strings.Join(fv.Values, ",")
}

// EnumsCSV is a `flag.Value` for comma-separated enum arguments.
// The value of the `Choices` field defines the valid choices.
// If `Accumulate` is set, the values of all instances of the flag are accumulated.
// The `Separator` field is used instead of the comma when set.
// If `CaseSensitive` is set to `true` (default `false`), the comparison is case-sensitive.
type EnumsCSV struct {
	Choices       []string
	Separator     string
	Accumulate    bool
	CaseSensitive bool

	Values []string
	Texts  []string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *EnumsCSV) Help() string {
	separator := ","
	if fv.Separator != "" {
		separator = fv.Separator
	}
	if fv.CaseSensitive {
		return fmt.Sprintf("%q-separated list of values from %v (case-sensitive)", separator, fv.Choices)
	}
	return fmt.Sprintf("%q-separated list of values from %v", separator, fv.Choices)
}

// Set is flag.Value.Set
func (fv *EnumsCSV) Set(v string) error {
	equal := strings.EqualFold
	if fv.CaseSensitive {
		equal = func(a, b string) bool { return a == b }
	}
	separator := fv.Separator
	if separator == "" {
		separator = ","
	}
	if !fv.Accumulate {
		fv.Values = fv.Values[:0]
	}
	parts := strings.Split(v, separator)
	for _, part := range parts {
		part = strings.TrimSpace(part)
		var ok bool
		var value string
		for _, c := range fv.Choices {
			if equal(c, part) {
				value = c
				ok = true
				break
			}
		}
		if !ok {
			return fmt.Errorf(`"%s" must be one of [%s]`, v, strings.Join(fv.Choices, " "))
		}
		fv.Values = append(fv.Values, value)
		fv.Texts = append(fv.Texts, part)
	}
	return nil
}

func (fv *EnumsCSV) String() string {
	return strings.Join(fv.Values, ",")
}
