package main

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	got := sum(numbers)
	want := 36

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)

	}
}
