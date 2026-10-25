[← Spotkania i decyzje](../README.md)

# ADR-004 Przepływ pracy i odbiór zadania

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** [issue decyzji, link po założeniu]
>
> **Status:** proponowana, do zatwierdzenia na warsztacie · **Data warsztatu:** [uzupełnijcie]

## Kontekst

Musimy ustalić, kiedy zadanie jest zakończone i kto to sprawdza. GitHub nie blokuje przejść między statusami, więc kontrolę trzeba oprzeć na przeglądzie i ochronie gałęzi.

## Warianty

| Wariant | Opis | Plusy | Minusy |
| :- | :- | :- | :- |
| A | Przegląd PR przez inną osobę i odbiór historyjki przez koordynatora | Wiedza krąży po zespole, a koordynator pilnuje kryteriów | Kolejka przeglądów |
| B | Bez osobnego przeglądu, odbiór tylko przez testy | Szybciej | Brak kontroli kryteriów biznesowych |

## Wybór i uzasadnienie

**Wariant A z przeglądem w ciągu 24 godzin roboczych.** Spór o wymagania rozstrzyga koordynator, a spór techniczny opiekun modułu. Samodzielne zamykanie zadań przez wykonawców przyspieszyłoby pracę, ale zwiększyłoby ryzyko przeoczenia błędów.

## Rozbieżności lub wynik głosowania

[uzupełnijcie po warsztacie]

## Wdrożenie

- Ochrona gałęzi głównej: zmiany tylko przez PR z zatwierdzeniem innej osoby i zielonym CI (D-07).
- Statusy Do zrobienia, W trakcie, Do weryfikacji, Do poprawy, Zakończone na [tablicy projektu](https://github.com/orgs/happy-voucher/projects/10/views/1).
- Lista kontrolna w szablonie PR i [definicja ukończenia](../../02-zespol-i-wspolpraca.md#definicja-ukończenia).

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
