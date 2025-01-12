package adapters

import (
	"boilerplates/libs/crypto"

	"github.com/golobby/container/v3"
)

func IoCCrypto() {
	container.Singleton(func() crypto.IPasswordEncoder {
		return crypto.NewPasswordEncoder()
	})
}
