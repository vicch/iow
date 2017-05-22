# In Other Words
Inline translator written in Go, based on Google Translation API

How to use:
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

Translate single word:
```sh
$ iow -s zh -t en [翻译]
translation
```

Translate sentence:
```sh
$ iow -s zh -t en "Fly me to the [月亮]. Let me [玩耍] among the stars."
Fly me to the moon. Let me play among the stars.
```

Translate without specifying source language (use auto detection):
```sh
$ iow -t en "Let me see what [春天] is like on [木星] and [火星]."
Let me see what spring is like on Jupiter and Mars.
```