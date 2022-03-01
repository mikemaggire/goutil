# Utils in go

NOTA: this repo is a merge of gomathutil and gotermutil

## Colors for the terminal

```go
var (
	Black   = Color("\033[1;30m%s\033[0m")
	Grey    = Color("\033[0;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
	Grey   		= Color("\x1b[38;5;245m%s\x1b[0m")
	Orange   	= Color("\x1b[38;5;208m%s\x1b[0m")
	Brown  		= Color("\x1b[38;5;172m%s\x1b[0m")
	LightBlue  	= Color("\x1b[38;5;123m%s\x1b[0m")
  )

// Add special terminal color character to a string
func Color(colorString string) func(...interface{}) string {
```

## Tab & TabDec

```go
// create a string of len length by adding spaces on the left or the right
func Tab(s string, lenght int, fLeft bool) string


// format a float to a fixed length string with separator for thouthands and thousandths, aligning the decimal at pos
// add trucation character in case of overflow
func TabDec(f float64, length int, pos int) string

```

# Math MinI & MaxI

`MinI(a int, b int) int` return the smallest integer

`MaxI(a int, b int) int` return the gretest integer

---

## Install 

Run `go get` command to use it

`go get github.com/mikemaggire/util@latest

## Publishing a new version

[Publish a new version](https://go.dev/doc/modules/publishing)

## Licence

[MIT](LICENCE)
