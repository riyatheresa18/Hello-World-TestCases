package main

import (
	"go/types"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandler(t *testing.T) {
	//test case for checking string hello world
	req, err := http.NewRequest("GET", "/getText/{text}", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(handler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestEmptyString(t *testing.T) {
	//test case to check whether the string is empty or not
	req, _ := http.NewRequest("	GET", "/getText/{text}", nil)
	//response := executeRequest(req)
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(handler)
	hf.ServeHTTP(recorder, req)
	if body := recorder.Body.String(); body == "" {
		t.Errorf("Got an empty string.  %s", body)
	}

}

func TestTypeOfString(t *testing.T) {
	//test case to check type of string
	req, _ := http.NewRequest("GET", "/getText/{text}", nil)
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(handler)
	hf.ServeHTTP(recorder, req)
	body := recorder.Body.String()
	if reflect.ValueOf(body).Kind() != reflect.Kind(types.String) {
		t.Errorf("Invalid Input  %s", body)
	}
}
