# Device api
Device api is capable of creating/CIDR device storage and CIDR validation checks.
using this api you can. 
-	CreateDevice
-	View DeviceList
-	View/Search Device By Ipaddress

## Installation
 go get github.com/saurabh2013/CRPCDemo

## Usage 

	1. Compile and Run
		go run main.go 
		Running binary will start api at port :8080
		
	2. Try running these urls.

	-	Crete New Device 
	
		http://localhost:8080/CreateDevice?name=printer&desc=device%20one&ipaddress=1.2.2.1
		http://localhost:8080/CreateDevice?name=printer&desc=device%20one&ipaddress=1.2.2.1
		
		Test validation
		http://localhost:8080/CreateDevice?name=printer&desc=device%20one
		http://localhost:8080/CreateDevice?name=printer&desc=device%20one&ipaddress=1.22
	
	-	View Devices
		http://localhost:8080/DeviceList
	
	-	Search Device
		http://localhost:8080/DeviceList?ipaddress=1.2.2.1

## Log
File log:
	Only device created successfully will be logged into file. currently a fixed name is given to this log file 'log.log', this file will be create at the root folder only.
	This log file will be created automatically on runtime.
	
Console Log:
	Other logs can be found as console outputs.

 