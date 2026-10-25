[← Dokumentacja](../README.md)

# Spotkania i decyzje

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** —

Tu jest rejestr decyzji zespołu D i notatki ze spotkań. Każda decyzja ma osobny plik w katalogu [decyzje/](decyzje/), a notatki są w [spotkania/](spotkania/README.md).

## Warsztat decyzji zespołowych

- **Data:** [uzupełnijcie]
- **Uczestnicy:** [uzupełnijcie]
- **Notatka:** [dodajcie po warsztacie jako `spotkania/RRRR-MM-DD-warsztat-decyzji.md`]

## Rejestr decyzji

| Nr | Zagadnienie | Wybór | Status | Data |
| :- | :- | :- | :- | :- |
| [ADR-001](decyzje/ADR-001-sposob-organizacji-pracy.md) | Sposób organizacji pracy | Scrum w sprintach dwutygodniowych, limit 2 zadań w toku na osobę | proponowana | [ ] |
| [ADR-002](decyzje/ADR-002-zakres-pierwszej-wersji.md) | Zakres pierwszej wersji | Zakres wspólnego scenariusza, bez voucherów kwotowych, zwrotów i zamykania okresów | proponowana | [ ] |
| [ADR-003](decyzje/ADR-003-priorytety-i-szacowanie.md) | Priorytety i szacowanie | P1 do P3, punkty Fibonacciego, planning poker. Priorytet ustala koordynator | proponowana | [ ] |
| [ADR-004](decyzje/ADR-004-przeplyw-pracy-i-odbior.md) | Przepływ pracy i odbiór zadania | Przegląd PR przez inną osobę w 24 h, odbiór historyjki przez koordynatora | proponowana | [ ] |
| [ADR-005](decyzje/ADR-005-podzial-odpowiedzialnosci.md) | Podział odpowiedzialności i zastępstwa | Opiekun i zastępca dla każdego modułu | proponowana | [ ] |
| [ADR-006](decyzje/ADR-006-organizacja-dokumentacji.md) | Organizacja dokumentacji | Katalog `docs/` w repozytorium zamiast wiki | proponowana | [ ] |
| [ADR-007](decyzje/ADR-007-zmiana-zakresu.md) | Zmiana zakresu i sytuacja problemowa | Praca na symulatorze partnera, nowe funkcje tylko w zamian za tę samą liczbę punktów | proponowana | [ ] |

## Zasady rejestru

- Jedna decyzja to jeden plik `decyzje/ADR-NNN-temat.md`. Numery są kolejne i nie używamy ich ponownie.
- Każda decyzja porównuje co najmniej dwa warianty. Zapisuje argumenty, wybór z uzasadnieniem, rozbieżności albo wynik głosowania i sposób wdrożenia.
- Status zmienia się z „proponowana” na „przyjęta” po warsztacie. Przyjętej decyzji nie edytujemy. Zmianę zapisujemy jako nową decyzję, a w starej zmieniamy tylko status na „zastąpiona przez ADR-NNN”.
- **Imienny wkład:** każdy uczestnik dodaje w issue decyzji komentarz z argumentem, który wpłynął na wybór, i z własnym działaniem, które z niego wynika. Link do komentarza wpisujemy w tabeli „Uczestnicy i ich wkład” w pliku decyzji.
