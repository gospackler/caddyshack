package caddyshack

import (
    "errors"
    "strings"
    "encoding/json"
    "io/ioutil"

    "github.com/bushwood/caddyshack/model"
    "github.com/bushwood/caddyshack/resource"
)

//
// Resource parsers
//

// ParseModelDir parses a json model files froma  directory into a map
func ParseRscDir(dir string) (map[string]resource.Definition, error) {
    rMap := make(map[string]resource.Definition)
    files, _ := ioutil.ReadDir(dir)
    for _, f := range files {
        model, fErr := ParseRscFile(dir + "/" + f.Name())
        if fErr != nil {
            return rMap, fErr
        }
        mName := strings.Replace(f.Name(), ".json", "", -1)
        rMap[mName] = model
    }
    return rMap, nil
}

// ParseRscFile parses a json resource file into an object
func ParseRscMapFile(path string) (map[string]resource.Definition, error) {
    dat, fErr := ioutil.ReadFile(path)
    if fErr != nil {
        return map[string]resource.Definition{}, fErr
    }

    return ParseRscMap(dat)
}

// ParseRscFile parses a json resource file into an object
func ParseRscFile(path string) (resource.Definition, error) {
    dat, fErr := ioutil.ReadFile(path)
    if fErr != nil {
        return resource.Definition{}, fErr
    }

    return ParseRsc(dat)
}

// ParseRscFile parses a json resource file into an object
func ParseRscMap(s []byte) (map[string]resource.Definition, error) {
    var rsc map[string]resource.Definition
    e := json.Unmarshal(s, &rsc)
    return rsc, e
}

// ParseRscFile parses a json resource file into an object
func ParseRsc(s []byte) (resource.Definition, error) {
    var rsc resource.Definition
    e := json.Unmarshal(s, &rsc)
    return rsc, e
}

//
// Model parsers
//

// ParseModelDir parses a json model files froma  directory into a map
func ParseModelDir(dir string) (map[string]model.Definition, error) {
    mMap := make(map[string]model.Definition)
    files, _ := ioutil.ReadDir(dir)
    for _, f := range files {
        model, fErr := ParseModelFile(dir + "/" + f.Name())
        mName := strings.Replace(f.Name(), ".json", "", -1)
        if fErr != nil {
            if model.Name == "" && model.Adapter != "" {
                model.Name = mName
            } else {
                return mMap, fErr
            }
        }
        mMap[mName] = model
    }
    return mMap, nil
}

// ParseModelMapFile parses a json model file into an object
func ParseModelMapFile(path string) (map[string]model.Definition, error) {
    dat, fErr := ioutil.ReadFile(path)
    if fErr != nil {
        return map[string]model.Definition{}, fErr
    }
    return ParseModelMap(dat)
}

// ParseModelFile parses a json model file into an object
func ParseModelFile(path string) (model.Definition, error) {
    dat, fErr := ioutil.ReadFile(path)
    if fErr != nil {
        return model.Definition{}, fErr
    }

    return ParseModel(dat)
}

// ParseModelMap parse a json byte model map into a model map
func ParseModelMap(s []byte) (map[string]model.Definition, error) {
    var models map[string]model.Definition
    e := json.Unmarshal(s, &models)
    for n, m := range models {
        if m.Name == "" {
            m.Name = n
        }
    }
    return models, e
}

// ParseModel parse a model json byte array into a model
func ParseModel(m []byte) (model.Definition, error) {
    var model model.Definition
    e := json.Unmarshal(m, &model)
    if e == nil && model.Name == "" {
        return model, errors.New("failed to parse model, name not set")
    }
    return model, e
}
