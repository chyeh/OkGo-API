package main

type APIAuthConfig struct {
	googlePlaceAPIKey     string
	bookingDotComUsername string
	bookingDotComPassword string
}

var APIConfig *APIAuthConfig = &APIAuthConfig{}

func InitAPIConfig(config *APIAuthConfig) {
	APIConfig = config
}
