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

    fmt.Println("Source is", args.Source)
    fmt.Println("Target is", args.Target)
    fmt.Println("List is", args.List)
    fmt.Println("Text is", args.Text)

    if (args.List) {
        ListLanguages()
        return
    }
    
    config := GetConfig()
    
    fmt.Println("Google API key is", config.ApiKey)
}

/********** Arguments **********/

type Args struct {
    Source string
    Target string
    Text   string
    List   bool
}

func GetArgs() *Args {
    args := new(Args)

    flag.StringVar(&args.Source, "s", "", "Source language")
    flag.StringVar(&args.Target, "t", "", "Target language")
    flag.BoolVar(&args.List, "l", false, "List supported languages")
    flag.Parse()

    args.Text = flag.Arg(0)

    return args
}

/********** Configs **********/

type Config struct {
    ApiKey string `json:"api_key"`
}

func GetConfig() *Config {
    if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
        SetupConfig()
    } else if err != nil {
        log.Fatalf("Error reading config file")
    }

    raw, err := ioutil.ReadFile(ConfigPath)
    if err != nil {
        log.Fatalf("Error reading config file")
    }

    config := new(Config)
    if err = json.Unmarshal(raw, config); err != nil {
        log.Fatalf("Error parsing config file")
    }

    return config
}

func SetupConfig() error {
    config := new(Config)
    config.ApiKey = GetInput("Google API Key: ")
    fmt.Println(config.ApiKey)

    buf, _ := json.Marshal(config)
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

/********** Auxiliary **********/

func ListLanguages() {
    fmt.Print(`Supported languages:
    af  Afrikaans
    am  Amharic
    ar  Arabic
    az  Azeerbaijani
    be  Belarusian
    bg  Bulgarian
    bn  Bengali
    bs  Bosnian
    ca  Catalan
    ceb Cebuano
    co  Corsican
    cs  Czech
    cy  Welsh
    da  Danish
    de  German
    el  Greek
    en  English
    eo  Esperanto
    es  Spanish
    et  Estonian
    eu  Basque
    fa  Persian
    fi  Finnish
    fr  French
    fy  Frisian
    ga  Irish
    gd  Scots Gaelic
    gl  Galician
    gu  Gujarati
    ha  Hausa
    haw Hawaiian
    hi  Hindi
    hmn Hmong
    hr  Croatian
    ht  Haitian Creole
    hu  Hungarian
    hy  Armenian
    id  Indonesian
    ig  Igbo
    is  Icelandic
    it  Italian
    iw  Hebrew
    ja  Japanese
    jw  Javanese
    ka  Georgian
    kk  Kazakh
    km  Khmer
    kn  Kannada
    ko  Korean
    ku  Kurdish
    ky  Kyrgyz
    la  Latin
    lb  Luxembourgish
    lo  Lao
    lt  Lithuanian
    lv  Latvian
    ma  Punjabi
    mg  Malagasy
    mi  Maori
    mk  Macedonian
    ml  Malayalam
    mn  Mongolian
    mr  Marathi
    ms  Malay
    mt  Maltese
    my  Burmese
    ne  Nepali
    nl  Dutch
    no  Norwegian
    ny  Chichewa
    pl  Polish
    ps  Pashto
    pt  Portuguese
    ro  Romanian
    ru  Russian
    sd  Sindhi
    si  Sinhala
    sk  Slovak
    sl  Slovenian
    sm  Samoan
    sn  Shona
    so  Somali
    sq  Albanian
    sr  Serbian
    st  Sesotho
    su  Sundanese
    sv  Swedish
    sw  Swahili
    ta  Tamil
    te  Telugu
    tg  Tajik
    th  Thai
    tl  Filipino
    tr  Turkish
    uk  Ukrainian
    ur  Urdu
    uz  Uzbek
    vi  Vietnamese
    xh  Xhosa
    yi  Yiddish
    yo  Yoruba
    zh  Chinese
    zu  Zulu
`)
}