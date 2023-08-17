package passwordhasher

import (
	"github.com/matthewhartstonge/argon2"
)

var argon = argon2.Config{
	HashLength:  40,
	SaltLength:  16,
	TimeCost:    8,
	MemoryCost:  64 * 1024,
	Parallelism: 2,
	Mode:        argon2.ModeArgon2id,
	Version:     19,
}

func HashPassword(p string) string {
	hashedPassword, err := argon.HashEncoded([]byte(p))

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
