package lib

import "log"

func LogFatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
