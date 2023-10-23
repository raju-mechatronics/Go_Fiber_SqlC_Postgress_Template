package utils

import "log"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func FetalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
