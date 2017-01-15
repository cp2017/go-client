package util


import (
    "log"
    "os"
    "strconv"
    "strings"
    "time"
    "net/http"

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

func Monitor(url, validate string, iterations,delay int, negate bool) error {
    log.Printf("Fetching '%s' %d times with a delay of %d; validate:'%v'\n", url, iterations, delay, validate)
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
    if validate != "" {
        validate := ValidateStatus(res, cnt, validate)
        if negate {
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
