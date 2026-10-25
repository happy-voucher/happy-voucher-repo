[← Spotkania i decyzje](../README.md)

# ADR-006 Organizacja dokumentacji

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** [issue decyzji, link po założeniu]
>
> **Status:** proponowana, do zatwierdzenia na warsztacie · **Data warsztatu:** [uzupełnijcie]

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

[uzupełnijcie po warsztacie]

## Wdrożenie

- Struktura, zasady nazywania i aktualizacji w [indeksie dokumentacji](../../README.md).
- Opiekunowie obszarów w pliku [CODEOWNERS](../../../.github/CODEOWNERS).

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
