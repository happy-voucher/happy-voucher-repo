[← Projekt rozwiązania](README.md)

# API D dla partnerów

> **Opiekun:** Szymon Piechutowski (BE) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** [D-18 Rejestrowanie wykorzystania pozycji (#100)](https://github.com/happy-voucher/happy-voucher-repo/issues/100), [D-04 Specyfikacja API D (#85)](https://github.com/happy-voucher/happy-voucher-repo/issues/85)
>
> **Status:** szkic. Ścieżki i nazwy kodów błędów to propozycja do specyfikacji OpenAPI w `api/openapi.yaml` (D-04) i do kontraktu (D-03)

Przez to API partnerzy A, B i C sprawdzają vouchery, blokują pozycje, potwierdzają realizację i pobierają swoje zestawienia. Operacje, które D wywołuje u partnerów, opisuje [kontrakt integracji](kontrakt-integracji.md). Reguły, na których opiera się API, są w [regułach biznesowych](../03-wymagania/reguly-biznesowe.md).

## Zasady wspólne

| Zasada | Opis | Reguła |
| :- | :- | :- |
| Uwierzytelnianie | Każde żądanie ma nagłówek `Authorization: Bearer <klucz API partnera>`. Brak klucza albo błędny klucz daje `401` | UPR-01 |
| Format błędu | Każdy błąd ma kod z kontraktu, komunikat i identyfikator żądania (przykład niżej) | WSP-02 |
| Idempotencja | Żądania zmieniające stan wymagają nagłówka `Idempotency-Key`. Powtórzenie z tym samym kluczem zwraca tę samą odpowiedź | WSP-03 |
| Czas | RFC 3339, na przykład `2026-11-14T15:30:00+01:00` | WSP-01 |
| Kwoty | Liczby całkowite w groszach | ROZ-01 |

```json
{
  "error": {
    "code": "ITEM_ALREADY_REDEEMED",
    "message": "Ta pozycja vouchera została już zrealizowana.",
    "request_id": "req_01JC8W3Q2V"
  }
}
```

## Operacje

| Operacja | Metoda i ścieżka (propozycja) | Historyjka |
| :- | :- | :- |
| [Weryfikacja kodu](#weryfikacja-kodu) | `GET /api/v1/vouchers/{code}` | D-16 |
| [Blokada pozycji](#blokada-pozycji) | `POST /api/v1/vouchers/{code}/holds` | D-17 |
| [Zwolnienie blokady](#zwolnienie-blokady) | `DELETE /api/v1/vouchers/{code}/items/{item_id}/hold` | D-17 |
| [Realizacja pozycji](#realizacja-pozycji) | `POST /api/v1/vouchers/{code}/items/{item_id}/redemption` | D-18, D-19 |
| [Zestawienie partnera](#zestawienie-partnera) | `GET /api/v1/settlements?month=RRRR-MM` | D-22 |

### Weryfikacja kodu

Zwraca status vouchera, termin ważności i pozycje widoczne dla wywołującego partnera: jego własne, a dla koordynatora pakietu wszystkie pozycje pakietu (BLK-01).

| Odpowiedź | Kiedy |
| :- | :- |
| `200` | Kod istnieje, a voucher nie wygasł i nie jest anulowany |
| `404 VOUCHER_NOT_FOUND` | Nieznany kod |
| `410 VOUCHER_EXPIRED` | Termin ważności minął |
| `410 VOUCHER_CANCELLED` | Voucher anulowany |

### Blokada pozycji

Partner, który przyjął rezerwację z kodem vouchera, zgłasza ją w D. Podaje usługę i identyfikator swojej rezerwacji, a D wybiera wolną pozycję tej usługi i oznacza ją jako zarezerwowaną (BLK-03).

```json
{ "service_id": "a-masaz-relaks-60", "partner_booking_ref": "A-2026-001234" }
```

| Odpowiedź | Kiedy |
| :- | :- |
| `201` | Pozycja zablokowana. Odpowiedź zawiera `item_id` i stan `booked` |
| `409 NO_AVAILABLE_ITEM` | W voucherze nie ma wolnej pozycji tej usługi |
| `403 SERVICE_NOT_OWNED` | Usługa należy do innego partnera |
| `404`, `410` | Jak przy weryfikacji kodu |

Ponowienie z tym samym `Idempotency-Key` nie blokuje drugiej pozycji (BLK-06).

### Zwolnienie blokady

Przywraca pozycji stan `available` (BLK-05). Pozycja, która nie jest zarezerwowana, daje `409 ITEM_NOT_BOOKED`, a pozycja innego partnera `403 ITEM_NOT_OWNED`.

### Realizacja pozycji

Partner potwierdza po wykonaniu usługi, że pozycja została wykorzystana. Dzięki temu klient nie może jej użyć ponownie, a partner dostaje należną kwotę w rozliczeniu (D-18).

**Żądanie**

```http
POST /api/v1/vouchers/VD-7K3M-Q9TX-2HWP/items/itm_102/redemption
Authorization: Bearer <klucz API partnera A>
Idempotency-Key: 5f0c1c9e-2b7a-4d8e-9a51-0e6f3b2d7c11
Content-Type: application/json

{ "partner_booking_ref": "A-2026-001234" }
```

**Odpowiedź `201`**

```json
{
  "item_id": "itm_102",
  "state": "redeemed",
  "redeemed_at": "2026-11-14T15:30:00+01:00",
  "voucher_status": "partially_used",
  "remaining_items": [
    { "item_id": "itm_103", "service_name": "Trening personalny 60 min", "state": "booked" }
  ]
}
```

| Odpowiedź | Kiedy | Reguła |
| :- | :- | :- |
| `201` | Własna pozycja w stanie `available` albo `booked`. Pozycja przechodzi w `redeemed`, powstaje wykorzystanie i wpis w historii vouchera | REA-02 |
| `201`, ta sama odpowiedź | Ponowienie z tym samym `Idempotency-Key`, bez drugiego wpisu | REA-04 |
| `409 ITEM_ALREADY_REDEEMED` | Pozycja już zrealizowana. Stan bazy się nie zmienia | REA-03 |
| `403 ITEM_NOT_OWNED` | Pozycja należy do innego partnera | REA-03 |
| `410 VOUCHER_EXPIRED`, `410 VOUCHER_CANCELLED` | Voucher po terminie albo anulowany | REA-03, WAZ-03 |
| `404 VOUCHER_NOT_FOUND` | Nieznany kod | REA-03 |

Z 20 równoległych realizacji tej samej pozycji dokładnie jedna dostaje `201`, a 19 dostaje `409` (REA-05). Gdy w voucherze zostały inne pozycje, `voucher_status` to `partially_used`, a po ostatniej `used` (REA-06). Zrealizowana pozycja trafia do zestawienia partnera za miesiąc realizacji.

### Zestawienie partnera

Zwraca pozycje zrealizowane przez wywołującego partnera w podanym miesiącu i sumę kwot. O miesiącu decyduje data realizacji w strefie Europe/Warsaw (ROZ-03). Partner widzi tylko własne zestawienie, a suma jest równa sumie w eksporcie CSV operatora (ROZ-05).

| Pole pozycji | Opis |
| :- | :- |
| `code` | Kod vouchera |
| `service_name` | Usługa |
| `redeemed_at` | Data realizacji |
| `partner_booking_ref` | Identyfikator rezerwacji partnera |
| `amount` | Kwota dla partnera w groszach |
