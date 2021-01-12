// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/dialect"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver used for executing database requests.
	driver dialect.Driver
	// debug enable a debug logging.
	debug bool
	// log used for logging on debug mode.
	log func(...interface{})
	// hooks to execute on mutations.
	hooks *hooks
}

// hooks per client, for fast access.
type hooks struct {
	Bloodtype       []ent.Hook
	Certificate     []ent.Hook
	CoveredPerson   []ent.Hook
	Department      []ent.Hook
	Diagnose        []ent.Hook
	Disease         []ent.Hook
	Doctor          []ent.Hook
	Drug            []ent.Hook
	Examinationroom []ent.Hook
	Fund            []ent.Hook
	Gender          []ent.Hook
	Medical         []ent.Hook
	Nurse           []ent.Hook
	Operative       []ent.Hook
	Operativerecord []ent.Hook
	Patient         []ent.Hook
	Prefix          []ent.Hook
	Prescription    []ent.Hook
	Rent            []ent.Hook
	Room            []ent.Hook
	Roomtype        []ent.Hook
	SchemeType      []ent.Hook
	Tool            []ent.Hook
}

// Options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...interface{})) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}