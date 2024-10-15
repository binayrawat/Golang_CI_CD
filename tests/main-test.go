package main

import "testing"

func TestMain(t *testing.T) {
    // This is a simple test
    if 1 != 1 {
        t.Errorf("Expected 1 to equal 1")
    }
}
