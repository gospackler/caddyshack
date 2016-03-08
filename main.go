package caddyshack

import (
	"errors"
	"fmt"

	"github.com/georgethomas111/caddyshack/adapter"
	"github.com/georgethomas111/caddyshack/collection"
	"github.com/georgethomas111/caddyshack/model"
	"github.com/georgethomas111/caddyshack/resource"
)

// Caddyshack is the core struct managing models and adapters
type Caddyshack struct {
	Models      map[string]*model.Definition
	Adapters    map[string]adapter.Definition    //interface values do not req pointers.
	Collections map[string]collection.Definition // also is an interface
}

// New creates a caddyshack object
// Initialization function
func New() (*Caddyshack, error) {
	cs := &Caddyshack{}
	cs.Models = make(map[string]*model.Definition)
	cs.Adapters = make(map[string]adapter.Definition)
	cs.Collections = make(map[string]collection.Definition)
	return cs, nil
}

// LoadModels loads a map of models into the instance
func (cs *Caddyshack) LoadModels(models map[string]*model.Definition) error {
	for _, v := range models {
		err := cs.LoadModel(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadModel loads a single model into the instance
func (cs *Caddyshack) LoadModel(model *model.Definition) error {
	if model.Name == "" {
		return errors.New("model .Name cannot be empty")
	}
	cs.Models[model.Name] = model
	return nil
}

// LoadAdapters loads a map of adapters into the instance
func (cs *Caddyshack) LoadAdapters(adps map[string]adapter.Definition, rscs map[string]*resource.Definition) error {
	for _, adp := range adps {
		name := adp.GetName()
		if name == "" {
			return errors.New("adapter .Name cannot be empty")
		}
		err := cs.LoadAdapter(adp, rscs[name])
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadAdapter loads a single adapter into the instance
func (cs *Caddyshack) LoadAdapter(adp adapter.Definition, rsc *resource.Definition) error {
	name := adp.GetName()
	if name == "" {
		return errors.New("adapter .Name cannot be empty")
	}
	cs.Adapters[name] = adp
	cs.Adapters[name].SetConfig(rsc)
	return nil
}

// Build generates collections based on the current state of the Caddyshack instance
func (cs *Caddyshack) Build() error {
	for _, m := range cs.Models {
		a := cs.Adapters[m.Adapter]
		if a == nil {
			return errors.New("model [" + m.Name + "] attempted to use unknown adapter [" + m.Adapter + "]")
		}

		err := cs.BuildCollection(m, a)
		if err != nil {
			return err
		}
	}
	return nil
}

// BuildCollection generates a single collection using the model and adapter provided
func (cs *Caddyshack) BuildCollection(m *model.Definition, a adapter.Definition) error {

	coll, err := a.BuildCollection(m)
	if err != nil {
		return err
	}

	cs.Collections[m.Name] = coll

	return nil
}

// Open a connection form the adapter connection pool to the specified collection
// FIXME The call is basically, Get Collection("collectionName") from the MAP
func (cs *Caddyshack) Open(collName string) (collection.Definition, error) {
	c := cs.Collections[collName]
	if c.GetName() == "" {
		return c, errors.New("attempted to open unknown collection [" + collName + "]")
	}

	return c, nil
}

// For debugging
func (cs *Caddyshack) String() string {
	str := "\n Models Map :" + fmt.Sprintf("%s", cs.Models)
	str = str + "\n Adapters Map :" + fmt.Sprintf("%s", cs.Adapters)
	str = str + "\n Collections Map :" + fmt.Sprintf("%s", cs.Collections)
	return str
}
