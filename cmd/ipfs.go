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

package cmd

import (
    "fmt"

    "github.com/urfave/cli"
    //"github.com/ipfs/go-ipfs-api"
    //"github.com/ethereum/go-ethereum/ethclient"
)

// MonitorIPFS a given service from ethereum
func MonitorIPFS(c *cli.Context) error {
    fmt.Println(c.FlagNames())
    //cli, _ := ethclient.Dial("http://localhost:8545")
    //_ = cli
    return nil
}