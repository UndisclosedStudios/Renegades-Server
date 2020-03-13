/*
A simple package which provides tools to help handling the errors.

It can be used anywhere inside the cmd/ and internal/ directories (since it is inside the internal/ folder).
*/
package error

import (
  "github.com/getsentry/sentry-go"
  log "github.com/sirupsen/logrus"
)

func Check(err error, params ...string) {
  if err != nil{
    log.WithError(err).WithFields(log.Fields{
      "Message": params,
    })
    sentry.CaptureException(err)
  }
  panic(err)
}

func Checks(err []error, params ...string) {
  if err != nil{
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
