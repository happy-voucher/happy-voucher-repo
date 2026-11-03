[← Spotkania i decyzje](../README.md)

# ADR-004 Przepływ pracy i odbiór zadania

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-11-04 · **Issue:** [Głosowanie nad decyzjami (#237)](https://github.com/happy-voucher/happy-voucher-repo/issues/237)
>
> **Status:** przyjęta · **Data:** 2026-11-04, głosowanie asynchroniczne w issue zamiast warsztatu

## Kontekst

Musimy ustalić, kiedy zadanie jest zakończone i kto to sprawdza. GitHub nie blokuje przejść między statusami, więc kontrolę trzeba oprzeć na przeglądzie i ochronie gałęzi.

## Warianty

| Wariant | Opis | Plusy | Minusy |
| :- | :- | :- | :- |
| A | Przegląd PR przez inną osobę i odbiór historyjki przez koordynatora | Wiedza krąży po zespole, a koordynator pilnuje kryteriów | Kolejka przeglądów |
| B | Bez osobnego przeglądu, odbiór tylko przez testy | Szybciej | Brak kontroli kryteriów biznesowych |

## Wybór i uzasadnienie

**Wariant A z przeglądem w ciągu 24 godzin roboczych.** Spór o wymagania rozstrzyga koordynator, a spór techniczny opiekun modułu. Samodzielne zamykanie zadań przez wykonawców przyspieszyłoby pracę, ale zwiększyłoby ryzyko przeoczenia błędów.

## Rozbieżności lub wynik głosowania

Głosowanie asynchroniczne w issue [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237) od 25 października do 4 listopada 2026, reakcjami pod [komentarzem tej decyzji](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6401145260). Wynik: 4 za, 1 przeciw, 2 osoby bez głosu (Krystian Inoch, Mikołaj Liszewski). Przeciw: Szymon Piechutowski, z komentarzem „Wariant B”, czyli odbiór tylko przez testy. Decyzja przyjęta, bo ma głosy za od 4 z 7 osób.

## Wdrożenie

- Ochrona gałęzi głównej: zmiany tylko przez PR z zatwierdzeniem innej osoby i zielonym CI (D-07).
- Statusy Do zrobienia, Gotowe do startu, W trakcie, Do weryfikacji, Do poprawy, Zakończone na [tablicy „Main”](https://github.com/orgs/happy-voucher/projects).
- Workflow [Odblokuj zależne](../../../.github/workflows/odblokuj-zalezne.yml): po zamknięciu issue przenosi odblokowane historyjki i ich zadania z „Do zrobienia” do „Gotowe do startu”.
- Lista kontrolna w szablonie PR i [definicja ukończenia](../../02-zespol-i-wspolpraca.md#definicja-ukończenia).

## Uczestnicy i ich wkład

**Moderator:** Alan Lubiński (TA) · **Autor notatki:** Krystian Inoch (FE), według podziału z [#237](https://github.com/happy-voucher/happy-voucher-repo/issues/237). Wyniki głosowania spisał Maciej Rożek (AN).

| Osoba | Rola | Głos | Argument, który wpłynął na wybór | Własne działanie | Link do komentarza |
| :- | :- | :- | :- | :- | :- |
| Jakub Sikora | KP | za | — | — | — |
| Maciej Rożek | AN | za | — | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6617612164) |
| Szymon Piechutowski | BE | przeciw (wariant B) | „Wariant B”, „Wśród chaosu jest też sposobność” | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6613221286) |
| Krystian Inoch | FE | brak głosu | — | — | — |
| Jakub Barej | DO | za | — | — | — |
| Alan Lubiński | TA | za | — (w komentarzu tylko „4.A”) | — | [komentarz](https://github.com/happy-voucher/happy-voucher-repo/issues/237#issuecomment-6547608578) |
| Mikołaj Liszewski | TM | brak głosu | — | — | — |
