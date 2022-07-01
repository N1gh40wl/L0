package database

import (
	"L0/models"
	"database/sql"
)

func updateDelivery(delivery models.Delivery, db *sql.DB, order_id string) {
	update := `update "deliveries" set "name"=$1, "phone"=$2, "zip"=$3, "city"=$4, "address"=$5, "region"=$6, "email"=$7 where "order_id"=$8`
	_, err := db.Exec(update, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email, order_id)
	CheckError(err)
}

func updatePayment(payment models.Payment, db *sql.DB) {
	update := `update "payments" set "request_id"=$2, "currency"=$3, "provider"=$4, "amount"=$5, "payment_dt"=$6, "bank"=$7, "delivery_cost"=$8, "goods_total"=$9, "custom_fee"=$10 where "transaction"=$1`
	_, err := db.Exec(update, payment.Transaction, payment.Request_id, payment.Currency, payment.Provider, payment.Amount, payment.Payment_dt, payment.Bank, payment.Delivery_cost, payment.Goods_total, payment.Custom_fee)
	CheckError(err)
}

func updateItems(items []models.Item, db *sql.DB) {
	update := `update "items" set "track_number"=$2, "price"=$3, "rid"=$4, "name"=$5, "sale"=$6, "size"=$7, "total_price"=$8, "nm_id"=$9, "brand"=$10, "status"=$11 where "chrt_id"=$1 `
	for _, i := range items {
		_, err := db.Exec(update, i.Chrt_id, i.Track_number, i.Price, i.Rid, i.Name, i.Sale, i.Size, i.Total_price, i.Nm_id, i.Brand, i.Status)
		CheckError(err)
	}
}

func updateItem(item models.Item, db *sql.DB) {
	update := `update "items" set "track_number"=$2, "price"=$3, "rid"=$4, "name"=$5, "sale"=$6, "size"=$7, "total_price"=$8, "nm_id"=$9, "brand"=$10, "status"=$11 where "chrt_id"=$1 `
	_, err := db.Exec(update, item.Chrt_id, item.Track_number, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.Total_price, item.Nm_id, item.Brand, item.Status)
	CheckError(err)
}

func UpdateOrder(order models.Order, db *sql.DB) {
	update := `update "orders" set "track_number"=$2, "entry"=$3, "locale"=$4, "internal_signature"=$5, "customer_id"=$6, "delivery_service"=$7, "shardkey"=$8, "sm_id"=$9, "date_created"=$10, "oof_shard"=$11 where "order_uid"=$1`
	_, err := db.Exec(update, order.Order_uid, order.Track_number, order.Entry, order.Locale, order.Internal_signature, order.Customer_id, order.Delivery_service, order.Shardkey, order.Sm_id, order.Date_created, order.Oof_shard)
	CheckError(err)
	if !IsDelivery(db, order.Order_uid) {
		updateDelivery(order.Delivery, db, order.Order_uid)
	} else {
		insertDelivery(order.Delivery, db, order.Order_uid)
	}
	if !IsPayment(db, order.Order_uid) {
		updatePayment(order.Payment, db)
	} else {
		insertPayments(order.Payment, db)
	}
	for _, v := range order.Items {
		if !IsItem(db, v.Chrt_id) {
			updateItem(v, db)
		} else {
			insertItem(v, db)
		}
	}

}
