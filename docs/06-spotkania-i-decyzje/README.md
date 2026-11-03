[← Dokumentacja](../README.md)

# Spotkania i decyzje

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-11-04 · **Issue:** [Głosowanie nad decyzjami (#237)](https://github.com/happy-voucher/happy-voucher-repo/issues/237)

Tu jest rejestr decyzji zespołu D i notatki ze spotkań. Każda decyzja ma osobny plik w katalogu [decyzje/](decyzje/), a notatki są w [spotkania/](spotkania/README.md).

## Głosowanie nad decyzjami zespołowymi

Zamiast warsztatu głosowanie odbyło się asynchronicznie w issue [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237): reakcja 👍 albo 👎 pod każdą decyzją i komentarz z argumentem.

- **Data:** od 2026-10-25 do 2026-11-04
- **Uczestnicy:** 5 z 7 osób: Jakub Sikora (KP), Maciej Rożek (AN), Szymon Piechutowski (BE), Jakub Barej (DO), Alan Lubiński (TA). Bez głosu: Krystian Inoch (FE), Mikołaj Liszewski (TM)
- **Wynik:** każda decyzja ma 4 głosy za, czyli większość zespołu, więc wszystkie są przyjęte. Jedyny głos przeciw dotyczy ADR-004. Szczegóły i linki do komentarzy są w plikach decyzji

## Rejestr decyzji

| Nr | Zagadnienie | Wybór | Status | Data |
| :- | :- | :- | :- | :- |
| [ADR-001](decyzje/ADR-001-sposob-organizacji-pracy.md) | Sposób organizacji pracy | Scrum w sprintach dwutygodniowych, limit 2 zadań w toku na osobę | przyjęta | 2026-11-04 |
| [ADR-002](decyzje/ADR-002-zakres-pierwszej-wersji.md) | Zakres pierwszej wersji | Zakres wspólnego scenariusza, bez voucherów kwotowych, zwrotów i zamykania okresów | przyjęta | 2026-11-04 |
| [ADR-003](decyzje/ADR-003-priorytety-i-szacowanie.md) | Priorytety i szacowanie | P1 do P3, punkty Fibonacciego, planning poker. Priorytet ustala koordynator | przyjęta | 2026-11-04 |
| [ADR-004](decyzje/ADR-004-przeplyw-pracy-i-odbior.md) | Przepływ pracy i odbiór zadania | Przegląd PR przez inną osobę w 24 h, odbiór historyjki przez koordynatora | przyjęta | 2026-11-04 |
| [ADR-005](decyzje/ADR-005-podzial-odpowiedzialnosci.md) | Podział odpowiedzialności i zastępstwa | Opiekun i zastępca dla każdego modułu | przyjęta | 2026-11-04 |
| [ADR-006](decyzje/ADR-006-organizacja-dokumentacji.md) | Organizacja dokumentacji | Katalog `docs/` w repozytorium zamiast wiki | przyjęta | 2026-11-04 |
| [ADR-007](decyzje/ADR-007-zmiana-zakresu.md) | Zmiana zakresu i sytuacja problemowa | Praca na symulatorze partnera, nowe funkcje tylko w zamian za tę samą liczbę punktów | przyjęta | 2026-11-04 |

## Zasady rejestru

- Jedna decyzja to jeden plik `decyzje/ADR-NNN-temat.md`. Numery są kolejne i nie używamy ich ponownie.
- Każda decyzja porównuje co najmniej dwa warianty. Zapisuje argumenty, wybór z uzasadnieniem, rozbieżności albo wynik głosowania i sposób wdrożenia.
- Status zmienia się z „proponowana” na „przyjęta” po warsztacie albo po głosowaniu w issue. Przyjętej decyzji nie edytujemy. Zmianę zapisujemy jako nową decyzję, a w starej zmieniamy tylko status na „zastąpiona przez ADR-NNN”.
- **Imienny wkład:** każdy uczestnik dodaje w issue decyzji komentarz z argumentem, który wpłynął na wybór, i z własnym działaniem, które z niego wynika. Link do komentarza wpisujemy w tabeli „Uczestnicy i ich wkład” w pliku decyzji.
