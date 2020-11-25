package models

import (
	"errors"
	"time"
)

type taskType string

const (
	wait taskType = "wait"
)

type task struct {
	Duration time.Time
	Type taskType
}

func (t task) validate() error {
	switch t.Type {
	case wait:
	default:
		return errors.New("invalid task type")
	}
	return nil
}