package sort

type API interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
