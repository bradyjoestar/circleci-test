package main

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Log("hello world")
}

func TestBugfix(t *testing.T) {
	fmt.Println("this is a test for bugfix")
	t.Log("bug fix")
}