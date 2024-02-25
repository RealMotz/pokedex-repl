package main

import (
	"errors"
)

func exitCommand() error {
	return errors.New("EOF")
}
