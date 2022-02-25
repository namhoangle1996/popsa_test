package main

import (
	 "context"
	 "fmt"
	 "github.com/joho/godotenv"
	 "github.com/spf13/cast"
	 "os"
	 "popsa_test/internal/usecase"

	 "popsa_test/config"
	 "popsa_test/model"
)

func init() {
	 godotenv.Load()
}

func main() {
	 //Step  1 : Read data from csv file and push to photo album model like this
	 var examplePhotos = model.PhotoAlbum{
		  Photos: []model.Photo{
				{cast.ToTime("2019-03-30 14:12:19"), 40.703717, -74.016094},
				{cast.ToTime("2019-03-31 12:18:04"), 40.703717, -74.016094},
				{cast.ToTime("2019-03-31 12:18:04"), 40.703717, -74.016094},
				{cast.ToTime("2019-03-30 11:18:04"), 40.703717, -74.016094},
				{cast.ToTime("2019-03-30 12:18:04"), 40.703717, -74.016094},
		  },
	 }

	 // Result
	 fmt.Println(output(examplePhotos.Photos))
}

func output(photos []model.Photo) string {
	 var configs config.Config
	 configs.GoogleApiKey = os.Getenv("google_api_key") // lets to put your google api key to config file

	 uc := usecase.NewUseCase(config.Config{})

	 var district, country string

	 //get District, Country
	 for _, v := range photos {
		  getDistrict, getCountry, err := uc.IThirdPartyReverseGeocode.GetLocationName(context.Background(), v.Latitude, v.Longitude)
		  if err == nil {
				district = getDistrict
				country = getCountry
		  }else {
		 		fmt.Println("err",err)
		  }
	 }

	 // todo : Calculate month the month base on timestamp ( using time.Month())
	 var month int

	 return uc.ISuggestion.Suggest(context.Background(), model.Suggestion{
		  Month:    month,
		  District: district,
		  Country:  country,
	 })
}
