//Package error contains a bunch of methods to help with errors handling
package error

import (
  "github.com/getsentry/sentry-go"
  log "github.com/sirupsen/logrus"
)

//Check handle the processing of errors by making usage of a third-party logger and sentry.
//err must be passed the error itself and params some optionals string to be printed in the console.
func Check(err error, params ...string) {
  if err != nil {
    log.WithError(err).WithFields(log.Fields{
      "Message": params,
    })
    sentry.CaptureException(err)
  }
  panic(err)
}

//Checks handle the processing of errors by making usage of a third-party logger and sentry.
//err must be passed the errors and params some optionals string to be printed in the console.
func Checks(err []error, params ...string) {
  if err != nil {
    switch params {
    case nil:
      for i := 0; i < len(err); i++ {
        log.WithError(err[i])
        sentry.CaptureException(err[i])
      }
    default:
      for i := 0; i < len(err); i++ {
        log.WithError(err[i]).WithFields(log.Fields{
          "Message": params[i],
        })
        sentry.CaptureException(err[i])
      }
    }
  }
}
