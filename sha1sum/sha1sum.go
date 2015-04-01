package main
import (
    "fmt"

    "github.com/souleiman/checksum_util/helper"
    "github.com/docopt/docopt-go"
    "github.com/souleiman/checksum"
    "crypto/sha1"
)

func main() {
    name := "sha1sum"
    usage := fmt.Sprintf(helper.Usage, name, name, name)

    args, _ := docopt.Parse(usage, nil, true, fmt.Sprintf("%s 1.0", name), false)

    files := args["<files>"].([]string)
    for _, file := range files {
        sum, err := checksum.Compute(sha1.New(), file)

        if err != nil {
            fmt.Println(err.Error())
            continue
        }

        fmt.Printf("%x\n\t%s", sum, file)
    }
}
