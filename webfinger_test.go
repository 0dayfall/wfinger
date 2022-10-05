package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	setup()
}

func TestFinger(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/.well-known/webfinger?resource=acct%3A%20%40test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	finger(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Subject":"acct: @test","Aliases":["pelle@pelle.com","www.facebook.com/pelle"],"Properties":{"@anyone":"GPS: 76.3434334,32.345435","@shiela":"Will be going from work at 16:30","GPS":"GPS 76.3434334,32.345435","job":"Male - emeperor","status":"On vacation in Philippines"},"Links":[{"Rel":"","Href":"http://www.substack.com/blog/pelle","Titles":{},"Properties":{}},{"Rel":"","Href":"spotify://link.to.pod","Titles":{},"Properties":{}},{"Rel":"","Href":"insta://link.to.pod","Titles":{},"Properties":{}}]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
