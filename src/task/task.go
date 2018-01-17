package task

import (
  "fmt"
  "sync"
  "github.com/satori/go.uuid"
)

func version() {
  fmt.Println("package task")
}

type Task struct {
  id uuid.UUID
  parents []*Task
  children []*Task
  name string
  msg string
  script string
  timeout float64
  status string
  priority uint64
  mux sync.Mutex
}

func (t *Task) hasParent(p *Task) {
  for _, _t := range t.Task {
    if _t.id == p.id {
      return true
    }
  }
  return false
}

func (t *Task) AddParent(p *Task) {
  t.mux.Lock()
  if ! t.hasParent(p) {
    t.parents.append(p)
    p.AddChildren(t)
  }
  t.mux.UnLock()
}

func (t *Task) hasChildren(c *Task) {
  for _, _t := range t.Task {
    if _t.id == c.id {
      return true
    }
  }
  return false
}

func (t *Task) AddChildren(c *Task) {
  t.mux.Lock()
  if ! t.hasChildren(c) {
    t.children.append(c)
    c.AddParent(t)
  }
  t.mux.UnLock()
}
