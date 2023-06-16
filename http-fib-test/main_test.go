package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleFib(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want []byte
	}{
		{name: "zero", num: 0, want: []byte("0")},
		{name: "one", num: 1, want: []byte("1")},
		{name: "two", num: 2, want: []byte("1")},
	}

	handler := http.HandlerFunc(handleFib)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/fib?num=%d", tc.num), nil)
			handler.ServeHTTP(rec, req)
			assert.Equal(t, tc.want, rec.Body.Bytes())
		})
	}
}
