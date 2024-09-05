package utils

import (
	"bytes"
	"encoding/gob"
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
