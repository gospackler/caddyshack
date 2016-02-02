package caddyshack

import (
    "strings"
    "encoding/json"
    "io/ioutil"
    // "github.com/johnhof/adapters/couchdb"
)

// ParseModelDir parses a json model files froma  directory into a map
func ParseModelDir(dir string) (map[string]CollectionModel, error) {
    mMap := make(map[string]CollectionModel)
    files, _ := ioutil.ReadDir(dir)
    for _, f := range files {
        model, fErr := ParseModelFile(dir + "/" + f.Name())
        if fErr != nil {
            return mMap, fErr
        }
        mName := strings.Replace(f.Name(), ".json", "", -1)
        mMap[mName] = model
    }
    return mMap, nil
}

// ParseModelFile parses a json model file into an object
func ParseModelMapFile(path string) (map[string]CollectionModel, error) {
    dat, fErr := ioutil.ReadFile(path)
    if fErr != nil {
        return map[string]CollectionModel{}, fErr
    }
    return ParseModelMap(dat)
}

// ParseModelFile parses a json model file into an object
func ParseModelFile(path string) (CollectionModel, error) {
    dat, fErr := ioutil.ReadFile(path)
    if fErr != nil {
        return CollectionModel{}, fErr
    }

    return ParseModel(dat)
}

// ParseModelMap parse a json byte model map into a model map
func ParseModelMap(s []byte) (map[string]CollectionModel, error) {
    var cs map[string]CollectionModel
    e := json.Unmarshal(s, &cs)
    return cs, e
}

// ParseModel parse a model json byte array into a model
func ParseModel(m []byte) (CollectionModel, error) {
    var cs CollectionModel
    e := json.Unmarshal(m, &cs)
    return cs, e
}

// LoadModel load a model map and return a map of initialized collections
func LoadModels(s map[string]CollectionModel) (map[string]Collection) {
    var c map[string]Collection
    // create connection requirements
    return c
}

// LoadModel load a model map and return a map of initialized collections
func LoadModel(s CollectionModel) (Collection) {
    var c map[string]Collection
    // create connection requirements
    return c
}

type CollectionModel struct {
    Adapter string                      `json:adapter`
    Properties map[string]PropertyModel `json:properties`
}

type PropertyModel struct {
    Type string   `json:type`
    Required bool `json:required`
}
