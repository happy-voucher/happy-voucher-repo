[← Spotkania i decyzje](../README.md)

# ADR-007 Zmiana zakresu i sytuacja problemowa

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** [issue decyzji, link po założeniu]
>
> **Status:** proponowana, do zatwierdzenia na warsztacie · **Data warsztatu:** [uzupełnijcie]

## Kontekst

Sytuacja przykładowa: partner B ogłasza opóźnienie swojego API. Część naszych historyjek (rezerwacje, scenariusz wspólny) zależy od tego API. Osobno musimy ustalić, co robimy z wartościowymi, ale nieplanowanymi funkcjami.

## Warianty

| Wariant | Opis | Plusy | Minusy |
| :- | :- | :- | :- |
| A | Czekamy na partnera i przesuwamy zależne historyjki | Brak podwójnej pracy | Przestój |
| B | Pracujemy na symulatorze partnera zgodnym z uzgodnionym kontraktem | Brak przestoju, a zmiana po stronie partnera dotyczy tylko jednej warstwy | Koszt utrzymania symulatora |

## Wybór i uzasadnienie

**Wariant B.** Nowe funkcje trafiają do backlogu jako issue z uzasadnieniem i szacunkiem. Do trwającego sprintu wchodzą tylko w zamian za historyjki o tej samej liczbie punktów, a decyduje koordynator.

## Rozbieżności lub wynik głosowania

[uzupełnijcie po warsztacie]

## Wdrożenie

- Etykieta `czeka-na-partnera` dla zadań, które czekają na A, B albo C.
- Symulator partnerów `cmd/partnersim` (D-09) i adaptery partnerów za interfejsem `PartnerClient` (D-10), opisane w [projekcie rozwiązania](../../05-projekt-rozwiazania/README.md).
- Zasada wymiany za tę samą liczbę punktów w [celu i zakresie](../../01-cel-i-zakres.md#poza-zakresem).

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
