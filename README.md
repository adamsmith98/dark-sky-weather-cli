# dark-sky-weather-cli
Command-line interface for Dark Sky weather API

## Usage
```
go build main.go geo.go forecast.go print.go
./main <options>
``` 
Example (get next 4 days forecast)
```
./main -days=4
```
Setup with key
```
./main -key=[API KEY]
```
## Development

### Next steps 
* Refactor
* Unit testing
* Include weather icons
