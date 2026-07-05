package weather

import "time"

type ErrorResponse struct {
	Type            string `json:"type"`
	Title           string `json:"title"`
	Status          int    `json:"status"`
	Detail          string `json:"detail"`
	Instance        string `json:"instance"`
	CorrelationId   string `json:"correlationId"`
	AdditionalProp1 struct {
	} `json:"additionalProp1"`
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

type PointsResponse struct {
	Id               string           `json:"@id"`
	Type             string           `json:"@type"`
	ForecastOffice   string           `json:"forecastOffice"`
	GridId           string           `json:"gridId"`
	GridX            int              `json:"gridX"`
	GridY            int              `json:"gridY"`
	Forecast         string           `json:"forecast"`
	ForecastHourly   string           `json:"forecastHourly"`
	ForecastGridData string           `json:"forecastGridData"`
	RelativeLocation RelativeLocation `json:"relativeLocation"`
	TimeZone         string           `json:"timeZone"`
	RadarStation     string           `json:"radarStation"`
	Geometry         string           `json:"geometry"`
}

type RelativeLocation struct {
	City     string   `json:"city"`
	State    string   `json:"state"`
	Geometry string   `json:"geometry"`
	Distance Distance `json:"distance"`
	Bearing  Bearing  `json:"bearing"`
}

type Distance struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

type Bearing struct {
	UnitCode string `json:"unitCode"`
	Value    int    `json:"value"`
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

type ForecastResponse struct {
	Context           any       `json:"@context"`
	Geometry          string    `json:"geometry"`
	Units             string    `json:"units"`
	ForecastGenerator string    `json:"forecastGenerator"`
	GeneratedAt       time.Time `json:"generatedAt"`
	UpdateTime        time.Time `json:"updateTime"`
	ValidTimes        string    `json:"validTimes"`
	Elevation         Elevation `json:"elevation"`
	Periods           []Period  `json:"periods"`
}

type Elevation struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

type Period struct {
	Number                     int                        `json:"number"`
	Name                       string                     `json:"name"`
	StartTime                  time.Time                  `json:"startTime"`
	EndTime                    time.Time                  `json:"endTime"`
	IsDaytime                  bool                       `json:"isDaytime"`
	Temperature                Temperature                `json:"temperature"`
	TemperatureTrend           any                        `json:"temperatureTrend"`
	ProbabilityOfPrecipitation ProbabilityOfPrecipitation `json:"probabilityOfPrecipitation"`
	WindSpeed                  string                     `json:"windSpeed"`
	WindDirection              string                     `json:"windDirection"`
	Icon                       string                     `json:"icon"`
	ShortForecast              string                     `json:"shortForecast"`
	DetailedForecast           string                     `json:"detailedForecast"`
}

type Temperature struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

type ProbabilityOfPrecipitation struct {
	UnitCode string `json:"unitCode"`
	Value    int    `json:"value"`
}
