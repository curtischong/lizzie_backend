package main

import (
	"bytes"
	"encoding/json"
	is "github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCardsAndPanelsCall(t *testing.T) {

	is := is.New(t)
	d := GetCardsAndPanelsObj{
		CardAmount:  3,
		CardOffset:  0,
		PanelAmount: 2,
		PanelOffset: 0,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(d)
	is.NoErr(err) // json.NewEncoder

	srv := server{
		db:    mockDatabase,
		email: mockEmailSender,
	}

	got, err := http.NewRequest(http.MethodPost, "/greet", &buf)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	is.Equal(w.StatusCode, http.StatusOK)
	want := "asd"
	is.NoErr(err)

}
