# caddyshack

Adapter based ORM in Golang

# DO NOT USE

Architecture not finalized && nothing is implemented...

## TODO

- add `Find` function to caddyshack
  - `cs.Find(map[string]interface, string)`
  - takes the notation `foo.bar[0].biz`
  - iterates through object and performs casting
  - look into viper as a reference

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
    cs.Build()
    SomeHandlerFunction(cs)
}

func SomeHandlerFunction (cs caddyshack.Caddyshack) {
  u := cs.Open("user") // get the user collection

  usr = u.FindOne(Query{})

  // ... some work ...
}
```
