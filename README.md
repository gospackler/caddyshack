# caddyshack

Adapter based ORM in Golang

# DO NOT USE

Architecture not finalized && nothing is implemented...

## Example

Assume directory structure:

```
- main.go
- config/
--- models/
----- user.json
----- note.json
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
    models, _ := caddyshack.ParseModelDir("./models")
    collections, _ := caddyshack.LoadModels(models)
    caddyshack.Connect(&collections)
    collections["user"].Create(SomeDefinition{})

    // OR

    cs, _ := caddyshack.New()
    models, _ := caddyshack.ParseModelDir("./models")
    cs.LoadModels(models)
    cs.Connect()
    cs.Create("user", SomeDefinition{})

    // OR

    cs, _ := caddyshack.New()
    models, _ := caddyshack.ParseModelDir("./models")
    cs.LoadModels(models)
    cs.Connect()
    cs.Collections["user"].Create("user", SomeDefinition{})


}
```
