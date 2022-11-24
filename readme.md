# Flight paths 
flight path calculates all possible paths for flying and costs from user inputted start and end destinations  

flight paths has an internal structure travelData which contains all the locations and costs this  service currently provides flight paths for. This structure is expandable new flights can be added in the same format as the exisiting data; the extra location costs will also need adding to all exisitings data. 


# Build and Run 
if no binary is avaliable run:
 ``` go build flightpaths.go ```

``` ./flightpaths "{depart}" "{destination}" ``` 

Current locations listed North to South: 
Castle Black
Winterfell
Riverrun
King's Landing

## Benchmark

