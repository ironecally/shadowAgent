package shadowAgent

import "log"

type errType struct {
	err error
	job Job
}

func errHandler(errObj errType) {
	log.Printf("[ShadowAgent] found error %+v", errObj.err)
}
