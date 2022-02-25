package usecaseInterface

import (
	 "context"
	 "popsa_test/model"
)

type ISuggestion interface {
	 Suggest(ctx context.Context, suggestion model.Suggestion) (o string)
}

type IThirdPartyReverseGeocode interface {
	 GetLocationName(ctx context.Context, lat, long float64) (district, country string, err error)
}
