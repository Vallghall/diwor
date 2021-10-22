package experiment

import (
	"encoding/hex"

	ini "gitlab.com/Valghall/diwor/internal/initial_data"
)

type experiment struct {
	initialData      *ini.InitialData
	cryptographyData *ini.CryptographyData
}

func (e *experiment) InitialData() *ini.InitialData {
	return e.initialData
}

func (e *experiment) SetInitialData(initialData *ini.InitialData) {
	e.initialData = initialData
}

func (e *experiment) EncryptedA() string {
	return e.cryptographyData.EncryptedA
}

func (e *experiment) EncryptedB() string {
	return e.cryptographyData.EncryptedB
}

/*
func (e *experiment) Decrypted() string {
	return e.decrypted
}
*/

func (e *experiment) SetEncrypted(encryptedA, encryptedB []byte) {
	e.cryptographyData = new(ini.CryptographyData)
	e.setEncryptedA(encryptedA)
	e.setEncryptedB(encryptedB)
}

func (e *experiment) setEncryptedA(encrypted []byte) {
	e.cryptographyData.EncryptedA = hex.EncodeToString(encrypted)
}

func (e *experiment) setEncryptedB(encrypted []byte) {
	e.cryptographyData.EncryptedB = hex.EncodeToString(encrypted)
}

func (e *experiment) SetDecrypted(decrypted []byte) {
	e.cryptographyData.EncryptedB = hex.EncodeToString(decrypted)
}
