[← Strona główna](../README.md)

# Dokumentacja systemu D

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** —

To strona nadrzędna całej dokumentacji. Dokumentację trzymamy w repozytorium, a nie w wiki, żeby zmieniała się razem z kodem i przechodziła przegląd w pull requestach ([ADR-006](06-spotkania-i-decyzje/decyzje/ADR-006-organizacja-dokumentacji.md)). Bieżąca praca (historyjki, zadania, błędy, dyskusje) jest w GitHub Issues i na [tablicy „Main”](https://github.com/orgs/happy-voucher/projects). Tutaj jest to, co opisuje system i obowiązuje dłużej.

## Obszary

| Obszar | Co zawiera | Opiekun |
| :- | :- | :- |
| [01 Cel i zakres](01-cel-i-zakres.md) | Cel systemu D, aktorzy, zakres pierwszej wersji, co jest poza zakresem | Maciej Rożek (AN) |
| [02 Zespół i współpraca](02-zespol-i-wspolpraca.md) | Role, zastępstwa, kontakty do A, B, C, zasady pracy, definicja ukończenia | Jakub Sikora (KP) |
| [03 Wymagania](03-wymagania/README.md) | Słownik pojęć, reguły biznesowe voucherów, stany vouchera, odnośniki do epików | Maciej Rożek (AN) |
| [04 Plan prac](04-plan-prac/README.md) | Plan sprintów, cele sprintów i ich podsumowania | Jakub Sikora (KP) |
| [05 Projekt rozwiązania](05-projekt-rozwiazania/README.md) | Architektura, model danych, API D, kontrakt integracji z A, B, C | Szymon Piechutowski (BE) |
| [06 Spotkania i decyzje](06-spotkania-i-decyzje/README.md) | Notatki ze spotkań i rejestr decyzji zespołowych (ADR) | Jakub Sikora (KP) |
| [07 Testy i odbiór](07-testy-i-odbior/README.md) | Plan testów, przypadki testowe (konflikt, brak miejsc, niedostępność partnera), raporty | Mikołaj Liszewski (TM) |
| [08 Materiały](08-materialy/README.md) | Opis projektów od prowadzącego, prezentacje, załączniki | Jakub Barej (DO) |

## Drzewo stron

```
README.md                          strona główna repozytorium
docs/
  README.md                        ta strona
  01-cel-i-zakres.md
  02-zespol-i-wspolpraca.md
  03-wymagania/
    README.md
    slownik-pojec.md
    reguly-biznesowe.md
    stany-vouchera-i-pozycji.md
  04-plan-prac/
    README.md
    sprint-1.md
  05-projekt-rozwiazania/
    README.md
    model-danych.md
    api-d.md
    kontrakt-integracji.md
  06-spotkania-i-decyzje/
    README.md                      rejestr decyzji
    decyzje/                       ADR-001 do ADR-007
    spotkania/
      README.md
  07-testy-i-odbior/
    README.md
    plan-testow.md
    raporty/
      README.md
  08-materialy/
    README.md
    opis-projektow.md
```

**Relacje między stronami:** ta strona jest nadrzędna dla wszystkich obszarów i linkuje do każdego z nich. Każdy obszar złożony z kilku plików ma własny `README.md`, który jest stroną nadrzędną swoich podstron. Każda podstrona w pierwszym wierszu linkuje w górę do strony nadrzędnej.

## Zasady

**Nazwy plików:** małe litery, łączniki zamiast spacji, bez polskich znaków. Notatki ze spotkań nazywamy `RRRR-MM-DD-temat.md`, a decyzje `ADR-NNN-temat.md`.

**Nagłówek strony:** każda strona zaczyna się od linku do strony nadrzędnej, opiekuna, daty ostatniej zmiany i powiązanych issue.

**Wersjonowanie i aktualizacja:**
- Wersjonowanie zapewnia git. Każda zmiana trafia przez pull request zatwierdzony przez inną osobę.
- Zmiana API albo zachowania systemu wymaga zmiany dokumentacji w tym samym PR.
- Przyjętej decyzji nie edytujemy. Zmianę zapisujemy jako nową decyzję, która zastępuje poprzednią.
- Opiekun obszaru raz na sprint przegląda swoje strony.

**Uprawnienia:** wszyscy członkowie zespołu mają rolę Admin (Mikołaj Liszewski: Maintain). Repozytorium jest publiczne, więc prowadzący widzi dokumentację bez osobnego dostępu. Gałąź główna jest chroniona: zmiany wchodzą tylko przez PR z jednym zatwierdzeniem innej osoby, także od administratorów. Plik [CODEOWNERS](../.github/CODEOWNERS) przypisuje opiekunów do katalogów dokumentacji, więc opiekun jest automatycznie proszony o przegląd zmian w swoim obszarze.

**Powiązanie zadań z dokumentacją:** issue linkuje do strony, której dotyczy, a strona podaje w nagłówku numer issue. PR zmieniający dokumentację ma w opisie odnośnik do issue. Przykład: historyjka D-18 i [api-d.md](05-projekt-rozwiazania/api-d.md).
