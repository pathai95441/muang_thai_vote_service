## Getting Started

how to run start server: 

```bash
# then clone git repo complete run install lib
make tidy

# then run start container 
make compose-dev 

# then run create database & migrate
make dbmate ARGS="create vote"
make dbmate ARGS+"migrate"

# for run unit-test
make unit-test
```
Server will start [http://localhost:8080]
