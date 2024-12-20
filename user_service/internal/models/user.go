package models

import "time"

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Login        string    `json:"login"`
	PaymentAcc   string    `json:"payment_acc"`
	ConnType     string    `json:"conn_type"` // fttx, adsl, gpon
	Balance      int64     `json:"balance"`
	WriteOffDate time.Time `json:"write_off_date"`
	Active       bool      `json:"active"`
	
	// TODO: plan_service
	PlanTitle   string `json:"plan_title"`
	PlanSeries  string `json:"plan_series"`
	PlanSubsFee uint32 `json:"plan_subs_fee"`

	LastPaymentSum  int64     `json:"last_payment_sum"`
	LastPaymentDate time.Time `json:"last_payment_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
