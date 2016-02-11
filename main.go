package caddyshack

import (
	"errors"

	"github.com/bushwood/caddyshack/adapter"
	"github.com/bushwood/caddyshack/collection"
	"github.com/bushwood/caddyshack/model"
	"github.com/bushwood/caddyshack/resource"
)

// Caddyshack is the core struct managing models and adapters
type Caddyshack struct {
	Models      map[string]model.Definition
	Adapters    map[string]adapter.Definition
	Collections map[string]collection.Definition
}

// New creates a caddyshack object
func New() (Caddyshack, error) {
	cs := Caddyshack{}
	cs.Models = make(map[string]model.Definition)
	cs.Adapters = make(map[string]adapter.Definition)
	cs.Collections = make(map[string]collection.Definition)
	return cs, nil
}

// LoadModels loads a map of models into the instance
func (cs *Caddyshack) LoadModels(models map[string]model.Definition) (Caddyshack, error) {
	for _, v := range models {
		_, err := cs.LoadModel(v)
		if err != nil {
			return *cs, err
		}
	}
	return *cs, nil
}

// LoadModel loads a single model into the instance
func (cs *Caddyshack) LoadModel(model model.Definition) (Caddyshack, error) {
	if model.Name == "" {
		return *cs, errors.New("model .Name cannot be empty")
	}
	cs.Models[model.Name] = model
	return *cs, nil
}

// LoadAdapters loads a map of adapters into the instance
func (cs *Caddyshack) LoadAdapters(adps map[string]adapter.Definition, rscs map[string]resource.Definition) (Caddyshack, error) {
	for _, adp := range adps {
		name := adp.GetName()
		if name == "" {
			return *cs, errors.New("adapter .Name cannot be empty")
		}
		_, err := cs.LoadAdapter(adp, rscs[name])
		if err != nil {
			return *cs, err
		}
	}
	return *cs, nil
}

// LoadAdapter loads a single adapter into the instance
func (cs *Caddyshack) LoadAdapter(adp adapter.Definition, rsc resource.Definition) (Caddyshack, error) {
	name := adp.GetName()
	if name == "" {
		return *cs, errors.New("adapter .Name cannot be empty")
	}
	cs.Adapters[name] = adp
	return *cs, nil
}

// Build generates collections based on the current state of the Caddyshack instance
func (cs *Caddyshack) Build() (Caddyshack, error) {
	for _, m := range cs.Models {
		a := cs.Adapters[m.Name]
		if a == nil {
			return *cs, errors.New("model [" + m.Name + "] attempted to use unknown adapter [" + m.Adapter + "]")
		}

		_, err := cs.BuildCollection(m, a)
		if err != nil {
			return *cs, err
		}
	}
	return *cs, nil
}

// BuildCollection generates a single collection using the model and adapter provided
func (cs *Caddyshack) BuildCollection(m model.Definition, a adapter.Definition) (Caddyshack, error) {
	cs.Collections[m.Name] = collection.Definition{}
	return *cs, nil
}

// Open a connection form the adapter connection pool to the specified collection
func (cs *Caddyshack) Open(collName string) (collection.Definition, error) {
	c := cs.Collections[collName]
	if c.Name == "" {
		return collection.Definition{}, errors.New("attempted to open unknown collection [" + collName + "]")
	}

	// c.Connect()
	return collection.Definition{}, nil
}
