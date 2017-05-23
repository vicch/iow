# In Other Words
Command line translator written in Go, based on Google Translation API

```
iow [OPTION]... [TEXT]

[OPTION]
    -s LANG  Set source language
    -t LANG  Set target language
    -l       List supported languages
    -h       Help

[TEXT]
    Use "[]" to wrap word(s) to translate. Example: "Translate this [word]."
```

It will ask you to set up API key and default settings on first execution:
```sh
$ iow
Google API Key: [API key]
Default source language: [Language code]
Default target language: [Language code]
```

> Note: It is recommended to leave source language empty and use Google Translate auto detection.

Translate single word:
```sh
$ iow -t en [翻译]
translation
$ iow -t ja [moon]
月
```

Translate sentence:
```sh
$ iow -s zh -t en "Fly me to the [月亮]. Let me [玩耍] among the stars."
Fly me to the moon. Let me play among the stars.
```

When source/target language is not specified, default settings are used:
```sh
$ iow "Let me see what [春天] is like on [木星] and [火星]."
Let me see what spring is like on Jupiter and Mars.
```