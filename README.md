# weather-service
This service servers the forecast weather

## Service flow
1. Accepts latitude and longitude coordinates
2. Returns the short forecast for that area for Today ("Partly Cloudy" etc)
3. Returns a characterization of whether the temperature is "hot", "cold", or "moderate" (use your discretion on mapping temperatures to each type)
4. Use the National Weather Service API Web Service as a data source.

## Execution
1. Execute go run main.go to startup http server which listen to port 8080.
2. Run `curl http://127.0.0.1:8080/forecast\?latitude\=39.7456\&longitude\=-97.0892\&units\=si`
3. Or run `curl http://127.0.0.1:8080/forecast\?latitude\=39.7456\&longitude\=-97.0892` (default is si)
4. You can replace with preferred geolocation.

## Note
1. Currently I only implement units=si which is metrics.