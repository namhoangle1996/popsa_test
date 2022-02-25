package suggestion

import (
	 "popsa_test/internal/usecaseInterface"
	 "popsa_test/suggestion/suggestionType1"
	 "popsa_test/suggestion/suggestionType2"
	 "popsa_test/suggestion/suggestionType3"
)



type SuggestionType int64

const (
	 SuggestionSentenceType1 SuggestionType = iota
	 SuggestionSentenceType2
	 SuggestionSentenceType3
)


func GetSuggestion(t SuggestionType) usecaseInterface.ISuggestion {
	 switch t {
	 case SuggestionSentenceType1:
		  return suggestionType1.NewSuggestionTyp1()
	 case SuggestionSentenceType2:
		  return suggestionType2.NewSuggestionType2()
	 case SuggestionSentenceType3:
		  return suggestionType3.NewSuggestionType3()
	 }
	 panic("not implement")
}
