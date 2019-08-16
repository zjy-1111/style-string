package main

import (
	"fmt"
	"style-string"
)

func main() {
	fmt.Println(style.StyleString("hello", style.Reset, style.Blue, style.Red), style.StyleString("world", style.Reset, style.Default, style.Green))
	fmt.Println(style.StyleString("hello", style.Bold, style.Default, style.Green), style.StyleString("world", style.Bold, style.Default, style.Yellow))
	fmt.Println(style.StyleString("hello", style.Faint, style.Default, style.Yellow), style.StyleString("world", style.Faint, style.Default, style.Blue))
	fmt.Println(style.StyleString("hello", style.Italic, style.Default, style.Blue), style.StyleString("world", style.Italic, style.Default, style.Puple))
	fmt.Println(style.StyleString("hello", style.Underline, style.Default, style.Puple), style.StyleString("world", style.Underline, style.Default, style.Cyan))
	fmt.Println(style.StyleString("hello", style.BlinkRapid, style.Default, style.Cyan), style.StyleString("world", style.BlinkRapid, style.Default, style.White))
	fmt.Println(style.StyleString("hello", style.ReverseVideo, style.Default, style.White), style.StyleString("world", style.ReverseVideo, style.Default, style.Red))

	fmt.Println(style.StyleString("hello world", style.CrossedOut, style.Blue, style.Default))
}
