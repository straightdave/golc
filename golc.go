/*
GOLC - GO List Case names
Dave Wu @ 2017
*/
package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

var caseRegex = regexp.MustCompile(`^func (?P<CaseName>Test.+?)\s*?\(.+?$`)
var cnames []string

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "No file names given, TODO: list cases in current folder\n")
        os.Exit(-1)
    }

    fname := os.Args[1]

    if f, err := os.Open(fname); err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
        os.Exit(-1)
    } else {
        defer f.Close()

        input := bufio.NewScanner(f)
        for input.Scan() {
            line := input.Text()

            if len(line) == 0 {
                continue
            }

            matches := caseRegex.FindStringSubmatch(line)

            if len(matches) > 0 {
                cnames = append(cnames, matches[1])
            }
        }
    }

    for _, s := range cnames {
        fmt.Println(s)
    }
}
