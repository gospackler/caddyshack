package caddyshack

import (
	"errors"
	"strings"
)

// Representation of a query.

type StoreQuery struct {
	Condition string
	Store     *TextStore
}

func (q *StoreQuery) SetCondition(cond string) {
	q.Condition = cond
}

func (q *StoreQuery) GetCondition() string {
	return q.Condition
}

func (q *StoreQuery) Execute() (err error, objList []StoreObject) {

	queryStr := q.GetCondition()
	keys := strings.Split(queryStr, ":")
	if len(keys) == 0 {
		return nil, nil
	}

	for _, key := range keys {
		err, obj := q.Store.ReadOne(key)
		if err != nil {
			return errors.New(err.Error()), objList
		} else {
			objList = append(objList, obj)
		}
	}
	return
}
