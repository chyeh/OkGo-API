package main

type HotelAvailabilityV2 struct {
	Longitude   string        `json:"longitude"`
	Latitude    string        `json:"latitude"`
	Hotels      []*Hotel      `json:"hotels"`
	GuestGroups []*GuestGroup `json:"guest_groups"`
	Checkin     string        `json:"checkin"`
	Checkout    string        `json:"checkout"`
	Radius      string        `json:"radius"`
}
type Hotel struct {
	HotelId           string   `json:"hotel_id"`
	ReviewScoreWord   string   `json:"review_score_word"`
	ReviewScore       string   `json:"review_score"`
	ReviewNr          int      `json:"review_nr"`
	Price             string   `json:"price"`
	HotelCurrencyCode string   `json:"hotel_currency_code"`
	Location          Location `json:"location"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GuestGroup struct {
	Guests   int      `json:"guests"`
	Children []string `json:"children"`
}
