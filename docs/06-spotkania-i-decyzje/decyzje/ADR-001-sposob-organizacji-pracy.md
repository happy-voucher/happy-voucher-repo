[← Spotkania i decyzje](../README.md)

# ADR-001 Sposób organizacji pracy

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-11-07 · **Issue:** [Głosowanie nad decyzjami (#237)](https://github.com/happy-voucher/happy-voucher-repo/issues/237)
>
> **Status:** przyjęta · **Data:** 2026-11-04, głosowanie asynchroniczne w issue zamiast warsztatu

## Kontekst

Projekt D zależy od trzech innych zespołów (A, B, C), które budują swoje systemy równolegle. Potrzebujemy rytmu pracy, który daje stałe punkty synchronizacji z partnerami i terminy przeglądów z prowadzącym.

## Warianty

| Wariant | Opis | Plusy | Minusy |
| :- | :- | :- | :- |
| A | Scrum ze sprintami dwutygodniowymi | Stałe punkty synchronizacji z A, B, C, naturalne terminy przeglądów | Koszt spotkań |
| B | Kanban z limitami pracy w toku | Elastyczność, mniej spotkań | Brak wspólnego rytmu testów integracji |

## Wybór i uzasadnienie

**Wariant A z limitem 2 zadań w toku na osobę, przejętym z wariantu B.** Integracja z trzema zespołami wymaga wspólnego rytmu. Koniec sprintu to wspólny test integracji i termin przeglądu z prowadzącym. Limit pracy w toku chroni przed zaczynaniem wielu rzeczy naraz.

## Rozbieżności lub wynik głosowania

Głosowanie asynchroniczne w issue [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237) od 25 października do 4 listopada 2026, reakcjami pod [komentarzem tej decyzji](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6401144192). Wynik: 6 za, 0 przeciw, 1 osoba bez głosu (Alan Lubiński). Bez rozbieżności. Decyzja przyjęta, bo ma głosy za od 6 z 7 osób.

## Wdrożenie

- Pole Iteration jako sprint na [tablicy „Main”](https://github.com/orgs/happy-voucher/projects).
- Plan sprintów w [04-plan-prac](../../04-plan-prac/README.md).
- Limit 2 zadań w statusie „W trakcie” w [zasadach pracy](../../02-zespol-i-wspolpraca.md#zasady-pracy).

## Uczestnicy i ich wkład

**Moderator:** Jakub Sikora (KP) · **Autor notatki:** Maciej Rożek (AN), według podziału z [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237). Wyniki głosowania spisał Maciej Rożek (AN).

| Osoba | Rola | Głos | Argument, który wpłynął na wybór | Własne działanie | Link do komentarza |
| :- | :- | :- | :- | :- | :- |
| Jakub Sikora | KP | za | — (w komentarzu tylko „1.A”) | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6546876873) |
| Maciej Rożek | AN | za | — | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6617612164) |
| Szymon Piechutowski | BE | za | — | — | — |
| Krystian Inoch | FE | za | — | — | — |
| Jakub Barej | DO | za | — | — | — |
| Alan Lubiński | TA | brak głosu | — | — | — |
| Mikołaj Liszewski | TM | za | — | — | — |
