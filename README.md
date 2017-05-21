# In Other Words
Inline translator written in Go, based on Google Translation API

```sh
# Word
$ iow -s zh -t en 月亮
moon

# Sentence
$ iow -s zh -t en "Fly me to the [月亮]. Let me [玩耍] among the stars."
Fly me to the moon. Let me play among the stars.

# Not specifying source language, use auto detection
$ iow -t en "Let me see what [春天] is like on [木星] and [火星]."
Let me see what spring is like on Jupiter and Mars.
```