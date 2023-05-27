package badwords

import (
	"io"
	"regexp"
	"strings"
)

type List struct {
	words []string
}

func ReadList(r io.Reader) (newList *List, err error) {
	buff := new(strings.Builder)

	_, err = io.Copy(buff, r)
	if err != nil {
		return nil, err
	}

	wordList := strings.Split(buff.String(), "\n")
	newList = &List{words: wordList}

	return
}

func (l *List) Filter(s string) (filtered string) {
	for _, word := range l.words {
		if strings.Contains(strings.ToLower(s), strings.ToLower(word)) {
			var replacer string
			for i := 0; i < len([]rune(word)); i++ {
				replacer += string('*')
			}
			re := regexp.MustCompile(`(?i)` + word)
			filtered = re.ReplaceAllString(s, replacer)
		}
	}
	return
}
