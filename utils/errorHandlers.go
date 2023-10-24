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

func HandleFunctionPanic[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func HandleFunctionFetal[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
