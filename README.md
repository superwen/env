## depends
```
github.com/joho/godotenv
```
## install
```
go get -u github.com/superwen/env
```
## usage
```
package main

import (
    "github.com/superwen/env"
    "fmt"
)

func main() {
    env.Load();
    str := env.StringDefault("KEY01", "default")
    fmt.Println(str)
}
```
## methods
- Load() error
- LoadMust()
- Has(key string) bool
- Must(key string)
- BoolDefault(key string, daf bool) bool
- StringDefault(key string, daf string) string
- StringMust(key string) string
- IntDefault(key string, daf int) int
- IntMust(key string) int
- Int64Default(key string, daf int64) int64
- Int64Must(key string) int64
- Float32Default(key string, daf float32) float32
- Float32Must(key string) float32
- Float64Default(key string, daf float64) float64
- Float64Must(key string) float64
