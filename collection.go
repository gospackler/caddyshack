package caddyshack

type Collection struct {
    Create func(doc interface{}) (interface{}, error)
    Find func(query interface{}) (interface{}, error)
    FindOne func(query interface{}) (interface{}, error)
    Update func(query interface{}, doc interface{}) (interface{}, error)
    Destroy func(query interface{}) (interface{}, error)
    Connect func() error
    Close func() error
}
