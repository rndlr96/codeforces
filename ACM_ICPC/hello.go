package main

import (
  "fmt"
  "strings"
  "strconv"
  "bufio"
  "os"
)

type List struct {
  list []int
}

func Input() int{
  input := -1

  for {
    fmt.Scanf("%d", &input)
    if input >= 1 && input <= 4096 {
      break;
    }
  }

  return input
}

func lineInput(val int) []*List {

  arraylist := []*List{}
  scanner := bufio.New(os.Stdin)

  for count := 0 ; count < val ; count++ {

    tmpList := new(List)
    var input_string string

    scanner.Scan()
    input_string = scanner.Text()

    splitString := strings.Split(input_string, " ")

    if len(splitString) != 12 {
      fmt.Println("Incorrect test case size")
      os.Exit(0)
    }

    for i := range splitString{
      value, _ := strconv.Atoi(splitString[i])
      tmpList.list = append(tmpList.list, value)
    }
    arraylist = append(arraylist, tmpList)
    tmpList = nil
  }

  return arraylist
}

func Check(arraylist []*List) {

  var largest int

  for i := 0 ; i < len(arraylist) ; i++ {

    isfalse := true
    for j := 0 ; j < 4 ; j++ {
      if j == 0 {
        largest = (*arraylist[i]).list[j]
      } else if largest > (*arraylist[i]).list[j] {
        fmt.Println("no")
        isfalse = false
        break
      } else {
        largest = (*arraylist[i]).list[j]
      }
    }

    if isfalse == true {
      for j := 4 ; j < 12 ; j++ {
        if largest > (*arraylist[i]).list[j] {
          fmt.Println("no")
          isfalse = false
        }
      }
      if isfalse == true {
        fmt.Println("yes")
      }
    }
  }
}

func main() {

  arraylist := []*List{}

  input := Input()

  arraylist = lineInput(input)

  Check(arraylist)

}
