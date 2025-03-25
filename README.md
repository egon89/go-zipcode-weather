# go zipcode weather

## Local environment
### Build image
```bash
docker build -t go-zipcode-weather .
```

### Run container with environment variables
```bash
docker run --name gzweather -p 8080:8080 \
  -e WEATHER_API_BASE_URL=http://api.weatherapi.com \
  -e WEATHER_API_KEY=weather_api_key \
  -e VIA_CEP_BASE_URL=https://viacep.com.br/ws \
  go-zipcode-weather:latest
```

### curl
```bash
curl --location 'http://localhost:8080/weather/{:zipcode}' --verbose
```
