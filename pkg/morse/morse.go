package morse

import "strings"

var morseTable = map[rune]string{
	'А': ".-", 'Б': "-...", 'В': ".--", 'Г': "--.", 'Д': "-..",
	'Е': ".", 'Ж': "...-", 'З': "--..", 'И': "..", 'Й': ".---",
	'К': "-.-", 'Л': ".-..", 'М': "--", 'Н': "-.", 'О': "---",
	'П': ".--.", 'Р': ".-.", 'С': "...", 'Т': "-", 'У': "..-",
	'Ф': "..-.", 'Х': "....", 'Ц': "-.-.", 'Ч': "---.", 'Ш': "----",
	'Щ': "--.-", 'Ъ': ".--.-.", 'Ы': "-.--", 'Ь': "-..-", 'Э': "..-..",
	'Ю': "..--", 'Я': ".-.-",
}

var textTable map[string]string

func init() {
	textTable = make(map[string]string)
	for k, v := range morseTable {
		textTable[v] = string(k)
	}
}

func ToMorse(text string) string {
	text = strings.ToUpper(text)
	var res []string
	for _, r := range text {
		if m, ok := morseTable[r]; ok {
			res = append(res, m)
		} else if r == ' ' {
			res = append(res, "/")
		}
	}
	return strings.Join(res, " ")
}

func ToText(code string) string {
	parts := strings.Split(code, " ")
	var res []string
	for _, p := range parts {
		if t, ok := textTable[p]; ok {
			res = append(res, t)
		} else if p == "/" {
			res = append(res, " ")
		}
	}
	return strings.Join(res, "")
}
