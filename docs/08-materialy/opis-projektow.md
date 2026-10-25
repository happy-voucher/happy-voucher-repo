[← Materiały](README.md)

# Opis projektów (materiał od prowadzącego)

> **Opiekun:** Jakub Barej (DO) · **Źródło:** opis przekazany przez prowadzącego, przepisany bez zmian treści · **Dodano:** 2026-10-25


## Projekt A. Obsługa firmy świadczącej usługi masażu

### Oferta usług i cennik
Zarządzanie rodzajami masaży, czasem trwania, cenami i opisami usług. Określanie, którzy pracownicy wykonują daną usługę i jakiego pomieszczenia ona wymaga.

### Pracownicy, gabinety i dostępność
Prowadzenie grafików pracowników oraz dostępności gabinetów. Uwzględnianie przerw między wizytami, dni wolnych i czasowej niedostępności pomieszczeń.

### Rezerwacje wizyt
Wyszukiwanie wolnych terminów, zapisywanie klientów, przekładanie i anulowanie wizyt. Zapobieganie jednoczesnej rezerwacji tego samego pracownika lub gabinetu.

### Obsługa klientów i wykonanych usług
Kartoteka klientów, historia wizyt i statusy realizacji. Oznaczanie wizyty jako wykonanej lub nieodbytej oraz zestawienia wykorzystania czasu pracowników i przychodów z usług.

### Własne API i współpraca
API udostępnia ofertę, wolne terminy i rezerwacje. Gabinet przyjmuje zapisy od B, C i D oraz korzysta z API D do sprawdzania i rozliczania voucherów.

---

## Projekt B. Obsługa studia fitness i treningów personalnych

### Oferta zajęć i treningów
Katalog zajęć grupowych, treningów indywidualnych i karnetów. Określanie poziomu zajęć, czasu trwania, ceny oraz limitu uczestników.

### Grafiki trenerów i sal
Planowanie pracy trenerów oraz wykorzystania sal. Kontrola nakładających się zajęć i prezentacja dostępnych terminów treningów indywidualnych.

### Zapisy i karnety
Rejestrowanie uczestników, zapisów i obecności. Kontrola ważności karnetu, liczby pozostałych wejść oraz obsługa listy rezerwowej na zajęcia.

### Oferta łączona z usługą masażu
Sprzedaż pakietu obejmującego trening i masaż u partnera A. Wyszukiwanie pasujących terminów obu usług oraz przechowywanie potwierdzeń rezerwacji partnera.

### Własne API i współpraca
API udostępnia zajęcia, terminy i operacje zapisu. Studio korzysta z API A do rezerwacji masażu, z D do obsługi voucherów i przyjmuje rezerwacje od hotelu C.

---

## Projekt C. Obsługa hotelu i pakietów wypoczynkowych

### Oferta pokoi i pobytów
Ewidencja pokoi, ich wyposażenia i cen. Wyszukiwanie dostępności dla wybranego okresu oraz obliczanie kosztu pobytu.

### Rezerwacje hotelowe
Przyjmowanie, zmiana i anulowanie rezerwacji. Rejestrowanie gości, liczby osób i statusu pobytu oraz kontrola dostępności pokoi.

### Pakiety wypoczynkowe
Tworzenie ofert łączących nocleg z masażem w A lub zajęciami w B. Określanie zawartości pakietu, ceny i liczby przysługujących usług.

### Koordynacja usług partnerów
Rezerwowanie masażu i treningu w terminach przypadających podczas pobytu. Prezentacja planu pobytu oraz obsługa sytuacji, gdy partner nie może potwierdzić wybranej usługi.

### Własne API i współpraca
API udostępnia ofertę, dostępność i rezerwacje pobytów. Hotel korzysta z API A i B do zamawiania usług oraz z D do weryfikacji voucherów na pobyt lub pakiet.

---

## Projekt D. Sprzedaż i obsługa voucherów

### Katalog usług partnerów
Pobieranie ofert gabinetu A, studia B i hotelu C. Publikowanie voucherów na konkretną usługę lub pakiet oraz powiązanie każdej oferty z właściwym usługodawcą.

### Sprzedaż i wydawanie voucherów
Obsługa zamówień klientów i symulowanego potwierdzenia płatności. Generowanie unikalnego kodu, określanie wartości lub zakresu usług oraz terminu ważności.

### Rezerwacja usługi u partnera
Sprawdzanie dostępności i rezerwowanie usługi przez API A, B lub C. Zapisywanie identyfikatora rezerwacji oraz prezentacja klientowi jej potwierdzenia.

### Realizacja i kontrola wykorzystania
Sprawdzanie ważności voucheru i rejestrowanie jego wykorzystania. Obsługa voucherów jednorazowych lub obejmujących kilka usług oraz zapobieganie ponownemu wykorzystaniu tej samej usługi.

### Własne API i rozliczenia partnerów
API umożliwia weryfikację kodu, potwierdzenie realizacji i sprawdzenie pozostałych usług. Firma przygotowuje zestawienia wykorzystanych voucherów oraz kwot do rozliczenia z A, B i C.

---

## Wspólny scenariusz działania

1. Klient kupuje w D voucher na pakiet: nocleg, masaż i trening.
2. D sprawdza terminy i rezerwuje pobyt w C.
3. C uzgadnia z klientem plan pobytu i rezerwuje masaż w A oraz trening w B.
4. A i B potwierdzają rezerwacje i blokują odpowiednie terminy.
5. Przy realizacji usługi każda firma sprawdza voucher przez API D.
6. Po wykonaniu usługi partner aktualizuje własną rezerwację i potwierdza wykorzystanie odpowiedniej części voucheru w D.
7. D prezentuje klientowi pozostałe usługi i przygotowuje zestawienie rozliczeń z partnerami.

Każda firma obsługuje również własnych klientów, zarządza własną ofertą i przechowuje dane w swojej bazie.

---

## Wymagania integracji

- Każdy projekt udostępnia własne API i korzysta z API co najmniej jednego innego zespołu.
- Każdy zespół przechowuje dane, za które odpowiada, w swojej bazie.
- Dostęp do danych partnerów odbywa się przez API.
- Zespoły uzgadniają wspólne identyfikatory, formaty danych i sposób zgłaszania błędów.
- API dostępne online mogą rozszerzać funkcje systemu.
- Testy obejmują poprawny przebieg oraz konflikt rezerwacji, brak miejsc i niedostępność partnera.
