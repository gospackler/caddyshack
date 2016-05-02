# What is caddyshack ?

 CaddyShack is an adapter based framework which allows different stores to be added to it.

# Why should I use it ?

 Benefit of caddyshack is that shifting between datastores is about changing the store that is loaded and the working should be seamless.

# Architecture

Not completed. Still work in progress.
## Current Architecture

Caddyshack Suports Stores and Queries On Stores.

#### What is a Store ?
A store can be anything that can be used to save data. Any Object supporting the following methods can be a store. The way the object is created does not matter but the presence of the following functions is needed.
``` go
GetName() string
SetName(string) error
Init(*model.Definition) (error, Store) //
Create(StoreObject) error
ReadOne(string) (error, StoreObject)
UpdateOne(StoreObject) error
DestroyOne(string) error
Read(Query) (error, []StoreObject)
```

To use a store it needs to be loaded to caddyshack 

``` go
 cs := New()
// From storedemo.go
textStore := NewTextStore()
err := cs.LoadStore(textStore) // The name of the store will be matched to get the right store Instance.
```

To get a better understanding, there is storedemo.go which is the implementation of an in memory key value store. (store in a map and get it from the map) 
Please have a look. 

#### What are queries ?
The basic need of any data store is the ability to be able to get the data when needed. Each store has its on query mechanism and the Store should be able to support that. 
``` go
// The condition can be independent for each of the Stores dealt with.
// couch db has one condition format for query
// Redis may have another.
// Queries can be of different types
type Query interface {
    Execute() (error, []StoreObject)
}
```

The Query implemented for a store Need to have an execute function and gives a StoreObject. 

#### What is a StoreObject ?

StoreObject is any object that can dumped to caddyshack, it contains two functions. 

``` go
type StoreObject interface {
    GetKey() string
    SetKey(string)
}
```
Anything which has a key and can be dumped and retrieved from the database is a StoreObject.

#### What is a model and why should models be loaded ?

A model is a way to identify a paticular instance of caddyshack. Caddyshack can have any number of models and adapters to it. An adapter is the representation of a store and model is the representation of an object. Together we get an instance of an adapter for a model which is called a *Caddy*

For the demo text store, the name of the adapter
``` go 
model := &model.Definition{
        Adapter: "text",
        Name:    "testModel",
}
cs.AddModel(model)
caddyName := model.Name + model.Adapter
err, caddy := cs.GetCaddy(caddyName)
```
## Running the default example

This example makes use of the storedemo.go to get the work done. 

``` bash
git clone git@github.com:gospackler/caddyshack.git
cd caddyshack
go test -v
```

Have a look at storedemo.go to get an idea of how an in memory key - value store can be made using caddyshack. 

Couch db implementation of caddyshack can be found below.

https://github.com/gospackler/caddyshack-couchdb

## TODO

* Support Jobs to run on Stores, Using Queries and a store.
* An ideal case is any job being able to run on any store. 


#### Steps once more. 
* Have the interface in store.go implemented by the newStore we are planning to load.
* Load the newStore Object to caddyshack.
* Add the model we need to deal with also to caddyshack.
* Get the instance of the store w.r.t model.
* Use the storeInstance methods defined in the interface to play with it.

## Files

* store.go - abstraction of store
* caddyshack.go - abstraction of caddies/ individual instances for one model and a store.
* factory.go - abstraction of config
* jobs.go - jobs that can run on caddyshack. #Any type of tags that are possible for it.
