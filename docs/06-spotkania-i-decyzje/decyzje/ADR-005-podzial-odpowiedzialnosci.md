[← Spotkania i decyzje](../README.md)

# ADR-005 Podział odpowiedzialności i zastępstwa

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-11-04 · **Issue:** [Głosowanie nad decyzjami (#237)](https://github.com/happy-voucher/happy-voucher-repo/issues/237)
>
> **Status:** przyjęta · **Data:** 2026-11-04, głosowanie asynchroniczne w issue zamiast warsztatu

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

Głosowanie asynchroniczne w issue [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237) od 25 października do 4 listopada 2026, reakcjami pod [komentarzem tej decyzji](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6401145604). Wynik: 4 za, 0 przeciw, 3 osoby bez głosu (Alan Lubiński, Krystian Inoch, Mikołaj Liszewski). Bez rozbieżności. Decyzja przyjęta, bo ma głosy za od 4 z 7 osób.

## Wdrożenie

- Tabela ról i zastępstw w [Zespół i współpraca](../../02-zespol-i-wspolpraca.md#role-i-zastępstwa).
- Plik [CODEOWNERS](../../../.github/CODEOWNERS) przypisuje opiekunów do obszarów dokumentacji.

## Uczestnicy i ich wkład

**Moderator:** Jakub Barej (DO) · **Autor notatki:** Szymon Piechutowski (BE), według podziału z [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237). Wyniki głosowania spisał Maciej Rożek (AN).

| Osoba | Rola | Głos | Argument, który wpłynął na wybór | Własne działanie | Link do komentarza |
| :- | :- | :- | :- | :- | :- |
| Jakub Sikora | KP | za | — | — | — |
| Maciej Rożek | AN | za | — | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6617612164) |
| Szymon Piechutowski | BE | za | — | — | — |
| Krystian Inoch | FE | brak głosu | — | — | — |
| Jakub Barej | DO | za | — | — | — |
| Alan Lubiński | TA | brak głosu | — | — | — |
| Mikołaj Liszewski | TM | brak głosu | — | — | — |
