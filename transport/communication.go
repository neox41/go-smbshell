package transport

import (
	"go-smbshell/config"

	"github.com/mervick/aes-everywhere/go/aes256"
)

func Decoder(encodedMessage string) (cleartextMessage string) {
	cleartextMessage = aes256.Decrypt(encodedMessage, config.Key)
	return
}

func Encoder(cleartextMessage string) (encodedMessage string) {
	encodedMessage = aes256.Encrypt(cleartextMessage, config.Key)
	return
}
