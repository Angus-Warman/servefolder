package main

import (
	"os"

	"github.com/Angus-Warman/httpmin"
)

func main() {
	key := "FOLDER"

	c := httpmin.New()

	if os.Getenv("HTTPS") != "" {
		c.UseSelfSignedHTTPSFromFolder("tls")
	}

	if os.Getenv("PUBLIC") != "" {
		c.PublicIP()
	}

	c.DefaultEnvVar(key, "./").
		ServeFolder(os.Getenv(key)).
		Run()
}
