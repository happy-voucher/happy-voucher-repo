package db

import (
	"fmt"
	"slices"
)

type OfferType string

const (
	OfferTypeService OfferType = "service"
	OfferTypePackage OfferType = "package"
)

type OfferState string

const (
	OfferStateDraft     OfferState = "draft"
	OfferStateActive    OfferState = "active"
	OfferStateWithdrawn OfferState = "withdrawn"
)

type OrderState string

const (
	OrderStatePending OrderState = "pending"
	OrderStatePaid    OrderState = "paid"
)

type ItemState string

const (
	ItemStateAvailable ItemState = "available"
	ItemStateBooked    ItemState = "booked"
	ItemStateRedeemed  ItemState = "redeemed"
	ItemStateExpired   ItemState = "expired"
	ItemStateCancelled ItemState = "cancelled"
)

type BookingState string

const (
	BookingStateConfirmed BookingState = "confirmed"
	BookingStateUncertain BookingState = "uncertain"
	BookingStateFailed    BookingState = "failed"
	BookingStateEnded     BookingState = "ended"
)

type Role string

const (
	RoleOperator Role = "operator"
	RoleAdmin    Role = "admin"
)

func OfferTypeKnown(v OfferType) error {
	return known(v, OfferTypeService, OfferTypePackage)
}

func OfferStateKnown(v OfferState) error {
	return known(v, OfferStateDraft, OfferStateActive, OfferStateWithdrawn)
}

func OrderStateKnown(v OrderState) error {
	return known(v, OrderStatePending, OrderStatePaid)
}

func ItemStateKnown(v ItemState) error {
	return known(v, ItemStateAvailable, ItemStateBooked, ItemStateRedeemed, ItemStateExpired, ItemStateCancelled)
}

func BookingStateKnown(v BookingState) error {
	return known(v, BookingStateConfirmed, BookingStateUncertain, BookingStateFailed, BookingStateEnded)
}

func RoleKnown(v Role) error {
	return known(v, RoleOperator, RoleAdmin)
}

func known[T ~string](v T, set ...T) error {
	if slices.Contains(set, v) {
		return nil
	}
	return fmt.Errorf("%q nie jest jedną z wartości %v", string(v), set)
}
