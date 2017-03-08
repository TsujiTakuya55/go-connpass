# go-connpass

go-connpass is a Go client library for accessing the [connpass API](https://connpass.com/about/api/).

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/TsujiTakuya55/go-connpass/blob/master/LICENSE)

## Requirement

- Go version 1.7 or greater.

## Installation

```shell
$ go get github.com/TsujiTakuya55/go-connpass/connpass
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/TsujiTakuya55/go-connpass/connpass"
)

func main() {

    param := &connpass.Param{
        EventId:2,
    }

    client := connpass.NewClient()
    connpass, _, err := client.Get(param)

    if err != nil {
        fmt.Println(fmt.Errorf("err! \n", err))
        return
    }

    for _, k := range *connpass.Events {
        fmt.Println("eventId :", *k.EventId)
        fmt.Println("title :", *k.Title)
        fmt.Println("address :", *k.Address)
        fmt.Println("startedAt :", *k.StartedAt)
    }
}
```

For details, please click [here](https://github.com/TsujiTakuya55/go-connpass/tree/master/examples)

## Anything Else

coming soon

## Author

[@Tsuji_Taku50](https://twitter.com/Tsuji_Taku50?lang=ja)

## License

Please see the [license file](https://github.com/TsujiTakuya55/go-connpass/blob/master/LICENSE) for the license of this library
