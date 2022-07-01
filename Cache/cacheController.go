package cache

import (
	"L0/models"
)

func New() *Cache {

	orders := make(map[string]models.Order)

	cache := Cache{
		orders: orders,
	}

	return &cache
}

func (c *Cache) Set(key string, order models.Order) {

	c.Lock()

	defer c.Unlock()

	c.orders[key] = order

}

func (c *Cache) Get(key string) (models.Order, bool) {

	c.RLock()

	defer c.RUnlock()

	item, found := c.orders[key]

	if !found {
		return item, false
	}

	return item, true
}
