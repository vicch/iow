package main

import (
    "bufio"
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "regexp"
    "strings"
)

const ApiUri = "https://www.googleapis.com/language/translate/v2"
const ConfigPath = "./.config"

func main() {
    args := GetArgs()

    if (args.List) {
        ListLanguages()
        return
    }
    
    before := FindWords(args.Text)

    // Nothing to translate
    if len(before) == 0 {
        fmt.Println(args.Text)
        return
    }

    after := TranslateWords(args.Source, args.Target, before)
    fmt.Println(ReplaceWords(args.Text, before, after))

    return
}

/********** Translation **********/

func FindWords(text string) []string {
    // Find all words to translate (in format "[...]")
    re := regexp.MustCompile("\\[[^\\[\\]]*\\]")
    return re.FindAllString(text, -1)
}

func TranslateWords(source, target string, words []string) []string {
    after := make([]string, len(words))

    body := CallApi(source, target, words)
    for i, translation := range body.Data.Translations {
        after[i] = translation.TranslatedText
    }

    return after
}

func ReplaceWords(text string, before, after []string) string {
    for i, word := range before {
        text = strings.Replace(text, word, after[i], -1)
    }
    return text
}

/********** API **********/

type Body struct {
    Data *TranslationsResponse `json:"data"`
}

type TranslationsResponse struct {
    Translations []*TranslationsResource `json:"translations,omitempty"`
}

type TranslationsResource struct {
    TranslatedText string `json:"translatedText,omitempty"`
}

func CallApi(source, target string, words []string) *Body {
    config := GetConfig()
    res, _ := http.Get(MakeApiUri(config.ApiKey, source, target, words))
    body := new(Body)
    json.NewDecoder(res.Body).Decode(&body)

    return body
}

func MakeApiUri(key, source, target string, words []string) string {
    uri := ApiUri
    uri += "?key=" + key
    uri += "&source=" + source
    uri += "&target=" + target
    for _, word := range words {
        // Remove "[" and "]"
        uri += "&q=" + word[1:len(word)-1]
    }
    return uri
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

    flag.StringVar(&args.Source, "s", "", "Set source language")
    flag.StringVar(&args.Target, "t", "", "Set target language")
    flag.BoolVar(&args.List, "l", false, "List supported languages")
    flag.Parse()

    args.Text = flag.Arg(0)

    return args
}

/********** Config **********/

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

    buf, _ := json.Marshal(config)
    return ioutil.WriteFile(ConfigPath, buf, 0644)
}

/********** Auxiliary **********/

func GetInput(prompt string) string {
    fmt.Print(prompt)

    reader := bufio.NewReader(os.Stdin)
    buf, err := reader.ReadString('\n')

    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }

    return strings.Trim(buf, "\n")
}

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