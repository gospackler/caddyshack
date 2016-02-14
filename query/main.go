package query

// Definition outlines a base
type Definition struct {
	Where map[string]interface{}
	Skip  int
	Limit int
	Sort  map[string]interface{}
}

// Sort defined the sorting structure
type Sort struct {
	Key  string
	Desc bool
}
