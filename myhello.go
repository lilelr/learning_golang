package main

import (
    "fmt"
    // "os"
    "strings"
    "math"
  "unicode"
  "unicode/utf8"
)

const (
  Cyan = iota
  Magenta
  Yello
)

type BitFlag int
const (
  Active BitFlag =1 << iota
  Send
  Receive
)

func (flag BitFlag) String() string {
  var flags []string
  if flag & Active == Active {
    flags = append(flags, "Active")
  }

  if flag & Send == Send {
    flags = append(flags, "Send")
  }

  if flag & Receive == Receive {
    flags = append(flags, "Receive")
  }

  if len(flags) > 0 {
    return fmt.Sprintf("%d(%s)",int(flag),strings.Join(flags,"|"))
  }

  return "0()"

}

func Uint8FromInt(x int)(uint8, error){
  if 0<=x && x<=math.MaxUint8{
    return uint8(x),nil
  }
  return 0,fmt.Errorf("%d is out of the uint8 range",x)
}

func EqualFloat(x,y,limit float64) bool {
  if limit <= 0.0 {
    limit = math.SmallestNonzeroFloat64
  }
  return math.Abs(x-y) <= (limit * math.Min(math.Abs(x), math.Abs(y)))
}

func EqualFloatPrec(x,y,float64,decimals int) bool{
  a := fmt.Sprintf("%.*f",decimals, x)
  b := fmt.Sprintf("%.*f",decimals, y)
  return len(a) == len(b) && a== b
}

func IntFromFloat64(x float64)int {
  if math.MinInt32 <= x && x <= math.MaxInt32 {
    whole, fraction := math.Modf(x)
    if fraction >= 0.5{
      whole++
    }
    return int(whole)
  }
  panic(fmt.Sprintf("%g is out of the int32 range",x))
}

func main() {
    // who := "lele"
    // if len(os.Args)> 1 {
    //   who = strings.Join(os.Args[1:]," ")
    // }
    //
    // fmt.Println("Hello",who)
    //
    // // flag := Active | Send
    // // fmt.Println(BitFlag(0), Active, Send, flag, Receive, flag|Receive)
    //
    // const factor = 3 // factor is compatible with any numeric type
    // i := 20000
    // i *= factor
    // j := int16(20)
    // i += int(j)
    // k := uint8(0)
    // k = uint8(i) // succeed, but k's value is truncated to 8 bits X
    // fmt.Println(i,j,k)

    book := "The Soprit" + "byee "
    book += "and kate Pickett"
    fmt.Print("Josey" < "jose", "josey" == "www")
    fmt.Println()

    phrase := "我爱北京天安门"
    fmt.Println(len(phrase))
    fmt.Printf("string: \"%s\"\n",phrase)
    //for index,char := range phrase {
    //  fmt.Printf("%-2d     %U    %c    %X\n",index,char,char, []byte(string(char)))
    //}
	//
    //line := "road og glue"
    //i := strings.Index(line," ")
    //firstWord := line[:i]
    //j := strings.LastIndex(line, " ")
    //lastWord := line[j+1:]
    //fmt.Println(firstWord, lastWord)

    line := "天空 是 蓝色的"
    i:=strings.IndexFunc(line,unicode.IsSpace)
    firstWord:= line[:i]
    j:= strings.LastIndexFunc(line,unicode.IsSpace)
    _,size := utf8.DecodeLastRuneInString(line[j:])
    lastWord := line[j+size:]
    fmt.Println(firstWord,lastWord)

    fmt.Printf("%t %t\n", true, false)
    fmt.Printf("|%b|%9b|%-9b|%09b|% 9b|\n",37,37,37,37,37)
}
