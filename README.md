# caddyshack

Adapter based ORM in Golang

# DO NOT USE

Architecture not finalized && nothing is implemented...

## Example

Assume directory structure:

```
- main.go
- config/
--- resources.json
--- models/
----- user.json
----- note.json
```

#### ./config/resources.json

```json
{
  "redis": {
    "host": "127.0.0.1",
    "port": "6379"
  },
  "couchdb": {
    "host": "127.0.0.1",
    "port": "5984"
  }
}
```

#### ./config/models/user.json

```json
{
  "adapter": "couchdb",
  "properties": {
    "email": {
      "type": "string",
      "required": true
    },
    "password": {
      "type": "string",
      "required": true
    }
  }
}
```

#### ./config/models/note.json

```json
{
  "adapter": "couchdb",
  "properties": {
    "title": {
      "type": "string",
      "required": true
    },
    "contents": {
      "type": "string",
      "required": true
    }
  }
}
```

#### ./main.go

```go
package main

import (
    "github.com/bushwood/caddyshack"
  	couch "github.com/bushwood/caddyshack-couchdb"
)

func main() {
    rscs, _ :=  caddyshack.ParseRscFile("./resources.json")
    models, _ := caddyshack.ParseModelDir("./models")

    cs, _ := caddyshack.New()
    cs.LoadModels(models)
    cs.LoadAdapter(couchAdp.Adapter, rscs["couchdb"])
    // cs.Init() // open connection pools to loaded adapter databases

    SomeHandlerFunction(cs)
}

func SomeHandlerFunction (cs caddyshack.Caddyshack) {
  User := cs.Open("User") // retrieve from connection pool

  // ... some work ...

  User["findbyid"](Query{})
}

func Open (name string) map[string]func {
  m := Model.New(Adapter[name])

}

func New(a Adapter) map[string]func {
  m := make(map[string]func)
  m.Connect()
  m.set["findBy" + model.getkey] -> findByProperties(model.getKey, connectiobn)
  return m
}


func  findByProperties(k string, c Connection){
  return func (query) {

  }
}
```

### process

#### Init

- load models
- load adapters
- build collections
  - for each model, get the adapters
  - call `cs.Collections[NAME] = cs[ADAPTER_NAME].BuildCollection(NAME, MODEL)`
    - internally
      - somehow map models with a dynamic connection....



#### Usage

- open the model
  - `user := cs.Open("User")`

#### cs.Open("User")

```go
func (cs * caddyshack.Definition) Open(name string) (caddyshack.Result) {
    conn := cs.Collection[name].Connect()

}
```  

#### Approach 1

- pros
  - leaves pool/connection implementation up to adapter
  - allows injection of methods into collections
- cons
  - forces user to manage connection teardown
  -

Caddyshack

```go
func (cs *Caddyshack) Open(model string) (map[string]caddychack.Method, caddyshack.Connection, error) {
  adp := cs.Adapters[model]
  conn := adp.Connect() // get connection from adapters pool
  methods := adp.BuildMethods()
  return methods, conn, nil
}
```

Adapter

```go
func (adp *Adapter) BuildMethods(conn caddyshack.Connection, m caddyshack.Model) (map[string]caddychack.Method, error){
  methods: := make(map[string]caddychack.Method)
  for _, m := range model.Properites {
    // foreach property, use the connection in a closure to create methods and return them
  }
  return methods, nil
}
```

Handler Usage

```go
func SomeHandler(cs Caddyshack)  {
  user, conn, _ := cs.Open("user")
  user["FindById"](Query{})
  conn.close()
}
```
