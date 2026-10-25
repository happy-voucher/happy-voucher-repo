[← Plan prac](README.md)

# Sprint 1

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** historyjki w tabeli niżej

**Cel:** wiemy, co budujemy, mamy bazę, szkielet i symulator.

**Daty:** [uzupełnijcie: początek i koniec sprintu, termin przeglądu z prowadzącym]

## Historyjki

| Historyjka | Epik | Punkty | Prowadzi |
| :- | :- | :-: | :-: |
| [D-01 Wymagania i reguły biznesowe voucherów (#72)](https://github.com/happy-voucher/happy-voucher-repo/issues/72) | D-E0 | 5 | AN |
| [D-02 Model pojęciowy danych i przypadki testowe (#83)](https://github.com/happy-voucher/happy-voucher-repo/issues/83) | D-E0 | 3 | AN |
| [D-03 Uzgodnienie kontraktu z zespołami A, B, C (#84)](https://github.com/happy-voucher/happy-voucher-repo/issues/84) | D-E1 | 5 | KP |
| [D-04 Specyfikacja API D (#85)](https://github.com/happy-voucher/happy-voucher-repo/issues/85) | D-E1 | 3 | AN |
| [D-05 Schemat bazy danych i migracje (#86)](https://github.com/happy-voucher/happy-voucher-repo/issues/86) | D-E2 | 5 | BE |
| [D-07 Repozytorium, CI i uruchamianie (#87)](https://github.com/happy-voucher/happy-voucher-repo/issues/87) | D-E3 | 3 | DO |
| [D-08 Szkielet serwera, uwierzytelnianie i błędy (#88)](https://github.com/happy-voucher/happy-voucher-repo/issues/88) | D-E3 | 5 | BE |
| [D-09 Symulator API partnerów (#89)](https://github.com/happy-voucher/happy-voucher-repo/issues/89) | D-E3 | 5 | DO |
| **Razem** | | **34** | |

Równolegle tester manualny przygotowuje przypadki testowe (D-02 T3), a frontend makiety sklepu i panelu. Szkielet aplikacji Next.js z logowaniem przez Firebase ([D-26, #90](https://github.com/happy-voucher/happy-voucher-repo/issues/90)) może zacząć się po D-07 i D-08.

## Kiedy sprint jest zakończony

- Serwer startuje z pustą bazą po migracjach.
- Symulator A dostaje `401` bez klucza i `404` dla nieznanego kodu.
- Cztery zespoły mają podpisany kontrakt integracji.

## Podsumowanie

[uzupełnijcie po sprincie: co zrobiliśmy, czego nie zrobiliśmy i dlaczego, wnioski z retrospektywy, zmiany na sprint 2]
