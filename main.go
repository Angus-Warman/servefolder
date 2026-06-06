package main

import (
	"os"

	"github.com/Angus-Warman/httpmin"
)

func main() {
	key := "FOLDER"
	httpmin.Setup().DefaultEnvVar(key, "./").ServeFolder(os.Getenv(key)).Run()
}
