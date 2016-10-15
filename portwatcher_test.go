package main

import (
	"reflect"
	"testing"
)

func TestSimple(t *testing.T) {
	result, err := expand("100")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, []int{100}) {
		t.Fatal("Unexpected result: %v", result)
	}
}

func TestComplex(t *testing.T) {
	result, err := expand("1,2,5-8,15")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, []int{1, 2, 5, 6, 7, 8, 15}) {
		t.Fatal("Unexpected result: %v", result)
	}
}

func TestEdgeFormatting(t *testing.T) {
	result, err := expand("1, 2  ,    5 -  8,15")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, []int{1, 2, 5, 6, 7, 8, 15}) {
		t.Fatal("Unexpected result: %v", result)
	}
}

func TestErrors(t *testing.T) {
	_, err := expand("1a")
	if err == nil {
		t.Fatal("No error")
	}

	_, err = expand("1a-2b")
	if err == nil {
		t.Fatal("No error")
	}
}
