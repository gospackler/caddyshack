package adapters

import (
    "../../caddyshack"
    "github.com/patrickjuchli/couch"
    "github.com/johnhof/caddyshack/adapters"
    "github.com/spf13/viper"
)
var CouchDb caddyshack.Adapter = caddyshack.Adapter{
    state: interface{},
    Connect: func(host string, username string, password, string) {

    },
    Create: func(doc map[sring]interface{}) {

    },
    Read: func(query map[sring]interface{}) {

    },
    Update: func(query map[sring]interface{}, doc map[sring]interface{}) {

    },
    Destroy :func(query map[sring]interface{}) {

    },
}
