# tesla

This library provides a wrapper around the API to easily query and command the a [https://www.teslamotors.com/models](Telsa Model S) remotely in Go.

![Go Tesla Gopher](https://dl.dropboxusercontent.com/u/25511/Images/gotesla.png)

## API Documentation

[View Tesla JSON API Documentation](http://docs.timdorr.apiary.io/)

This is unofficial documentation of the Tesla JSON API used by the iOS and Android apps. The API provides functionality to monitor and control the Model S (and future Tesla vehicles) remotely. The project provides both a documention of the API and a Go library for accessing it.

## Installation

```
go get github.com/jsgoecke/tesla
```

## Tokens

[Your Current client_id and client_secret when logged into My Tesla Motors](http://pastebin.com/fX6ejAHd)

## Usage

Here's an example (more in the /examples project directory):

```go
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
	fmt.Println(vehicle.HonkHorn())
}
```

## ToDo

* Provide access to the streaming API and a means to process data coming from it.
* Implement the summon features.

## Credits

## Copyright & License

Copyright (c) 2016 Jason Goecke. Released under the terms of the
MIT license. See LICENSE for details.