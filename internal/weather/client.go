package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	agent        = "myweatherapp.com, contact@myweatherapp.com"
	accept       = "application/ld+json"
	contentType  = "application/json"
	feature      = "forecast_temperature_qv"
	pointsUrl    = "https://api.weather.gov/points/%f,%f"
	forecastUrl  = "https://api.weather.gov/gridpoints/%s/%d,%d/forecast"
	errorRespFmt = "{\"statusCode\": %d, \"reason\": \"%s\", \"details\": \"%s\"}"
)

type Client struct {
	Client *http.Client
}

func New() *Client {
	return &Client{
		Client: &http.Client{},
	}
}

func (c *Client) Forecast(latitude, longitude, units string) (*ForecastResponse, error) {
	var resp any
	var err error
	lat, err := strconv.ParseFloat(latitude, 64)
	lng, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse latitude or longitude: %w", err)
	}
	resp, err = c.getPoints(lat, lng)
	if err != nil {
		return nil, err
	}

	var errorResponse *ErrorResponse
	var pointsResponse *PointsResponse
	switch resp.(type) {
	case *ErrorResponse:
		// got the error response
		errorResponse = resp.(*ErrorResponse)
	case *PointsResponse:
		// got the points response
		pointsResponse = resp.(*PointsResponse)
	}
	if errorResponse != nil {
		return nil, fmt.Errorf(errorRespFmt, errorResponse.Status, errorResponse.Title, errorResponse.Detail)
	}

	resp, err = c.getForecast(pointsResponse.GridId, pointsResponse.GridX, pointsResponse.GridY, units)
	if err != nil {
		return nil, err
	}
	var forecastResponse *ForecastResponse
	switch resp.(type) {
	case *ErrorResponse:
		errorResponse = resp.(*ErrorResponse)
	case *ForecastResponse:
		forecastResponse = resp.(*ForecastResponse)
	}
	if errorResponse != nil {
		return nil, fmt.Errorf(errorRespFmt, errorResponse.Status, errorResponse.Title, errorResponse.Detail)
	}
	return forecastResponse, nil
}

func (c *Client) do(url string, enableFeature bool, units string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("User-Agent", agent)
	req.Header.Set("Accept", accept)
	req.Header.Set("Content-Type", contentType)
	if enableFeature {
		req.Header.Set("Feature-Flags", feature)
		//append query parameter
		q := req.URL.Query()
		q.Add("units", units)
		req.URL.RawQuery = q.Encode()
		//log.Printf("Request URL: %s", req.URL.String())
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not do request: %w", err)
	}
	return resp, nil
}

func (c *Client) getPoints(lat, lng float64) (any, error) {
	var resp *http.Response
	var err error
	if resp, err = c.do(fmt.Sprintf(pointsUrl, lat, lng), false, ""); err != nil {
		return nil, err
	}
	defer c.close(resp)

	var bs []byte

	bs, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return c.unmarshalErrorResponse(bs)
	}

	var p PointsResponse
	err = json.Unmarshal(bs, &p)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}
	return &p, nil
}

func (c *Client) getForecast(officeId string, gridX, gridY int, units string) (any, error) {
	var resp *http.Response
	var err error
	if resp, err = c.do(fmt.Sprintf(forecastUrl, officeId, gridX, gridY), true, units); err != nil {
		return nil, err
	}
	defer c.close(resp)

	var bs []byte
	bs, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return c.unmarshalErrorResponse(bs)
	}
	var f ForecastResponse
	err = json.Unmarshal(bs, &f)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}
	return &f, nil
}

func (c *Client) close(resp *http.Response) {
	err := resp.Body.Close()
	if err != nil {
		log.Printf("error closing response body: %v", err)
	}
}

func (c *Client) unmarshalErrorResponse(bs []byte) (*ErrorResponse, error) {
	var errResp ErrorResponse
	err := json.Unmarshal(bs, &errResp)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %w", err)
	}
	return &errResp, nil
}
