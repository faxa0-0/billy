package models

import (
	"time"
)

type Plan struct {
	ID              int      `json:"id"`
	Title           string   `json:"title"`
	Series          string   `json:"series"`
	Corp            bool     `json:"corp"`
	SubscriptionFee int      `json:"subscription_fee"`
	ConnType        []string `json:"conn_type"`
	Speed           []int    `json:"speed"`
	TasixSpeed      []int    `json:"tasix_speed"`
	LimitMb         uint     `json:"limit_mb"`
	AdditionalInfo  struct {
		Ru  string `json:"ru"`
		En  string `json:"en"`
		Uzb string `json:"uzb"`
	} `json:"additional_info"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
