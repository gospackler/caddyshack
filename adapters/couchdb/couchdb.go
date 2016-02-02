package adapters

import (
    "time"

    "github.com/johnhof/caddyshack"
    "github.com/spf13/viper"
    "github.com/satori/go.uuid"
    couch "github.com/rhinoman/couchdb-go"
)

const DBNAME string = "_CADDYSHACK"

var CouchDB caddyshack.Adapter = caddyshack.Adapter{
    Connect: func(config caddyshack.Resource) (caddyshack.ORM, error) {
        h := config.Host
        p := atoi(config.Port)
        t := time.Duration(config.Timeout)
        uname := config.Username
        pass := config.Password
        conn, err := couch.NewConnection(h, p, t)
        auth := couch.BasicAuth{Username: uname, Password: pass}
        return NewCouchORM(conn.SelectDB(DBNAME, &auth))
    },
}

func NewCouchORM(db couch.Database) (caddyshack.ORM, error) {
    return caddyshack.ORM{
        Create: func(doc interface{}) (interface{}, error) {
            id := uuid.NewV4().String()
            db.Save(doc, id, "")
            return interface{}, nil
        },
        Find: func(query interface{}) (interface{}, error) {
            db.Read()
            return interface{}, nil
        },
        FindOne: func(query interface{}) (interface{}, error) {
            db.Read()
            return interface{}, nil
        },
        Update: func(query interface{}, doc interface{}) (interface{}, error){
            db.Save()
            return interface{}, nil
        },
        Destroy: func(query interface{}) (interface{}, error) {
            db.Delete()
            return interface{}, nil
        },
        Connect func() (error) {
            return nil
        }
        Close: func() (error) {
            return nil
        }
    }
}
