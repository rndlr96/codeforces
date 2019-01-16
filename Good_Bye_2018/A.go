package main

import (
  "fmt"
  "math"
)

func main(){

  var y, b, r int

  fmt.Scanf("%d %d %d", &y, &b, &r);

  min := int(math.Min(math.Min(float64(y),float64(b)),float64(r)))

  for i := min ; i >= 0 ; i-- {
    if y >= i && b >= i+1 && r >= i+2 {
      fmt.Println((i+1)*3)
      break;
    }
  }
}
