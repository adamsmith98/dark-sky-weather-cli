# dark-sky-weather-cli
Command-line interface for Dark Sky weather API

## Usage
Create API_KEY file containing your Dark Sky API key in the same directory as main.go
```
go build main.go geo.go
./main <options>
``` 
Example (get next 4 days forecast)
```
./main -days=4
```
## Development

### Next steps 
* Refactor
* Unit testing
* Include weather icons
