package regression

import "log"

func PanicCheck(err interface{}) {
	if err != nil {
		log.Panic(err)
	}
}
