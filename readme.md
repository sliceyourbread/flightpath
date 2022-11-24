# Flight paths 
flight path calculates all possible paths and costs from a user inputted start and end destination  

flight paths has an internal structure travelData which contains all the locations and costs this  service currently provides flight paths for. This structure is expandable, new flights can be added in the same format as the exisiting data; the extra location costs will also need adding to all exisitings data. 

Current locations listed North to South: 
Castle Black,
Winterfell,
Riverrun,
King's Landing,

##  Build and Run 
if no binary is avaliable run:
 ``` go build flightpaths.go ```

``` ./flightpaths "{depart}" "{destination}" ``` 

## Testing
for unit test run: 
``` go test```

for benchmarks run
``` go test -bench=.```


## Benchmark
goos: darwin
goarch: amd64
pkg: loveHolidayFlightPaths
cpu: Intel(R) Core(TM) i5-8257U CPU @ 1.40GHz
BenchmarkFlightPaths4x4-8          91372             12068 ns/op
BenchmarkFlightPaths8x8-8          81765             14520 ns/op
BenchmarkCostcalculator4x4-8    35541592                33.40 ns/op
BenchmarkStopCalculator4x4-8     3328588               359.3 ns/op
