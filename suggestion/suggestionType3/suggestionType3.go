package suggestionType3

import (
	 "context"
	 "github.com/spf13/cast"
	 "popsa_test/internal/usecaseInterface"
	 "popsa_test/model"
)

type suggestionType3 struct {
}

func (s suggestionType3) Suggest(ctx context.Context, suggestion model.Suggestion) (o string) {
	 return suggestion.District + " in " + cast.ToString(suggestion.Month)
}

func NewSuggestionType3() usecaseInterface.ISuggestion {
	 return &suggestionType3{}
}
