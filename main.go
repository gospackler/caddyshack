package caddyshack

import (
	"errors"
	//	"fmt"

	"github.com/gospackler/caddyshack/model"
)

var (
	STORENOTFOUND = errors.New("Store needed for the model not loaded")
)

// Caddyshack is the core struct managing models and adapters
type Caddyshack struct {
	Caddies map[string]*Caddies
	Config  *Config
	//	Query   *query Add later
}

// New creates a caddyshack object
// Initialization function
func New() (cs *Caddyshack) {
	// Fixme May need to extend later
	cs = &Caddyshack{
		Config:  NewConfig(),
		Caddies: make(map[string]*Caddies),
	}
	return cs
}

func (cs *Caddyshack) GetCaddy(id string) (err error, caddy *Caddies) {
	caddy, status := cs.Caddies[id]
	if status == false {
		err = errors.New("Error while loading store : " + id)
	}
	return
}

// LoadStore loads a single store into the instance
func (cs *Caddyshack) LoadStore(store Store) error {
	status := cs.Config.AddStore(store)
	if status == false {
		return errors.New("Error while creating store")
	}
	return nil
}

func (cs *Caddyshack) AddModel(model *model.Definition) (err error) {

	status, storeIns := cs.Config.GetConfig(model.Adapter)
	if status == false {
		return STORENOTFOUND
	}
	err, newCaddy := NewCaddy(model, storeIns)
	if err == nil {
		cs.Caddies[newCaddy.Id] = newCaddy
	}
	return
}
