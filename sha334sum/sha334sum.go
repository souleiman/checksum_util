package main
import (
    "fmt"

    "github.com/souleiman/checksum_util/helper"
    "github.com/docopt/docopt-go"
    "github.com/souleiman/checksum"
    "crypto/sha512"
)

func main() {
    name := "sha334sum"
    usage := fmt.Sprintf(helper.Usage, name, name, name)

    args, _ := docopt.Parse(usage, nil, true, fmt.Sprintf("%s 1.0", name), false)

    files := args["<files>"].([]string)
    for _, file := range files {
        sum, err := checksum.Compute(sha512.New384(), file)

        if err != nil {
            fmt.Println(err.Error())
            continue
        }

        fmt.Printf("%x\t%s\n", sum, file)
    }
}
