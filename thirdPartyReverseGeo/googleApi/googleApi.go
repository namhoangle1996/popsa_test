package googleApi

import (
	 "context"
	 "encoding/json"
	 "fmt"
	 "io/ioutil"
	 "log"
	 "net/http"
	 "os"
	 "popsa_test/internal/usecaseInterface"
	 "popsa_test/model"
)

type googleApiImplement struct {
	 apikey string `json:"apikey"`
}

const url = "https://revgeocode.search.hereapi.com/v1/revgeocode?apiKey="

func (i googleApiImplement) GetLocationName(ctx context.Context, lat, long float64) (district, country string, err error) {
	 x := os.Getenv("ENV")
	 if x == "TEST" {
		return "HANOI", "VIETNAM", err
	 }

	 url := url + i.apikey + "&at=" + fmt.Sprint(lat) + "," + fmt.Sprint(long)

	 res, err := http.Get(url)
	 if err != nil {
		  log.Fatalln(err)
	 }
	 body, err := ioutil.ReadAll(res.Body)
	 if err != nil {
		  log.Fatalln(err)
		  return "", "", err
	 }

	 var data model.GoogleReverseGeoApiRes
	 json.Unmarshal(body, &data)
	 for _, v := range data.Items {
		  district = v.Address.District
		  country = v.Address.CountryName
		  return district, country, nil
	 }

	 return
}

func NewGoggleImpl(apiKey string) usecaseInterface.IThirdPartyReverseGeocode {
	 return &googleApiImplement{apikey: apiKey}
}
