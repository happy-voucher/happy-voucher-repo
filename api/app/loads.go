package app

import (
	"context"
	"crypto/sha256"
	"errors"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/Piechutowski/volt"
	"github.com/Piechutowski/volt/nao/rt"

	"github.com/happy-voucher/happy-voucher-repo/db"
)

var FirebaseAuth *auth.Client

func bearerToken(r volt.Request) string {
	token, ok := strings.CutPrefix(r.Header().Get("Authorization"), "Bearer ")
	if !ok {
		return ""
	}
	return strings.TrimSpace(token)
}

func CurrentPartner(r volt.Request) volt.Load[db.Partner] {
	key := bearerToken(r)
	if key == "" {
		return volt.Halt[db.Partner](Err(r, http.StatusUnauthorized, "UNAUTHENTICATED", "brak klucza API"))
	}
	hash := sha256.Sum256([]byte(key))
	return volt.LoadWith(db.PartnerGetByAPIKeyHash(hash[:])).Rescue(func(err error) volt.Response {
		if errors.Is(err, rt.ErrNotFound) {
			return Err(r, http.StatusUnauthorized, "UNAUTHENTICATED", "błędny klucz API")
		}
		return volt.Fail(err)
	})
}

func FirebaseUser(r volt.Request) (*auth.Token, error) {
	token := bearerToken(r)
	if token == "" {
		return nil, errors.New("brak tokenu")
	}
	return FirebaseAuth.VerifyIDToken(context.Background(), token)
}

func CurrentStaff(r volt.Request) volt.Load[db.Staff] {
	user, err := FirebaseUser(r)
	if err != nil {
		return volt.Halt[db.Staff](Err(r, http.StatusUnauthorized, "UNAUTHENTICATED", "zaloguj się"))
	}
	return volt.LoadWith(db.StaffGet(user.UID)).Rescue(func(err error) volt.Response {
		if errors.Is(err, rt.ErrNotFound) {
			return Err(r, http.StatusForbidden, "FORBIDDEN", "konto nie ma roli operator ani admin")
		}
		return volt.Fail(err)
	})
}
