package main

import "testing"

func TestHelloWorld(t *testing.T) {
    want := "Hello, World!"
    got := HelloWorld()
    if got != want {
        t.Errorf("HelloWorld() 1  = %v, want %v", got, want)
    }
}

func HelloWorld() string {
    return "Hello, World!"
}