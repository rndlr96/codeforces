package main

import (
  "fmt"
 "sort"
)

func main(){

  var input int
  var now int
  m := make(map[int]bool)
  total := 0
  fmt.Scanf("%d", &input)

  for i:=1 ; i < input+1 ; i++ {
    total = 1
    now = 1+i

    for (now%input) != 1 {
      if now > input {
        now = now%input
        total += now
        now += i
      } else {
        total += now
        now += i
      }
    }

    m[total] = true
  }

  result := make([]int, 0, len(m))
  for key, _ := range m {
    result = append(result, key)
  }
  sort.Ints(result)

  for _, val := range result {
    fmt.Printf("%d ", val)
  }

}
