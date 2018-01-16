package task

import (
  "fmt"
)

func version() {
  fmt.Println("package task")
}

type Task struct {
  depends []*Task
  name string
  msg string
  script string
  timeout float64
  status string
  priority uint64
}
