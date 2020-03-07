/*
A simple package which provides tools to help handling the errors.

It can be used anywhere inside the cmd/ and internal/ directories (since it is inside the internal/ folder).
*/
package error

import log "github.com/sirupsen/logrus"

func Check(err error, params ...string) {
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"Message": params,
		})
		panic(err)
	}
}
