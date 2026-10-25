[← Wymagania](README.md)

# Słownik pojęć i aktorzy

> **Opiekun:** Maciej Rożek (AN) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** D-01 Wymagania i reguły biznesowe voucherów, zadanie T1
>
> **Status:** szkic do przeglądu (D-01 T4 BE), zatwierdzenie na spotkaniu D-01 T5 (KP)

Ta strona ustala wspólne słownictwo zespołu D. Tych samych słów używamy w issue, dokumentacji, kodzie i rozmowach z zespołami A, B i C. Jeśli ktoś w rozmowie rozumie pojęcie inaczej niż tutaj, poprawiamy tę stronę albo rozmowę, a nie tworzymy synonimu.

Kolumna „Nazwa w kodzie” to propozycja dla BE i specyfikacji API (D-04, D-05). Ostateczne nazwy tabel i pól ustala schemat bazy w [model-danych.md](../05-projekt-rozwiazania/model-danych.md). Encje modelu pojęciowego z D-02 są oznaczone gwiazdką (★).

## Aktorzy

### Osoby

| Aktor | Kto to jest | Gdzie pracuje | Co robi w systemie D |
| :- | :- | :- | :- |
| **Klient** | Osoba, która kupuje voucher | Sklep www | Przegląda katalog, składa zamówienie, opłaca je (płatność symulowana) i dostaje kody voucherów |
| **Posiadacz vouchera** | Osoba, która zna kod vouchera: klient albo obdarowany | Sklep www | Ogląda voucher i pozostałe usługi, wybiera termin i rezerwuje usługę u partnera przez D |
| **Operator D** | Pracownik obsługi firmy D | Panel obsługi (guigui) | Rejestruje partnerów i ich klucze, tworzy i publikuje oferty, wyszukuje vouchery, przegląda historię i kolejkę spraw, przygotowuje rozliczenia |

### Partnerzy (systemy zewnętrzne z własnym API)

| Aktor | Kto to jest | Co robi w systemie D |
| :- | :- | :- |
| **Partner A** | Gabinet masażu | Wykonuje masaże z voucherów |
| **Partner B** | Studio fitness | Wykonuje treningi z voucherów |
| **Partner C** | Hotel | Zapewnia pobyty z voucherów, a w pakiecie z noclegiem jest koordynatorem pakietu |
| **Koordynator pakietu** | Rola hotelu C w pakiecie z noclegiem | Widzi wszystkie pozycje pakietu, nie tylko własne, i planuje pobyt: rezerwuje masaż w A i trening w B |

Każdy partner korzysta z API D: weryfikuje kod, blokuje i zwalnia pozycje, potwierdza realizację i pobiera swoje zestawienie. Każdy udostępnia też D własne API: listę usług, dostępność, rezerwację, status rezerwacji i anulowanie (D-03).

### Systemy pomocnicze

| Aktor | Co robi |
| :- | :- |
| **Symulator płatności** | Zwraca wynik płatności wybrany przy testach: sukces, odmowę albo przekroczenie czasu. Prawdziwych płatności nie ma |
| **Harmonogram zadań** | Uruchamia zadania cykliczne: synchronizację usług partnerów co godzinę, wygaszanie voucherów raz na dobę, sprawdzanie rezerwacji „do wyjaśnienia” |

Symulator API partnerów (`cmd/partnersim`, D-09) nie jest aktorem. To narzędzie testowe, które udaje A, B i C zgodnie z kontraktem.

## Pojęcia

### Partnerzy i oferta

| Pojęcie | Definicja | Nazwa w kodzie |
| :- | :- | :- |
| **Partner** ★ | Firma A, B lub C zarejestrowana w D z adresem API i kluczami | `partner` |
| **Usługa partnera** ★ | Pojedyncza usługa pobrana z API partnera: rodzaj masażu, trening, pobyt. Ma identyfikator u partnera, nazwę, cenę u partnera i znacznik aktywności | `partner_service` |
| **Synchronizacja usług** | Pobranie aktualnej listy usług partnera, co godzinę i na żądanie. Usługi, których partner już nie zwraca, stają się nieaktywne. Każda synchronizacja trafia do historii synchronizacji | `sync`, `sync_log` |
| **Oferta** ★ | Produkt, który D sprzedaje w katalogu: nazwa, opis, cena, okres ważności w dniach i pozycje oferty. Ma stan: robocza, aktywna albo wycofana | `offer` |
| **Pozycja oferty** ★ | Jedna usługa partnera w ofercie, z liczbą sztuk i kwotą dla partnera | `offer_item` |
| **Voucher na usługę** | Oferta z jedną usługą jednego partnera, na przykład „Masaż relaksacyjny 60 min w A” | rodzaj oferty `service` |
| **Pakiet** | Oferta z kilkoma usługami, także od różnych partnerów, na przykład „Nocleg w C + masaż w A + trening w B” | rodzaj oferty `package` |
| **Pakiet z noclegiem** | Pakiet, w którym jest pobyt w hotelu C. C jest wtedy koordynatorem pakietu | pole `coordinator_partner_id` |
| **Katalog** | Aktywne oferty w sklepie, podzielone na masaż, fitness, hotel i pakiety, z filtrami po partnerze i rodzaju | `catalog` |

### Zamówienie i voucher

| Pojęcie | Definicja | Nazwa w kodzie |
| :- | :- | :- |
| **Zamówienie** ★ | Zakup oferty przez klienta: oferta, liczba sztuk, e-mail klienta i opcjonalnie imię obdarowanego. Cena jest zapamiętana w chwili zakupu | `order` |
| **Płatność** ★ | Próba opłacenia zamówienia w symulatorze płatności, z wynikiem sukces, odmowa albo przekroczenie czasu | `payment` |
| **Voucher** ★ | Uprawnienie do usług z jednej oferty, wydane po udanej płatności. Jedna sztuka w zamówieniu to jeden voucher z własnym kodem. Status vouchera wynika ze stanów jego pozycji | `voucher` |
| **Kod vouchera** | Unikalny kod w formacie `VD-XXXX-XXXX-XXXX` z alfabetu bez mylących znaków, losowany z `crypto/rand`, z sumą kontrolną. Kto zna kod i e-mail z zamówienia, ma dostęp do vouchera, więc kod jest poufny | `code` |
| **Pozycja vouchera** ★ | Jedna usługa do wykorzystania dokładnie raz, u jednego partnera. Ma właściciela (partnera), kwotę dla partnera i stan | `voucher_item` |
| **Okres ważności** | Liczba dni ważności zapisana w ofercie | `validity_days` |
| **Termin ważności** | Ostatni dzień, w którym voucher jest ważny. Liczymy go z okresu ważności oferty od dnia opłacenia | `valid_until` |
| **Obdarowany** | Osoba, której klient daje voucher. System zna najwyżej jej imię | `recipient_name` |
| **Strona vouchera** | Widok w sklepie, w którym posiadacz po podaniu kodu i e-maila widzi status, termin ważności i pozycje z terminami rezerwacji | `voucher_page` |

### Rezerwacja, blokada i realizacja

| Pojęcie | Definicja | Nazwa w kodzie |
| :- | :- | :- |
| **Dostępność** | Wolne terminy usługi, które partner zwraca przez swoje API. D pyta o zakres do 14 dni | `availability` |
| **Rezerwacja u partnera** ★ | Termin usługi zarezerwowany u partnera: przez D na prośbę posiadacza albo przez samego partnera i zgłoszony w D blokadą | `partner_booking` |
| **Identyfikator rezerwacji partnera** | Identyfikator nadany rezerwacji w systemie partnera. D go zapisuje, żeby obie strony mówiły o tej samej wizycie | `partner_booking_ref` |
| **Do wyjaśnienia** | Status rezerwacji, gdy D nie wie, czy partner ją przyjął, na przykład po przekroczeniu czasu odpowiedzi. Zadanie cykliczne sprawdza go u partnera | `uncertain` |
| **Blokada** | Zgłoszenie przez partnera w API D rezerwacji, którą przyjął z kodem vouchera (na przykład klient zadzwonił do gabinetu albo hotel C zarezerwował masaż w A). D wybiera wolną pozycję tej usługi i oznacza ją jako zarezerwowaną | `hold` |
| **Zwolnienie blokady** | Cofnięcie blokady przez partnera. Pozycja wraca do stanu dostępna | `release` |
| **Weryfikacja kodu** | Sprawdzenie przez partnera, czy kod istnieje i jest ważny oraz jakie pozycje widzi ten partner | `verify` |
| **Realizacja** | Operacja w API D: partner potwierdza, że wykonał usługę z pozycji | `redeem` |
| **Wykorzystanie** ★ | Zapis w bazie powstały przy realizacji: pozycja, data, identyfikator rezerwacji partnera. Na pozycję przypada najwyżej jedno aktywne wykorzystanie | `redemption` |
| **Klucz idempotencji** | Identyfikator żądania w nagłówku `Idempotency-Key`. Powtórzenie żądania z tym samym kluczem zwraca tę samą odpowiedź i niczego nie zmienia drugi raz | `idempotency_key` |
| **Klucz API partnera** | Sekret, którym partner uwierzytelnia się w API D. D przechowuje tylko jego skrót | `api_key_hash` |
| **Bezpiecznik** | Mechanizm, który po 5 kolejnych błędach partnera wstrzymuje wywołania tego partnera na 30 s (ang. *circuit breaker*) | `breaker` |

### Rozliczenia i obsługa

| Pojęcie | Definicja | Nazwa w kodzie |
| :- | :- | :- |
| **Kwota dla partnera** | Część ceny, którą D płaci partnerowi za zrealizowaną pozycję. Wszystkie kwoty to liczby całkowite w groszach | `partner_amount` |
| **Marża D** | Cena oferty minus suma kwot dla partnerów | `margin` |
| **Rozliczenie** ★ | Zestawienie pozycji zrealizowanych przez partnera w miesiącu kalendarzowym z sumą kwot. O miesiącu decyduje data realizacji w strefie Europe/Warsaw | `settlement` |
| **Kolejka spraw** | Lista spraw do wyjaśnienia przez operatora: rezerwacje „do wyjaśnienia” i nieudane synchronizacje | `case_queue` |
| **Historia vouchera** | Zapis każdej zmiany stanu pozycji: kto, kiedy, z jakiego stanu na jaki i w wyniku jakiej operacji | `voucher_event` |

Statusy vouchera i stany pozycji oraz przejścia między nimi opisuje [stany-vouchera-i-pozycji.md](stany-vouchera-i-pozycji.md).
