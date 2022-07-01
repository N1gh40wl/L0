package database

import (
	"L0/models"
	"database/sql"
)

func insertDelivery(delivery models.Delivery, db *sql.DB, order_id string) {

	insert := `insert into "deliveries"("name", "phone","zip", "city", "address", "region", "email", "order_id") values($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Exec(insert, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email, order_id)
	CheckError(err)
	//fmt.Println("Delivery insert")
}

func insertPayments(payment models.Payment, db *sql.DB) {
	insert := `insert into "payments"("transaction", "request_id","currency", "provider", "amount", "payment_dt", "bank", "delivery_cost","goods_total","custom_fee" ) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := db.Exec(insert, payment.Transaction, payment.Request_id, payment.Currency, payment.Provider, payment.Amount, payment.Payment_dt, payment.Bank, payment.Delivery_cost, payment.Goods_total, payment.Custom_fee)
	CheckError(err)
	//fmt.Println("Payment insert")
}

func insertItems(items []models.Item, db *sql.DB) {

	insert := `insert into "items"("chrt_id", "track_number","price", "rid", "name", "sale", "size", "total_price","nm_id","brand","status" ) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	for _, i := range items {

		_, err := db.Exec(insert, i.Chrt_id, i.Track_number, i.Price, i.Rid, i.Name, i.Sale, i.Size, i.Total_price, i.Nm_id, i.Brand, i.Status)
		CheckError(err)
		//fmt.Println("Item insert")
	}
}

func insertItem(item models.Item, db *sql.DB) {
	insert := `insert into "items"("chrt_id", "track_number","price", "rid", "name", "sale", "size", "total_price","nm_id","brand","status" ) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := db.Exec(insert, item.Chrt_id, item.Track_number, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.Total_price, item.Nm_id, item.Brand, item.Status)
	CheckError(err)
	//fmt.Println("Item insert")

}

func InsertOrders(order models.Order, db *sql.DB) {
	insert := `insert into "orders"("order_uid", "track_number","entry", "locale", "internal_signature", "customer_id", "delivery_service", "shardkey", "sm_id", "date_created", "oof_shard") values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := db.Exec(insert, order.Order_uid, order.Track_number, order.Entry, order.Locale, order.Internal_signature, order.Customer_id, order.Delivery_service, order.Shardkey, order.Sm_id, order.Date_created, order.Oof_shard)
	CheckError(err)
	insertDelivery(order.Delivery, db, order.Order_uid)
	insertPayments(order.Payment, db)
	insertItems(order.Items, db)
	//fmt.Println("Order insert")
}
