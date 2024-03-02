package persian

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

type CharForms struct {
	Independent rune
	Initial     rune
	Medial      rune
	Final       rune
}

type letterGroup struct {
	backLetter  rune
	letter      rune
	frontLetter rune
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

var glyphs = map[rune]CharForms{
	'آ': {'\u0622', '\u0622', '\u0622', '\ufe82'},
	'ا': {'\u0627', '\ufe8e', '\ufe8e', '\ufe8e'},
	'ب': {'\u0628', '\ufe91', '\ufe92', '\ufe90'},
	'ت': {'\u062a', '\ufe97', '\ufe98', '\ufe96'},
	'ث': {'\u062b', '\ufe9b', '\ufe9c', '\ufe9a'},
	'ج': {'\u062c', '\ufe9f', '\ufea0', '\ufe9e'},
	'ح': {'\u062d', '\ufea3', '\ufea4', '\ufea2'},
	'چ': {'\u0686', '\ufb7c', '\ufb7d', '\ufb7b'},
	'خ': {'\u062e', '\ufea7', '\ufea8', '\ufea6'},
	'د': {'\u062f', '\u062f', '\u062f', '\ufeaa'},
	'ذ': {'\u0630', '\u0630', '\u0630', '\ufeac'},
	'ر': {'\u0631', '\u0631', '\u0631', '\ufeae'},
	'ز': {'\u0632', '\u0632', '\u0632', '\ufeb0'},
	'س': {'\u0633', '\ufeb3', '\ufeb4', '\ufeb2'},
	'ش': {'\u0634', '\ufeb7', '\ufeb8', '\ufeb6'},
	'ص': {'\u0635', '\ufebb', '\ufebc', '\ufeba'},
	'ض': {'\u0636', '\ufebf', '\ufec0', '\ufebe'},
	'ط': {'\u0637', '\ufec3', '\ufec4', '\ufec2'},
	'ظ': {'\u0638', '\ufec7', '\ufec8', '\ufec6'},
	'ع': {'\u0639', '\ufecb', '\ufecc', '\ufeca'},
	'غ': {'\u063a', '\ufecf', '\ufed0', '\ufece'},
	'ک': {'\u06a9', '\ufedb', '\ufedc', '\ufeda'},
	'گ': {'\u06af', '\ufb94', '\ufb95', '\ufb93'},
	'ل': {'\u0644', '\ufedf', '\ufee0', '\ufede'},
	'ن': {'\u0646', '\ufee7', '\ufee8', '\ufee6'},
	'ه': {'\u0647', '\ufeeb', '\ufeec', '\ufeea'},
	'و': {'\u0648', '\u0648', '\u0648', '\ufeee'},
	'ف': {'\u0641', '\ufed3', '\ufed4', '\ufed2'},
	'ق': {'\u0642', '\ufed7', '\ufed8', '\ufed6'},
	'ی': {'\u06cc', '\ufef3', '\ufef4', '\ufef2'},
	'م': {'\u0645', '\ufee3', '\ufee4', '\ufee2'},
	'پ': {'\ufb56', '\ufb58', '\ufb59', '\ufb57'},
}

func ReShape(input string) string {
	var langSections []string
	var continousLangAr string
	var continousLangLt string

	for _, letter := range input {
		if IsPersianLetter(letter) {
			if len(continousLangLt) > 0 {
				langSections = append(langSections, strings.TrimSpace(continousLangLt))
			}
			continousLangLt = ""
			continousLangAr += string(letter)
		} else {
			if len(continousLangAr) > 0 {
				langSections = append(langSections, strings.TrimSpace(continousLangAr))
			}
			continousLangAr = ""
			continousLangLt += string(letter)
		}
	}
	if len(continousLangLt) > 0 {
		fmt.Println(continousLangLt)
		langSections = append(langSections, strings.TrimSpace(continousLangLt))
	}
	if len(continousLangAr) > 0 {
		langSections = append(langSections, strings.TrimSpace(continousLangAr))
	}

	var shapedSentence []string
	for _, section := range langSections {
		if IsPersian(section) {
			for _, word := range strings.Fields(section) {
				shapedSentence = append(shapedSentence, shapeWord(word))
			}
		} else {
			shapedSentence = append(shapedSentence, section)
		}
	}
	return strings.Join(ReverseContinuousPersian(shapedSentence), " ")
}

func ReverseContinuousPersian(words []string) []string {
	var result []string
	var persianWords []string

	for _, word := range words {
		if IsPersian(word) {
			persianWords = append(persianWords, word)
		} else {
			if len(persianWords) > 0 {
				for i := len(persianWords) - 1; i >= 0; i-- {
					result = append(result, persianWords[i])
				}
				persianWords = nil
			}
			result = append(result, word)
		}
	}

	for i := len(persianWords) - 1; i >= 0; i-- {
		result = append(result, persianWords[i])
	}

	return result
}

func isAlwaysInitial(letter rune) bool {
	alwaysInitial := [13]rune{'\u0627', '\u0623', '\u0622', '\u0625', '\u0649', '\u0621', '\u0624', '\u0629', '\u062f', '\u0630', '\u0631', '\u0632', '\u0648'}
	for _, item := range alwaysInitial {
		if item == letter {
			return true
		}
	}
	return false
}

func IsDigit(letter rune) bool {
	return letter >= 0x6f0 && letter <= 0x6f9
}

func IsWordDigit(word string) bool {
	var isPersian = false
	for _, v := range word {
		if IsDigit(v) {
			isPersian = true
			break
		}
	}
	return isPersian
}

func shapeWord(input string) string {
	if !IsPersian(input) {
		return input
	}

	var shapedInput bytes.Buffer

	inputRunes := []rune(input)
	for i := range inputRunes {
		var backLetter, frontLetter rune
		if i-1 >= 0 {
			backLetter = inputRunes[i-1]
		}
		if i != len(inputRunes)-1 {
			frontLetter = inputRunes[i+1]
		}
		if _, ok := glyphs[inputRunes[i]]; ok {
			adjustedLetter := adjustLetter(letterGroup{backLetter, inputRunes[i], frontLetter})
			shapedInput.WriteRune(adjustedLetter)
		} else {
			shapedInput.WriteRune(inputRunes[i])
		}
	}
	if len([]rune(shapedInput.String())) == len([]rune(input)) && !IsWordDigit(shapedInput.String()) {
		return reverse(shapedInput.String())
	}

	var shapedInputTashkeel bytes.Buffer
	inputTashkeelRunes := []rune(input)

	letterIndex := 0
	for i := range inputTashkeelRunes {
		if _, ok := glyphs[inputTashkeelRunes[i]]; ok {
			shapedInputTashkeel.WriteRune([]rune(shapedInput.String())[letterIndex])
			letterIndex++
		} else {
			shapedInputTashkeel.WriteRune(inputTashkeelRunes[i])
		}
	}

	if !IsWordDigit(shapedInputTashkeel.String()) {
		return reverse(shapedInputTashkeel.String())
	} else {
		return shapedInputTashkeel.String()
	}
}

func adjustLetter(g letterGroup) rune {
	switch {
	case g.backLetter > 0 && g.frontLetter > 0:
		if isAlwaysInitial(g.backLetter) {
			return glyphs[g.letter].Initial
		}
		return glyphs[g.letter].Medial

	case g.backLetter == 0 && g.frontLetter > 0:
		return glyphs[g.letter].Initial

	case g.backLetter > 0 && g.frontLetter == 0:
		if isAlwaysInitial(g.backLetter) {
			return glyphs[g.letter].Independent
		}
		return glyphs[g.letter].Final

	default:
		return glyphs[g.letter].Independent
	}
}

func IsPersianLetter(ch rune) bool {
	return (ch >= 0x600 && ch <= 0x6FF) || (ch >= 0xfb56 && ch <= 0xfef4)
}

func IsPersian(input string) bool {
	var isPersian = true
	for _, v := range input {
		if !unicode.IsSpace(v) && !IsPersianLetter(v) {
			isPersian = false
			break
		}
	}
	return isPersian
}
