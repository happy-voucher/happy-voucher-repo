// Trasy systemu D: API dla partnerów, API sklepu i API panelu.
package app

import (
	db
)

pipeline partner {
	use volt.RequestID
}

pipeline web {
	use volt.RequestID
	use Cors
}

// API dla partnerów A, B, C (api-d.md), klucz w Authorization: Bearer.
scope /api/v1 [pipe: partner, error_handler: ApiErrors, name: partner] {
	load partner CurrentPartner

	get    /vouchers/:code                                Partner.Verify
	post   /vouchers/:code/holds                          Partner.Hold     [body: HoldForm]
	delete /vouchers/:code/items/:item_id(int)/hold       Partner.Release
	post   /vouchers/:code/items/:item_id(int)/redemption Partner.Redeem   [body: RedeemForm]
	get    /settlements                                   Partner.Settlement
}

// API sklepu. Logowanie jest opcjonalne.
scope /api/shop [pipe: web, error_handler: ApiErrors, name: shop] {
	get  /offers                                       Shop.Offers
	get  /offers/:id(int)                              Shop.Offer
	post /orders                                       Shop.CreateOrder [body: OrderForm]
	post /orders/:id(int)/pay                          Shop.Pay         [body: PayForm]
	get  /my_vouchers                                  Shop.MyVouchers
	get  /vouchers/:code                               Shop.Voucher
	get  /vouchers/:code/items/:item_id(int)/slots     Shop.Slots
	post /vouchers/:code/items/:item_id(int)/booking   Shop.Book        [body: BookForm]
}

// API panelu operatora, tylko role operator i admin.
scope /api/admin [pipe: web, error_handler: ApiErrors, name: admin] {
	load staff CurrentStaff

	get   /partners                    Admin.Partners
	post  /partners                    Admin.CreatePartner [body: PartnerForm]
	post  /partners/:id(int)/sync      Admin.Sync
	get   /services                    db.PartnerServiceList

	get   /offers                      Admin.Offers
	post  /offers                      Admin.CreateOffer   [body: OfferForm]
	patch /offers/:id(int)             Admin.UpdateOffer   [body: OfferForm]
	post  /offers/:id(int)/publish     Admin.Publish
	post  /offers/:id(int)/withdraw    Admin.Withdraw

	get   /vouchers                    Admin.Vouchers
	get   /vouchers/:id(int)           Admin.Voucher
	post  /vouchers/:id(int)/cancel    Admin.CancelVoucher [body: ReasonForm]
	post  /bookings/:id(int)/cancel    Admin.CancelBooking [body: ReasonForm]
	post  /redemptions/:id(int)/revert Admin.Revert        [body: ReasonForm]

	get   /queue                       Admin.Queue
	get   /settlements                 Admin.Settlements
	get   /settlements/csv             Admin.SettlementsCsv
}
