package nlpt_detect

import (
	"github.com/jbowles/cld2_nlpt"
	"log"
)

// Detect is the default function for detecting a language using the cld2_nlpt wrapper. It requires the text to be detected, the format of the output (code ['en'], name ['ENGLISH'], declared name ['ENGLISH']), the size of the buffer, the ranked choice, the percent, and the normal score.
// NOTE: scored indices are inversely ranked. That is, if you want accuracy you should index the 4th element.
// For exmaple: Detect(english_text, "name", len(english_text), 3, 3, 3)
// From the description of CLD2:
//   language3 is an array of the top 3 languages or UNKNOWN_LANGUAGE
//   percent3 is an array of the text percentages 0..100 of the top 3 languages
// CLD2 returns the 3 highest ranked languages, the 3 best percentages, and the 3 best normalized scores... all of which are returned from CLD2 as arrays. The integer value here is the index of the array to return.
//
//
func Detect(s string, format string, buffer_length, rank, percent, normal_score int) string {
	lang, err := cld2_nlpt.DetectExtendedLanguage(s, format, buffer_length, rank, percent, normal_score)
	if err != nil {
		log.Fatal(err)
	}
	return string(cld2_nlpt.Language(lang))
}

func StaticDetect(s string) string {
	lang, err := cld2_nlpt.SimpleDetect(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(cld2_nlpt.Language(lang))
}

// GetLanguageName returns the the name('ENGLISH') of detected text.
// It should be used for testing but not the greatest amount of accuracy.
//
func GetLanguageName(s string) string {
	lang, err := cld2_nlpt.DetectLanguage(len(s), s, "name")
	if err != nil {
		log.Fatal(err)
	}
	return string(cld2_nlpt.Language(lang))
}

// GetLanguageCode returns the the name('en') of detected text.
// It should be used for testing but not the greatest amount of accuracy.
//
func GetLanguageCode(s string) string {
	lang, err := cld2_nlpt.DetectLanguage(len(s), s, "code")
	if err != nil {
		log.Fatal(err)
	}
	return string(cld2_nlpt.Language(lang))
}

// GetLanguageDeclaredName returns the the name('ENGLISH') of detected text.
// It should be used for testing but not the greatest amount of accuracy.
//
func GetLanguageDeclaredName(s string) string {
	lang, err := cld2_nlpt.DetectLanguage(len(s), s, "declname")
	if err != nil {
		log.Fatal(err)
	}
	return string(cld2_nlpt.Language(lang))
}
