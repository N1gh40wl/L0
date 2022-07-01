package controllers

import (
	"L0/models"
	"encoding/json"
)

func UnmarshalJson(message string) (models.Order, error) {
	var order models.Order
	err := json.Unmarshal([]byte(message), &order)
	if err != nil {
		return order, err
	}
	return order, nil
}

func MarshalJson(order models.Order) ([]byte, error) {
	data, err := json.Marshal(order)
	if err != nil {
		return data, err
	}
	return data, nil
}
