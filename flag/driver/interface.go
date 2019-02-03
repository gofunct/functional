package driver

type Flag interface {
	Set(string) error
	String() string
	HasChanged() bool
	Name() string
	ValueString() string
	ValueType() string
}

