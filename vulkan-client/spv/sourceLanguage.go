package spv

type SourceLanguage uint32

//goland:noinspection GoUnusedConst
const (
	SourceLanguageUnknown      SourceLanguage = 0
	SourceLanguageESSL         SourceLanguage = 1
	SourceLanguageGLSL         SourceLanguage = 2
	SourceLanguageOpenClC      SourceLanguage = 3
	SourceLanguageOpenClCpp    SourceLanguage = 4
	SourceLanguageHLSL         SourceLanguage = 5
	SourceLanguageCppForOpenCl SourceLanguage = 6
	SourceLanguageSYCL         SourceLanguage = 7
	SourceLanguageMax          SourceLanguage = 0x7fffffff
)
