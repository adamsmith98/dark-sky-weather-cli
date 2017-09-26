# dark-sky-weather-cli
Command-line interface for Dark Sky weather API

## Usage
Create API_KEY file containing your Dark Sky API key in the same directory as main.go
```
go build main.go geo.go
./main <options>
``` 
Example (show temperatures in degrees celsius)
```
./main -units=C
```
## Development

### Next steps 
* Add ability to request different times
* Include weather icons
