package jsontoenv

import (
	"encoding/json"
	"os"
	"strings"
)

// Opts is a struct that contains options for the environment
type Opts struct {
	UseUpperCase bool
}

type Env interface {
	// FromBytes reads a JSON object and sets the environment variables
	FromBytes(data []byte) error
	// Omit is a list of keys to omit from the environment
	OmitKeys(keys ...string)
}

type env struct {
	omit []string
	opts Opts
}

// New creates a new instance of the Env interface
// Opts is a struct that contains options for the environment
func New(opts Opts) Env {
	return &env{
		opts: opts,
	}
}

// OmitKeys Omit is a list of keys to omit from the environment
func (e *env) OmitKeys(keys ...string) {
	e.omit = append(e.omit, keys...)
}

// IsOmitted checks if a key is omitted from the environment
func (e *env) IsOmitted(key string) bool {
	for _, k := range e.omit {
		if k == key {
			return true
		}
	}
	return false
}

// FromBytes reads a JSON object and sets the environment variables
func (e *env) FromBytes(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	for key, value := range raw {
		if e.IsOmitted(key) {
			continue
		}

		values, err := parseValuesWithObjects(key, value)
		if err != nil {
			return err
		}

		for k, v := range values {
			if e.opts.UseUpperCase {
				k = strings.ToUpper(k)
			}
			if err := os.Setenv(k, v); err != nil {
				return err
			}
		}
	}

	return nil
}
