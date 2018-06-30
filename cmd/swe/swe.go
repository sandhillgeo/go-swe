package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

import (
	//"github.com/pkg/errors"
)

import (
	"github.com/sandhillgeo/go-swe/swe"
)

var GO_SWE_VERSION = "0.0.1"

func main() {

	start := time.Now()

	var output_uri string
	//var verbose bool
	var version bool
	var help bool

	flag.StringVar(&output_uri, "output_uri", "", "The output uri expression to evaulate")

	flag.BoolVar(&version, "version", false, "Prints version to stdout")
	flag.BoolVar(&help, "help", false, "Print help")

	flag.Parse()

	if help {
		fmt.Println("Usage: swe -output_uri OUTPUT [-version] [-help]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	} else if len(os.Args) == 1 {
		fmt.Println("Error: Provided no arguments.")
		fmt.Println("Run \"dfl -help\" for more information.")
		os.Exit(0)
	} else if len(os.Args) == 2 && os.Args[1] == "help" {
		fmt.Println("Usage: swe -output_uri OUTPUT [-version] [-help]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if version {
		fmt.Println(GO_SWE_VERSION)
		os.Exit(0)
	}

	sgpkg := swe.NewSensorGeoPackage(output_uri);

	err := sgpkg.Init()
	if err != nil {
		panic(err)
	}

	sgpkg.AutoMigrate()

	err = sgpkg.Close()
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	fmt.Println("Done in " + elapsed.String())
}
