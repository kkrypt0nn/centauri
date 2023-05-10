package discord

// Locale represents the different locales that are available on Discord
// https://discord.com/developers/docs/reference#locales
type Locale string

// LanguageName returns the language name of the locale
func (l Locale) LanguageName() string {
	if name, ok := LocaleLanguageNames[l]; ok {
		return name
	}
	return LocaleLanguageNames[LocaleUnknown]
}

// NativeName returns the language name of the locale in the language itself
func (l Locale) NativeName() string {
	if name, ok := LocaleNativeName[l]; ok {
		return name
	}
	return LocaleNativeName[LocaleUnknown]
}

const (
	LocaleIndonesian   Locale = "id"
	LocaleDanish       Locale = "da"
	LocaleGerman       Locale = "de"
	LocaleEnglishGB    Locale = "en-GB"
	LocaleEnglishUS    Locale = "en-US"
	LocaleSpanishES    Locale = "es-ES"
	LocaleFrench       Locale = "fr"
	LocaleCroatian     Locale = "hr"
	LocaleItalian      Locale = "it"
	LocaleLithuanian   Locale = "lt"
	LocaleHungarian    Locale = "hu"
	LocaleDutch        Locale = "nl"
	LocaleNorwegian    Locale = "no"
	LocalePolish       Locale = "pl"
	LocalePortugueseBR Locale = "pt-BR"
	LocaleRomanian     Locale = "ro"
	LocaleFinnish      Locale = "fi"
	LocaleSwedish      Locale = "sv-SE"
	LocaleVietnamese   Locale = "vi"
	LocaleTurkish      Locale = "tr"
	LocaleCzech        Locale = "cs"
	LocaleGreek        Locale = "el"
	LocaleBulgarian    Locale = "bg"
	LocaleRussian      Locale = "ru"
	LocaleUkrainian    Locale = "uk"
	LocaleHindi        Locale = "hi"
	LocaleThai         Locale = "th"
	LocaleChineseCN    Locale = "zh-CN"
	LocaleJapanese     Locale = "ja"
	LocaleChineseTW    Locale = "zh-TW"
	LocaleKorean       Locale = "ko"
	LocaleUnknown      Locale = "" // Not sure if that will ever be the case, nice to have
)

var LocaleLanguageNames = map[Locale]string{
	LocaleIndonesian:   "Indonesian",
	LocaleDanish:       "Danish",
	LocaleGerman:       "German",
	LocaleEnglishGB:    "English, UK",
	LocaleEnglishUS:    "English, US",
	LocaleSpanishES:    "Spanish",
	LocaleFrench:       "French",
	LocaleCroatian:     "Croatian",
	LocaleItalian:      "Italian",
	LocaleLithuanian:   "Lithuanian",
	LocaleHungarian:    "Hungarian",
	LocaleDutch:        "Dutch",
	LocaleNorwegian:    "Norwegian",
	LocalePolish:       "Polish",
	LocalePortugueseBR: "Portuguese, Brazilian",
	LocaleRomanian:     "Romanian, Romania",
	LocaleFinnish:      "Finnish",
	LocaleSwedish:      "Swedish",
	LocaleVietnamese:   "Vietnamese",
	LocaleTurkish:      "Turkish",
	LocaleCzech:        "Czech",
	LocaleGreek:        "Greek",
	LocaleBulgarian:    "Bulgarian",
	LocaleRussian:      "Russian",
	LocaleUkrainian:    "Ukrainian",
	LocaleHindi:        "Hindi",
	LocaleThai:         "Thai",
	LocaleChineseCN:    "Chinese, China",
	LocaleJapanese:     "Japanese",
	LocaleChineseTW:    "Chinese, Taiwan",
	LocaleKorean:       "Korean",
	LocaleUnknown:      "",
}

var LocaleNativeName = map[Locale]string{
	LocaleIndonesian:   "Bahasa Indonesia",
	LocaleDanish:       "Dansk",
	LocaleGerman:       "Deutsch",
	LocaleEnglishGB:    "English, UK",
	LocaleEnglishUS:    "English, US",
	LocaleSpanishES:    "Español",
	LocaleFrench:       "Français",
	LocaleCroatian:     "Hrvatski",
	LocaleItalian:      "Italiano",
	LocaleLithuanian:   "Lietuviškai",
	LocaleHungarian:    "Magyar",
	LocaleDutch:        "Nederlands",
	LocaleNorwegian:    "Norsk",
	LocalePolish:       "Polski",
	LocalePortugueseBR: "Português do Brasil",
	LocaleRomanian:     "Română",
	LocaleFinnish:      "Suomi",
	LocaleSwedish:      "Svenska",
	LocaleVietnamese:   "Tiếng Việt",
	LocaleTurkish:      "Türkçe",
	LocaleCzech:        "Čeština",
	LocaleGreek:        "Ελληνικά",
	LocaleBulgarian:    "български",
	LocaleRussian:      "Pусский",
	LocaleUkrainian:    "Українська",
	LocaleHindi:        "हिन्दी",
	LocaleThai:         "ไทย",
	LocaleChineseCN:    "中文",
	LocaleJapanese:     "日本語",
	LocaleChineseTW:    "繁體中文",
	LocaleKorean:       "한국어",
	LocaleUnknown:      "",
}
