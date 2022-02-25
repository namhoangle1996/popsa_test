package suggestionType1

import (
	 "context"
	 "popsa_test/internal/usecaseInterface"
	 "popsa_test/model"
)

type suggestionType1 struct {
}

func (s suggestionType1) Suggest(ctx context.Context, suggestion model.Suggestion) (o string) {
	return "A weekend in " + suggestion.Country // or suggestion.District (will implement later )
}

func NewSuggestionTyp1() usecaseInterface.ISuggestion {
	 return &suggestionType1{}
}
