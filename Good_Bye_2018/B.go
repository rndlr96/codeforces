package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

type pair struct {
  list []int
}

func input(totalX *int, totalY *int) {

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  input_string := scanner.Text()
  splitString := strings.Split(input_string, " ")

  for i := range splitString{
    value, _ := strconv.Atoi(splitString[i])
    if i == 0 {
      *totalX += value
    } else if i == 1 {
      *totalY += value
    }
  }
}

func main(){
  var line int
  totalX := 0;
  totalY := 0;
  fmt.Scanf("%d", &line)

  for i := 0 ; i < (line*2)+1 ; i++ {
    input(&totalX, &totalY)
  }

  fmt.Println(totalX/line, totalY/line)

}
