package main

import "testing"

func TestConn(t *testing.T) {
	if ans := Conn(); ans != nil {
		t.Error("error")
	}
}