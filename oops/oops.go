package oops

import "github.com/go-errors/errors"

func Wrap(err error, msg string) error {
	return errors.WrapPrefix(err, msg, 0)
}

func Err(err error) error {
	return errors.Wrap(err, 1)
}
