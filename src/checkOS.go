package app

import (
	"log"
	"os"
	"runtime"
)

func checkOStype() {
	if runtime.GOOS == "linux" {
		os.Setenv("QT_QPA_PLATFORM", "xcb") // Sets x11 for linux
		log.Println("Running on Linux")
	} else if runtime.GOOS == "windows" {
		log.Println("Running on Windows")
	} else if runtime.GOOS == "darwin" {
		log.Println("Running on MacOS")
	} else {
		log.Println("System type cannot be detected")
	}
}
