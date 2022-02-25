package hereApi

import (
	 "context"
	 "popsa_test/internal/usecaseInterface"
)

type hereApiImplement struct {
}

func (h hereApiImplement) GetLocationName(ctx context.Context, lat, long float64) (district, country string, err error) {
	 panic("Let s's integrate Here api here to get location")
}

func NewHereApiImpl() usecaseInterface.IThirdPartyReverseGeocode {
	 return &hereApiImplement{}
}
