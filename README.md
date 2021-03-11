# Getting Started

### Requirement

* Linux kernel (version 5.8.0-44-generic) or macOS 11.0+ or OS X 10.8+
* docker (version 20.10.5)
* docker-compose  (version 1.25.0)
* make (GNU Make 4.2.1)
* golang (1.15 optional)

### Quickstart

```bash
> git clone https://github.com/versus/goUploadService.git
.......
> cd  goUploadService
> make docker-run
```
* waiting for start the service
```bash
> curl  http://localhost:8080/health
* Connect checker: OK
 ```
* try to connect to the service
```bash
> curl  http://localhost:8080
Use `curl -F 'file=@nameFile.ext' http://localhost:8080/upload` and see file into /tmp

```
* try to upload a file upload.jar to the service
```bash
> curl -F 'file=@steam_latest.deb' http://localhost:8080/upload
File steam_latest.deb has been successfully uploaded

```

* check the file in /tmp directory in the service images

```bash
> docker exec gouploadservice_go-upload_1 ls /tmp/steam_latest.deb
/tmp/steam_latest.deb

```

* get Bearer token to access to prometeus metrics 
```bash
> curl  http://localhost:8080/token
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY1NjUwMTcsImlzcyI6IklmY29uZmlnY28iLCJuYmYiOjE2MTUwMjkwMTd9.bfCk_76yTozw4LjQ9cH4Ig1GsaRZZPZUbxScY8iFryA

```


* check prometeus metrics for the service
```bash
> curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY1NjUwMTcsImlzcyI6IklmY29uZmlnY28iLCJuYmYiOjE2MTUwMjkwMTd9.bfCk_76yTozw4LjQ9cH4Ig1GsaRZZPZUbxScY8iFryA"  http://localhost:8080/metrics

```

### Use make commands

#### Start the service in local environment
```bash
cp example.env .env
..... edit .env file.....
make run
```

####  Build linux/darwin file of the service

```bash
make linux
make darwin
.....
ls ./dist
upload-0.0.2-darwin-amd64  upload-0.0.2-linux-amd64

```

####  Test build and work for Dockerfile and docker-compose.yml files 
```bash
make docker-test
```

<script id="asciicast-emNJPUoKSGJTsFRSe95RFAkDe" src="https://asciinema.org/a/emNJPUoKSGJTsFRSe95RFAkDe.js" async></script>

####  Build docker image versus/go-upload:latest of the service
```bash
make docker-build
```

####  Build docker image and start project in docker environment
```bash
make docker-run
```

### Issues
Please open *issues* here: [New Issue](https://github.com/versus/goUploadService/issues)

### Suggestions and improvements are welcome.

-Valentyn Nastenko(c 2021) https://github.com/versus