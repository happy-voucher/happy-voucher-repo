[← Projekt rozwiązania](README.md)

# Kontrakt integracji z A, B, C

> **Opiekun:** Szymon Piechutowski (BE) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** [D-03 Uzgodnienie kontraktu z zespołami A, B, C (#84)](https://github.com/happy-voucher/happy-voucher-repo/issues/84) (T1 propozycja AN, T2 spotkanie KP, T4 zapis AN)
>
> **Status:** propozycja zespołu D do uzgodnienia na spotkaniu czterech zespołów (D-03 T2)

Kontrakt ustala, jak systemy A, B, C i D rozmawiają ze sobą. Po podpisaniu przez cztery zespoły kopia trafia do repozytorium każdego zespołu, a zmiany zgłaszamy według zasad na dole strony.

## Wspólne formaty

| Element | Propozycja D | Uzgodnione |
| :- | :- | :- |
| Identyfikatory partnerów | `A`, `B`, `C`, `D` | [ ] |
| Czas | RFC 3339 ze strefą, na przykład `2026-11-14T15:30:00+01:00` | [ ] |
| Kwoty | Liczby całkowite w groszach, waluta PLN | [ ] |
| Format błędu | `{"error": {"code", "message", "request_id"}}`, przykład w [api-d.md](api-d.md#zasady-wspólne) | [ ] |
| Kody błędów | Lista kodów wspólna dla czterech zespołów [do uzgodnienia] | [ ] |
| Uwierzytelnianie | Klucz API w nagłówku `Authorization: Bearer`, osobny klucz dla każdej pary zespołów | [ ] |
| Idempotencja | Nagłówek `Idempotency-Key` w żądaniach zmieniających stan, w obu kierunkach | [ ] |

## Operacje, które D wywołuje u partnerów

| Operacja | Po co D jej potrzebuje | Historyjka |
| :- | :- | :- |
| Lista usług | Synchronizacja usług co godzinę i na żądanie | D-10 |
| Dostępność | Wolne terminy usługi w zakresie do 14 dni | D-20 |
| Rezerwacja | Rezerwacja terminu z kodem vouchera i kluczem idempotencji. Partner zwraca identyfikator rezerwacji | D-20 |
| Status rezerwacji | Wyjaśnienie rezerwacji o niepewnym wyniku („do wyjaśnienia”) | D-21 |
| Anulowanie | Anulowanie rezerwacji złożonej przez D | D-20 |

**Limity po stronie D, o których informujemy partnerów:** D czeka na odpowiedź najwyżej 5 s i ponawia żądanie do 3 razy. Po 5 kolejnych błędach przestaje wywoływać partnera na 30 s (D-21).

## Operacje, które partnerzy wywołują w D

Weryfikacja kodu, blokada i zwolnienie pozycji, realizacja i zestawienie partnera. Opis jest w [api-d.md](api-d.md).

## Zasady zmian

- Zmianę kontraktu zgłasza się przez koordynatora zespołu D (Jakub Sikora) [kanał do uzgodnienia].
- Każda zmiana podnosi numer wersji kontraktu i jest opisana w tabeli wersji.
- Zmiana, która psuje zgodność, wymaga zgody wszystkich czterech zespołów i terminu wdrożenia.

## Wersje i akceptacja

| Wersja | Data | Zmiana |
| :- | :- | :- |
| 0.1 | 2026-10-25 | Propozycja zespołu D |

| Zespół | Osoba | Data | Akceptacja |
| :- | :- | :- | :- |
| A | [ ] | [ ] | [ ] |
| B | [ ] | [ ] | [ ] |
| C | [ ] | [ ] | [ ] |
| D | Jakub Sikora (KP) | [ ] | [ ] |
