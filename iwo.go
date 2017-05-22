package main

import (
    "flag"
    "fmt"
)

func GetArgs() (map[string]string) {
    source := flag.String("s", "", "Source language")
    target := flag.String("t", "", "Target language")
    flag.Parse()
    text := flag.Arg(0)

    return map[string]string {
        "source": *source,
        "target": *target,
        "text":   text,
    }
}

func main() {
    args := GetArgs()
    fmt.Println("Source is", args["source"])
    fmt.Println("Target is", args["target"])
    fmt.Println("Text is", args["text"])
}