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

// New returns a new instance of query defintion
func New() Definition {
	var q Definition
	q.Where = make(map[string]interface{})
	q.Sort = make(map[string]interface{})
	return q
}
