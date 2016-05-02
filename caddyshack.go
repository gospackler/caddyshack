package caddyshack

import (
	"github.com/gospackler/caddyshack/model"
)

type Caddies struct {
	Id       string
	StoreIns Store
}

func NewCaddy(model *model.Definition, store Store) (err error, caddy *Caddies) {

	err, StoreIns := store.Init(model)
	if err != nil {
		return
	}

	caddy = &Caddies{
		StoreIns: StoreIns,
		Id:       model.Name + store.GetName(),
	}

	return
}
