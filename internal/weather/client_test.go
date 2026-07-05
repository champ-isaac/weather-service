package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointsRequest_Success(t *testing.T) {
	client := New()
	resp, err := client.getPoints(39.7456, -97.0892)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	rp := resp.(*PointsResponse)
	assert.Equal(t, "TOP", rp.GridId)
	assert.Equal(t, 32, rp.GridX)
	assert.Equal(t, 81, rp.GridY)
}

func TestPointsRequest_ErrorResponse(t *testing.T) {
	client := New()
	resp, err := client.getPoints(-97.0892, 39.7456)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	rp := resp.(*ErrorResponse)
	assert.Equal(t, 404, rp.Status)
}

func TestForecastsRequest_Success(t *testing.T) {
	client := New()
	resp, err := client.getForecast("TOP", 32, 81, "si")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fp := resp.(*ForecastResponse)
	assert.Equal(t, "si", fp.Units)
}

func TestForecastsRequest_ErrorResponse(t *testing.T) {
	client := New()
	resp, err := client.getForecast("TOP", -32, 81, "si")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	rp := resp.(*ErrorResponse)
	assert.Equal(t, 404, rp.Status)

	resp, err = client.getForecast("TOP", 0, 0, "si")
	rp = resp.(*ErrorResponse)
	assert.Equal(t, 404, rp.Status)

	resp, err = client.getForecast("TOP", 0, 0, "metrics")
	rp = resp.(*ErrorResponse)
	assert.Equal(t, 400, rp.Status)
}

func TestClient_Forecast_Success(t *testing.T) {
	client := New()
	resp, err := client.Forecast(39.7456, -97.0892, "si")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.IsType(t, &ForecastResponse{}, resp)
}

func TestClient_Forecast_ErrorResponse(t *testing.T) {
	client := New()
	resp, err := client.Forecast(-97.0892, 39.7456, "si")
	assert.Error(t, err)
	assert.Nil(t, resp)
}
