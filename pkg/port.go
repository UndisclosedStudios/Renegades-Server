package pkg

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"runtime"
	errhandler "undisclosedstudios.infrandom.net/renegades/gameservers/go/internal/error"
)

//File Represent a file red by the system
type File struct {
	file []byte //Content of the file
	err  error //Errors while trying to read the files
}


 // GetAvailablePort returns an available port to the caller as an int
 // - Linux:
 // Read the /proc/net/tcp file and grab a free port to bind on
 // - Windows:
 // TODO: Determine a way to get a free port
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
