package utils

import (
	"bytes"
	"encoding/gob"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var logFn = log.Panic

func FromBytes(i interface{}, data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(decoder.Decode(i))
	//decoder.Decode(i)
}

func HandleErr(err error) {
	if err != nil {
		logFn(err)
	}
}

func Hash(value string, cost int) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(value), cost)
	return string(passwordHash), err
}

func CheckPassword(password1 string, password2 string) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(password1),
		[]byte(password2),
	)

	return err
}
