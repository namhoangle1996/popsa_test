package model

type GoogleReverseGeoApiRes struct {
	 Items []struct {
		  Title   string `json:"title"`
		  Address struct {
				District    string
				CountryName string
		  }
	 } `json:"items"`
}
