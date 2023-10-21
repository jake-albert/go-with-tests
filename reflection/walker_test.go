package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	testCases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one field",
			struct {
				Name string
			}{"Joey"},
			[]string{"Joey"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			var got []string
			walk(testCase.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, testCase.ExpectedCalls) {
				t.Errorf("got %v want %v", got, testCase.ExpectedCalls)
			}
		})
	}
}
