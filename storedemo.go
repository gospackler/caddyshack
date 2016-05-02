package caddyshack

import (
	"errors"

	log "github.com/Sirupsen/logrus"
	"github.com/gospackler/caddyshack/model"
)

type TextStore struct {
	name       string
	model      *model.Definition
	dictionary map[string]StoreObject
}

func NewTextStore() (textStore *TextStore) {
	textStore = &TextStore{
		dictionary: make(map[string]StoreObject),
		name:       "text",
	}
	return
}

func (t *TextStore) Init(model *model.Definition) (error, Store) {
	t.model = model
	return nil, t
}

func (t *TextStore) GetName() string {

	return t.name
}

func (t *TextStore) SetName(name string) error {

	t.name = name
	return nil
}

// TODO : This method could be part of the interface in general which can be overridden
// Does it work that way ??
func (t *TextStore) verify(obj StoreObject) {

}

func (t *TextStore) Create(obj StoreObject) error {

	log.Debug("Got the request to save ", obj)
	log.Debug("TODO : Validate obj with the model dealt with")
	//private verify method exists which does not do anything right now. Use laterr
	t.dictionary[obj.GetKey()] = obj
	return nil
}

func (t *TextStore) ReadOne(key string) (err error, storeObj StoreObject) {

	storeObj, status := t.dictionary[key]
	if status == false {
		err = errors.New("Error while reading from the store")
	}
	return
}

func (t *TextStore) Read(query Query) (err error, objList []StoreObject) {
	return query.Execute()
}

func (t *TextStore) UpdateOne(obj StoreObject) (err error) {

	_, status := t.dictionary[obj.GetKey()]
	if status == false {
		err = errors.New("Object not found in the store")
	}

	t.dictionary[obj.GetKey()] = obj
	return
}

// Generally delets succeed
func (t *TextStore) DestroyOne(key string) error {
	delete(t.dictionary, key)
	return nil
}
