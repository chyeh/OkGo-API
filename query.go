package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"googlemaps.github.io/maps"

	"github.com/dghubble/sling"
	"github.com/kr/pretty"
)

func GetPlaceDetails(req *maps.PlaceDetailsRequest) maps.PlaceDetailsResult {
	c, err := maps.NewClient(maps.WithAPIKey(APIConfig.googlePlaceAPIKey))
	if err != nil {
		logger.Fatalln(err)
	}
	resp, err := c.PlaceDetails(context.Background(), req)
	if err != nil {
		logger.Errorln("PlaceDetails error:", err)
	}
	pretty.Println(resp)
	return resp
}

func GetPlaceNearbySearch(req *maps.NearbySearchRequest) maps.PlacesSearchResponse {
	c, err := maps.NewClient(maps.WithAPIKey(APIConfig.googlePlaceAPIKey))
	if err != nil {
		logger.Fatalln(err)
	}
	resp, err := c.NearbySearch(context.Background(), req)
	if err != nil {
		logger.Errorln("NearbySearch error:", err)
	}

	pretty.Println(resp)
	return resp
}

func GetAttractions(placeID string) maps.PlacesSearchResponse {
	d := GetPlaceDetails(&maps.PlaceDetailsRequest{
		PlaceID: placeID,
	})
	location := d.Geometry.Location
	r := GetPlaceNearbySearch(&maps.NearbySearchRequest{
		Location: &location,
		Radius:   500,
		Type:     "museum|night_club|cafe",
	})

	return r
}

type GetHotalAvailabilityParam struct {
	Checkin   string  `url:"checkin,omitempty"`
	Checkout  string  `url:"checkout,omitempty"`
	Latitude  float64 `url:"latitude,omitempty"`
	Longitude float64 `url:"longitude,omitempty"`
	Radius    string  `url:"radius,omitempty"`
	Room1     string  `url:"room1,omitempty"`
	Output    string  `url:"output,omitempty"`
}

func GetHotelsForAttractions(checkin string, checkout string, pIDs string) HotelAvailabilityV2 {
	placeIDs := strings.Split(pIDs, ",")
	var results []maps.PlaceDetailsResult
	sumLat := float64(0.0)
	sumLng := float64(0.0)
	for _, id := range placeIDs {
		r := GetPlaceDetails(&maps.PlaceDetailsRequest{
			PlaceID: id,
		})
		sumLat = sumLat + r.Geometry.Location.Lat
		sumLng = sumLng + r.Geometry.Location.Lng
		results = append(results, r)
	}
	avgLat := sumLat / float64(len(placeIDs))
	avgLng := sumLng / float64(len(placeIDs))

	var hotelAv HotelAvailabilityV2
	req, err :=
		sling.New().SetBasicAuth(APIConfig.bookingDotComUsername, APIConfig.bookingDotComPassword).
			Get("https://distribution-xml.booking.com/json/getHotelAvailabilityV2").
			QueryStruct(&GetHotalAvailabilityParam{
				Checkin:   checkin,
				Checkout:  checkout,
				Latitude:  avgLat,
				Longitude: avgLng,
				Radius:    "5",
				Room1:     "A",
				Output:    "hotel_details",
			}).
			Request()

	if err != nil {
		logger.Infoln("[sling] request error:", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Infoln("[booking.com] API call error:", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Infoln("[booking.com] API response decode error:", err)
	}
	bodyString := string(bodyBytes)
	logger.Infoln(bodyString)
	err = json.Unmarshal(bodyBytes, &hotelAv)
	if err != nil {
		logger.Infoln("JSON decode error:", err)
	}
	logger.Debugf("%+v\n", hotelAv.Hotels)

	return hotelAv
}
