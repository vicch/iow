package main

import (
    "bufio"
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strings"
)

const ConfigPath = "./.config"

func main() {
    args := GetArgs()
    fmt.Println("Source is", args["source"])
    fmt.Println("Target is", args["target"])
    fmt.Println("Text is", args["text"])
    
    configs := GetConfigs()
    fmt.Println("Google API key is", configs["key"])
}

/********** Arguments **********/

func GetArgs() map[string]string {
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

/********** Configs **********/

func GetConfigs() map[string]string {
    _, err := os.Stat(ConfigPath)
    if os.IsNotExist(err) {
        SetupConfigs()
    } else if err != nil {
        log.Fatalf("Error reading config file")
    }

    raw, err := ioutil.ReadFile(ConfigPath)
    if err != nil {
        log.Fatalf("Error reading config file")
    }

    var configs map[string]string
    err = json.Unmarshal(raw, &configs)
    if err != nil {
        log.Fatalf("Error parsing config file")
    }

    return configs
}

func SetupConfigs() error {
    configs := map[string]string {
        "key": GetInput("Google API Key: "),
    }

    buf, _ := json.Marshal(configs)
    return ioutil.WriteFile(ConfigPath, buf, 0644)
}

/********** IO **********/

func GetInput(prompt string) string {
    fmt.Print(prompt)

    reader := bufio.NewReader(os.Stdin)
    buf, err := reader.ReadString('\n')

    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }

    return strings.Trim(buf, "\n")
}