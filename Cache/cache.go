package cache

import (
	"L0/models"
	"sync"
)

type Cache struct {
	sync.RWMutex
	orders map[string]models.Order
}
