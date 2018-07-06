[![Build Status](https://travis-ci.org/sandhillgeo/go-swe.svg)](https://travis-ci.org/sandhillgeo/go-swe) [![GoDoc](https://godoc.org/github.com/sandhillgeo/go-swe?status.svg)](https://godoc.org/github.com/sandhillgeo/go-swe)

# go-swe

# Description

**go-swe** is a Go library supporting the OGC's [Sensor Web Enablement](http://www.opengeospatial.org/ogc/markets-technologies/swe) (SWE) standards.

# Usage

**Go**

You can import **go-swe** as a library with:

```go
import (
  "github.com/sandhillgeo/go-swe/swe"
)
```

**Android**

The `go-swe` and `go-gpkg` code are available under `com.sandhillgeo.swe` and `com.sandhillgeo.gpkg` respectively.  For example,

```java
import com.sandhillgeo.gpkg.GeoPackage;
import com.sandhillgeo.gpkg.Gpkg;
import com.sandhillgeo.gpkg.Gpkg.*;

import com.sandhillgeo.swe.SensorGeoPackage;
import com.sandhillgeo.swe.WifiNetwork;
import com.sandhillgeo.swe.WifiNetworkList;
import com.sandhillgeo.swe.SensorList;
import com.sandhillgeo.swe.SensorMapByWifiNetwork;
import com.sandhillgeo.swe.Swe;
```

**CLI**

You can use the command line tool to create a new GeoPackage to support sensor data.

```
Usage: swe -output_uri OUTPUT_URI [-version] [-help]
Options:
  -help
    	Print help
  -output_uri string
    	The output uri of the sensor GeoPackage.
  -version
    	Prints version to stdout
```

# Building

The `build_cli.sh` script is used to build executables for Linux and Windows.  The `build_android.sh` script is used to build an [Android Archive](https://developer.android.com/studio/projects/android-library) (AAR) file and associated Javadocs.  Given the current limit to [1 go AAR file](https://github.com/golang/go/issues/15956) in an Android application, `build_android.sh` builds the bindings for `go-swe` and `go-gpkg`.

# Contributing

[Sand Hill Geographic](http://sandhillgeo.com/) is currently accepting pull requests for this repository.  We'd love to have your contributions!

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
