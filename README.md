# Rakuten Appication: Api server in golang

## Dependencies

    Database:           MariaDB/MySQL
    ORM:                gorm
    Mngmt Routers:      Gorilla/mux
    Mngmt environments: GoDotEnv
    Source of truth:    https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml

## Note

I used **Debug()** function provided by gorm for the **debbugin\***

## Instructions for run the code.

### 1. Clone this repository:

```bash
git clone https://github.com/Rommel96/euro-exchange-rates-api-server-golang
```

### 2. Create a **database** and set values on .env file

    PORT = SERVER_PORT
    DB_TYPE= DB_TYPE
    DB_USER= DB_USER
    DB_PASS= DB_PASS
    DB_HOST= DB_HOST
    DB_PORT= DB_PORT
    DB_NAME= DB_NAME

### 3. Run server

```bash
go run main.go
```
