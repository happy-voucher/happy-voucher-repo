[← Spotkania i decyzje](../README.md)

# ADR-003 Priorytety i szacowanie

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** [issue decyzji, link po założeniu]
>
> **Status:** proponowana, do zatwierdzenia na warsztacie · **Data warsztatu:** [uzupełnijcie]

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

[uzupełnijcie po warsztacie]

## Wdrożenie

- Pola Priorytet i Szacunek na [tablicy projektu](https://github.com/orgs/happy-voucher/projects/10/views/1).
- Skala szacowania i historyjka wzorcowa D-14 w [planie prac](../../04-plan-prac/README.md#skala-szacowania).

## Uczestnicy i ich wkład

**Moderator:** [uzupełnijcie] · **Autor notatki:** [uzupełnijcie]

| Osoba | Rola | Argument, który wpłynął na wybór | Własne działanie | Link do komentarza |
| :- | :- | :- | :- | :- |
| Jakub Sikora | KP | | | |
| Maciej Rożek | AN | | | |
| Szymon Piechutowski | BE | | | |
| Krystian Inoch | FE | | | |
| Jakub Barej | DO | | | |
| Alan Lubiński | TA | | | |
| Mikołaj Liszewski | TM | | | |
