[← Spotkania i decyzje](../README.md)

# ADR-002 Zakres pierwszej wersji

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-11-04 · **Issue:** [Głosowanie nad decyzjami (#237)](https://github.com/happy-voucher/happy-voucher-repo/issues/237)
>
> **Status:** przyjęta · **Data:** 2026-11-04, głosowanie asynchroniczne w issue zamiast warsztatu

## Kontekst

Czas projektu jest ograniczony, a odbiór ocenia, czy wspólny scenariusz czterech zespołów działa od zakupu vouchera do rozliczenia.

## Warianty

| Wariant | Opis | Plusy | Minusy |
| :- | :- | :- | :- |
| A | Pełny zakres: także vouchery kwotowe, zwroty, zamykanie okresów rozliczeń | Kompletność | Ryzyko, że nic nie zadziała od początku do końca |
| B | Zakres wspólnego scenariusza: voucher na usługę i pakiet, rezerwacja, realizacja, API, zestawienie rozliczeń | Pokrywa oceniany scenariusz, mniejsze ryzyko | Mniej funkcji |

## Wybór i uzasadnienie

**Wariant B.** Scenariusz wspólny jest kryterium odbioru.

## Rozbieżności lub wynik głosowania

Głosowanie asynchroniczne w issue [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237) od 25 października do 4 listopada 2026, reakcjami pod [komentarzem tej decyzji](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6401144530). Wynik: 4 za, 0 przeciw, 3 osoby bez głosu (Alan Lubiński, Krystian Inoch, Mikołaj Liszewski). Jedna niejasność: komentarz Jakuba Sikory to „2.A”, a reakcja jest za proponowanym wariantem B. Liczymy reakcje, bo tak działa głosowanie w #237. Decyzja przyjęta, bo ma głosy za od 4 z 7 osób.

## Wdrożenie

- Priorytety P1 do P3 w backlogu.
- Zakres i lista funkcji poza zakresem w [01-cel-i-zakres.md](../../01-cel-i-zakres.md).
- Reguła ZAK-01 w [regułach biznesowych](../../03-wymagania/reguly-biznesowe.md#zak-poza-zakresem).

## Uczestnicy i ich wkład

**Moderator:** Maciej Rożek (AN) · **Autor notatki:** Jakub Sikora (KP), według podziału z [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237). Wyniki głosowania spisał Maciej Rożek (AN).

| Osoba | Rola | Głos | Argument, który wpłynął na wybór | Własne działanie | Link do komentarza |
| :- | :- | :- | :- | :- | :- |
| Jakub Sikora | KP | za | — (w komentarzu „2.A”, reakcja za wariantem B) | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6546876873) |
| Maciej Rożek | AN | za | Wspólny scenariusz jest kryterium odbioru, a pełny zakres grozi tym, że nic nie zadziała od początku do końca | Trzymanie wymagań i reguł biznesowych w [03-wymagania](../../03-wymagania/README.md) w granicach scenariusza wspólnego | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6617612164) |
| Szymon Piechutowski | BE | za | „Mniej znaczy więcej” | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6613125651) |
| Krystian Inoch | FE | brak głosu | — | — | — |
| Jakub Barej | DO | za | — | — | — |
| Alan Lubiński | TA | brak głosu | — | — | — |
| Mikołaj Liszewski | TM | brak głosu | — | — | — |
