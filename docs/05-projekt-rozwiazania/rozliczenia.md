[← Projekt rozwiązania](README.md)

# Rozliczenia z partnerami

> **Opiekun:** Szymon Piechutowski (BE) · **Autor:** Maciej Rożek (AN) · **Ostatnia zmiana:** 2026-11-23 · **Issue:** [D-22/T1 Reguły rozliczeń (#211)](https://github.com/happy-voucher/happy-voucher-repo/issues/211)
>
> **Status:** szkic do przeglądu przez BE. Zbiera reguły [ROZ](../03-wymagania/reguly-biznesowe.md#roz-rozliczenia) w jednym miejscu.

Strona opisuje, co wchodzi do rozliczenia partnera, jak wyznaczamy miesiąc i jak traktujemy cofnięte oraz wygasłe pozycje. API zestawienia jest w [api-d.md](api-d.md), a tabela `redemptions` w [model-danych.md](model-danych.md).

## Co wchodzi do rozliczenia

- Partnerowi należy się **kwota dla partnera za każdą zrealizowaną pozycję, której realizacji nie cofnięto** (ROZ-02).
- Kwota to `partner_amount` skopiowana z oferty przy wydaniu vouchera (POZ-02, OFE-07), liczba całkowita w groszach (ROZ-01).
- **Wygasłe, niezrealizowane pozycje nie są rozliczane** (ROZ-06).
- **Cofnięta realizacja nie wchodzi:** wykorzystanie zostaje w bazie z `reverted_at` i jest pomijane w sumie (REA-07, ROZ-02).

## Miesiąc rozliczenia

- O miesiącu decyduje **data realizacji w strefie Europe/Warsaw** (ROZ-03).
- Pole `redemptions.month` to `RRRR-MM` wyliczone z `redeemed_at` (zapisanego w UTC) po przeliczeniu na Europe/Warsaw.
- **Granica miesiąca:** realizacja 31.10 o 23:30 UTC to już 1.11 o 00:30 w Warszawie, więc wchodzi do listopada. Dlatego miesiąc liczymy zawsze z czasu lokalnego, nie z UTC. Ta sama zasada obowiązuje przy zmianie czasu letni/zimowy.
- Okresów nie zamykamy (ZAK-01), więc zestawienie danego miesiąca liczymy na bieżąco: cofnięcie realizacji po fakcie po prostu znika z sumy tego miesiąca.

## Zawartość zestawienia (ROZ-04)

- **Pozycja zestawienia:** kod vouchera, usługa, data realizacji, identyfikator rezerwacji partnera, kwota.
- **Raport operatora:** przychód, kwoty dla A, B i C oraz marża D (cena oferty minus suma kwot dla partnerów).

## Dostęp i spójność (ROZ-05)

- Partner pobiera przez API **tylko własne** zestawienie (`GET /api/v1/settlements?month=RRRR-MM`).
- Operator eksportuje zestawienia do CSV (`GET /api/admin/settlements/csv?month=RRRR-MM`).
- Sumy w API i w CSV muszą być równe.

## Jak to liczymy

Rozliczenie nie ma osobnej tabeli — to zapytanie po `redemptions` (patrz [model-danych.md](model-danych.md)):

- filtr `reverted_at IS NULL`, `partner_id` i `month`,
- suma `amount`,
- wspiera je indeks `(partner_id, month)`.

## Otwarte pytania

- **P-02:** czy wygasłe, niezrealizowane pozycje są rozliczane? Propozycja AN: nie (ROZ-06). Pytanie do A, B, C (D-22 T5).
- **P-14:** czy realizację można cofnąć? Propozycja AN: tak, tylko operator, a cofnięte wykorzystanie nie wchodzi do rozliczenia (REA-07).
