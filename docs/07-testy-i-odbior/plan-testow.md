[← Testy i odbiór](README.md)

# Plan testów

> **Opiekun:** Mikołaj Liszewski (TM) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** D-02 T3 lista przypadków testowych (TM), D-02 T4 oznaczenie przypadków automatycznych (TA)
>
> **Status:** szkielet. Przypadki opracowuje TM w zadaniu D-02 T3

## Rodzaje testów

| Rodzaj | Kto | Gdzie i kiedy | Historyjka |
| :- | :- | :- | :- |
| Jednostkowe i integracyjne w Go | TA | `go test ./...` w CI przy każdym PR | D-07 |
| Testy API na podstawie specyfikacji | TA | Na serwerze z symulatorem partnerów | D-04 T3 |
| Automatyczny test scenariusza wspólnego | TA | Na symulatorze w CI, a po podaniu adresów także na prawdziwych API | D-24 |
| Testy ręczne sklepu i panelu | TM | Na danych przykładowych z `cmd/seed` | D-06, D-12, D-15, D-23 |
| Testy z prawdziwymi partnerami i odbiór | TM | Wspólne sesje z zespołami A, B, C | D-25 |

## Środowisko

- Symulator partnerów `cmd/partnersim` (D-09) udaje A, B i C. Tryb wybiera się flagą: `ok`, `conflict`, `full`, `down`, `slow`.
- Dane przykładowe ładuje `cmd/seed` (D-06): trzech partnerów, kilka usług, dwie oferty i trzy vouchery w różnych stanach.
- Oczekiwane zachowanie opisują [reguły biznesowe](../03-wymagania/reguly-biznesowe.md) i [skutki operacji w każdym stanie](../03-wymagania/stany-vouchera-i-pozycji.md#skutek-operacji-w-każdym-stanie).

## Przypadki testowe

Kolumnę „Rodzaj” uzupełnia TA w zadaniu D-02 T4 (automatyczny albo ręczny).

| Nr | Przypadek | Reguły | Rodzaj | Stan |
| :- | :- | :- | :- | :- |
| TC-01 | Przebieg poprawny: zakup pakietu, rezerwacja, realizacja, zestawienie | ZAM-05, REZ-03, REZ-10, REA-02, ROZ-02 | | do opracowania |
| TC-02 | Konflikt rezerwacji | REZ-05 | | do opracowania |
| TC-03 | Brak miejsc | REZ-05 | | do opracowania |
| TC-04 | Niedostępność partnera | REZ-06, REZ-07, REZ-09, ZAM-08 | | do opracowania |
| TC-05 | Podwójna realizacja, także równoległa | REA-03, REA-04, REA-05 | | do opracowania |
| TC-06 | Voucher po terminie | WAZ-03, WAZ-04, REA-03 | | do opracowania |

Każdy przypadek opisujemy krokami, danymi wejściowymi i wynikiem oczekiwanym.

## Zgłaszanie błędów

- Błąd zgłaszamy szablonem issue „Błąd”, z krokami odtworzenia, wynikiem oczekiwanym i rzeczywistym.
- Priorytet P0 dostają tylko błędy, które blokują integrację albo demo.
- Przypadek z planu, który wykrył błąd, dostaje w raporcie link do tego błędu.

## Warunki odbioru

- Scenariusz wspólny jest wykonany ręcznie z prawdziwymi systemami A, B i C, a wynik jest w raporcie.
- Każdy przypadek z tego planu ma wynik i link do błędu, jeśli błąd wystąpił.
- Wszystkie błędy P0 i P1 są naprawione przed odbiorem.
- Raport końcowy jest w [raporty/](raporty/README.md).
