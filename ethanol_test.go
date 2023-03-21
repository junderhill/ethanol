package ethanol

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNoRoutesReturns404(t *testing.T) {
	eth := &EthanolRouter{}
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(eth.ServeHTTP)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
