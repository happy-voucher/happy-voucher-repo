[← Spotkania i decyzje](../README.md)

# ADR-003 Priorytety i szacowanie

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-11-07 · **Issue:** [Głosowanie nad decyzjami (#237)](https://github.com/happy-voucher/happy-voucher-repo/issues/237)
>
> **Status:** przyjęta · **Data:** 2026-11-04, głosowanie asynchroniczne w issue zamiast warsztatu

## Kontekst

Musimy ustalić, kto decyduje o kolejności prac i jak oceniamy wielkość historyjek. Duża część pracy zależy od API partnerów, których jeszcze nie znamy.

## Warianty

| Wariant | Opis | Plusy | Minusy |
| :- | :- | :- | :- |
| A | Priorytety P1 do P3 i punkty w ciągu Fibonacciego, szacowane w planning pokerze | Punkty uwzględniają niepewność, a poker angażuje wszystkich | Na początku punkty są abstrakcyjne |
| B | Szacowanie w godzinach, a priorytet ustala sam koordynator | Intuicyjne | Godziny przy nieznanych API partnerów to zgadywanie |

## Wybór i uzasadnienie

**Wariant A.** Priorytet ustala koordynator, a szacunek cały zespół. Przy ustalaniu priorytetu koordynator bierze pod uwagę kolejno: czy funkcja jest potrzebna do scenariusza wspólnego, czy jej brak blokuje inny zespół i jaką ma wartość dla klienta. Głosowanie całego zespołu nad priorytetami odrzuciliśmy, bo przy remisie ktoś i tak musi zdecydować.

## Rozbieżności lub wynik głosowania

Głosowanie asynchroniczne w issue [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237) od 25 października do 4 listopada 2026, reakcjami pod [komentarzem tej decyzji](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6401144856). Wynik: 6 za, 0 przeciw, 1 osoba bez głosu (Alan Lubiński). Bez rozbieżności. Decyzja przyjęta, bo ma głosy za od 6 z 7 osób.

## Wdrożenie

- Pola Priorytet i Szacunek na [tablicy „Main”](https://github.com/orgs/happy-voucher/projects).
- Skala szacowania i historyjka wzorcowa D-14 w [planie prac](../../04-plan-prac/README.md#skala-szacowania).

## Uczestnicy i ich wkład

**Moderator:** Jakub Sikora (KP) · **Autor notatki:** Mikołaj Liszewski (TM), według podziału z [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237). Wyniki głosowania spisał Maciej Rożek (AN).

| Osoba | Rola | Głos | Argument, który wpłynął na wybór | Własne działanie | Link do komentarza |
| :- | :- | :- | :- | :- | :- |
| Jakub Sikora | KP | za | — (w komentarzu tylko „3.A”) | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6546876873) |
| Maciej Rożek | AN | za | — | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6617612164) |
| Szymon Piechutowski | BE | za | „Tak jak woda nie ma stałego kształtu, tak wojna nie ma stałych warunków” | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6613175555) |
| Krystian Inoch | FE | za | — | — | — |
| Jakub Barej | DO | za | — | — | — |
| Alan Lubiński | TA | brak głosu | — | — | — |
| Mikołaj Liszewski | TM | za | — | — | — |
