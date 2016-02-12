package collection

// Definition contains the base struct for the collection
type Definition interface {
	GetName() string
	Create() error
	Read() error
	ReadOne(string) error
	Update() error
	UpdateOne(string) error
	Destroy() error
	DestroyOne(string) error
}
