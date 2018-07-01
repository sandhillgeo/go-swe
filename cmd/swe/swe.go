package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/sandhillgeo/go-swe/swe"
)

var GO_SWE_VERSION = "0.0.1"

func printUsage() {
	fmt.Println("Usage: swe -output_uri OUTPUT_URI [-version] [-help]")
}

func main() {

	start := time.Now()

	var output_uri string
	//var verbose bool
	var version bool
	var help bool

	flag.StringVar(&output_uri, "output_uri", "", "The output uri of the sensor GeoPackage.")

	flag.BoolVar(&version, "version", false, "Prints version to stdout")
	flag.BoolVar(&help, "help", false, "Print help")

	flag.Parse()

	if help {
		printUsage()
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	} else if len(os.Args) == 1 {
		fmt.Println("Error: Provided no arguments.")
		fmt.Println("Run \"swe -help\" for more information.")
		os.Exit(0)
	} else if len(os.Args) == 2 && os.Args[1] == "help" {
		printUsage()
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if version {
		fmt.Println(GO_SWE_VERSION)
		os.Exit(0)
	}

	if len(output_uri) == 0 {
		fmt.Println("Error: Provided no -output_uri.")
		fmt.Println("Run \"swe -help\" for more information.")
		os.Exit(1)
	}

	sgpkg := swe.NewSensorGeoPackage(output_uri)

	err := sgpkg.Init()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Error initializing sensor GeoPackage"))
		os.Exit(1)
	}

	err = sgpkg.AutoMigrate()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Error auto migrating sensor GeoPackage"))
		os.Exit(1)
	}

	err = sgpkg.Close()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Error closing sensor GeoPackage"))
		os.Exit(1)
	}

	elapsed := time.Since(start)
	fmt.Println("Done in " + elapsed.String())
}
