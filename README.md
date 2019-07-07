## FootballApiV2
This library provides functionality to connect and retrieve data from http://football-data.org/ 

### How to use?
#### Get API KEY
To be able to get data from http://football-data.org/ you should get a valid API key: [here](https://www.football-data.org/client/register)

#### Install footballapiv2 package
`go get github.com/rocklviv/fooballapiv2`

#### Usage

Example 1. Using free plan
```golang
package main

import (
    "github.com/rocklviv/footballapiv2"
    "fmt"
)

func main() {
    // Setting API KEY for library to be able to talk with API.
    client := NewClient("<your_api_key>")

    res, err := client.GetCompetitions(nil)
    if err != nil {
        panic(err)
    }

    // Library returns a data struct that will be easy to access.
    fmt.Println(res[0].CurrentSeason.ID)
}
```

Example 2. Setting up client with maximum requests per minute.
```golang
package main

import(
    "github.com/rocklviv/footballapiv2"
    "fmt"
)

func main() {
    // In case if you are using Standart/Advanced/Pro plan 
    // you can specify RPM that related to your plan.
    client := NewClient("<your_api_key>").SetMaxRPM(30)
    res, err := client.GetCompetitions(nil)
    ...
}
```

Example 3. Using client with API filrers
```golang
package main

import(
    "github.com/rocklviv/footballapiv2"
    "fmt"
)

type Filter struct {
    DateFrom string
    DateTo   string
}


func main() {
    // In case if you are using Standart/Advanced/Pro plan 
    // you can specify RPM that related to your plan.
    client := NewClient("<your_api_key>").SetMaxRPM(30)
    filter := Filter{
        DateFrom: "2019-06-07",
        DateTo: "2019-06-08",
    }
    res, err := client.GetCompetitions(&filter)
    ...
}

```

### Contribution
1. Fork project
2. Make changes and add tests
3. Make a pull request

### Author
Denys Chekirda aka Rocklviv

### License
