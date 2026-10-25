[← Dokumentacja](README.md)

# Cel i zakres

> **Opiekun:** Maciej Rożek (AN) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** D-01 Wymagania i reguły biznesowe voucherów, D-01 T5 decyzja o zakresie

## Cel

System D sprzedaje vouchery na usługi gabinetu masażu (A), studia fitness (B) i hotelu (C) oraz obsługuje je od zakupu do rozliczenia. Klient kupuje voucher na usługę albo pakiet i dostaje kod. Posiadacz vouchera rezerwuje terminy u partnerów. Partnerzy sprawdzają vouchery i potwierdzają ich wykorzystanie przez API D, a D przygotowuje zestawienia do rozliczeń.

Każda usługa z vouchera ma być wykorzystana dokładnie raz i tylko w terminie ważności.

## Aktorzy

Klient, posiadacz vouchera, operator D, partnerzy A, B i C oraz hotel C w roli koordynatora pakietu. Pełne opisy są w [słowniku pojęć](03-wymagania/slownik-pojec.md#aktorzy).

## Zakres pierwszej wersji

Zakres obejmuje wspólny scenariusz czterech zespołów, bo to on jest kryterium odbioru ([ADR-002](06-spotkania-i-decyzje/decyzje/ADR-002-zakres-pierwszej-wersji.md)).

| Obszar | Co wchodzi | Epik |
| :- | :- | :- |
| Analiza i kontrakt | Reguły biznesowe, model danych, kontrakt z A, B, C, specyfikacja API D | D-E0, D-E1 |
| Fundamenty | Baza SQLite z migracjami, CI, szkielet serwera, symulator partnerów | D-E2, D-E3 |
| Katalog i sprzedaż | Pobieranie usług partnerów, oferty na usługę i pakiet, katalog, zamówienie, symulowana płatność, kod vouchera, strona vouchera | D-E4, D-E5 |
| API dla partnerów | Weryfikacja kodu, blokada pozycji, realizacja, ochrona przed ponownym wykorzystaniem, wygasanie | D-E6 |
| Rezerwacje | Dostępność i rezerwacja u A, B, C, konflikt terminu, brak miejsc, niedostępność partnera | D-E7 |
| Rozliczenia i obsługa | Zestawienia wykorzystanych voucherów i kwot, panel obsługi w guigui | D-E8 |
| Odbiór | Automatyczny test scenariusza wspólnego, testy z prawdziwymi partnerami | D-E9 |

## Poza zakresem

- Vouchery kwotowe, zwroty i zamykanie okresów rozliczeniowych ([ADR-002](06-spotkania-i-decyzje/decyzje/ADR-002-zakres-pierwszej-wersji.md)).
- Faktury i rabaty (D-E5).
- Prawdziwe płatności: płatność jest symulowana (D-13).

Nowe funkcje trafiają do backlogu. Do trwającego sprintu wchodzą tylko w zamian za historyjki o tej samej liczbie punktów, a decyduje koordynator ([ADR-007](06-spotkania-i-decyzje/decyzje/ADR-007-zmiana-zakresu.md)).
