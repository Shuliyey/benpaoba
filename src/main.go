package main

import (
  "fmt"
  "os"
  "./lang"
  "github.com/urfave/cli"
)

func main() {
  info := lang.Info{Name: "奔跑吧(bēn pǎo ba)", Usage: "脚本执行", Version: "0.0.1"}
  app := cli.NewApp()
  app.Name = info.Name
  app.Usage = info.Usage
  app.Version = info.Version
  app.Action = func(c *cli.Context) error {
    fmt.Println("伐木累")
    return nil
  }

  app.Run(os.Args)
}
