package main

import (
  "fmt"
  "bufio"
  "os"
  "regexp"
  "strings"
  "strconv"
  "math"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  r, err := regexp.Compile(`\d+\. `)

  if err != nil {
    fmt.Printf("There is a problem with you regexp.\n")
    return
  }

  var (
    least_dif float64 = 1000.0
    team_name = ""
  )

  for scanner.Scan() {
    if r.MatchString(scanner.Text()) == true {
      s := strings.Fields(scanner.Text())

      f_goals, err := strconv.ParseFloat(s[6],64)
      if err != nil {
        fmt.Println(err)
      }

      a_goals, err := strconv.ParseFloat(s[8],64)
      if err != nil {
        fmt.Println(err)
      }
      
      dif := math.Abs(f_goals - a_goals)

      if dif < least_dif {
        least_dif = dif
        team_name = s[1]
      }
    }
  }
  
  fmt.Println(team_name)

  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
  }
}
