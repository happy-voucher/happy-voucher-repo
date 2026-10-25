[← Dokumentacja](README.md)

# Zespół i współpraca

> **Opiekun:** Jakub Sikora (KP) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** [D-07 Repozytorium, CI i uruchamianie (#87)](https://github.com/happy-voucher/happy-voucher-repo/issues/87)

## Role i zastępstwa

Każde zadanie ma dokładnie jedną osobę odpowiedzialną. Inni mogą pomagać, ale za wynik odpowiada ta jedna ([ADR-005](06-spotkania-i-decyzje/decyzje/ADR-005-podzial-odpowiedzialnosci.md)).

| Rola | Skrót | Osoba | GitHub | Za co odpowiada | Zastępca |
| :- | :- | :- | :- | :- | :- |
| Koordynator projektu (Product Owner) | KP | Jakub Sikora | @JSikoraVistula | Plan sprintów, priorytety, odbiór historyjek, kontakt z A, B, C i prowadzącym, rejestr decyzji | AN (odbiór i priorytety) |
| Analityk | AN | Maciej Rożek | @kroliktroll123-create | Wymagania, reguły biznesowe, model pojęciowy danych, kontrakt API, kryteria akceptacji, dokumentacja | [uzupełnijcie] |
| Backend | BE | Szymon Piechutowski | @Piechutowski | Serwer Go, schemat i migracje SQLite, logika voucherów, API D, klienci API partnerów | TA (naprawy w kodzie) |
| Frontend | FE | Krystian Inoch | @xwolter | Aplikacja Next.js: sklep (`/`) i panel operatora (`/panel`), logowanie przez Firebase | BE |
| DevOps | DO | Jakub Barej | @Ksionk | Repozytorium, CI, budowanie i uruchamianie, symulator partnerów, środowisko testowe, kopie bazy | BE (administracja) |
| Tester automatyczny | TA | Alan Lubiński | @Alanlub01 | Testy jednostkowe i integracyjne w Go, testy API, scenariusz wspólny, testy współbieżności | [uzupełnijcie] |
| Tester manualny | TM | Mikołaj Liszewski | @Mikson098 | Przypadki testowe, testy sklepu i panelu, testy odbiorcze z partnerami, zgłaszanie błędów | [uzupełnijcie] |

Zastępca robi co najmniej jeden przegląd PR na sprint w module osoby, którą zastępuje.

## Uprawnienia w GitHubie

| Kto | Uprawnienia |
| :- | :- |
| DevOps | Admin: konfiguracja repozytorium i projektu |
| Backend | Admin jako zastępca DevOps |
| Pozostali członkowie zespołu | Write |
| Prowadzący | Read albo Triage |

Issue nie usuwamy, tylko zamykamy jako nieplanowane, bo usunięcia nie da się cofnąć.

## Kontakty do zespołów A, B, C

| Zespół | System | Osoba kontaktowa | Kanał |
| :- | :- | :- | :- |
| A | Gabinet masażu | [uzupełnijcie] | [uzupełnijcie] |
| B | Studio fitness | [uzupełnijcie] | [uzupełnijcie] |
| C | Hotel | [uzupełnijcie] | [uzupełnijcie] |

Za kontakt z zespołami odpowiada koordynator. Zmiany kontraktu zgłaszamy według zasad w [kontrakcie integracji](05-projekt-rozwiazania/kontrakt-integracji.md).

## Zasady pracy

- Uproszczony Scrum w dwutygodniowych sprintach, prowadzony w [GitHub Projects, na tablicy „Main”](https://github.com/orgs/happy-voucher/projects). Sprint to pole Iteration ([ADR-001](06-spotkania-i-decyzje/decyzje/ADR-001-sposob-organizacji-pracy.md)).
- Każda osoba ma najwyżej 2 zadania w statusie „W trakcie”.
- Statusy: Do zrobienia → Gotowe do startu → W trakcie → Do weryfikacji → Zakończone, a z Do weryfikacji także Do poprawy. Epik ma uproszczony przepływ: Do zrobienia, W trakcie, Zakończone.
- Historyjka i jej zadania przechodzą do „Gotowe do startu” automatycznie, gdy zamknięte zostaną wszystkie issue, które je blokują (workflow [Odblokuj zależne](../.github/workflows/odblokuj-zalezne.yml)). Epiki mają historyjki jako sub-issues, a historyjki mają zadania jako sub-issues.
- Priorytety: P0 krytyczny (blokuje integrację albo demo, tylko błędy), P1 musi być, P2 powinno być, P3 jeśli starczy czasu. Priorytet ustala koordynator, a szacunek cały zespół ([ADR-003](06-spotkania-i-decyzje/decyzje/ADR-003-priorytety-i-szacowanie.md)).
- Blokady oznaczamy relacją „blocked by” i etykietą `zablokowane`. Oczekiwanie na A, B albo C oznaczamy dodatkowo etykietą `czeka-na-partnera`.
- Pull request przegląda inna osoba w ciągu 24 godzin roboczych. Spór o wymagania rozstrzyga koordynator, a spór techniczny opiekun modułu ([ADR-004](06-spotkania-i-decyzje/decyzje/ADR-004-przeplyw-pracy-i-odbior.md)).
- Zadanie, które od 3 dni stoi w „W trakcie” bez zmian, Scrum Master omawia z wykonawcą, zaczynając od pytania, czy potrzebuje pomocy.

## Definicja ukończenia

1. Kod jest w gałęzi głównej przez PR zatwierdzony przez inną osobę.
2. Testy automatyczne przechodzą.
3. Nowa logika ma testy, a naprawiony błąd ma test, który bez poprawki by nie przeszedł.
4. Zmiana API albo zachowania jest opisana w dokumentacji w tym samym PR.
5. Historyjka: koordynator projektu sprawdził każde kryterium akceptacji i zostawił komentarz.
6. Nie ma otwartych sub-issues ani nierozwiązanych komentarzy w PR.
