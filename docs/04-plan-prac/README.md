[← Dokumentacja](../README.md)

# Plan prac

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** —

Pracujemy w dwutygodniowych sprintach ([ADR-001](../06-spotkania-i-decyzje/decyzje/ADR-001-sposob-organizacji-pracy.md)). Sprint to pole Iteration na [tablicy projektu](https://github.com/orgs/happy-voucher/projects/10/views/1). Każdy sprint ma tu własną stronę z celem, historyjkami i podsumowaniem.

## Plan sprintów

| Sprint | Cel | Historyjki | Punkty | Najbardziej obciążone osoby | Strona |
| :-: | :- | :- | :-: | :- | :- |
| 1 | Wiemy, co budujemy, mamy bazę, szkielet i symulator | D-01, D-02, D-03, D-04, D-05, D-07, D-08, D-09 | 34 | AN, BE, DO | [sprint-1.md](sprint-1.md) |
| 2 | Da się kupić voucher i dostać kod, a partner może zweryfikować kod | D-06, D-10, D-11, D-12, D-13, D-14, D-15, D-16 | 29 | BE, FE | — |
| 3 | Partner blokuje i realizuje, klient rezerwuje, sytuacje błędne są obsłużone | D-17, D-18, D-19, D-20, D-21 | 26 | BE, TA | — |
| 4 | Rozliczenia, panel, testy z partnerami i odbiór | D-22, D-23, D-24, D-25 | 23 | FE, TA, TM, KP | — |

Razem 112 punktów w 25 historyjkach.

W sprintach 1 i 2 tester manualny przygotowuje przypadki testowe i testuje przyrosty. Frontend w sprincie 1 przygotowuje makiety sklepu i szkielet panelu guigui, żeby w sprincie 2 od razu podpiąć API.

## Kolejność

Kolejność wynika z zależności. Nie da się napisać migracji bez modelu danych ani modelu bez wymagań. Nie da się też pisać API D bez kontraktu z partnerami. Dlatego sprint 1 to prawie wyłącznie analiza, kontrakt, baza i szkielet.

## Skala szacowania

- Punkty w ciągu Fibonacciego: 1, 2, 3, 5, 8, 13. Historyjka wzorcowa to D-14 Generowanie kodu i pozycji vouchera, za 3 punkty.
- Punkty oznaczają złożoność i niepewność, a nie godziny.
- Historyjkę za 13 punktów dzielimy przed wejściem do sprintu.
- Szacujemy w planning pokerze. Przy rozbieżności osoby ze skrajnymi ocenami uzasadniają swoje zdanie i głosujemy ponownie ([ADR-003](../06-spotkania-i-decyzje/decyzje/ADR-003-priorytety-i-szacowanie.md)).
