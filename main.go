package caddyshack

import (
    "errors"

    "github.com/bushwood/caddyshack/adapter"
    "github.com/bushwood/caddyshack/model"
    "github.com/bushwood/caddyshack/resource"
    log "github.com/Sirupsen/logrus"
)

type Caddyshack struct {
    Models map[string]model.Definition
    Adapters map[string]adapter.Definition
}

// New creates a caddyshack object
func New() (Caddyshack, error){
    cs := Caddyshack{}
    cs.Models = make(map[string]model.Definition)
    cs.Adapters = make(map[string]adapter.Definition)
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
    log.Debug(cs.Adapters[name])
    log.Debug("FUCK")
    return *cs, nil
}
