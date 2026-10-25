[← Spotkania i decyzje](../README.md)

# ADR-005 Podział odpowiedzialności i zastępstwa

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** [issue decyzji, link po założeniu]
>
> **Status:** proponowana, do zatwierdzenia na warsztacie · **Data warsztatu:** [uzupełnijcie]

## Kontekst

Zespół liczy siedem osób o różnych rolach, doświadczeniu i dostępności. Musimy wiedzieć, kto odpowiada za każdy moduł i kto go przejmie, gdy opiekun jest niedostępny.

## Warianty

| Wariant | Opis | Plusy | Minusy |
| :- | :- | :- | :- |
| A | Opiekun i zastępca dla każdego modułu | Jasna odpowiedzialność | Ryzyko, że wiedzę o module ma jedna osoba |
| B | Wszyscy robią wszystko | Elastyczność | Nikt nie odpowiada za spójność modułu |

## Wybór i uzasadnienie

**Wariant A.** Każde zadanie ma dokładnie jedną osobę odpowiedzialną. Żeby wiedza nie zostawała u jednej osoby, zastępca robi co najmniej jeden przegląd PR na sprint w module, za który zastępuje. Przed sprintem każdy deklaruje dostępność, a osoby z mniejszym doświadczeniem pracują w parze z opiekunem modułu albo biorą zadania z gotowym wzorcem.

## Rozbieżności lub wynik głosowania

[uzupełnijcie po warsztacie]

## Wdrożenie

- Tabela ról i zastępstw w [Zespół i współpraca](../../02-zespol-i-wspolpraca.md#role-i-zastępstwa).
- Plik [CODEOWNERS](../../../.github/CODEOWNERS) przypisuje opiekunów do obszarów dokumentacji.

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
