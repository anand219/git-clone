package random

import (
	"fmt"
	"sync/atomic"
)

var (
	uidGen uint64
)

// Generator struct
type Generator struct {
	uid uint64
}

// New creates a new generator
func New() *Generator {
	return &Generator{
		uid: atomic.AddUint64(&uidGen, 1),
	}
}

// Name returns a name
func (g *Generator) Name() string {
	return fmt.Sprintf("User %d", g.uid)
}

// Company returns the company name
func (g *Generator) Company() string {
	return fmt.Sprintf("Company %d PVT LTD", g.uid)
}

// Email returns the email
func (g *Generator) Email() string {
	return fmt.Sprintf("email%d@yopmail.com", g.uid)
}

// Password generates a password for the i the time
func (g *Generator) Password(i uint64) string {
	return fmt.Sprintf("Password%d%d", g.uid, i)
}
