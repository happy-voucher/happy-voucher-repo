package app

import (
	"net/http"

	"github.com/Piechutowski/volt"

	"github.com/happy-voucher/happy-voucher-repo/db"
)

func todo(r volt.Request) volt.Response {
	return Err(r, http.StatusNotImplemented, "NOT_IMPLEMENTED", "jeszcze nie zrobione")
}

type Partner struct{}

func (Partner) Verify(r volt.Request, partner db.Partner, code string) volt.Response {
	return todo(r)
}

func (Partner) Hold(r volt.Request, partner db.Partner, code string, body HoldForm) volt.Response {
	return todo(r)
}

func (Partner) Release(r volt.Request, partner db.Partner, code string, itemID int64) volt.Response {
	return todo(r)
}

func (Partner) Redeem(r volt.Request, partner db.Partner, code string, itemID int64, body RedeemForm) volt.Response {
	return todo(r)
}

func (Partner) Settlement(r volt.Request, partner db.Partner) volt.Response {
	return todo(r)
}

type Shop struct{}

func (Shop) Offers(r volt.Request) volt.Response {
	return todo(r)
}

func (Shop) Offer(r volt.Request, id int64) volt.Response {
	return todo(r)
}

func (Shop) CreateOrder(r volt.Request, body OrderForm) volt.Response {
	return todo(r)
}

func (Shop) Pay(r volt.Request, id int64, body PayForm) volt.Response {
	return todo(r)
}

func (Shop) MyVouchers(r volt.Request) volt.Response {
	return todo(r)
}

func (Shop) Voucher(r volt.Request, code string) volt.Response {
	return todo(r)
}

func (Shop) Slots(r volt.Request, code string, itemID int64) volt.Response {
	return todo(r)
}

func (Shop) Book(r volt.Request, code string, itemID int64, body BookForm) volt.Response {
	return todo(r)
}

type Admin struct{}

func (Admin) Partners(r volt.Request, staff db.Staff) volt.Response {
	return todo(r)
}

func (Admin) CreatePartner(r volt.Request, staff db.Staff, body PartnerForm) volt.Response {
	return todo(r)
}

func (Admin) Sync(r volt.Request, staff db.Staff, id int64) volt.Response {
	return todo(r)
}

func (Admin) Offers(r volt.Request, staff db.Staff) volt.Response {
	return todo(r)
}

func (Admin) CreateOffer(r volt.Request, staff db.Staff, body OfferForm) volt.Response {
	return todo(r)
}

func (Admin) UpdateOffer(r volt.Request, staff db.Staff, id int64, body OfferForm) volt.Response {
	return todo(r)
}

func (Admin) Publish(r volt.Request, staff db.Staff, id int64) volt.Response {
	return todo(r)
}

func (Admin) Withdraw(r volt.Request, staff db.Staff, id int64) volt.Response {
	return todo(r)
}

func (Admin) Vouchers(r volt.Request, staff db.Staff) volt.Response {
	return todo(r)
}

func (Admin) Voucher(r volt.Request, staff db.Staff, id int64) volt.Response {
	return todo(r)
}

func (Admin) CancelVoucher(r volt.Request, staff db.Staff, id int64, body ReasonForm) volt.Response {
	return todo(r)
}

func (Admin) CancelBooking(r volt.Request, staff db.Staff, id int64, body ReasonForm) volt.Response {
	return todo(r)
}

func (Admin) Revert(r volt.Request, staff db.Staff, id int64, body ReasonForm) volt.Response {
	return todo(r)
}

func (Admin) Queue(r volt.Request, staff db.Staff) volt.Response {
	return todo(r)
}

func (Admin) Settlements(r volt.Request, staff db.Staff) volt.Response {
	return todo(r)
}

func (Admin) SettlementsCsv(r volt.Request, staff db.Staff) volt.Response {
	return todo(r)
}
