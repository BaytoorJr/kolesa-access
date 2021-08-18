package domain

import "time"

type Car struct {
	ID               int     `json:"id"`
	Link             string  `json:"link"`
	Title            string  `json:"title"`
	ShortDescription string  `json:"short_description"`
	Price            int     `json:"price"`
	AvgPrice         int     `json:"avg_price"`
	PostDate         string  `json:"post_date"`
	ViewsCount       int     `json:"views_count"`
	City             string  `json:"city"`
	CarMark          string  `json:"car_mark"`
	CarModel         string  `json:"car_model"`
	CarYear          int     `json:"car_year"`
	CarType          string  `json:"car_type"`
	EngineVolume     float64 `json:"engine_volume"`
	Transmission     string  `json:"transmission"`
	WheelLocation    string  `json:"wheel_location"`
	Color            string  `json:"color"`
	WheelDrive       string  `json:"wheel_drive"` // привод машины
	IsKz             bool    `json:"rastamozhen"` // растаможен
	FullDescription  string  `json:"full_description"`
	Promotions       string  `json:"promotions"`
	PromotionsPrice  int     `json:"promotions_price"`
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
}
