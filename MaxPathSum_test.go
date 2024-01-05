package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleMaxPathSum(t *testing.T) {
	testCases := []struct {
		Name       string
		Input      string
		Expected   int
		StatusCode int
	}{
		{
			Name: "Test Case 1",
			Input: `{
				"nodes": [
					{"id": "1", "left": "2", "right": "3", "value": 1},
					{"id": "3", "left": "6", "right": "7", "value": 3},
					{"id": "7", "left": null, "right": null, "value": 7},
					{"id": "6", "left": null, "right": null, "value": 6},
					{"id": "2", "left": "4", "right": "5", "value": 2},
					{"id": "5", "left": null, "right": null, "value": 5},
					{"id": "4", "left": null, "right": null, "value": 4}
				],
				"root": "1"
			}`,
			Expected:   18,
			StatusCode: http.StatusOK,
		},
		
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/maxPathSum", bytes.NewBufferString(tc.Input))
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()
			handleMaxPathSum(rec, req)

			if rec.Code != tc.StatusCode {
				t.Errorf("Expected status code %d, but got %d", tc.StatusCode, rec.Code)
			}

			var result MaxPathSumResponse
			err = json.Unmarshal(rec.Body.Bytes(), &result)
			if err != nil {
				t.Fatal(err)
			}

			if result.MaxPathSum != tc.Expected {
				t.Errorf("Expected maxPathSum %d, but got %d", tc.Expected, result.MaxPathSum)
			}
		})
	}
}
