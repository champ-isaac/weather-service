## Forecast

1. https://api.weather.gov/points/{latitude},{longitude}
   For example: https://api.weather.gov/points/39.7456,-97.0892
2. https://api.weather.gov/gridpoints/{office}/{gridX},{gridY}/forecast
   For example: https://api.weather.gov/gridpoints/TOP/31,80/forecast

## Note:
This will provide the grid forecast endpoints for three format options in these properties:

forecast - forecast for 12h periods over the next seven days
forecastHourly - forecast for hourly periods over the next seven days
forecastGridData - raw forecast data over the next seven days
   