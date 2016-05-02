// This file contains the abstraction of the store entities.
package caddyshack

import (
	"github.com/gospackler/caddyshack/model"
)

// The condition can be independent for each of the Stores dealt with.
// couch db has one condition format for query
// Redis may have another.
// Queries can be of different types
type Query interface {
	Execute() (error, []StoreObject)
}

type StoreObject interface {
	GetKey() string
	SetKey(string)
}

// Store represents an abstraction(interface) to any store.
// Configuration varies and let's leave that at the jurisdiction of each Store.
type Store interface {
	GetName() string
	SetName(string) error
	Init(*model.Definition) (error, Store)
	Create(StoreObject) error
	ReadOne(string) (error, StoreObject)
	UpdateOne(StoreObject) error
	DestroyOne(string) error
	Read(Query) (error, []StoreObject)
	//Update(interface{}) error
	//Destroy(interface{}) error
}
