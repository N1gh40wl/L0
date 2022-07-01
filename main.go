package main

import (
	cache "L0/Cache"
	"L0/controllers"
	"L0/database"
	"L0/models"
	"L0/routes"
	"time"

	"github.com/gin-gonic/gin"
	stan "github.com/nats-io/stan.go"
)

func main() {
	preTime, _ := time.ParseDuration("20m")
	cache := cache.New() //кэш
	var del []models.Order
	db := database.Connection() //дб
	defer db.Close()
	del = database.GetOrders(db)
	for _, v := range del {
		cache.Set(v.Order_uid, v)
	}
	sc, _ := stan.Connect("test-cluster", "test")
	defer sc.Close()
	sc.Subscribe("foo", controllers.Sub(cache, db), stan.StartAtTimeDelta(preTime))

	r := gin.Default() // сервер
	r.GET("/get/:uid", routes.GetDataByUid(cache))
	r.Run()
}
