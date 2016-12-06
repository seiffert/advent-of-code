package main

import "testing"

func TestGetPassword(t *testing.T) {
	if pwd := GetPassword("abc"); pwd != "05ace8e3" {
		t.Errorf("Invalid password: %q", pwd)
	}
}
