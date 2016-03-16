package caddyshack

import (
	"github.com/bushwood/caddyshack/model"
	"testing"
)

// Create a compatable storeObject
type TestObj struct {
	Name  string
	Value string
}

func (t *TestObj) GetKey() string {
	return t.Name
}

func (t *TestObj) SetKey(name string) {
	t.Name = name
}

var Caddy *Caddies
var TextStoreObj *TextStore

func TestInit(t *testing.T) {

	// Add model definition in future to it.
	cs := New()

	// From storedemo.go
	textStore := NewTextStore()
	err := cs.LoadStore(textStore)

	if err != nil {
		t.Error("Error while loading a store.")
	}

	model := &model.Definition{
		Adapter: "text",
		Name:    "testModel",
	}
	err = cs.AddModel(model)
	if err != nil {
		t.Error("Error while building caddyshack", err)
	}

	testObj := &TestObj{
		Name:  "abcd",
		Value: "1234",
	}
	caddyName := model.Name + model.Adapter
	err, caddy := cs.GetCaddy(caddyName)
	if err != nil {
		t.Error("Error while retreiving caddy ", err)
	}

	err = caddy.StoreIns.Create(testObj)
	if err != nil {

		t.Error("Error creating object in the test Store")
	}

	err, obj := caddy.StoreIns.ReadOne(testObj.GetKey())
	if err != nil {
		t.Error("Error while retreiving object")
	}

	if obj.GetKey() != testObj.GetKey() {
		t.Error("Retreived wrong object")
	}
	Caddy = caddy
	TextStoreObj = textStore

}

func TestRead(t *testing.T) {
	query := &StoreQuery{Condition: "abcd:abcd", Store: TextStoreObj}
	err, objects := Caddy.StoreIns.Read(query)
	if err != nil {
		t.Error("Error while reading query ", query)
	} else {
		t.Log("Read", objects)
	}
}
