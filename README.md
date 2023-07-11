# Estimator

## Dev Startup

1. Export the following variables:

```
DB_HOST=localhost
DB_PORT=3306
DB_NAME=estimator
DB_USER=root
DB_PASS=*
```

2. Run `docker-compose up` in separate terminal.

3. Browse to `cmd/api` and run `go run main.go`:

## Access MySQL

1. Call `docker ps` to get list of running containers:

`docker ps`

2. Spawn a shell on the container:

`docker exec -it estimator-db-1 /bin/bash`

3. Connect to MySQL using client:

`mysql -u root -p`

4. Enter password.