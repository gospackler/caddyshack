package caddyshack

import (
    "strings"
    "encoding/json"
    "io/ioutil"
    // "github.com/johnhof/adapters/couchdb"
)

// ParseModelDir parses a json model files froma  directory into a map
func ParseModelDir(dir string) (map[string]CollectionSchema, error) {
    mMap := make(map[string]CollectionSchema)
    files, _ := ioutil.ReadDir(dir)
    for _, f := range files {
        schema, fErr := ParseModelFile(dir + "/" + f.Name())
        if fErr != nil {
            return mMap, fErr
        }
        mName := strings.Replace(f.Name(), ".json", "", -1)
        mMap[mName] = schema
    }
    return mMap, nil
}

// ParseModelFile parses a json model file into an object
func ParseModelMapFile(path string) (map[string]CollectionSchema, error) {
    dat, fErr := ioutil.ReadFile(path)
    if fErr != nil {
        return map[string]CollectionSchema{}, fErr
    }
    return ParseModelMap(dat)
}

// ParseModelFile parses a json model file into an object
func ParseModelFile(path string) (CollectionSchema, error) {
    dat, fErr := ioutil.ReadFile(path)
    if fErr != nil {
        return CollectionSchema{}, fErr
    }

    return ParseModel(dat)
}

// ParseModelMap parse a json byte model map into a schema map
func ParseModelMap(s []byte) (map[string]CollectionSchema, error) {
    var cs map[string]CollectionSchema
    e := json.Unmarshal(s, &cs)
    return cs, e
}

// ParseModel parse a model json byte array into a schema
func ParseModel(m []byte) (CollectionSchema, error) {
    var cs CollectionSchema
    e := json.Unmarshal(m, &cs)
    return cs, e
}

// LoadSchema load a schema map and return a map of initialized collections
func LoadSchema(s map[string]CollectionSchema) (map[string]Collection) {
    var c map[string]Collection
    // create connection requirements
    return c
}

type CollectionSchema struct {
    Adapter string                       `json:adapter`
    Properties map[string]PropertySchema `json:properties`
}

type PropertySchema struct {
    Type string   `json:type`
    Required bool `json:required`
}
