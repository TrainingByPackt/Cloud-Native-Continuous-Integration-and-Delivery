# Cloud-native Book Server Microservice 
* REST API Server that works with any SQL database
* Cloud ready and all steps in Docker
* Gitlab CI Pipeline ready

## Requirements
* Docker (> version 17.05)
* GNU make
	
## Testing
All testing levels are implemented:
```
make static-code-check smoke-test unit-test integration-test
```

## Build
Production ready Docker container:
```
make prod
```

## Dependency Management
* [govendor](https://github.com/kardianos/govendor) is used for dependency management.
* Fixed versions can be checked from [vendor.json](vendor/vendor.json)
