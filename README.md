# processing-api

## Configuration

The application uses environment variables for configuration. Create a `.env` file in the parent directory (`../`) with the following variables:

```
WEATHER_API_KEY=your_weatherapi_key_here
PORT=8080
```

## Running with Make

Build and run the container (automatically loads `.env` file):

```bash
make run
```

This will:
1. Build the container image
2. Run the container with environment variables from `../.env`
3. Expose the application on port 8080

## API Examples

Test the weather API:

```bash
curl http://localhost:8080/api/v1/weather/Czechia/Brno
```
