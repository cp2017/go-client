// Copyright © 2017 Christian Kniep <christian@qnib.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
  //"fmt"
  "os"

  "github.com/cp2017/go-client/cmd"
  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "go-monitor"
  app.Version = "1.1.0"

  app.Commands = []cli.Command{{
        Name:    "ipfs",
        Aliases: []string{"m"},
        Usage:   "Monitor a given service provided via IPFS",
        Flags:   []cli.Flag {
            cli.StringFlag{
                Name: "ipfs-id, id",
                Usage: "HASH of service within IPFS",
            },
        },
        Action:  func(c *cli.Context) error {
            return cmd.MonitorIPFS(c)
            },
    },{
        Name:    "http",
        Aliases: []string{"m"},
        Usage:   "Monitor a given HTTP service",
        Flags:   []cli.Flag {
            cli.StringFlag{
                Name: "url",
                Usage: "check http-status codes of given URL",
            },
            cli.StringFlag{
                Name: "iterations,i",
                Usage: "Amount of queries to be made",
            },
            cli.StringFlag{
                Name: "delay,d",
                Value: "1000",
                Usage: "Delay between queries in milliseconds",
            },
            cli.StringFlag{
                Name: "validate",
                Usage: "Comma separated list of code:percentage pairs. No validation if empty.",
            },
            cli.BoolFlag{
                Name: "negate-validate",
                Usage: "If set negate the result of the validation.",
            },
        },
        Action:  func(c *cli.Context) error {
            return cmd.MonitorHTTP(c)
        },
    },}
    app.Run(os.Args)
}
