package test

import (
	"example/line-bot-ledger/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterWithSuccess(t *testing.T, NODE_ENV string) {
	r := api.Router(NODE_ENV)
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/admin")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Status code for /home is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	res, err = http.Get(ts.URL + "/chatbot")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Status code for /home is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	res, err = http.Get(ts.URL + "/dishes")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Status code for /home is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	res, err = http.Get(ts.URL + "/line")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Status code for /home is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}
