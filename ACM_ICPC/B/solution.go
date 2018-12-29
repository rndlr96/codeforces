package main

import (
  "fmt"
  "os"
  "bufio"
  "log"
  "strings"
  "strconv"
)

func Check(length int, scanner *bufio.Scanner) {
  count := 0
  var map1 map[int]int
  var map2 map[int]int

  map1 = make(map[int]int)
  map2 = make(map[int]int)

  scanner.Scan()
  map1[0] = strings.Count(scanner.Text(), "0")
  map1[1] = strings.Count(scanner.Text(), "1")

  scanner.Scan()
  map2[0] = strings.Count(scanner.Text(), "0")
  map2[1] = strings.Count(scanner.Text(), "1")

  for ; map1[0] != 0 && map2[1] != 0 ; count++{
    map1[0] -= 1
    map2[1] -= 1
  }
  for ; map2[0] != 0 && map1[1] != 0 ; count++{
    map2[0] -= 1
    map1[1] -= 1
  }
  fmt.Println(count)
}

func main() {

  file, err := os.Open("hamming.in")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Scan()
  testcase, _ := strconv.Atoi(scanner.Text())

  if testcase < 1 || testcase > 512 {
    os.Exit(0)
  }

  for i := 0 ; i < testcase ; i++ {
    scanner.Scan()
    val,_ := strconv.Atoi(scanner.Text())
    Check(val, scanner)
  }
}
