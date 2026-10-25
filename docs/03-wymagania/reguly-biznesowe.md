[← Wymagania](README.md)

# Reguły biznesowe

> **Opiekun:** Maciej Rożek (AN) · **Ostatnia zmiana:** 2026-10-25 · **Issue:** D-01 Wymagania i reguły biznesowe voucherów, zadanie T2
>
> **Status:** szkic do przeglądu (D-01 T4 BE), zatwierdzenie na spotkaniu D-01 T5 (KP)

Reguły opisują, co system D musi zapewnić niezależnie od tego, jak to zaprogramujemy. Powołują się na nie kryteria akceptacji, testy i ekrany. Pojęcia znaczą to, co w [słowniku pojęć](slownik-pojec.md), a stany i przejścia opisuje [stany-vouchera-i-pozycji.md](stany-vouchera-i-pozycji.md).

**Jak czytać:**
- Każda reguła ma stały numer: skrót grupy i kolejny numer, na przykład `REA-03`. Po zatwierdzeniu numerów nie zmieniamy i nie używamy ponownie. Reguła wycofana zostaje w tabeli z dopiskiem „wycofana”.
- Kolumna „Źródło” wskazuje historyjkę z backlogu, decyzję zespołu albo założenie. **Założenie** to propozycja analityka, która czeka na odpowiedź na pytanie z sekcji [Otwarte pytania](#otwarte-pytania).
- Kody HTTP podajemy tam, gdzie wymagają ich kryteria akceptacji. Pełną listę kodów błędów ustalą kontrakt (D-03) i specyfikacja API D (D-04).

## Pokrycie kryteriów D-01

| Kryterium historyjki D-01 | Gdzie jest spełnione |
| :- | :- |
| (1) Aktorzy: klient, posiadacz vouchera, operator D, partnerzy A, B, C | [Słownik pojęć, sekcja Aktorzy](slownik-pojec.md#aktorzy) |
| (2) Unikalny kod | KOD-01, KOD-02, KOD-03 |
| (2) Termin ważności | WAZ-01 do WAZ-06 |
| (2) Każda usługa z vouchera wykorzystana dokładnie raz | POZ-01, REA-03, REA-04, REA-05 |
| (2) Realizuje tylko partner, do którego należy usługa | POZ-02, REA-01, REA-03 |
| (2) Kwoty w groszach | ROZ-01 |
| (3) Dokument zatwierdzony przez cały zespół | Spotkanie D-01 T5, wpis w [rejestrze decyzji](../06-spotkania-i-decyzje/README.md) |

## Grupy reguł

| Skrót | Grupa |
| :- | :- |
| [OFE](#ofe-oferty-i-katalog) | Oferty i katalog |
| [SYN](#syn-partnerzy-i-synchronizacja-usług) | Partnerzy i synchronizacja usług |
| [ZAM](#zam-zamówienia-i-płatność) | Zamówienia i płatność |
| [KOD](#kod-voucher-i-kod) | Voucher i kod |
| [WAZ](#waz-ważność) | Ważność |
| [POZ](#poz-pozycje-i-jednorazowość) | Pozycje i jednorazowość |
| [BLK](#blk-weryfikacja-i-blokada) | Weryfikacja i blokada |
| [REZ](#rez-rezerwacje-przez-d) | Rezerwacje przez D |
| [REA](#rea-realizacja) | Realizacja |
| [ROZ](#roz-rozliczenia) | Rozliczenia |
| [UPR](#upr-uprawnienia) | Uprawnienia |
| [WSP](#wsp-zasady-wspólne-api) | Zasady wspólne API |
| [ZAK](#zak-poza-zakresem) | Poza zakresem |

## OFE Oferty i katalog

| Nr | Reguła | Źródło |
| :- | :- | :- |
| OFE-01 | Ofertę tworzy, edytuje, publikuje i wycofuje operator D. Oferta ma nazwę, opis, cenę, okres ważności w dniach, rodzaj (voucher na usługę albo pakiet) i pozycje. | D-11 |
| OFE-02 | Każda pozycja oferty wskazuje usługę jednego partnera, liczbę sztuk i kwotę dla partnera za jedną sztukę. | D-11, założenie (P-12) |
| OFE-03 | Oferta ma stan robocza, aktywna albo wycofana. Klient widzi i kupuje tylko oferty aktywne. | D-11, D-12 |
| OFE-04 | Oferty nie można opublikować, jeśli którakolwiek jej usługa jest nieaktywna albo suma kwot dla partnerów przekracza cenę. Różnica między ceną a sumą kwot to marża D. | D-11, D-22 |
| OFE-05 | Pakiet z noclegiem w hotelu C ma C jako koordynatora pakietu. | D-11 |
| OFE-06 | Wycofana oferta znika z katalogu, a sprzedane z niej vouchery działają dalej bez zmian. | D-11, założenie |
| OFE-07 | Zmiana ceny lub pozycji oferty dotyczy tylko nowych zamówień. Złożone zamówienia i wydane vouchery zostają bez zmian. | D-13, założenie |
| OFE-08 | Gdy synchronizacja oznaczy usługę jako nieaktywną, oferty z tą usługą są wycofywane z katalogu, a operator dostaje sprawę w kolejce. Wydane vouchery zostają bez zmian. | Założenie (P-04) |
| OFE-09 | Katalog dzieli oferty na masaż, fitness, hotel i pakiety oraz pozwala filtrować po partnerze i rodzaju. Szczegóły oferty pokazują każdą jej usługę. | D-12 |

## SYN Partnerzy i synchronizacja usług

| Nr | Reguła | Źródło |
| :- | :- | :- |
| SYN-01 | Operator rejestruje partnera z adresem API i kluczami. | D-10 |
| SYN-02 | D pobiera usługi każdego partnera przez jego API co godzinę i na żądanie operatora. | D-10 |
| SYN-03 | Nowe i zmienione usługi są zapisywane. Usługi nieobecne w odpowiedzi partnera są oznaczane jako nieaktywne, a nie usuwane. | D-10 |
| SYN-04 | Błąd partnera nie kasuje wcześniej pobranych danych. Jest zapisany w historii synchronizacji i trafia do kolejki spraw. | D-10, D-23 |
| SYN-05 | Zmiana ceny usługi u partnera nie zmienia automatycznie ceny oferty ani kwoty dla partnera. Decyzję podejmuje operator. | Założenie (P-10) |

## ZAM Zamówienia i płatność

| Nr | Reguła | Źródło |
| :- | :- | :- |
| ZAM-01 | Zamówienie zawiera ofertę, liczbę sztuk, e-mail klienta i opcjonalnie imię obdarowanego. Cena jest zapamiętana w chwili zakupu. | D-13 |
| ZAM-02 | Zamówić można tylko aktywną ofertę. | D-12, D-13 |
| ZAM-03 | Jedno zamówienie obejmuje od 1 do 10 sztuk. | Założenie (P-06) |
| ZAM-04 | Płatność jest symulowana i kończy się sukcesem, odmową albo przekroczeniem czasu. | D-13 |
| ZAM-05 | Sukces oznacza zamówienie jako opłacone i tworzy vouchery. Odmowa nie tworzy niczego. Przebieg po przekroczeniu czasu i ponowienie płatności opisze specyfikacja D-13 T1. | D-13 |
| ZAM-06 | Ponowne wysłanie tej samej płatności nie tworzy drugiego kompletu voucherów. | D-13 |
| ZAM-07 | Każda sztuka w zamówieniu to osobny voucher z własnym kodem. | D-13, D-14 |
| ZAM-08 | Sprzedaż voucherów działa także wtedy, gdy któryś partner nie odpowiada. | D-21 |

## KOD Voucher i kod

| Nr | Reguła | Źródło |
| :- | :- | :- |
| KOD-01 | Kod vouchera ma format `VD-XXXX-XXXX-XXXX` z alfabetu bez mylących znaków, jest losowany z `crypto/rand` i ma sumę kontrolną. | D-14 |
| KOD-02 | Kod jest unikalny w całym systemie, co wymusza ograniczenie UNIQUE w bazie. Przy kolizji kod jest losowany ponownie. | D-05, D-14 |
| KOD-03 | Kodu nie da się wyliczyć z danych zamówienia ani odgadnąć na podstawie innych kodów. | D-14 |
| KOD-04 | Przy wydaniu vouchera z każdej pozycji oferty powstaje tyle pozycji vouchera, ile sztuk tej usługi jest w pozycji oferty. Każda pozycja vouchera dostaje stan **dostępna**, właściciela i kwotę dla partnera. | D-14, założenie (P-12) |
| KOD-05 | Voucher i wszystkie jego pozycje powstają w jednej transakcji. | D-14 |
| KOD-06 | Po udanej płatności klient widzi kody na stronie po zakupie. | D-15 |
| KOD-07 | Posiadacz ogląda voucher po podaniu kodu i adresu e-mail z zamówienia. Widzi status, termin ważności oraz każdą pozycję z jej stanem i terminem rezerwacji. | D-15, założenie (P-03) |
| KOD-08 | Błędny e-mail daje taki sam komunikat jak nieznany kod. Po 10 nieudanych próbach z jednego adresu kolejne próby są odrzucane przez 15 minut. | D-15 |

## WAZ Ważność

| Nr | Reguła | Źródło |
| :- | :- | :- |
| WAZ-01 | Okres ważności jest określony w ofercie w dniach. Termin ważności vouchera liczymy z tego okresu od dnia opłacenia zamówienia. | D-11, D-14 |
| WAZ-02 | Voucher jest ważny do 23:59:59 ostatniego dnia ważności w strefie Europe/Warsaw. | Założenie |
| WAZ-03 | Weryfikacja, blokada, rezerwacja i realizacja sprawdzają ważność w chwili operacji, niezależnie od zadania cyklicznego. | D-19 |
| WAZ-04 | Zadanie cykliczne raz na dobę oznacza niezrealizowane pozycje przeterminowanych voucherów jako **wygasłe**. | D-19 |
| WAZ-05 | D pokazuje i rezerwuje tylko terminy przed końcem ważności vouchera. | D-20 |
| WAZ-06 | W pierwszej wersji ważności nie można przedłużyć. | Założenie (P-07) |

## POZ Pozycje i jednorazowość

| Nr | Reguła | Źródło |
| :- | :- | :- |
| POZ-01 | Każda usługa z vouchera, czyli każda pozycja, jest wykorzystana najwyżej raz. Baza danych nie pozwala, żeby pozycja miała dwa aktywne wykorzystania (częściowy indeks UNIQUE). | D-01, D-05, D-19 |
| POZ-02 | Pozycja należy do jednego partnera: tego, który wykonuje usługę. Właściciela nie można zmienić. | D-01, D-18 |
| POZ-03 | Pozycja ma zawsze dokładnie jeden stan: dostępna, zarezerwowana, zrealizowana, wygasła albo anulowana. Dozwolone przejścia opisuje [stany-vouchera-i-pozycji.md](stany-vouchera-i-pozycji.md). | D-02 |
| POZ-04 | Zrealizowana, wygasła i anulowana to stany końcowe. Wyjątkiem może być cofnięcie realizacji. | D-02, założenie (P-14) |
| POZ-05 | Statusu vouchera nie ustawia się ręcznie. Zawsze wynika ze stanów jego pozycji. | D-18 |
| POZ-06 | Każda zmiana stanu pozycji trafia do historii vouchera: kto, kiedy, z jakiego stanu na jaki i w wyniku jakiej operacji. | D-18, D-23 |
| POZ-07 | Voucher anuluje operator. Wszystkie niezrealizowane pozycje przechodzą wtedy w stan **anulowana**, a D anuluje ich rezerwacje u partnerów. Zrealizowane pozycje zostają i są rozliczane. | Założenie (P-13) |

## BLK Weryfikacja i blokada

| Nr | Reguła | Źródło |
| :- | :- | :- |
| BLK-01 | Weryfikacja zwraca status vouchera, ważność i pozycje widoczne dla wywołującego partnera: jego własne, a dla koordynatora pakietu wszystkie pozycje pakietu. | D-16 |
| BLK-02 | Nieznany kod daje 404, a voucher wygasły albo anulowany daje 410 z odpowiednim kodem błędu z kontraktu. | D-16 |
| BLK-03 | Partner, który przyjął rezerwację z kodem vouchera, blokuje pozycję: podaje identyfikator usługi i swojej rezerwacji. D wybiera wolną (dostępną) pozycję tej usługi, zmienia jej stan na **zarezerwowana** i zwraca jej identyfikator. | D-17 |
| BLK-04 | Gdy nie ma wolnej pozycji tej usługi, blokada kończy się błędem 409. | D-17 |
| BLK-05 | Zwolnienie blokady przywraca pozycji stan **dostępna**. | D-17 |
| BLK-06 | Ponowienie blokady z tym samym kluczem idempotencji nie blokuje drugiej pozycji i zwraca tę samą odpowiedź. | D-17 |
| BLK-07 | Blokuje i zwalnia tylko właściciel pozycji. Koordynator pakietu nie blokuje pozycji innych partnerów. | Założenie (P-09) |

## REZ Rezerwacje przez D

| Nr | Reguła | Źródło |
| :- | :- | :- |
| REZ-01 | Posiadacz wybiera termin w sklepie. D pobiera dostępność z API partnera dla zakresu do 14 dni i pokazuje tylko terminy przed końcem ważności. | D-20 |
| REZ-02 | Rezerwować można tylko pozycję dostępną. | D-20 |
| REZ-03 | D wysyła rezerwację z kodem vouchera i kluczem idempotencji. Po potwierdzeniu w jednej transakcji zapisuje identyfikator rezerwacji partnera i zmienia stan pozycji na **zarezerwowana**. Klient widzi potwierdzenie z terminem. | D-20 |
| REZ-04 | Każda próba rezerwacji jest zapisana z wynikiem, także nieudana. | Założenie |
| REZ-05 | Zajęty termin daje komunikat i listę najbliższych wolnych terminów, a pozycja zostaje dostępna. Brak miejsc daje propozycję innego zakresu dat. | D-21 |
| REZ-06 | Wywołanie partnera ma limit czasu 5 s i do 3 ponowień. Po 5 kolejnych błędach bezpiecznik wstrzymuje wywołania tego partnera na 30 s. | D-21 |
| REZ-07 | Rezerwacja o niepewnym wyniku dostaje status „do wyjaśnienia”. Zadanie cykliczne sprawdza ją u partnera, a niewyjaśnione sprawy trafiają do kolejki spraw operatora. Po wyjaśnieniu pozycja jest dostępna albo zarezerwowana, zgodnie ze stanem u partnera. | D-21, D-23 |
| REZ-08 | Dopóki rezerwacja pozycji ma status „do wyjaśnienia”, tej pozycji nie da się zarezerwować ani zablokować drugi raz (409). | Założenie (P-11) |
| REZ-09 | Posiadacz nie traci vouchera ani pozycji z powodu błędu partnera. | D-21 |
| REZ-10 | W pakiecie z noclegiem D rezerwuje pobyt w C, a C jako koordynator rezerwuje masaż w A i trening w B. A i B blokują wtedy swoje pozycje w D. | D-20, wspólny scenariusz |
| REZ-11 | To C pilnuje, żeby terminy masażu i treningu mieściły się w pobycie. D tego nie sprawdza. | Założenie (P-05) |
| REZ-12 | Rezerwację złożoną przez D anuluje operator na prośbę posiadacza. D anuluje ją w API partnera, a pozycja wraca do stanu **dostępna**. | D-03, D-20, założenie (P-15) |

## REA Realizacja

| Nr | Reguła | Źródło |
| :- | :- | :- |
| REA-01 | Realizację potwierdza przez API D tylko partner, do którego należy pozycja, po wykonaniu usługi. | D-01, D-18 |
| REA-02 | Zrealizować można własną pozycję w stanie dostępna albo zarezerwowana. Odpowiedź to 201, pozycja przechodzi w stan **zrealizowana**, powstaje wykorzystanie z datą i identyfikatorem rezerwacji partnera oraz wpis w historii vouchera. | D-18 |
| REA-03 | Pozycja już zrealizowana daje 409, pozycja innego partnera 403, voucher po terminie 410, a nieznany kod 404. W żadnym z tych przypadków dane się nie zmieniają. | D-18 |
| REA-04 | Ponowienie z tym samym kluczem idempotencji zwraca tę samą odpowiedź i nie tworzy drugiego wpisu. | D-18 |
| REA-05 | Z 20 równoległych realizacji tej samej pozycji dokładnie jedna kończy się 201, a 19 kończy się 409. | D-19 |
| REA-06 | Gdy w voucherze zostały inne pozycje, ma on status **częściowo wykorzystany**, a po realizacji ostatniej status **wykorzystany**. | D-18 |
| REA-07 | Realizację może cofnąć tylko operator, z podaniem powodu. Wykorzystanie zostaje w bazie oznaczone jako cofnięte i nie wchodzi do rozliczenia, a pozycja wraca do stanu **dostępna**. | Założenie (P-14) |

## ROZ Rozliczenia

| Nr | Reguła | Źródło |
| :- | :- | :- |
| ROZ-01 | Wszystkie kwoty są liczbami całkowitymi w groszach, w złotych (PLN). | D-01, D-05, założenie (P-08) |
| ROZ-02 | Partnerowi należy się kwota dla partnera za każdą zrealizowaną pozycję, której realizacji nie cofnięto. | D-22 |
| ROZ-03 | O miesiącu rozliczenia decyduje data realizacji w strefie Europe/Warsaw. | D-22 |
| ROZ-04 | Zestawienie zawiera kod vouchera, usługę, datę realizacji, identyfikator rezerwacji partnera i kwotę. Raport pokazuje przychód, kwoty dla A, B i C oraz marżę D. | D-22 |
| ROZ-05 | Partner pobiera przez API tylko własne zestawienie, a operator eksportuje zestawienia do CSV. Sumy w API i w CSV są równe. | D-22 |
| ROZ-06 | Wygasłe, niezrealizowane pozycje nie są rozliczane z partnerem. | Założenie (P-02) |

## UPR Uprawnienia

| Nr | Reguła | Źródło |
| :- | :- | :- |
| UPR-01 | Partner uwierzytelnia się w API D kluczem API, a D przechowuje tylko skrót klucza. Żądanie bez klucza albo z błędnym kluczem dostaje 401. | D-08 |
| UPR-02 | Operator loguje się do panelu obsługi, a sesja wygasa po 8 godzinach. Panel korzysta wyłącznie z API administracyjnego. | D-23 |
| UPR-03 | Operator zarządza partnerami i ich kluczami oraz ofertami. Przegląda też vouchery, ich historię, kolejkę spraw i rozliczenia. | D-10, D-11, D-22, D-23 |
| UPR-04 | Posiadacz vouchera nie ma konta. Dostęp do vouchera daje kod razem z adresem e-mail z zamówienia. | D-15, założenie (P-03) |

### Kto co może

✓ oznacza „może”, a „—” oznacza „nie może”.

| Operacja | Klient lub posiadacz | Partner (właściciel pozycji) | Koordynator pakietu (C) | Operator |
| :- | :-: | :-: | :-: | :-: |
| Przeglądanie katalogu | ✓ | — | — | ✓ |
| Zamówienie i płatność | ✓ | — | — | — |
| Podgląd vouchera | ✓ kod + e-mail | — | — | ✓ |
| Wybór terminu i rezerwacja przez D | ✓ | — | — | — |
| Weryfikacja kodu | — | ✓ własne pozycje | ✓ cały pakiet | — |
| Blokada i zwolnienie pozycji | — | ✓ | tylko własne | — |
| Realizacja pozycji | — | ✓ | tylko własne | — |
| Własne zestawienie przez API | — | ✓ | ✓ | — |
| Anulowanie rezerwacji, anulowanie vouchera, cofnięcie realizacji | — | — | — | ✓ |
| Oferty, partnerzy i klucze API | — | — | — | ✓ |
| Wyszukiwanie, historia, kolejka spraw | — | — | — | ✓ |
| Rozliczenia wszystkich partnerów, eksport CSV | — | — | — | ✓ |

## WSP Zasady wspólne API

| Nr | Reguła | Źródło |
| :- | :- | :- |
| WSP-01 | Czas zapisujemy w formacie RFC 3339. | D-03 |
| WSP-02 | Każdy błąd API ma kod z kontraktu, komunikat i identyfikator żądania. | D-03, D-08 |
| WSP-03 | Żądania zmieniające stan przyjmują nagłówek `Idempotency-Key`. Powtórzenie z tym samym kluczem zwraca tę samą odpowiedź i niczego nie zmienia drugi raz. | D-08 |

## ZAK Poza zakresem

| Nr | Reguła | Źródło |
| :- | :- | :- |
| ZAK-01 | Vouchery kwotowe, zwroty i zamykanie okresów rozliczeniowych są poza zakresem pierwszej wersji. | [ADR-002](../06-spotkania-i-decyzje/decyzje/ADR-002-zakres-pierwszej-wersji.md) |
| ZAK-02 | Faktury i rabaty są poza zakresem. | D-E5 |

## Otwarte pytania

Pytania do prowadzącego i partnerów zbiera koordynator (D-01 T3) i wpisuje odpowiedzi do kolumny „Odpowiedź”. Do czasu odpowiedzi obowiązuje propozycja analityka. Gdy odpowiedź jest inna, poprawiamy wskazane reguły w tym samym PR.

| Nr | Pytanie | Propozycja AN | Reguły | Kogo pytamy | Odpowiedź |
| :- | :- | :- | :- | :- | :- |
| P-01 | Czy partner może potwierdzić realizację po końcu ważności, jeśli usługa odbyła się jeszcze w terminie? | Nie. Partner potwierdza realizację w dniu usługi | REA-03, WAZ-03 | Zespoły A, B, C | |
| P-02 | Czy wygasłe, niezrealizowane pozycje są rozliczane z partnerem? | Nie | ROZ-06 | Zespoły A, B, C (D-22 T5) | |
| P-03 | Którym adresem e-mail posiadacz potwierdza dostęp do vouchera, skoro obdarowany może nie znać e-maila klienta? | E-mailem z zamówienia. Klient przekazuje obdarowanemu kod razem z tym adresem | KOD-07, UPR-04 | Zespół D, prowadzący | |
| P-04 | Co dzieje się z ofertami i wydanymi voucherami, gdy usługa partnera stanie się nieaktywna? | Oferty są wycofywane z katalogu, wydane vouchery zostają bez zmian, operator dostaje sprawę | OFE-08 | Zespoły A, B, C | |
| P-05 | Kto pilnuje, żeby masaż i trening w pakiecie mieściły się w terminie pobytu? | Hotel C | REZ-11 | Zespół C | |
| P-06 | Ile sztuk można kupić w jednym zamówieniu? | Od 1 do 10 | ZAM-03 | Zespół D | |
| P-07 | Czy operator może przedłużyć ważność vouchera? | Nie w pierwszej wersji | WAZ-06 | Prowadzący | |
| P-08 | Czy obsługujemy tylko złote? | Tak | ROZ-01 | Prowadzący | |
| P-09 | Czy koordynator pakietu (C) może blokować albo realizować pozycje A i B? (D-18 T1 pyta też o „pozycję zarezerwowaną przez innego partnera”) | Nie, tylko je widzi | BLK-07 | Zespoły A, B, C | |
| P-10 | Czy zmiana ceny usługi u partnera zmienia automatycznie kwotę dla partnera w ofertach D? | Nie, decyduje operator | SYN-05 | Zespoły A, B, C | |
| P-11 | Czy partner może zablokować pozycję, której rezerwacja w D ma status „do wyjaśnienia”? | Nie. Dostaje 409 do czasu wyjaśnienia | REZ-08 | Zespoły A, B, C | |
| P-12 | Czy pozycja oferty z liczbą 2 daje dwie pozycje vouchera? D-14 mówi o jednej pozycji vouchera na pozycję oferty, ale jednej pozycji nie da się wykorzystać dwa razy (D-05). | Tak, jedna pozycja vouchera na każdą sztukę, a kwota dla partnera jest podana za sztukę | OFE-02, KOD-04 | BE (D-01 T4) | |
| P-13 | Kto i kiedy anuluje voucher? D-02 i D-16 przewidują stan „anulowany”, ale żadna historyjka nie opisuje anulowania. | Operator w panelu. Niezrealizowane pozycje stają się anulowane, rezerwacje są anulowane u partnerów, zrealizowane pozycje są rozliczane. Zwrotu pieniędzy nie ma (ZAK-01) | POZ-07 | Zespół D, prowadzący | |
| P-14 | Czy realizację można cofnąć? D-05 mówi o „aktywnych” wykorzystaniach, a D-22 T1 o „cofniętych realizacjach”. | Tak, tylko operator i z powodem. Wykorzystanie zostaje jako cofnięte, nie wchodzi do rozliczenia, a pozycja wraca do stanu dostępna | POZ-04, REA-07, ROZ-02 | Zespoły A, B, C, D | |
| P-15 | Kto anuluje rezerwację złożoną przez D? D-03 i D-20 przewidują anulowanie w API partnerów. | Operator na prośbę posiadacza. W pierwszej wersji posiadacz nie anuluje sam | REZ-12 | Zespół D | |
