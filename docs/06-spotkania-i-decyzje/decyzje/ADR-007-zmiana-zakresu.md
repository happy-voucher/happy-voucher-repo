[← Spotkania i decyzje](../README.md)

# ADR-007 Zmiana zakresu i sytuacja problemowa

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-11-07 · **Issue:** [Głosowanie nad decyzjami (#237)](https://github.com/happy-voucher/happy-voucher-repo/issues/237)
>
> **Status:** przyjęta · **Data:** 2026-11-04, głosowanie asynchroniczne w issue zamiast warsztatu

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

Głosowanie asynchroniczne w issue [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237) od 25 października do 4 listopada 2026, reakcjami pod [komentarzem tej decyzji](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6401146510). Wynik: 6 za, 0 przeciw, 1 osoba bez głosu (Szymon Piechutowski). Bez rozbieżności. Pytanie Szymona Piechutowskiego o symulator ma odpowiedź w [komentarzu](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6617612164): chodzi o symulator partnerów `cmd/partnersim` z D-09. Decyzja przyjęta, bo ma głosy za od 6 z 7 osób.

## Wdrożenie

- Etykieta `czeka-na-partnera` dla zadań, które czekają na A, B albo C.
- Symulator partnerów `cmd/partnersim` (D-09) i adaptery partnerów za interfejsem `PartnerClient` (D-10), opisane w [projekcie rozwiązania](../../05-projekt-rozwiazania/README.md).
- Zasada wymiany za tę samą liczbę punktów w [celu i zakresie](../../01-cel-i-zakres.md#poza-zakresem).

## Uczestnicy i ich wkład

**Moderator:** Szymon Piechutowski (BE) · **Autor notatki:** Alan Lubiński (TA), według podziału z [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237). Wyniki głosowania spisał Maciej Rożek (AN).

| Osoba | Rola | Głos | Argument, który wpłynął na wybór | Własne działanie | Link do komentarza |
| :- | :- | :- | :- | :- | :- |
| Jakub Sikora | KP | za | — | — | — |
| Maciej Rożek | AN | za | Symulator zgodny z kontraktem pozwala pracować i testować w CI, zanim partnerzy udostępnią API, więc opóźnienie partnera nie zatrzymuje pracy | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6617612164) |
| Szymon Piechutowski | BE | brak głosu | — (pytanie „Jaki symulator?”) | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6613277944) |
| Krystian Inoch | FE | za | — | — | — |
| Jakub Barej | DO | za | — | — | — |
| Alan Lubiński | TA | za | — | — | — |
| Mikołaj Liszewski | TM | za | — | — | — |
