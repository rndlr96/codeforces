package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
  "strconv"
)

type pair struct {
  start int
  end int
  mid int
}

func makePair(n int) []pair{
  var pairs []pair

  for i := 0 ; i < n-1 ; i++ {
    for j := i+1 ; j < n ; j++ {
      if (j-i+1)%2 == 0 {
        tmp := pair{i,j,(i+j+1)/2}
        pairs = append(pairs, tmp)
      }
    }
  }

  return pairs
}

func check(candidate pair, a []int64) bool {
  left := a[candidate.start]
  right := a[candidate.end]

  for i := candidate.start+1 ; i < candidate.mid ; i++ {
    left = left^a[i]
  }
  for i := candidate.mid ; i < candidate.end ; i++ {
    right = right^a[i]
  }
  if left == right {
    return true
  } else {
    return false
  }
}

func main(){

  reader := bufio.NewScanner(os.Stdin)
  var n int
  count := 0

  fmt.Scanf("%d", &n)

  a := make([]int64, 0, n)
  reader.Scan()
  input := reader.Text()
  inputs := strings.Split(input, " ")

  for val := range inputs {
    tmp, _ := strconv.ParseInt(inputs[val], 10, 64)
    a = append(a, tmp)
  }

  pairs := makePair(n)

  for _,candidate := range pairs {
    if check(candidate, a) == true {
      count++
    }
  }

  fmt.Println(count)
}
