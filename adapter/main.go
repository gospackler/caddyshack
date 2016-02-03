package adapter

import (
    "github.com/bushwood/caddyshack/resource"
)

type Definition interface {
    SetName(string)
    GetName() (string)
    GetConfig() (resource.Definition)
    SetConfig(resource.Definition) (error)
}
