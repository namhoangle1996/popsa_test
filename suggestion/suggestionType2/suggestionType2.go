package suggestionType2

import (
	 "context"
	 "popsa_test/internal/usecaseInterface"
	 "popsa_test/model"
)

type suggestionType2 struct {
}

func (s suggestionType2) Suggest(ctx context.Context, suggestion model.Suggestion) (o string) {
	 return "A trip to " + suggestion.Country // or suggestion.District (will implement later )
}

func NewSuggestionType2() usecaseInterface.ISuggestion {
	 return &suggestionType2{}
}
