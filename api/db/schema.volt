// Schemat bazy systemu D (D-05). Kwoty w groszach (ROZ-01), czas w UTC.
package db

type OfferType  string
type OfferState string
type OrderState string
type ItemState  string
type BookingState string
type Role       string

// Partnerzy A, B, C (SYN-01). api_key_hash to SHA-256 klucza, którym
// partner wywołuje D (UPR-01). outbound_api_key to klucz, którym D
// wywołuje partnera.
table partners {
  id               int    [pk]
  code             string [unique]
  name             string
  api_url          string
  api_key_hash     []byte [unique]
  outbound_api_key string
}

// Usługi pobrane z API partnera (SYN-02, SYN-03). external_id to
// identyfikator u partnera. Usług nie usuwamy, tylko wyłączamy.
table partner_services {
  id          int    [pk]
  partner_id  int    [ref: partners.id]
  external_id string
  name        string
  price       int
  active      bool   [default: true]

  indexes {
    (partner_id, external_id) [unique]
  }
}

// Historia synchronizacji (SYN-04). Wpis z błędem i bez handled_at jest
// w kolejce spraw operatora.
table sync_log {
  id         int       [pk]
  partner_id int       [ref: partners.id]
  at         time.Time [default: `CURRENT_TIMESTAMP`]
  ok         bool
  error      string    [null]
  handled_at time.Time [null]
}

// Oferty (OFE-01 do OFE-05).
table offers {
  id                     int        [pk]
  name                   string
  description            string     [default: ""]
  kind                   OfferType
  status                 OfferState [default: "draft"]
  price                  int
  validity_days          int
  coordinator_partner_id int        [null, ref: partners.id]

  validators {
    OfferTypeKnown(kind)
    OfferStateKnown(status)
  }
}

// Pozycja oferty: usługa, liczba sztuk, kwota dla partnera za sztukę.
table offer_items {
  id             int [pk]
  offer_id       int [ref: offers.id]
  service_id     int [ref: partner_services.id]
  quantity       int
  partner_amount int
}

// Zamówienie (ZAM-01). Płatność jest symulowana, więc jej wynik to po
// prostu status zamówienia. Zamówienie opłacone drugi raz nie tworzy
// nowych voucherów (ZAM-06).
table orders {
  id             int        [pk]
  offer_id       int        [ref: offers.id]
  quantity       int
  unit_price     int
  email          string     [index]
  recipient_name string     [null]
  status         OrderState [default: "pending"]
  created_at     time.Time  [default: `CURRENT_TIMESTAMP`]
  paid_at        time.Time  [null]

  validators {
    OrderStateKnown(status)
  }
}

// Voucher (KOD-01, KOD-02, WAZ-01). valid_until to ostatni dzień
// ważności w strefie Europe/Warsaw, zapisany jako RRRR-MM-DD. Status
// vouchera nie jest zapisywany, wynika z pozycji (POZ-05).
table vouchers {
  id           int       [pk]
  code         string    [unique]
  order_id     int       [ref: orders.id]
  valid_until  string
  cancelled_at time.Time [null]
}

// Pozycja vouchera: jedna sztuka usługi do wykorzystania raz (KOD-04).
// Partner i kwota są kopiowane z oferty przy wydaniu (POZ-02, OFE-07).
table voucher_items {
  id             int       [pk]
  voucher_id     int       [ref: vouchers.id]
  partner_id     int       [ref: partners.id]
  service_id     int       [ref: partner_services.id]
  partner_amount int
  state          ItemState [default: "available"]

  validators {
    ItemStateKnown(state)
  }
}

// Rezerwacja u partnera: złożona przez D albo zgłoszona blokadą
// (REZ-03, BLK-03). Zapisujemy też nieudane próby (REZ-04). Na pozycję
// przypada najwyżej jedna trwająca rezerwacja (REZ-08).
table bookings {
  id                  int          [pk]
  item_id             int          [ref: voucher_items.id]
  by_partner          bool
  status              BookingState
  partner_booking_ref string       [null]
  starts_at           time.Time    [null]
  idempotency_key     string       [null, unique]
  created_at          time.Time    [default: `CURRENT_TIMESTAMP`]

  indexes {
    item_id [unique, name: "ux_bookings_one_live", where: `status IN ('confirmed', 'uncertain')`]
  }

  validators {
    BookingStateKnown(status)
  }
}

// Wykorzystanie (REA-02). Na pozycję przypada najwyżej jedno aktywne,
// czyli niecofnięte, wykorzystanie (POZ-01). month to RRRR-MM daty
// realizacji w strefie Europe/Warsaw (ROZ-03).
table redemptions {
  id                  int       [pk]
  item_id             int       [ref: voucher_items.id]
  partner_id          int       [ref: partners.id]
  partner_booking_ref string    [null]
  amount              int
  redeemed_at         time.Time [default: `CURRENT_TIMESTAMP`]
  month               string
  idempotency_key     string    [unique]
  reverted_at         time.Time [null]
  revert_reason       string    [null]

  indexes {
    item_id [unique, name: "ux_redemptions_one_active", where: `reverted_at IS NULL`]
    (partner_id, month)
  }
}

// Historia vouchera (POZ-06). actor to na przykład "partner:A",
// "operator:<uid>" albo "system".
table voucher_events {
  id         int       [pk]
  voucher_id int       [ref: vouchers.id]
  item_id    int       [null, ref: voucher_items.id]
  at         time.Time [default: `CURRENT_TIMESTAMP`]
  actor      string
  operation  string
  from_state string    [null]
  to_state   string    [null]
  note       string    [null]

  indexes {
    (voucher_id, at)
  }
}

// Konta panelu (UPR-02). Konto Firebase bez wiersza tutaj dostaje 403.
table staff {
  uid  string    [pk]
  role Role

  validators {
    RoleKnown(role)
  }
}
