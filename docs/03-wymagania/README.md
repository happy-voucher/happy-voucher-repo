[← Dokumentacja](../README.md)

# Wymagania

> **Opiekun:** Maciej Rożek (AN) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** D-01 Wymagania i reguły biznesowe voucherów, D-02 Model pojęciowy danych i przypadki testowe

Ten obszar opisuje, **co** system D ma robić i według jakich zasad: pojęcia, aktorów, reguły biznesowe i stany vouchera. Jak to zbudujemy, opisuje [projekt rozwiązania](../05-projekt-rozwiazania/README.md). Na te strony powołują się kryteria akceptacji historyjek, testy i ekrany.

## Strony

| Strona | Zawartość | Zadanie | Stan |
| :- | :- | :- | :- |
| [slownik-pojec.md](slownik-pojec.md) | Aktorzy i słownik pojęć z proponowanymi nazwami w kodzie | D-01 T1 | szkic do przeglądu |
| [reguly-biznesowe.md](reguly-biznesowe.md) | Numerowane reguły, tabela „kto co może” i otwarte pytania do prowadzącego i partnerów | D-01 T2 | szkic do przeglądu |
| [stany-vouchera-i-pozycji.md](stany-vouchera-i-pozycji.md) | Diagramy stanów pozycji i vouchera, przejścia, skutki operacji | D-02 T2 | szkic |
| wymagania-niefunkcjonalne.md | Bezpieczeństwo kodów, wydajność, audyt zmian | brak zadania w backlogu, do ustalenia z KP | do zrobienia |

## Epiki

Reguły odnoszą się do historyjek z backlogu. Linki do issue dodamy, gdy epiki i historyjki zostaną założone na GitHubie.

| Epik | Cel |
| :- | :- |
| D-E0 Analiza i wymagania | Cały zespół rozumie, co ma powstać |
| D-E1 Kontrakt integracji z A, B, C | Cztery zespoły mają jeden uzgodniony sposób komunikacji |
| D-E2 Baza danych | Spójny schemat SQLite, na którym pracują wszystkie moduły |
| D-E3 Infrastruktura i szkielet aplikacji | Każdy może uruchomić i przetestować system w minutę |
| D-E4 Katalog usług i ofert | Klient widzi aktualne vouchery zbudowane z usług A, B i C |
| D-E5 Sprzedaż i wydawanie voucherów | Klient kupuje voucher online i od razu dostaje kod |
| D-E6 API D: weryfikacja i realizacja | Partner weryfikuje kod, blokuje pozycję i potwierdza realizację, a usługa jest wykorzystana dokładnie raz |
| D-E7 Rezerwacja usług u partnerów | Posiadacz vouchera rezerwuje termin u partnera z poziomu D |
| D-E8 Rozliczenia i panel obsługi | D rozlicza się z partnerami, a operator wyjaśnia sprawy klientów |
| D-E9 Testy integracyjne i odbiór | Scenariusz wspólny działa z prawdziwymi partnerami, także w sytuacjach błędnych |

## Jak zmieniamy wymagania

- Każda zmiana przechodzi przez pull request. Zmianę reguły przegląda BE pod kątem wykonalności (jak w D-01 T4), a zmianę zakresu zatwierdza KP.
- Spór o wymagania rozstrzyga koordynator projektu ([ADR-004](../06-spotkania-i-decyzje/decyzje/ADR-004-przeplyw-pracy-i-odbior.md)).
- Po zatwierdzeniu numerów reguł i pytań nie zmieniamy ani nie używamy ponownie. Reguła wycofana zostaje w tabeli z dopiskiem „wycofana” i informacją, co ją zastępuje.
- Gdy odpowiedź na otwarte pytanie zmienia regułę, w tym samym PR poprawiamy regułę i wpisujemy odpowiedź.
- Nowe pojęcie najpierw trafia do słownika, a dopiero potem do reguł, kodu i API.
