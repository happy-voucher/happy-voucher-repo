[← Spotkania i decyzje](../README.md)

# ADR-001 Sposób organizacji pracy

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** [issue decyzji, link po założeniu]
>
> **Status:** proponowana, do zatwierdzenia na warsztacie · **Data warsztatu:** [uzupełnijcie]

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

[uzupełnijcie po warsztacie]

## Wdrożenie

- Pole Iteration jako sprint na [tablicy projektu](https://github.com/orgs/happy-voucher/projects/10/views/1).
- Plan sprintów w [04-plan-prac](../../04-plan-prac/README.md).
- Limit 2 zadań w statusie „W trakcie” w [zasadach pracy](../../02-zespol-i-wspolpraca.md#zasady-pracy).

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
