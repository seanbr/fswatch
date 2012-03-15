package main

import "fmt"
import "strings"

func decorate(cmd command, exec func(command)) {
  cmdStrings := append([]string{cmd.name}, cmd.args...)
  fmt.Fprintf(cmd.outPipe, "%c[34;4m%s%c[0m\n", 27, strings.Join(cmdStrings, " "), 27)
  exec(cmd)
  fmt.Fprintln(cmd.outPipe)
}