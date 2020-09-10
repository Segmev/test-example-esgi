package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMultiplyByTwo(t *testing.T) {
	channelIn, channelOut := make(chan int), make(chan int)
	go multiplyByTwo(channelIn, channelOut)
	channelIn <- 23
	res := <-channelOut
	if res != (23 * 2) {
		t.Errorf("multiplyByTwo expected %d, got: %d", 23, res)
	}

	channelIn <- 12
	res = <-channelOut
	if res != (12 * 2) {
		t.Errorf("multiplyByTwo expected %d, got: %d", 23, res)
	}
}

func TestHandler(t *testing.T) {
	channelIn, channelOut := make(chan int), make(chan int)
	go multiplyByTwo(channelIn, channelOut)

	req, err := http.NewRequest("GET", "/data", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(generateHandler(channelIn, channelOut))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"value": 4}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
