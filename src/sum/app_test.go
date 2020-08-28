package main

import "testing"

func TestSumFive(t *testing.T) {
    got := sum(5, 5)
    if got != 10 {
        t.Errorf("sum(5, 5) = %d; want 10", got)
    }
}

func TestSumOne(t *testing.T) {
    got := sum(1, 1)
    if got != 2 {
        t.Errorf("sum(1, 1) = %d; want 2", got)
    }
}
