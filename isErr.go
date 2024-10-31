package main

import "log"

func IsErrNo(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
