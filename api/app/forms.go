package app

type HoldForm struct {
	ServiceID         string `json:"service_id"`
	PartnerBookingRef string `json:"partner_booking_ref"`
}

type RedeemForm struct {
	PartnerBookingRef string `json:"partner_booking_ref"`
}

type OrderForm struct {
	OfferID       int64  `json:"offer_id"`
	Quantity      int64  `json:"quantity"`
	Email         string `json:"email"`
	RecipientName string `json:"recipient_name"`
}

type PayForm struct {
	Result string `json:"result"`
}

type BookForm struct {
	Email  string `json:"email"`
	SlotID string `json:"slot_id"`
}

type PartnerForm struct {
	Code           string `json:"code"`
	Name           string `json:"name"`
	APIURL         string `json:"api_url"`
	OutboundAPIKey string `json:"outbound_api_key"`
}

type OfferItemForm struct {
	ServiceID     int64 `json:"service_id"`
	Quantity      int64 `json:"quantity"`
	PartnerAmount int64 `json:"partner_amount"`
}

type OfferForm struct {
	Name                 string          `json:"name"`
	Description          string          `json:"description"`
	Kind                 string          `json:"kind"`
	Price                int64           `json:"price"`
	ValidityDays         int64           `json:"validity_days"`
	CoordinatorPartnerID int64           `json:"coordinator_partner_id"`
	Items                []OfferItemForm `json:"items"`
}

type ReasonForm struct {
	Reason string `json:"reason"`
}
