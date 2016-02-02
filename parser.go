package caddyshack

import (
    "encoding/json"
    // "github.com/johnhof/adapters/couchdb"
)

func ParseSchema(s []byte) (map[string]CollectionSchema, error) {
    var cs map[string]CollectionSchema
    e := json.Unmarshal(s, &cs)
    return cs, e
}

func ParseSchemaStr(s string) (map[string]CollectionSchema, error) {
    var cs map[string]CollectionSchema
    e := json.Unmarshal([]byte(s), &cs)
    return cs, e
}

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
