# tesla
<p align="center">
  <img src="https://raw.githubusercontent.com/rdbell/tesla/master/images/gotesla.png" width="200">
</p>

This library provides a golang wrapper around the Tesla API to easily query and command your Tesla Vehicle (https://www.tesla.com/) remotely in Go.

## Library Documentation

[https://godoc.org/github.com/rdbell/tesla](https://godoc.org/github.com/rdbell/tesla)

## API Documentation

[View Tesla JSON API Documentation](http://docs.timdorr.apiary.io/)

This is unofficial documentation of the Tesla JSON API used by the iOS and Android apps. The API provides functionality to monitor and control Tesla vehicles remotely. The project provides both a documention of the API and a Go library for accessing it.

## Installation

```
go get github.com/jsgoecke/rdbell
```

## Tokens

```
TESLA_CLIENT_ID=e4a9949fcfa04068f59abb5a658f2bac0a3428e4652315490b659d5ab3f35a9e
TESLA_CLIENT_SECRET=c75f14bbadc8bee3a7594412c31416f8300256d7668ea7e6e7f06727bfb9d220
```

## Usage

Here's an example (more in the /examples project directory):

```go
package main

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"github.com/rdbell/tesla"
)

func main() {
	client, err := tesla.NewClient(
		&tesla.Auth{
			ClientID:     os.Getenv("TESLA_CLIENT_ID"),
			ClientSecret: os.Getenv("TESLA_CLIENT_SECRET"),
			Email:        os.Getenv("TESLA_USERNAME"),
			Password:     os.Getenv("TESLA_PASSWORD"),
		})
	if err != nil {
		panic(err)
	}

	vehicles, err := client.Vehicles()
	if err != nil {
		panic(err)
	}

	vehicle := vehicles[0]
	status, err := vehicle.MobileEnabled()
	if err != nil {
		panic(err)
	}

	fmt.Println(status)
	//fmt.Println(vehicle.HonkHorn())

	// Autopark
	// Use with care, as this will move your car
	// vehicle.AutoparkForward()
	// vehicle.AutoparkReverse()
	// Use with care, as this will move your car

	// Stream vehicle events
	eventChan, errChan, err := vehicle.Stream()

	for {
		select {
		case event := <-eventChan:
			pp.Print(event)
		case err = <-errChan:
			fmt.Println(err)
			if err.Error() == "HTTP stream closed" {
				fmt.Println("Reconnecting!")
				eventChan, errChan, err = vehicle.Stream()
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
```

## Credits

This repo was forked from [https://github.com/jsgoecke/tesla](https://github.com/jsgoecke/tesla)

Thank you to [Tim Dorr](https://github.com/timdorr) who did the heavy lifting to document the Tesla API and also created the [model-s-api Ruby Gem](https://github.com/timdorr/model-s-api).


## Copyright & License

Copyright (c) 2016-Present Jason Goecke. Released under the terms of the MIT license. See LICENSE for details.
