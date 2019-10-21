package main

import (
	"fmt"
	"os"
)

const (
	banner = `
                 _     _             
  ___  __ _  ___| |__ | |_ _ __ __ _ 
 / _ \/ _' |/ __| '_ \| __| '__/ _' |
|  __/ (_| | (__| | | | |_| | | (_| |
 \___|\__,_|\___|_| |_|\__|_|  \__,_|`

	version = "0.0.1 - Just getting started"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
