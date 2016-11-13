package main

import (
  "os"

  "github.com/urfave/cli"

  api "app/api"
)

func main() {
  // Flags variables
  var appName string
  var port int
  var message string

  // Command line config
  app := cli.NewApp()
  app.Usage = "api usage"
  app.Version = "0.0.1"
  // Flags config
  app.Flags = []cli.Flag {
    cli.IntFlag{
      Name:        "port",
      Value:       9000,
      Usage:       "port for listen",
      Destination: &port,
    },
    cli.StringFlag{
      Name:        "message",
      Value:       "Server running",
      Usage:       "message for server",
      Destination: &message,
    },
    cli.StringFlag{
      Name:        "appName",
      Value:       "app",
      Usage:       "App Name",
      Destination: &appName,
    },
  }

  app.Action = func(c *cli.Context) error {
    app.Name = appName
    api.Run(message, port)
    return nil
  }
  app.Run(os.Args)
}