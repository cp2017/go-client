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
    "log"
    "strconv"
    "time"
    "net/http"

    "github.com/urfave/cli"
)

// MonitorHTTP a given service from ethereum
func MonitorHTTP(c *cli.Context) error {
    url := c.String("url")
    iterations := c.Int("iterations")
    delay := c.Int("delay")
    log.Printf("Fetching '%s' %d times with a delay of %d\n", url, iterations, delay)
    res := map[string]int{}
    cnt := 0
    for cnt < iterations {
        cnt++
        resp, err := http.Get(url)
        if err != nil {
            panic(err.Error())
        }
        log.Printf("[%3d] %s\n", cnt, resp.Status)
        key := strconv.Itoa(resp.StatusCode)
        if _, ok := res[key]; ! ok {
            res[key] = 0
        }
        res[key]++
        time.Sleep(time.Duration(delay) * time.Millisecond)
    }
    for key, value := range res {
        log.Println("Status:", key, "Count:", value)
    }
    return nil
}
