# GatorRenter-District-API-Golang
An API interface of popular housing districts in San Francisco. To run locally on localhost 3005, make sure you have 
Golang install. You can download here https://golang.org/dl/

## Endpoints
### /districts
return all relevant districts
![](/screenshot/api_1.jpg)

### /districts/{portion} 
return all relevant districts in the requested portion
Relevant portion includes (not case sensitive): 
* North
* NorthWest
* NorthEast
* Central
* West
* CentralWest
* East
* CentralEast
* South
* SouthWest
* SouthEast
![](/screenshot/api_2.jpg)

### query parameter lte={distance from sfsu}
return districts that is less than or equal to the requested distance. 
![](/screenshot/api_3.jpg)