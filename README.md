# Go Web Template
## Development
### Setup a local config file
Make a copy of `config.yaml` and name it `config.dev.yaml`.  The new file will need the following properties in addition to what's already present.
* **csrfTokenKey:** Take the value returned by the script below
    ```
    go run scripts/generate_32_byte_key/generate_32_byte_key.go
    ```
* **hmacKey:** Take the value returned by the script below
    ```
    go run scripts/generate_32_byte_key/generate_32_byte_key.go
    ```
* **baseURL:** See the section below on
* **mysqlUserName:** Take from `docker-compose.yaml`
* **mysqlPassword:** Take from `docker-compose.yaml`
* **mysqlHost:** `localhost`
* **mysqlPath:** The database path you'd like to create for your application
* **redisHost:** `localhost`
### Start MySQL & Redis containers
```
cd server
docker-compose up
```
### Setup [ngrok](https://ngrok.com/) tunnel
```
ngrok http 8080
```
From the output, copy the https forwarding URL to your clipboard, and paste it into `config.dev.yaml`
### Start server
```
go run --tags dev server.go
```
### Update web hook registration in Verity
```
go run scripts/update_webhook/update_webhook.go
```
### Start the UI
```
cd client
gatsby develop
```
Once started, UI should be available at [http://localhost:8000/](http://localhost:8080/)
