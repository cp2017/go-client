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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ipfs/go-ipfs-api"
	"github.com/urfave/cli"

	"github.com/cp2017/go-client/util"
)

// MonitorIPFS a given service from ethereum
func MonitorIPFS(c *cli.Context) error {
	log.Println("Start IPFS Monitor")
	swagger_path := c.String("swagger-path")
	var url string
	schemes := []string{"http"}
	if swagger_path != "" {
		s := shell.NewShell("http://localhost:5001")
		hash, _ := s.AddDir(swagger_path)
		fmt.Println(hash)
		var swag util.SwaggerJSON
		file, err := ioutil.ReadFile(swagger_path)
		if err != nil {
			fmt.Printf("File error: %v\n", err)
			os.Exit(1)
		}
		err = json.Unmarshal(file, &swag)
		if err != nil {
			log.Println("error:", err)
		}
		url = fmt.Sprintf("%s%s", swag.Host, swag.BasePath)
	} else {
		url = c.String("url")
	}
	validate := c.String("validate")
	iterations := c.Int("iterations")
	delay := c.Int("delay")
	negate := c.Bool("negate-validate")
	for _, scheme := range schemes {
		util.Monitor(fmt.Sprintf("%s://%s", scheme, url), validate, iterations, delay, negate)
	}
	return nil
}
