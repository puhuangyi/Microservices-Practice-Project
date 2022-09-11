# Back end API document

- - -

## registered
for user register
+ **post** : *http://server_ip:8080/registered*
+ #### Request parameters
|Parameter names|types|mandatory|sample value|instructions|
|:---:|:---:|:---:|:---:|:---:|
|userID|string|true|marry003|ID for login|
|email|string|true|xxxxx@ncl.ac.uk|user email|
|password|string|true|asbdhdjsaawbadkj|Please encrypt it before sending it|

+ #### Response
|Status code|Content format|
|:---:|:---:|
|200|json|

+ #### Response sample
1. Success Eample(200):
```json
{
  "status": "true",
  "message": "user created successful"
}
```
2. Exception Example(200):
```json
{
  "status": "false",
  "message": "user have been create"
}
```
<br>

## login
for user login. When user login successful, back end will return a token.
+ **post** : *http://server_ip:8080/login*
+ #### Request parameters
|Parameter names|types|mandatory|sample value|instructions|
|:---:|:---:|:---:|:---:|:---:|
|userID|string|true|marry003|user ID which entered during user registration|
|password|string|true|adbasjdasjkdwqiuasdasd|user account password (Encrypt then send)|

+ #### Response
|Status code|Content format|
|:---:|:---:|
|200|json|

+ #### Response sample
1. Success Eample(200):
```json
{
  "status": "ture",
  "message": "login successful",
  "token": "juasgbduyahdvbwuayhdvbuiayshvdbauisydvbiyw"
}
```

2. Exception Example(200):
```json
{
  "status": "false",
  "message": "User does not exist",
  "token": ""
}
```
<br>

## payment
pay the ticket
+ **post** : *http://server_ip:8080/payment/horsepay*
  
+ #### Request parameters
|Parameter names|location|types|mandatory|sample value|instructions|
|:---:|:---:|:---:|:---:|:---:|:---:|
|userID|form-data|string|true|marry001|user ID|
|amount|form-data|double|true|35.6|price|
|start|form-data|string|true|Newcastle|'Newcastle' only now|
|destination|form-data|string|true|Alnwick Castle|castle name|
|number|form-data|string|true|3|number of person|
|date|form-data|string|true|2022-05-25|travelling date|
|type|form-data|string|true|return|ticket type|
|routerDetailID|form-data|int|true|1|router Detail ID|
|cardNumber|form-data|int|false|1418635131321968|bank card number|
|securityCode|form-data|int|false|364|bank card security code|
|Authorization|header|string|true|akdjbd4a8w6dad56|token|

+ #### Response
|Status code|Content format|
|:---:|:---:|
|200|json|
|401|text|
|406|text|

+ #### Response sample
1. Success Eample(200):
```json
{
  "status": "true",
  "reason": ""
}
```

2. Exception Example(200):
 ```json
{
  "status": "false",
  "reason": "bank refuse transfer"
}
 ```
3. Exception Example(401):
```text
Please provide token
```
4. Exception Example(406):
```text
Please provide full parameter
```
<br>


## location
Queries the location of all buses on a specified bus route
+ **Get** : *http://server_ip:8080/query/location/{busName}*
  
+ #### Request parameters
|Parameter names|location|types|mandatory|sample value|instructions|
|:---:|:---:|:---:|:---:|:---:|:---:|
|busName|path|string|true|12A|Bus name|
|noc|query|string|true|GNEL|The national operator code that operates the bus line|
|Authorization|header|string|true|akdjbd4a8w6dad56|token|
sample: http://server_ip:8080/query/location/12A?noc=GNEL


+ #### Response
|Status code|Content format|
|:---:|:---:|
|200|json|
|401|text|
|406|text|

+ #### Response sample
1. Success Eample(200):
```json
{
  //if seatFree or wheelFree equal -1
  //means that could not found seat and wheel data
  "info": [
    {
      "busRef": "6071",
      "longitude": -1.776278,
      "latitude": 54.961771,
      "seatFree": 57,
      "wheelFree": 1
    },
    {
      "busRef": "6084",
      "longitude": -1.649175,
      "latitude": 54.958566,
      "seatFree": 60,
      "wheelFree": 1 
    },
    {
      "busRef": "3962",
      "longitude": -1.743088,
      "latitude": 54.972376,
      "seatFree": 45,
      "wheelFree": 0
    },
    {
      "busRef": "6050",
      "longitude": -1.617413,
      "latitude": 54.973636,
      "seatFree": 67,
      "wheelFree": 1
    }
  ]
}
```
2. Exception Example(401):
```text
please provide token
```
3. Exception Example(406):
```text
Please provide full parameter
```
<br>


## Order
Queries the order information
+ **Get** : *http://server_ip:8080/query/order*
  
+ #### Request parameters
|Parameter names|location|types|mandatory|sample value|instructions|
|:---:|:---:|:---:|:---:|:---:|:---:|
|userID|query|string|true|marry006|user ID|
|Authorization|header|string|ture|bhg3vf12uy3v213vujh|token|

+ #### Response
|Status code|Content format|
|:---:|:---:|
|200|json|
|401|text|
|406|text|

+ #### Response sample
1. Success Eample(200):
```json
{
  "orders": [
    {
      "orderID": 1,
      "amount": 385.5,
      "start": "Newcastle",
      "destination": "Alnwick Castle",
      "number": 5,
      "routeDetailID": 2,
      "date": "2022-05-20",
      "type": "return",
      "time": "2022-05-04 21:34:09.99"
    },
    {
      "orderID": 22,
      "amount": 163.8,
      "start": "Newcastle",
      "destination": "Alnwick Castle",
      "number": 3,
      "routeDetailID": 2,
      "date": "2022-06-29",
      "type": "return",
      "time": "2022-05-06 19:24:27.18"
    }
  ]
}
```
2. Exception Example(401):
```text
please provide token
```
3. Exception Example(406):
```text
Please provide full parameter
```

<br>

## Timetable (developing)
Queries the timetable if specified bus route
+ **Get** : *http://server_ip:8080/query/timetable/{busName}*
  
+ #### Request parameters
|Parameter names|location|types|mandatory|sample value|instructions|
|:---:|:---:|:---:|:---:|:---:|:---:|
|busName|path|string|ture|X21|Bus name|
|noc|query|string|ture|GNEL|National operator code|
|day|query|int|true|1|day of week. 1 is Monday.|
|Authorization|header|string|ture|d1as60da78dw8|token|

+ #### Response
|Status code|Content format|
|:---:|:---:|
|200|json|
|401|text|
|406|text|

+ #### Response sample
1. Success Eample(200):
```json
{
  
}
```
2. Exception Example(401):
```text
please provide token
```
3. Exception Example(406):
```text
Please provide full parameter
```

<br>

## Route
Queries the route to castle
+ **Get** : *http://server_ip:8080/query/route*
  
+ #### Request parameters
|Parameter names|location|types|mandatory|sample value|instructions|
|:---:|:---:|:---:|:---:|:---:|:---:|
|original|query|string|true|Newcastle|Start address. only "Newcastle" now.|
|destination|query|string|true|Auckland Castle|Castle name|
|Authorization|header|string|true|0asd186a5110da|token|

+ #### Response
|Status code|Content format|
|:---:|:---:|
|200|json|
|401|text|
|406|text|

+ #### Response sample
1. Success Eample(200):
```json
{
  "routeID": 3,
  "info": [
    {
      "original": "Newcastle Eldon Square - Stand F",
      "destination": "Market Place",
      "type": "bus",
      "transportName": "X21",
      "noc": "GNEL",
      "step": 1,
      "comment": "get off at Market Place",
      "time": 84
    },
    {
      "original": "Market Place",
      "destination": "Auckland Castle",
      "type": "walk",
      "transportName": "walk",
      "noc": "",
      "step": 2,
      "comment": "Head east from Market Pl, walk 300m",
      "time": 4
    }
  ]
}
```
2. Exception Example(401):
```text
please provide token
```
3. Exception Example(406):
```text
Please provide full parameter
```
<br>



## RoutePlus
Queries the route to castle and show all information of this route
+ **Get** : *http://server_ip:8080/query/route*
  
+ #### Request parameters
|Parameter names|location|types|mandatory|sample value|instructions|
|:---:|:---:|:---:|:---:|:---:|:---:|
|original|query|string|true|Newcastle|Start address. only "Newcastle" now.|
|destination|query|string|true|Auckland Castle|Castle name|
|day|query|int|true|1|1 is Monday... 5 is Friday|
|Authorization|header|string|true|0asd186a5110da|token|

+ #### Response
|Status code|Content format|
|:---:|:---:|
|200|json|
|401|text|
|406|text|

+ #### Response sample
1. Success Eample(200):
```json
{

  "routeID": 4,
  "original": "Newcastle",
  "destination": "Bamburgh Castle",
  "info": [
    {
      "original": "Newcastle Haymarket Bus Station",
      "destination": "Belford Fire Station",
      "type": "bus",
      "transportName": "X15",
      "noc": "ANUM",
      "step": 1,
      "comment": "Take X15 from Newcastle Haymarket Bus Station and get off at Belford Fire Station"
    },
    {
      "original": "Belford Fire Station",
      "destination": "Bamburgh Lord Crewe Hotel",
      "type": "bus",
      "transportName": "X18",
      "noc": "ANUM",
      "step": 2,
      "comment": "Take X18 from Belford Fire Station and get off at Bamburgh Lord Crewe Hotel"
    },
    {
      "original": "Bamburgh Lord Crewe Hotel",
      "destination": "Bamburgh Castle",
      "type": "walk",
      "transportName": "walk",
      "noc": " ",
      "step": 3,
      "comment": "walk for 8 minutes from Lord Crewe Hotel to Bamburgh Castle"
    }
  ],
  "time": [
    "07:23",
    "06:43",
    "07:23",
    "08:38",
    "09:38",
    "10:38",
    "11:38",
    "12:38",
    "13:38",
    "14:43",
    "15:48",
    "16:48",
    "17:53",
    "18:58",
    "19:58"
  ],
  "transportFare": [
    {
      "name": "X15",
      "type": "bus",
      "fare": 4.7
    },
    {
      "name": "X18",
      "type": "bus",
      "fare": 4.7
    }
  ]
}
```
2. Exception Example(401):
```text
please provide token
```
3. Exception Example(406):
```text
Please provide full parameter
```
4. Exception Example(406):
```text
please provide vaild parameter
```
5. Exception Example(406):
```text
no result, please check you parameter is correct!
```
<br>



## Fare of castle
Queries the fare of castle
+ **Get** : *http://server_ip:8080/query/fare/castle*
  
+ #### Request parameters
|Parameter names|location|types|mandatory|sample value|instructions|
|:---:|:---:|:---:|:---:|:---:|:---:|
|castle|query|string|true|Auckland Castle|Castle name|
|Authorization|header|string|true|0asd186a5110da|token|

+ #### Response
|Status code|Content format|
|:---:|:---:|
|200|json|
|401|text|
|406|text|

+ #### Response sample
1. Success Eample(200):
```json
{
  "fare": 35
}
```
2. Exception Example(401):
```text
please provide token
```
3. Exception Example(406):
```text
Please provide full parameter
```
<br>



## Fare of train
Queries the fare of train
+ **Get** : *http://server_ip:8080/query/fare/train/{trainName}*
  
+ #### Request parameters
|Parameter names|location|types|mandatory|sample value|instructions|
|:---:|:---:|:---:|:---:|:---:|:---:|
|trainName|path|string|true|LNER|Train name|
|Authorization|header|string|true|0asd186a5110da|token|

+ #### Response
|Status code|Content format|
|:---:|:---:|
|200|json|
|401|text|
|406|text|

+ #### Response sample
1. Success Eample(200):
```json
{
  "fare": 15.5
}
```
2. Exception Example(401):
```text
please provide token
```
3. Exception Example(406):
```text
Please provide full parameter
```
<br>



## Fare of bus
Queries the fare of bus
+ **Get** : *http://server_ip:8080/query/fare/bus/{busName}*
  
+ #### Request parameters
|Parameter names|location|types|mandatory|sample value|instructions|
|:---:|:---:|:---:|:---:|:---:|:---:|
|busName|path|string|true|X21|Bus name|
|Authorization|header|string|true|0asd186a5110da|token|

+ #### Response
|Status code|Content format|
|:---:|:---:|
|200|json|
|401|text|
|406|text|

+ #### Response sample
1. Success Eample(200):
```json
{
  "fare": 4.7
}
```
2. Exception Example(401):
```text
please provide token
```
3. Exception Example(406):
```text
Please provide full parameter
```
<br>



## Start time
Queries start time of route
+ **Get** : *http://server_ip:8080/query/info/startTime*
  
+ #### Request parameters
|Parameter names|location|types|mandatory|sample value|instructions|
|:---:|:---:|:---:|:---:|:---:|:---:|
|routeID|query|int|true|1|Route ID which in route response|
|day|query|int|true|1|1 is Monday... 7 is Sunday|
|Authorization|header|string|true|0asd186a5110da|token|

+ #### Response
|Status code|Content format|
|:---:|:---:|
|200|json|
|401|text|
|406|text|

+ #### Response sample
1. Success Eample(200):
```json
{
  "time": [
    "07:30",
    "08:00",
    "08:30",
    "09:30",
    "10:00",
    "10:30",
    "11:00",
    "11:30",
    "12:00",
    "12:30",
    "13:00",
    "13:30",
    "14:00"
  ]
}
```
2. Exception Example(401):
```text
please provide token
```
3. Exception Example(406):
```text
Please provide full parameter
```
<br>



## Castle Info
Queries information of castle, such as opening time
+ **Get** : *http://server_ip:8080/query/info/castle*
  
+ #### Request parameters
|Parameter names|location|types|mandatory|sample value|instructions|
|:---:|:---:|:---:|:---:|:---:|:---:|
|castle|query|string|true|Auckland Castle|Castle name|
|day|query|int|true|1|1 is Monday... 7 is Sunday|
|Authorization|header|string|true|0asd186a5110da|token|

+ #### Response
|Status code|Content format|
|:---:|:---:|
|200|json|
|401|text|
|406|text|

+ #### Response sample
1. Success Eample(200):
```json
{
  "isOpen": "true",
  "openTime": "11:00",
  "description": "Auckland Castle was established by Bishop Hugh Pudsey in 12th century,  in the reopening of 2019 as the Auckland project, the castle have been remodelled for making Durham more than a town to a touris"
}
```
2. Exception Example(401):
```text
please provide token
```
3. Exception Example(406):
```text
Please provide full parameter
```
<br>
