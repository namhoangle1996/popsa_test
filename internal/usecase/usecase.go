package usecase

import (
	 "popsa_test/config"
	 "popsa_test/internal/usecaseInterface"
	 "popsa_test/suggestion"
	 "popsa_test/thirdPartyReverseGeo"
)

type UseCase struct {
	 usecaseInterface.ISuggestion
	 usecaseInterface.IThirdPartyReverseGeocode
}

func NewUseCase(config config.Config) *UseCase {
	return &UseCase{
		 ISuggestion:               suggestion.GetSuggestion(suggestion.SuggestionSentenceType1),
		 IThirdPartyReverseGeocode: thirdPartyReverseGeo.GetProviderApiService(config, thirdPartyReverseGeo.GoogleMaps),
	}
}
