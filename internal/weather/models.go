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
	Id                  string           `json:"@id"`
	Type                string           `json:"@type"`
	Cwa                 string           `json:"cwa"`
	Type1               string           `json:"type"`
	ForecastOffice      string           `json:"forecastOffice"`
	GridId              string           `json:"gridId"`
	GridX               int              `json:"gridX"`
	GridY               int              `json:"gridY"`
	Forecast            string           `json:"forecast"`
	ForecastHourly      string           `json:"forecastHourly"`
	ForecastGridData    string           `json:"forecastGridData"`
	ObservationStations string           `json:"observationStations"`
	RelativeLocation    RelativeLocation `json:"relativeLocation"`
	ForecastZone        string           `json:"forecastZone"`
	County              string           `json:"county"`
	FireWeatherZone     string           `json:"fireWeatherZone"`
	TimeZone            string           `json:"timeZone"`
	RadarStation        string           `json:"radarStation"`
	AstronomicalData    AstronomicalData `json:"astronomicalData"`
	Nwr                 Nwr              `json:"nwr"`
	Context             []string         `json:"@context"`
	Geometry            string           `json:"geometry"`
}
type RelativeLocation struct {
	Context    []string `json:"@context"`
	Id         string   `json:"id"`
	Type       string   `json:"type"`
	Geometry   Geometry `json:"geometry"`
	Properties `json:"properties"`
}
type Geometry struct {
	Type        string `json:"type"`
	Coordinates []int  `json:"coordinates"`
	Bbox        []int  `json:"bbox"`
}
type Properties struct {
	City     string   `json:"city"`
	State    string   `json:"state"`
	Distance Distance `json:"distance"`
	Bearing  Bearing  `json:"bearing"`
}
type Distance struct {
	Value          int    `json:"value"`
	MaxValue       int    `json:"maxValue"`
	MinValue       int    `json:"minValue"`
	UnitCode       string `json:"unitCode"`
	QualityControl string `json:"qualityControl"`
}
type Bearing struct {
	Value          int    `json:"value"`
	MaxValue       int    `json:"maxValue"`
	MinValue       int    `json:"minValue"`
	UnitCode       string `json:"unitCode"`
	QualityControl string `json:"qualityControl"`
}
type AstronomicalData struct {
	Sunrise                   time.Time `json:"sunrise"`
	Sunset                    time.Time `json:"sunset"`
	Transit                   time.Time `json:"transit"`
	CivilTwilightBegin        time.Time `json:"civilTwilightBegin"`
	CivilTwilightEnd          time.Time `json:"civilTwilightEnd"`
	NauticalTwilightBegin     time.Time `json:"nauticalTwilightBegin"`
	NauticalTwilightEnd       time.Time `json:"nauticalTwilightEnd"`
	AstronomicalTwilightBegin time.Time `json:"astronomicalTwilightBegin"`
	AstronomicalTwilightEnd   time.Time `json:"astronomicalTwilightEnd"`
}
type Nwr struct {
	Transmitter    string `json:"transmitter"`
	SameCode       string `json:"sameCode"`
	AreaBroadcast  string `json:"areaBroadcast"`
	PointBroadcast string `json:"pointBroadcast"`
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

type ForecastResponse struct {
	Context           Context   `json:"@context"`
	Geometry          string    `json:"geometry"`
	Units             string    `json:"units"`
	ForecastGenerator string    `json:"forecastGenerator"`
	GeneratedAt       time.Time `json:"generatedAt"`
	UpdateTime        time.Time `json:"updateTime"`
	ValidTimes        string    `json:"validTimes"`
	Elevation         Elevation `json:"elevation"`
	Periods           []Period  `json:"periods"`
}

type Context struct {
	Version string `json:"@version"`
	Wx      string `json:"wx"`
	Geo     string `json:"geo"`
	Unit    string `json:"unit"`
	Vocab   string `json:"@vocab"`
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
