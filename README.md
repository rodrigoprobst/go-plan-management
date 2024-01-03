# go-plan-management

## Description
Testing development with go and gin to create a microservice to manage plans

## Local Dependencies
- [go](https://go.dev/dl/) 1.21.3
- [make](https://www.gnu.org/software/make/) >= 4.3
- [docker](https://docs.docker.com/engine/install/linux-postinstall/) >= 24.0.7


## Setup
### GO
#### local GO build
```bash
make build/api
```
#### local GO run
```bash
make run/api
```
#### local GO run with live reload
```bash
make up
```
#### Test + sonarqube analysis (sonar container required) (sonarqube localhost:9000)
```bash
make sonar
```

#### Remove SonarQube container and network
```bash
make sonar-prune
```

#### Sync dependencies and check code format
```bash
make tidy
```

## Utilization
* Local http server available on 4600 port.
* SonarQube available on 9000 port.
* CodeCoverage for new code must match 100%.
