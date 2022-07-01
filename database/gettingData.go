package database

import (
	"L0/models"
	"database/sql"
)

func GetDelivery(db *sql.DB, order_id string) models.Delivery {
	rows, err := db.Query(`SELECT * FROM "deliveries" WHERE order_id ='` + order_id + `'`)
	var delivery models.Delivery
	CheckError(err)
	defer rows.Close()
	if rows.Next() {
		var Order_uid string
		err = rows.Scan(&delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region, &delivery.Email, &Order_uid)
		CheckError(err)
	}
	CheckError(err)
	return delivery
}

func GetItems(db *sql.DB, track_number string) []models.Item {
	rows, err := db.Query(`SELECT * FROM "items" WHERE track_number ='` + track_number + `'`)
	items := make([]models.Item, 1)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		var i models.Item
		err = rows.Scan(&i.Chrt_id, &i.Track_number, &i.Price, &i.Rid, &i.Name, &i.Sale, &i.Size, &i.Total_price, &i.Nm_id, &i.Brand, &i.Status)
		CheckError(err)
		items = append(items, i)
	}
	CheckError(err)
	return items
}

func GetPayment(db *sql.DB, order_uid string) models.Payment {
	rows, err := db.Query(`SELECT * FROM "payments" WHERE transaction ='` + order_uid + `'`)
	var payment models.Payment
	CheckError(err)
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&payment.Transaction, &payment.Currency, &payment.Provider, &payment.Amount, &payment.Payment_dt, &payment.Bank, &payment.Delivery_cost, &payment.Goods_total, &payment.Custom_fee, &payment.Request_id)
		CheckError(err)
	}
	CheckError(err)
	return payment
}

func GetOrder(db *sql.DB, order_uid string) models.Order {
	rows, err := db.Query(`SELECT * FROM "orders" WHERE order_uid ='` + order_uid + `'`)
	var order models.Order
	CheckError(err)
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&order.Order_uid, &order.Track_number, &order.Entry, &order.Locale, &order.Internal_signature, &order.Customer_id, &order.Delivery_service, &order.Sm_id, &order.Date_created, &order.Oof_shard, &order.Shardkey)
		CheckError(err)

	}
	order.Delivery = GetDelivery(db, order_uid)
	order.Items = GetItems(db, order.Track_number)
	order.Payment = GetPayment(db, order_uid)
	CheckError(err)
	return order
}

func GetOrders(db *sql.DB) []models.Order {
	rows, err := db.Query(`SELECT * FROM "orders" `)
	o := make([]models.Order, 1)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		var order models.Order
		err = rows.Scan(&order.Order_uid, &order.Track_number, &order.Entry, &order.Locale, &order.Internal_signature, &order.Customer_id, &order.Delivery_service, &order.Sm_id, &order.Date_created, &order.Oof_shard, &order.Shardkey)
		CheckError(err)
		order.Delivery = GetDelivery(db, order.Order_uid)
		order.Items = GetItems(db, order.Track_number)
		order.Payment = GetPayment(db, order.Order_uid)
		o = append(o, order)
	}
	CheckError(err)
	return o
}

func IsDelivery(db *sql.DB, order_id string) bool {
	rows, err := db.Query(`SELECT * FROM "deliveries" WHERE order_id ='` + order_id + `'`)
	var delivery models.Delivery
	CheckError(err)
	defer rows.Close()
	if rows.Next() {
		var Order_uid string
		err = rows.Scan(&delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region, &delivery.Email, &Order_uid)
		CheckError(err)
	}
	CheckError(err)
	var delivery2 models.Delivery
	return delivery == delivery2
}

func IsPayment(db *sql.DB, order_uid string) bool {
	rows, err := db.Query(`SELECT * FROM "payments" WHERE transaction ='` + order_uid + `'`)
	var payment models.Payment
	CheckError(err)
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&payment.Transaction, &payment.Currency, &payment.Provider, &payment.Amount, &payment.Payment_dt, &payment.Bank, &payment.Delivery_cost, &payment.Goods_total, &payment.Custom_fee, &payment.Request_id)
		CheckError(err)
	}
	CheckError(err)
	var payment2 models.Payment
	return payment == payment2
}

func IsItem(db *sql.DB, chrt_id int) bool {
	rows, err := db.Query(`SELECT * FROM "items" WHERE chrt_id =$1`, chrt_id)
	var item models.Item
	CheckError(err)
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&item.Chrt_id, &item.Track_number, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.Total_price, &item.Nm_id, &item.Brand, &item.Status)
		CheckError(err)
	}
	CheckError(err)
	var item2 models.Item
	return item == item2
}
