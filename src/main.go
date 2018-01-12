package main

import (
  "fmt"
  "os"
  "./lang"
  "github.com/urfave/cli"
  "github.com/cloudfoundry-attic/jibber_jabber"
)

func main() {
  var info lang.Info
  if _, err := jibber_jabber.DetectLanguage(); err == nil {
    info = lang.InfoCN
  }
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
