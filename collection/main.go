package collection

import "github.com/bushwood/caddyshack/query"

// Definition contains the base struct for the collection
type Definition interface {
	GetName() string
	Create(interface{}) error
	Read(query.Definition) (interface{}, error)
	ReadOne(string) (interface{}, error)
	Update(interface{}) error
	UpdateOne(interface{}) error
	Destroy(interface{}) error
	DestroyOne(interface{}) error
}
