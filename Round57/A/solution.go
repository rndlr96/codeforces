package main

import (
  "fmt"
  "os"
)

type Pair struct {
  x int
  y int
}

func main(){
  var input_num int
  var tmpPair Pair;
  fmt.Scanf("%d", &input_num)

  if input_num < 1 && input_num > 1000 { os.Exit(1) }

  input_list := []Pair{}
  result_list := []Pair{}

  for i := 0 ; i < input_num ; i++ {
    fmt.Scanf("%d %d", &tmpPair.x, &tmpPair.y)
    input_list = append(input_list, tmpPair)
  }

  var j int
  for i := 0 ; i < input_num ; i++ {
    for j = 2 ; input_list[i].x * j > input_list[i].y ; j++ {
      if input_list[i].x * j > input_list[i].y{
        break
      } else {
        tmpPair.x = input_list[i].x
        tmpPair.y = input_list[i].x * j
        result_list = append(result_list, tmpPair)
        break
      }
    }
  }

  for i := 0 ; i < result_list.length ; i++ {
    fmt.Println(result_list[i].x, result_list[i].y)
  }
}
