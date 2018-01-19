package task

import (
  "fmt"
  "sync"
  "github.com/satori/go.uuid"
  "os/exec"
)

func version() {
  fmt.Println("package task")
}

type script struct {
  path string
  essential bool
}

func newscript(path string) *script {
  return &script{
    path: path
    essential: true
  }
}

type Task struct {
  id uuid.UUID
  parents []*Task
  children []*Task
  name string
  msg string
  scripts []script
  timeout float64
  status string
  retries: uint
  priority uint
  mux sync.Mutex
}

func (t *Task) hasParent(p *Task) {
  for _, _t := range t.parents {
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
  for _, _t := range t.children {
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

func (t *Task) runScripts(cmd *exec.Cmd) bool {
  // needs implementation
  for
  return true
}

func (t *Task) runScript(cmd *exec.Cmd, i uint) bool {
  // needs implementation
  return true
}
