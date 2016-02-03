package adapter

import (
    "github.com/bushwood/caddyshack/resource"
)

type Definition interface {
    GetName() (string)
    SetName(string) (error)
    GetConfig() (resource.Definition)
    SetConfig(resource.Definition) (error)
    Open() (error)
    Close() (error)
}
