package thirdPartyReverseGeo

import (
	 "popsa_test/config"
	 "popsa_test/internal/usecaseInterface"
	 "popsa_test/thirdPartyReverseGeo/googleApi"
	 "popsa_test/thirdPartyReverseGeo/hereApi"
)


type ThirdPartyReverseGeocodeServiceName int64

const (
	 GoogleMaps ThirdPartyReverseGeocodeServiceName = iota
	 Here
)

func GetProviderApiService(config config.Config, t ThirdPartyReverseGeocodeServiceName) usecaseInterface.IThirdPartyReverseGeocode {
	 switch t {
	 case GoogleMaps:
		  return googleApi.NewGoggleImpl(config.GoogleApiKey)
	 case Here:
		  return hereApi.NewHereApiImpl()
	 }
	 panic("Wrong service provider !")
}
