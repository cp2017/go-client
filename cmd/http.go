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
    "os"
    "strconv"
    "strings"
    "time"
    "net/http"

    "github.com/urfave/cli"
)

// ValidateStatus checks if validation string is met
func ValidateStatus(res map[string]int, count int, validate string) bool {
    for _, vPair := range strings.Split(validate, ",") {
        vS := strings.Split(vPair, ":")
        vCount,_ := strconv.Atoi(vS[1])
        percent := (float32(res[vS[0]])/float32(count))*100
        if percent != float32(vCount) {
            log.Printf("Expected %s:%s, found %.f", vS[0], vS[1], percent)
            return false
        }
    }

    return true
}

// MonitorHTTP a given service from ethereum
func MonitorHTTP(c *cli.Context) error {
    url := c.String("url")
    iterations := c.Int("iterations")
    delay := c.Int("delay")
    log.Printf("Fetching '%s' %d times with a delay of %d; validate:'%v'\n", url, iterations, delay, c.String("validate"))
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
    if c.String("validate") != "" {
        validate := ValidateStatus(res, cnt, c.String("validate"))
        if c.Bool("negate-validate") {
            log.Printf("Validation was %v, but as 'negate-validate' was set turn it around\n", validate)
            if validate {
                os.Exit(1)
            }
        } else if ! validate {
            log.Println("Validation failed")
            os.Exit(1)
        }
    }
    return nil
}
