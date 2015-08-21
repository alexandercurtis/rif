package main

import (
  "fmt"
  "os"
  "strconv"
)

type stack []float64

func (s stack) Empty() bool { return len(s) == 0 }
func (s stack) Peek() float64   { return s[len(s)-1] }
func (s *stack) Put(i float64)  { (*s) = append((*s), i) }
func (s *stack) Pop() float64 {
  d := (*s)[len(*s)-1]
  (*s) = (*s)[:len(*s)-1]
  return d
}
func (s *stack) Swap() {
  d := (*s)[len(*s)-1]
  (*s)[len(*s)-1] = (*s)[len(*s)-2]
  (*s)[len(*s)-2] = d
}


func (s *stack) Sum() float64 {
  res := 0.0
  for !(*s).Empty() {
    res += (*s).Pop()
  }
  return res
}

func (s *stack) Product () float64 {
  res := 1.0
  for !(*s).Empty() {
    res *= (*s).Pop()
  }
  return res
}

func main() {
  argsWithoutProg := os.Args[1:]

  var st stack;

  for _,s := range argsWithoutProg {
    switch s {
    case "help": fmt.Println("Functions are + - x / p base %")
    case "-": st.Put( 0 - st.Pop() )
    case "x","*": st.Put( st.Product() )
    case "/": {
      st.Swap()
      st.Put( st.Pop() / st.Pop() )
    }
    case "+": {
      st.Put( st.Sum() )
    }
    case "p": {
      fmt.Println( st )
    }
    case "base": {
      base := int(st.Pop())
      value := int(st.Pop())
      for value > 0 {
        rem := value % base
        st.Put( float64(rem) )
        value = (value-rem) / base
      }
      fmt.Println( st )
    }
    case "%": {
      a := int(st.Pop())
      b := int(st.Pop())
      st.Put( float64(b % a) )
    }
    default:
      var i,_ = strconv.ParseFloat( s, 64 )
      st.Put( i )
    }
  }
  fmt.Println(st.Sum())
}


