package auth

import (
	"fmt"
	"runtime"

	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) (string, error) {
	params := &argon2id.Params{
		Memory:      128 * 1024,
		Iterations:  4,
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := argon2id.CreateHash(password, params)
	if err != nil {
		return "", fmt.Errorf("couldn't hash password: %w", err)
	}

	return hash, nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, fmt.Errorf("invalid hash: %w", err)
	}

	return match, nil
}
