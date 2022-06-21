package models

type CHargeJSON struct {
	Amount       int64  `json:"amount" bson:"amount"`
	ReceiptEmail string `json:"receipt_email" bson:"receipt_email"`
}
