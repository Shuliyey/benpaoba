package main

import (
  "fmt"
  "os"
  // "./task"
  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "奔跑吧(bēn pǎo ba)"
  app.Usage = "脚本执行"
  app.Version= "0.0.1"
  app.Action = func(c *cli.Context) error {
    fmt.Println("伐木累")
    return nil
  }

  app.Run(os.Args)
}
