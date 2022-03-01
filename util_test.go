package util

import (
	"fmt"
	"testing"
)

func TestColor1(t * testing.T) {
	strGreen := Green("This is green")
	wanted := "\x1b[0;32mThis is green\x1b[0m"
	if strGreen != wanted {
		t.Errorf("Color fail")	
	}
}

func TestTab(t * testing.T) {
	
	test := Tab("Hello", 10, false)
	if test != "Hello     " {
		t.Errorf("Tab string to the left failed")			
	}

	test = Tab("Hello 🍏", 8, false)
	if test != "Hello 🍏 " {
		t.Errorf("Tab string with long rune failed")			
	}
	fmt.Println(test + "|")

	test = Tab("Hello|XYZabcdef 🍏", 21, false)
	if test != "Hello|XYZabcdef 🍏    " {
		t.Errorf("Tab string with long rune failed")			
	}
	fmt.Println(test + "|")

	test = Tab("Hello|XYZabcdef", 22, false)
	fmt.Println(test + "|")

	test = Tab("Hello", 10, true)
	if test != "     Hello" {
		t.Errorf("Tab string to the right failed")			
	}

	test = Tab("Hello World", 6, false)
	if test != "Hello~" {
		t.Errorf("Tab truncate right failed")			
	}

	test = Tab("Hello World", 6, true)
	if test != "~World" {
		t.Errorf("Tab truncate left failed")			
	}

	test = TabDec(0, 1, 0)
	if test != "\x1b[0;32m0\x1b[0m" {
		t.Errorf("TabDec 0 failed")			
	}

	test = TabDec(0, 5, 2)
	if test != "\x1b[0;32m 0.00\x1b[0m" {
		t.Errorf("TabDec 0.00 failed")			
	}

	test = TabDec(1234.45676, 12, 4)
	if test != "\x1b[0;32m 1 234.456 8\x1b[0m" {
		t.Errorf("TabDec 1234.45678 failed")			
	}

}

func TestColor2(t * testing.T) {

	fmt.Println(Black("black"))
	fmt.Println(Red("red"))
	fmt.Println(Green("green"))
	fmt.Println(Yellow("yellow"))
	fmt.Println(Blue("purple"))
	fmt.Println(Magenta("magenta"))
	fmt.Println(Cyan("teal"))
	fmt.Println(White("white"))
	fmt.Println(Grey("grey"))
	fmt.Println(Orange("orange"))
	fmt.Println(Brown("brown"))
	fmt.Println(LightBlue("lightblue"))

}

func TestMinMax(t * testing.T) {
	if MinI(-1,2) != -1 {
		t.Errorf("MinI error")	
	}
	if MaxI(-1,2) != 2 {
		t.Errorf("MaxI error")	
	}
}
