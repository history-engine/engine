package readability

import "history-engine/engine/setting"

func Parser() Readability {
	switch setting.Readability.Parser {
	case "mozilla":
		return MozillaParser
	default:
		return MozillaParser
	}
}
