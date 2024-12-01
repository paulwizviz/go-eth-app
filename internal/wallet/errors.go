package wallet

import "errors"

var (
	ErrKeystoreReadFile    = errors.New("reading keystore file error")
	ErrKeystoreDecrypt     = errors.New("decrypt keystore error")
	ErrKeystoreGenerateKey = errors.New("generating key in keystore error")
	ErrKeystoreCreate      = errors.New("create keystore error")
)
