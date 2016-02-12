package adapter

import (
	"github.com/bushwood/caddyshack/collection"
	"github.com/bushwood/caddyshack/model"
	"github.com/bushwood/caddyshack/resource"
)

// Definition specifies the adapter interface
type Definition interface {
	GetName() string
	SetName(string) error
	GetConfig() resource.Definition
	SetConfig(resource.Definition) error
	BuildCollection(model.Definition) (collection.Definition, error)
}
