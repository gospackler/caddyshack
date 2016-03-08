package adapter

import (
	"github.com/georgethomas111/caddyshack/collection"
	"github.com/georgethomas111/caddyshack/model"
	"github.com/georgethomas111/caddyshack/resource"
)

// Definition specifies the adapter interface
type Definition interface {
	GetName() string
	SetName(string) error
	GetConfig() *resource.Definition
	SetConfig(*resource.Definition) error
	BuildCollection(*model.Definition) (collection.Definition, error)
}
