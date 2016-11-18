package main

import (
  "os"

  "github.com/urfave/cli"

  api "app/api"
)

func main() {
  // Flags variables
  var port int
  var api_prefix string

  // Instans command line
  app := cli.NewApp()
  
  // Command line config
  app.Usage = "api usage"
  app.Version = "0.0.1"
  app.Name = "MyApp"

  // Flags config
  app.Flags = []cli.Flag {
    cli.IntFlag{
      Name:        "port",
      Value:       9000,
      Usage:       "port for listen",
      Destination: &port,
    },
    cli.StringFlag{
      Name:        "api_prefix",
      Value:       "api/v1",
      Usage:       "api prefix",
      Destination: &api_prefix,
    },
  }

  app.Action = func(c *cli.Context) error {
    // Run API
    api.Run(port, api_prefix)
    return nil
  }

  // Run command line App
  app.Run(os.Args)
}