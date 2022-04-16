package main

import "testing"

func TestSayHell(t *testing.T) {
	got := sayHello()
	want := "Hello World"

	if got != want {
		t.Errorf("Got %s, but want %s", got, want)
	}
}
