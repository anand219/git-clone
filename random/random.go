package random

import (
	cryptoRand "crypto/rand"
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
	return fmt.Sprintf("User %d-%s", g.uid, GenerateRandomBytes(16))
}

// Company returns the company name
func (g *Generator) Company() string {
	return fmt.Sprintf("Company %d-%s PVT LTD", g.uid, GenerateRandomBytes(16))
}

// Email returns the email
func (g *Generator) Email() string {
	return fmt.Sprintf("email-%d-%s-@yopmail.com", g.uid, GenerateRandomBytes(16))
}

// Password generates a password for the i the time
func (g *Generator) Password(i uint64) string {
	return fmt.Sprintf("Password%d%d", g.uid, i)
}

func (g *Generator) PhoneNumber(i uint64) string {
	return fmt.Sprintf("%d-%s", i, GenerateRandomBytes(9))
}

func GenerateRandomBytes(byteCount int) string {

	b := make([]byte, byteCount)
	if _, err := cryptoRand.Read(b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)
}
