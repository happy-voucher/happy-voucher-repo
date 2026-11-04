[← Spotkania i decyzje](../README.md)

# ADR-006 Organizacja dokumentacji

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-11-04 · **Issue:** [Głosowanie nad decyzjami (#237)](https://github.com/happy-voucher/happy-voucher-repo/issues/237)
>
> **Status:** przyjęta · **Data:** 2026-11-04, głosowanie asynchroniczne w issue zamiast warsztatu

## Kontekst

Zamiast Confluence używamy GitHuba. Musimy wybrać, gdzie trzymać dokumentację systemu: architekturę, API, wymagania, decyzje i testy.

## Warianty

| Wariant | Opis | Plusy | Minusy |
| :- | :- | :- | :- |
| A | Katalog `docs/` w repozytorium | Wersjonowanie razem z kodem, przegląd zmian przez PR | Osoby nietechniczne muszą znać PR |
| B | GitHub Wiki | Szybka edycja | Brak przeglądu zmian, łatwo o nieaktualne strony |

## Wybór i uzasadnienie

**Wariant A.** Dokumentacja API musi zmieniać się razem z kodem. W issue zapisujemy to, co dotyczy jednej zmiany: kryteria, dyskusję i przegląd. W `docs/` zapisujemy to, co opisuje system i obowiązuje dłużej. Aktualność zapewnia zmiana dokumentacji w tym samym PR co zmiana systemu.

## Rozbieżności lub wynik głosowania

Głosowanie asynchroniczne w issue [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237) od 25 października do 4 listopada 2026, reakcjami pod [komentarzem tej decyzji](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6401145942). Wynik: 4 za, 0 przeciw, 3 osoby bez głosu (Alan Lubiński, Krystian Inoch, Mikołaj Liszewski). Bez rozbieżności. Decyzja przyjęta, bo ma głosy za od 4 z 7 osób.

## Wdrożenie

- Struktura, zasady nazywania i aktualizacji w [indeksie dokumentacji](../../README.md).
- Opiekunowie obszarów w pliku [CODEOWNERS](../../../.github/CODEOWNERS).

## Uczestnicy i ich wkład

**Moderator:** Maciej Rożek (AN) · **Autor notatki:** Jakub Barej (DO), według podziału z [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237). Wyniki głosowania spisał Maciej Rożek (AN).

| Osoba | Rola | Głos | Argument, który wpłynął na wybór | Własne działanie | Link do komentarza |
| :- | :- | :- | :- | :- | :- |
| Jakub Sikora | KP | za | — | — | — |
| Maciej Rożek | AN | za | — | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6617612164) |
| Szymon Piechutowski | BE | za | — | — | — |
| Krystian Inoch | FE | brak głosu | — | — | — |
| Jakub Barej | DO | za | — | — | — |
| Alan Lubiński | TA | brak głosu | — | — | — |
| Mikołaj Liszewski | TM | brak głosu | — | — | — |
