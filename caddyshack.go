package caddyshack

import (
	"errors"
	"github.com/georgethomas111/caddyshack/adapter"
	"github.com/georgethomas111/caddyshack/collection"
	"github.com/georgethomas111/caddyshack/model"
	"github.com/georgethomas111/caddyshack/query"
)

type CaddyshackNew struct {
	ModelsObj     *Models
	AdapterObj    *Adapters
	CollectionObj *Collection
	QueryObj      *Query // For doing the actual querying.
}

// Create methods for it to make it easy in future. TODO

func ParseConfigs(resourceJson string, modelJson string) (error, *CaddyshackNew) {
	// Load the file
	// Build the Json
	// Do the Parse.
	return nil, &CaddyshackNew{}

}

type Query struct {
	query.Definition
}

// Proxy it for the time being
type ModelDefinition struct {
	model.Definition
}

type Models struct {
	Models map[string]*ModelDefinition
}

func (mod *Models) GetModel(key string) (err error, model *ModelDefinition) {
	obj, exists := mod.Models[key]
	if exists {
		model = obj
	} else {
		err = errors.New("Model object with name " + key + " not available.")
	}
	return
}

// Proxy it for the time being
type AdapterDefinition struct {
	adapter.Definition
}

type Adapters struct {
	Adapters map[string]*AdapterDefinition
}

func (ada *Adapters) GetAdapter(key string) (err error, adapter *AdapterDefinition) {

	obj, exists := ada.Adapters[key]
	if exists {
		adapter = obj
	} else {
		err = errors.New("Adapter object with name " + key + " not available.")
	}
	return
}

type CollDefinition struct {
	collection.Definition
}

type Collection struct {
	Collections map[string]*CollDefinition
}
