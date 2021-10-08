package textproc

import (
	"hash/fnv"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// runes group by types, used for checking character type (Vietnamese alphabet)
var (
	numerics    = "0123456789"
	lowerAlphas = "aàáãạảăắằẳẵặâấầẩẫậbcdđeèéẹẻẽêếềểễệfghiìíĩỉị" +
		"jklmnoòóõọỏôốồổỗộơớờởỡợpqrstuùúũụủưứừửữựvwxyýỳỵỷỹz"
	upperAlphas = strings.ToUpper(lowerAlphas)

	AlphaNumeric = toMapRunes(numerics + lowerAlphas + upperAlphas)

	AlphaNumericList   = []rune(numerics + lowerAlphas + upperAlphas)
	AlphaNumericEnList = []rune(
		"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	AlphaEnList = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() { rand.Seed(time.Now().UnixNano()) }

func toMapRunes(s string) map[rune]bool {
	ret := make(map[rune]bool, len(s))
	for _, char := range []rune(s) {
		ret[char] = true
	}
	return ret
}

// RemoveRedundantSpace replaces continuous spaces with one space
func RemoveRedundantSpace(text string) string {
	tokens := strings.FieldsFunc(text, checkIsSpace)
	text = strings.Join(tokens, " ")
	lines := strings.Split(text, "\n")
	builder := strings.Builder{}
	builder.Grow(len(text))
	for i, line := range lines {
		tmp := strings.TrimSpace(line)
		if tmp != "" {
			builder.WriteString(tmp)
			if i != len(lines)-1 {
				builder.WriteString("\n")
			}
		}
	}
	return builder.String()
}

// checkIsSpace returns false for newline
func checkIsSpace(char rune) bool {
	switch char {
	case '\t', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	default:
		return false
	}
}

// checkIsSpaceNL returns true for newline
func checkIsSpaceNL(char rune) bool {
	switch char {
	case '\t', '\v', '\f', '\r', ' ', 0x85, 0xA0, '\n':
		return true
	default:
		return false
	}
}

// HashTextToInt is a unique and fast hash func
func HashTextToInt(word string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(word))
	return int64(h.Sum64())
}

func GenRandomWord(minLen int, maxLen int, charList []rune) string {
	if minLen <= 0 {
		minLen = 0
	}
	if maxLen < minLen {
		maxLen = minLen
	}
	wordLen := minLen + rand.Intn(maxLen+1-minLen)
	builder := strings.Builder{}
	builder.Grow(3 * wordLen) // UTF8
	for i := 0; i < wordLen; i++ {
		builder.WriteRune(charList[rand.Intn(len(charList))])
	}
	return builder.String()
}

// GenRandomVarName returns an alpha numeric string, first char is a letter
func GenRandomVarName(wordLen int) string {
	if wordLen <= 0 {
		return ""
	}
	builder := strings.Builder{}
	builder.Grow(wordLen)
	builder.WriteRune(AlphaEnList[rand.Intn(len(AlphaEnList))])
	for i := 1; i < wordLen; i++ {
		builder.WriteRune(AlphaNumericEnList[rand.Intn(len(AlphaNumericEnList))])
	}
	return builder.String()
}

// TextToWords splits a text to list of words (punctuations removed)
func TextToWords(text string) []string {
	ret := make([]string, 0)
	wordsWithPun := strings.FieldsFunc(text, checkIsSpaceNL)
	for _, wordWP := range wordsWithPun {
		runes := []rune(wordWP)
		firstAlphaNumeric := -1
		for i, r := range runes {
			if AlphaNumeric[r] {
				firstAlphaNumeric = i
				break
			}
		}
		if firstAlphaNumeric == -1 {
			continue
		}
		lastAlphaNumeric := len(runes)
		for i := len(runes) - 1; i >= 0; i-- {
			if AlphaNumeric[runes[i]] {
				lastAlphaNumeric = i
				break
			}
		}
		if lastAlphaNumeric == len(runes) {
			continue
		}
		ret = append(ret, string(runes[firstAlphaNumeric:lastAlphaNumeric+1]))
	}
	return ret
}

// WordsToNGrams creates a set of n-gram from input words,
// (A n-gram is a contiguous sequence of n words)
func WordsToNGrams(words []string, n int) map[string]int {
	result := make(map[string]int, len(words))
	for i := 0; i < len(words)-n+1; i++ {
		nGram := strings.Join(words[i:i+n], " ")
		result[nGram] += 1
	}
	return result
}

// TextToNGrams creates a set of n-gram (lowercase) from input text
func TextToNGrams(text string, n int) map[string]int {
	text = strings.ToLower(text)
	words := TextToWords(text)
	return WordsToNGrams(words, n)
}

// There are often several ways to represent the same string. For example,
// an "é" can be represented in a string as a single rune ("\u00e9")
// or an "e" followed by an acute accent ("e\u0301").
// They should be treated as equal in text processing.
// Vietnamese text has an extra problem: diacritic position,
// example: old style: òa, óa, ỏa, õa, ọa; new style: oà, oá, oả, oã, oạ
func NormalizeText(text string) string {
	transformer := transform.Chain(norm.NFKD, norm.NFKC)
	ret, _, _ := transform.String(transformer, text)
	// TODO: Vietnamese diacritic position
	return ret
}

func removeVietnamDiacritic(char rune) rune {
	switch char {
	case 'à', 'á', 'ã', 'ạ', 'ả', 'ă', 'ắ', 'ằ', 'ẳ', 'ẵ', 'ặ', 'â', 'ấ', 'ầ', 'ẩ', 'ẫ', 'ậ':
		return 'a'
	case 'đ':
		return 'd'
	case 'è', 'é', 'ẹ', 'ẻ', 'ẽ', 'ê', 'ế', 'ề', 'ể', 'ễ', 'ệ':
		return 'e'
	case 'ì', 'í', 'ĩ', 'ỉ', 'ị':
		return 'i'
	case 'ò', 'ó', 'õ', 'ọ', 'ỏ', 'ô', 'ố', 'ồ', 'ổ', 'ỗ', 'ộ', 'ơ', 'ớ', 'ờ', 'ở', 'ỡ', 'ợ':
		return 'o'
	case 'ù', 'ú', 'ũ', 'ụ', 'ủ', 'ư', 'ứ', 'ừ', 'ử', 'ữ', 'ự':
		return 'u'
	case 'ý', 'ỳ', 'ỵ', 'ỷ', 'ỹ':
		return 'y'
	case 'À', 'Á', 'Ã', 'Ạ', 'Ả', 'Ă', 'Ắ', 'Ằ', 'Ẳ', 'Ẵ', 'Ặ', 'Â', 'Ấ', 'Ầ', 'Ẩ', 'Ẫ', 'Ậ':
		return 'A'
	case 'Đ', 'Ð':
		return 'D'
	case 'È', 'É', 'Ẹ', 'Ẻ', 'Ẽ', 'Ê', 'Ế', 'Ề', 'Ể', 'Ễ', 'Ệ':
		return 'E'
	case 'Ì', 'Í', 'Ĩ', 'Ỉ', 'Ị':
		return 'I'
	case 'Ò', 'Ó', 'Õ', 'Ọ', 'Ỏ', 'Ô', 'Ố', 'Ồ', 'Ổ', 'Ỗ', 'Ộ', 'Ơ', 'Ớ', 'Ờ', 'Ở', 'Ỡ', 'Ợ':
		return 'O'
	case 'Ù', 'Ú', 'Ũ', 'Ụ', 'Ủ', 'Ư', 'Ứ', 'Ừ', 'Ử', 'Ữ', 'Ự':
		return 'U'
	case 'Ý', 'Ỳ', 'Ỵ', 'Ỷ', 'Ỹ':
		return 'Y'
	default:
		return char
	}
}

// example: Đào => Dao
func RemoveVietnamDiacritic(text string) string {
	transformer := transform.Chain(
		norm.NFKD, runes.Remove(runes.In(unicode.Mn)), norm.NFKC)
	text, _, _ = transform.String(transformer, text)
	chars := make([]rune, 0)
	for _, r := range []rune(text) {
		chars = append(chars, removeVietnamDiacritic(r))
	}
	return string(chars)
}
