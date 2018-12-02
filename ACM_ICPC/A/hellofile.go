package main

import (
  "fmt"
  "strings"
  "strconv"
  "bufio"
  "os"
  "log"
)

type List struct {
  list []int
}

func lineInput(text string) *List {

  tmpList := new(List)

  splitString := strings.Split(text, " ")

  if len(splitString) != 12 {
    fmt.Println("Incorrect test case size")
    os.Exit(0)
  }

  for i := range splitString{
    value, _ := strconv.Atoi(splitString[i])
    tmpList.list = append(tmpList.list, value)
  }

  return tmpList
}

func Range(val int) bool {
  if val < 1 || val > 100 {
    return false
  } else {
    return true
  }
}

func Check(arraylist []*List) {

  var largest int
  var value int
  duplicate := make(map[int]bool)

  for i := 0 ; i < len(arraylist) ; i++ {

    isfalse := true
    for j := 0 ; j < 4 ; j++ {

      value = (*arraylist[i]).list[j]

      if j == 0 {
        isfalse = Range(value)
        duplicate[value] = true
        largest = value
      } else if largest > value {
        fmt.Println("no")
        isfalse = false
        break
      } else {
        if _, ok := duplicate[value]; ok {
          fmt.Println("no")
          isfalse = false
          break
        }
        isfalse = Range(value)
        largest = value
        duplicate[value] = true
      }
    }

    if isfalse == true {
      for j := 4 ; j < 12 ; j++ {
        value = (*arraylist[i]).list[j]
        if duplicate[value] == true || largest > value {
          isfalse = false
          break
        }
        isfalse = Range(value)
        duplicate[value] = true
      }
      if isfalse == true {
        fmt.Println("yes")
      } else {
        fmt.Println("no")
      }
    }
    for k := range duplicate {
      delete(duplicate, k)
    }
  }
}

func main() {

  arraylist := []*List{}
  file, err := os.Open("hello.in")
  scanner := bufio.NewScanner(file)

  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner.Scan()
  input, _ := strconv.Atoi(scanner.Text())

  for i := 0 ; i < input ; i++ {
    scanner.Scan()
    arraylist = append(arraylist, lineInput(scanner.Text()))
  }

  Check(arraylist)

}
