package main

import (
	"regexp"
	"testing"
)

func TestCleanUserInput(t *testing.T) {
	inputs := []string{
		"help ",
		"  help",
		"help",
	}

	want := regexp.MustCompile(`^help$`)

	for _, input := range inputs {
		msg := clean(input)
		if !want.MatchString(msg) {
			t.Fatalf(`clean(`+input+`) = %q, want match for %#q`, msg, want)
		}
	}
}
