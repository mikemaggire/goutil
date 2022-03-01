package util

import (
	"strings"
	"fmt"
	"strconv"
	"log"
	"math"
)

const stroverflow = "~"

var (
	Black   	= Color("\x1b[0;30m%s\x1b[0m")
	Red     	= Color("\x1b[0;31m%s\x1b[0m")
	Green   	= Color("\x1b[0;32m%s\x1b[0m")
	Yellow  	= Color("\x1b[0;33m%s\x1b[0m")
	Blue	  	= Color("\x1b[0;34m%s\x1b[0m")
	Magenta 	= Color("\x1b[0;35m%s\x1b[0m")
	Cyan    	= Color("\x1b[0;36m%s\x1b[0m")
	White   	= Color("\x1b[0;37m%s\x1b[0m")
	Grey   		= Color("\x1b[38;5;245m%s\x1b[0m")
	Orange   	= Color("\x1b[38;5;208m%s\x1b[0m")
	Brown  		= Color("\x1b[38;5;172m%s\x1b[0m")
	LightBlue  	= Color("\x1b[38;5;123m%s\x1b[0m")
)

// Add special terminal color character to a string
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
		fmt.Sprint(args...))
	}
	return sprint
}

// create a string of len length by adding spaces on the left or the right, or truncate it
func Tab(s string, length int, fRight bool) string {
	o := s; l := len([]rune(o))

	if fRight { // right tab
		if length > l {
			o = strings.Repeat(" ", length - l) + o
		} else if l > length {
			o = stroverflow + o[l-length+1:]
		}
	}else{ // left tab
		if length > l {
			o += strings.Repeat(" ", length - l)
		} else if l > length {
			o = o[:length-1] + stroverflow
		}
	}
	return o
}

// format a float to a fixed length string with separator for thouthands and thousandths, aligning the decimal at pos
// add trucation character in case of overflow
func TabDec(f float64, length int, pos int) string {
	if pos < 0 || length < 1 || length < pos || (pos > 0 && length <= pos) {
		panic("TabDec: illegal Parameters")
	}
	// basic format
	sfmt := "%" + strconv.Itoa(length) + "." + strconv.Itoa(pos) + "f"
	
	// build a string with a length multiple of 3, on the int part
	ent := strings.Split(fmt.Sprintf(sfmt, f), ".")
	len3 := ((len(ent[0]) / 3) +1) *3
	if len3 > len(ent[0]) {
		ent[0] = strings.Repeat(" ", len3-len(ent[0])) + ent[0]
	}

	// add a space every 3 chars
	o := ""; space :=""
	for i := 0; i<len3 ; i+=3 {
		o += space + ent[0][i:i+3]
		space = " "
	}

	// format decimals
	if pos > 0 {
		o += "."
		space = ""
		len1 := len(ent[1])
		for j := 0; j<len1 ; j+=3 {
			o += space + ent[1][j:MinI(j+3, len1)]
			space = " "
		}
	}

	// check overflow
	l := len(o)
	o = o[l-length:]
	fCheck, _ := strconv.ParseFloat(strings.Replace(o, " ", "", -1), 64)
	if math.Abs(fCheck - f) > 1 {
		o = stroverflow + o[len(o)-length+1:]
		log.Println("TabDec: overflow")
	}
	return Green(o)
}
