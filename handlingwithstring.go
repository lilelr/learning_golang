package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
	//"debug/dwarf"
)



func main()  {
	line := "我爱中国科学院 I like the institute of Chinese Academy."
	i := strings.IndexFunc(line, unicode.IsSpace)
	firstWord := line[:i]
	j := strings.LastIndexFunc(line, unicode.IsSpace)
	_, size := utf8.DecodeRuneInString(line[j:])
	lastWord := line[j+size:]
	fmt.Println(firstWord,lastWord)

	//var aa = -1
	//var bb="e"
	//fmt.Println(bb)
	//var myi uint8 =11
	//myi := 11
	//for ; myi>=0 ;myi--  {
	//	fmt.Println(myi)
	//}
  var u uint=10
  var u2 uint = 42
  fmt.Println(u - u2)
  fmt.Println(u2-u)

  //var myj int = 22
  //fmt.Println(u - myj)
  //fmt.Println(myj - u)

}