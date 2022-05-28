package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
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

// *** Mock tests *** //

// smsServiceMock
type smsServiceMock struct {
	mock.Mock
}

// Our mocked smsService method
func (m *smsServiceMock) SendChargeNotification(value int) error {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", value)
	// this records that the method was called and passes in the value
	// it was called with
	_ = m.Called(value)
	// it then returns whatever we tell it to return
	// in this case true to simulate an SMS Service Notification
	// sent out
	return nil
}

// we need to satisfy our MessageService interface
// which sadly means we have to stub out every method
// defined in that interface
func (m *smsServiceMock) DummyFunc() {
	fmt.Println("Dummy")
}

// TestChargeCustomer is where the magic happens
// here we create our SMSService mock
func TestChargeCustomer(t *testing.T) {
	smsService := new(smsServiceMock)

	// we then define what should be returned from SendChargeNotification
	// when we pass in the value 100 to it. In this case, we want to return
	// true as it was successful in sending a notification
	smsService.On("SendChargeNotification", 100).Return(true)

	// next we want to define the service we wish to test
	myService := MyService{smsService}
	// and call said method
	myService.ChargeCustomer(100)

	// at the end, we verify that our myService.ChargeCustomer
	// method called our mocked SendChargeNotification method
	smsService.AssertExpectations(t)
}
