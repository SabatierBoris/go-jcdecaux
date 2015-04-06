# Go-JCDecaux

JCDecaux API wrapper for Go (Golang).
Under heavy development and in dire need of work.

API structure and quite a bit of code has been shamelessly ripped off from
[go-trakt](https://github.com/42minutes/go-trakt); a lovely TraktTV API
wrapper.

## Current Features

* List all bike JCDecaux's Contracts
* List all bike JCDecaux's stations
* List all bike JCDecaux's stations for a contract
* Get the status of a station

## Unit Test

This wrapper have some unit test
To run it, you should add a testing.config with your API Key.

```json
{
	"JCDecaux_APIKey": "YOUR_API_KEY"
}
```

## Example

```go
package main

import (
  "fmt"

  "github.com/SabatierBoris/go-jcdecaux"
)

func main() {
  client := jcdecaux.NewClient(
    "API_KEY",
  )

  contracts, result := client.Contracts().All()
  fmt.Println(contracts)
  fmt.Println(result.Error()) // Should be nil

  stations, result := client.Stations().All()
  fmt.Println(stations)
  fmt.Println(result.Error()) // Should be nil

  stations, result = client.Stations().AllFor("Paris")
  fmt.Println(stations)
  fmt.Println(result.Error()) // Should be nil

  station, result := client.Stations().InformationsFor(31705, "Paris")
  fmt.Println(station)
  fmt.Println(result.Error()) // Should be nil
}
```
