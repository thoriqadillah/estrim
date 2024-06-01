package common

import (
	"strconv"
	"time"
)

type Parser interface {
	String(defaults ...string) string
	Int(defaults ...int) int
	Duration(defaults ...string) time.Duration
}

type stringParser struct {
	value string
}

func ParseString(v string) Parser {
	return &stringParser{
		value: v,
	}
}

func (s *stringParser) String(defaults ...string) string {
	if s.value != "" {
		return s.value
	}

	if s.value == "" && len(defaults) > 0 {
		return defaults[0]
	}

	return ""
}

func (s *stringParser) Int(defaults ...int) int {
	n, err := strconv.Atoi(s.value)
	if err == nil {
		return n
	}

	if len(defaults) > 0 {
		return defaults[0]
	}

	return 0
}

func (s *stringParser) Duration(defaults ...string) time.Duration {
	duration, err := time.ParseDuration(s.value)
	if err == nil {
		return duration
	}

	if len(defaults) > 0 {
		duration, err := time.ParseDuration(defaults[0])
		if err != nil {
			return 0
		}

		return duration
	}

	return 0
}
