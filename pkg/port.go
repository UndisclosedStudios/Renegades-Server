package pkg

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"runtime"
	errhandler "undisclosedstudios.infrandom.net/renegades/gameservers/go/internal/error"
)

type File struct {
	file []byte
	err  error
}

func GetAvailablePort() {
	var file File
	switch os := runtime.GOOS; os {
	case "windows":
		panic("Not yet implemented")
	case "linux":
		file.file, file.err = ioutil.ReadFile("/proc/net/tcp")
	}
	errhandler.Check(file.err, "An error occured while trying to fetch an available port")
	log.Print(file)
}
