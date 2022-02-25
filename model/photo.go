package model

import "time"

type PhotoAlbum struct {
	Photos []Photo
}
type Photo struct {
	 Timestamp time.Time
	 Latitude float64
	 Longitude float64
}
