// Package caddyshack is an adapter based ORM with CRUD actions
package couchdb

import (
    "errors"
    "github.com/bushwood/caddyshack/adapters"
)

var Adapters map[string]Adapter = []Adapter{
    "CouchDB": adapters.CouchDB
}

type Resource struct {
    Host string     `json:"host"`
    Port string     `json:"port"`
    Username string `json:"username"`
    password string `json:"password"`
}

type Adapter struct {
    state interface{}
    func Connect(host string, username string, password, string)
    func Create(doc map[sring]interface{})
    func Read(query map[sring]interface{})
    func Update(query map[sring]interface{}, doc map[sring]interface{})
    func Destroy(query map[sring]interface{})
}
