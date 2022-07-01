package routes

import (
	cache "L0/Cache"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDataByUid(cc *cache.Cache) gin.HandlerFunc {
	f := func(c *gin.Context) {
		uid := c.Param("uid")
		fmt.Println(c.Param("uid"))
		cacheOrder, status := cc.Get(uid)
		//order, _ := controllers.MarshalJson(cacheOrder)
		if status {
			c.JSON(http.StatusOK, cacheOrder)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Нет записей",
			})
		}

	}
	return gin.HandlerFunc(f)
}
