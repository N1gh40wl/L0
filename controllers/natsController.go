package controllers

import (
	cache "L0/Cache"
	"L0/database"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/nats-io/stan.go"
)

func Sub(cache *cache.Cache, db *sql.DB) func(m *stan.Msg) {
	f := func(m *stan.Msg) {
		json := string(m.Data)
		order, err := UnmarshalJson(json)
		if err != nil {
			fmt.Println(err)
		} else {
			if _, state := cache.Get(order.Order_uid); !state {
				cache.Set(order.Order_uid, order)
				database.InsertOrders(order, db)
			} else {
				if reflect.DeepEqual(order, database.GetOrder(db, order.Order_uid)) {
					fmt.Println("duplicate key: ", order.Order_uid)
				} else {
					database.UpdateOrder(order, db)
				}

			}

		}

	}
	return f
}
