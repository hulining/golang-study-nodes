package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "runtime"
)

var (
    Version      = "UNKNOWN"
    gitCommit    string
    gitTreeState = ""                     // state of git tree, either "clean" or "dirty"
    buildDate    = "1970-01-01T00:00:00Z" // build date, output of $(date +'%Y-%m-%dT%H:%M:%S')
)

func parseFlags() (string, error) {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        fmt.Fprint(os.Stderr, "http server\n")
        flag.PrintDefaults()
    }

    var (
        port       string
        versionPrt bool
    )

    flag.StringVar(&port, "port", "0.0.0.0:80", "port used ex: 0.0.0.0:80")
    flag.BoolVar(&versionPrt, "version", false, "print version info and exit")
    flag.Parse()

    if versionPrt {
        versionPrint()
        os.Exit(0)
    }

    return port, nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World")
}

func main() {
    port, _ := parseFlags()
    http.HandleFunc("/", IndexHandler)
    log.Fatal(http.ListenAndServe(port, nil))
}

func versionPrint() {
    fmt.Printf(`Name: http server
Version: %s
CommitID: %s
GitTreeState: %s
BuildDate: %s
GoVersion: %s
Compiler: %s
Platform: %s/%s
`, Version, gitCommit, gitTreeState, buildDate, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
}
