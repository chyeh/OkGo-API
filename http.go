package main

import (
	commonGin "github.com/Cepave/open-falcon-backend/common/gin"
	"github.com/Cepave/open-falcon-backend/common/gin/mvc"
	"googlemaps.github.io/maps"

	gin "gopkg.in/gin-gonic/gin.v1"
)

var router *gin.Engine = nil
var GinConfig *commonGin.GinConfig = &commonGin.GinConfig{}

func InitGin(config *commonGin.GinConfig) {
	if router != nil {
		return
	}

	router = commonGin.NewDefaultJsonEngine(config)

	logger.Infof("Going to start web service. Listen: %s", config)

	initApi()

	go commonGin.StartServiceOrExit(router, config)

	*GinConfig = *config
}

func initApi() {
	mvcBuilder := mvc.NewMvcBuilder(mvc.NewDefaultMvcConfig())
	router.GET("/place/nearbysearch", mvcBuilder.BuildHandler(getPlaceNearbySearch))
	router.GET("/place/details", mvcBuilder.BuildHandler(getPlaceDetails))
	router.GET("/attractions", mvcBuilder.BuildHandler(getAttractions))
	router.GET("/attractions/hotels", mvcBuilder.BuildHandler(getHotelsForAttractions))
}

func getPlaceDetails(
	p *struct {
		PlaceID  string `mvc:"query[place_id]"`
		Language string `mvc:"query[language]"`
	},
) mvc.OutputBody {
	return mvc.JsonOutputBody(GetPlaceDetails(&maps.PlaceDetailsRequest{
		PlaceID: p.PlaceID,
	}))
}

func getPlaceNearbySearch(
	p *struct {
		Location  string `mvc:"query[location]"`
		Radius    uint   `mvc:"query[radius]"`
		RankBy    string `mvc:"query[rankby]"`
		Keyword   string `mvc:"query[keyword]"`
		Language  string `mvc:"query[language]"`
		MinPrice  string `mvc:"query[minprice]"`
		MaxPrice  string `mvc:"query[maxprice]"`
		Name      string `mvc:"query[name]"`
		OpenNow   bool   `mvc:"query[opennow]"`
		Type      string `mvc:"query[types]"`
		PageToken string `mvc:"query[pagetoken]"`
	},
) mvc.OutputBody {
	laglng, err := maps.ParseLatLng(p.Location)
	if err != nil {
		logger.Fatalln("laglng error:", err)
	}
	return mvc.JsonOutputBody(GetPlaceNearbySearch(&maps.NearbySearchRequest{
		Location:  &laglng,
		Radius:    p.Radius,
		Keyword:   p.Keyword,
		Language:  p.Language,
		MinPrice:  maps.PriceLevel(p.MinPrice),
		MaxPrice:  maps.PriceLevel(p.MaxPrice),
		Name:      p.Name,
		OpenNow:   p.OpenNow,
		RankBy:    maps.RankBy(p.RankBy),
		Type:      maps.PlaceType(p.Type),
		PageToken: p.PageToken,
	}))
}

func getAttractions(
	p *struct {
		PlaceID string `mvc:"query[place_id]"`
	},
) mvc.OutputBody {
	logger.Debugln("[/attractions] got place_id:", p.PlaceID)
	return mvc.JsonOutputBody(GetAttractions(p.PlaceID))
}

func getHotelsForAttractions(
	p *struct {
		Checkin  string `mvc:"query[checkin]"`
		Checkout string `mvc:"query[checkout]"`
		PlaceIDs string `mvc:"query[place_ids]"`
	},
) mvc.OutputBody {
	logger.Debugln("[/attractions/hotels] got checkin:", p.Checkin)
	logger.Debugln("[/attractions/hotels] got checkout:", p.Checkout)
	logger.Debugln("[/attractions/hotels] got place_ids:", p.PlaceIDs)
	return mvc.JsonOutputBody(GetHotelsForAttractions(p.Checkin, p.Checkout, p.PlaceIDs))
}
