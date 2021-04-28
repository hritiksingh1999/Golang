package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	record := httptest.NewRecorder()
	handler := http.HandlerFunc(users)
	handler.ServeHTTP(record, req)
	if status := record.Code; status != http.StatusOK {
		t.Errorf("handler returned worng status code : got %v want %v", status, http.StatusOK)
	}
	expected := `"hritik"`
	if reflect.TypeOf(record.Body.String()) != reflect.TypeOf(expected) {
		t.Errorf("handler returned  unexpected body : got %v want %vuh", reflect.TypeOf(record.Body.String()), reflect.TypeOf(expected))
	}
}

func TestGetpost(t *testing.T) {
	req, err := http.NewRequest("GET", "/getpost/{hritik}", nil)
	if err != nil {
		t.Fatal(err)
	}

	record := httptest.NewRecorder()
	t.Error("Record", record.Body)

	handler := http.HandlerFunc(getpost)
	handler.ServeHTTP(record, req)
	if status := record.Code; status != http.StatusOK {
		t.Errorf("handler returned worng status code : got %v want %v", status, http.StatusOK)
	}
	expected := `user :hritik
	post  :-hey there 
	post  :-hey there how are you 
	post  :-hey there how are you `
	if record.Body.String() != expected {
		t.Errorf("handler returned unexpected body : got %v want %v", record.Body.String(), reflect.ValueOf(expected))
	}
}
