package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddTwo(t *testing.T) {
	if AddTwo(2) != 4 {
		t.Error("expected 2+2 to equal 4")
	}
}

func TestTableAddTwo(t *testing.T) {
	var tests = []struct {
		input    float64
		expected float64
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{99, 101},
	}

	for _, test := range tests {
		if output := AddTwo(test.input); output != test.expected {
			t.Errorf("input: %f, result: %f, expected: %f", test.input, output, test.expected)
		}
	}
}

type AddResult struct {
	x        float64
	y        float64
	expected float64
}

func TestTableAdd(t *testing.T) {
	tests := []AddResult{
		{2, 3, 5},
		{1.2, 2.1, 3.3},
		{-2, -5, -7},
		{1, 0, 1},
		{9, 9, 18},
	}

	for _, test := range tests {
		if output := Add(test.x, test.y); output != test.expected {
			t.Errorf("inputs: %f, %f, result: %f, expected: %f", test.x, test.y, output, test.expected)
		}
	}
}

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("test.data")
	if err != nil {
		t.Error("could not open file")
	}
	if string(data) != "hello world" {
		t.Error("string content do not match expected")
	}
}

func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		payload, _ := json.Marshal(map[string]string{
			"is_it_mocked_response": "yes",
		})
		io.WriteString(w, string(payload))
	}
	req := httptest.NewRequest("GET", "http://google.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if resp.StatusCode != 200 {
		t.Error("status code is not 200, it is: ", resp.StatusCode)
	}
}
