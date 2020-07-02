# GatorRenter-District-API-Golang
An API interface of popular housing districts in San Francisco. To run locally on localhost 3005, make sure you have 
Golang install. You can download here https://golang.org/dl/

## Endpoints
### /districts
return all relevant districts
![](/screenshot/picture1.jpg)

### /districts/{portion} 
return all relevant districts in the requested portion (North, NorthWest, NorthEast, West, East, South, SouthWest, SouthEast are all relevant endpoints) 
![](/screenshot/picture2.jpg)
![](/screenshot/picture3.jpg)

### query parameter lte={distance from sfsu}
return districts that is less than or equal to the requested distance. 
![](/screenshot/picture4.jpg)
![](/screenshot/picture5.jpg)