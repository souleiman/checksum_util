package main
import (
    "fmt"

    //"github.com/souleiman/checksum"
    "github.com/souleiman/checksum_util/helper"
    "github.com/docopt/docopt-go"
)

func main() {
    name := "md5sum"
    usage := fmt.Sprintf(helper.Usage, name, name, name)

    args, _ := docopt.Parse(usage, nil, true, "md5sum 1.0", false)

    files := args["<files>"]
    fmt.Println(files)
}
