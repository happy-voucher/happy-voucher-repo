# Happy Voucher: system D

> **Opiekun:** Jakub Sikora (KP), instrukcja uruchomienia: Jakub Barej (DO) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** D-07 Repozytorium, CI i uruchamianie

System D sprzedaje i obsługuje vouchery na usługi trzech partnerów: gabinetu masażu (A), studia fitness (B) i hotelu (C). Klient kupuje w D voucher na usługę albo pakiet, rezerwuje terminy u partnerów, a partnerzy sprawdzają i realizują vouchery przez API D. D rozlicza się potem z partnerami.

Projekt zespołowy, grupa zajęciowa 2. Pracujemy w uproszczonym Scrumie w dwutygodniowych sprintach ([ADR-001](docs/06-spotkania-i-decyzje/decyzje/ADR-001-sposob-organizacji-pracy.md)).

## Rola D we wspólnym scenariuszu

1. Klient kupuje w D voucher na pakiet: nocleg, masaż i trening.
2. D sprawdza terminy i rezerwuje pobyt w C.
3. C uzgadnia z klientem plan pobytu i rezerwuje masaż w A oraz trening w B.
4. A i B potwierdzają rezerwacje i blokują odpowiednie terminy.
5. Przy realizacji usługi każda firma sprawdza voucher przez API D.
6. Po wykonaniu usługi partner aktualizuje własną rezerwację i potwierdza w D wykorzystanie odpowiedniej części vouchera.
7. D pokazuje klientowi pozostałe usługi i przygotowuje zestawienie rozliczeń z partnerami.

## Z czego składa się system

| Część | Technologia | Dla kogo |
| :- | :- | :- |
| Serwer i API D | Go (`net/http`), baza SQLite z migracjami | Partnerzy A, B, C, sklep, panel |
| Sklep www | Szablony `html/template`, samodzielne pliki bez CDN | Klient i posiadacz vouchera |
| Panel obsługi | Aplikacja desktopowa w guigui, korzysta tylko z API administracyjnego | Operator D |
| Symulator partnerów | `cmd/partnersim`, tryby `ok`, `conflict`, `full`, `down`, `slow` | Zespół D, testy |

Szczegóły są w [projekcie rozwiązania](docs/05-projekt-rozwiazania/README.md).

## Jak uruchomić

System jest w budowie. Instrukcję uruchomienia serwera, symulatora i panelu dopiszemy w historyjce D-07 (zadanie T4).

## Linki

- [Tablica projektu w GitHub Projects](https://github.com/orgs/happy-voucher/projects/10/views/1)
- [Dokumentacja: indeks](docs/README.md)
- [Rejestr decyzji zespołu](docs/06-spotkania-i-decyzje/README.md)
- [Opis API D dla partnerów](docs/05-projekt-rozwiazania/api-d.md)

## Zespół

| Rola | Skrót | Osoba | GitHub |
| :- | :- | :- | :- |
| Koordynator projektu (Product Owner) | KP | Jakub Sikora | [@JSikoraVistula](https://github.com/JSikoraVistula) |
| Analityk | AN | Maciej Rożek | [@kroliktroll123-create](https://github.com/kroliktroll123-create) |
| Backend | BE | Szymon Piechutowski | [@Piechutowski](https://github.com/Piechutowski) |
| Frontend | FE | Krystian Inoch | [@xwolter](https://github.com/xwolter) |
| DevOps | DO | Jakub Barej | [@Ksionk](https://github.com/Ksionk) |
| Tester automatyczny | TA | Alan Lubiński | [@Alanlub01](https://github.com/Alanlub01) |
| Tester manualny | TM | Mikołaj Liszewski | [@Mikson098](https://github.com/Mikson098) |

Role, zastępstwa i zasady pracy opisuje [Zespół i współpraca](docs/02-zespol-i-wspolpraca.md).
