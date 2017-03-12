# Introduction

This API module provide APIs with parameters:

1. `/place/nearbysearch`

	Just like this [Google place API](https://developers.google.com/places/web-service/search)

2. `/place/details`

	Just like this [Google place API](https://developers.google.com/places/web-service/details)

3. `/attractions`

	`place_id`: The `place_id` provided by **Google place API**

4. `/attractions/hotels`

	`place_ids`: A set of `place_id` divided by `,`

	`checkin`: A date in this format `2017-06-07`

	`checkout`: A date in this format `2017-08-07`

# Get and build the code

Get and build the code

```shell
go get github.com/chyeh/OkGo-API
cd $GOPATH/src/github.com/chyeh/OkGo-API
go build
```

Create the configuration file `cfg.json`. Take a look at `cfg.example.json`:

```json
{
	"restful" : {
		"listen" : {
			"host": "0.0.0.0",
			"port": 8080
		}
	},
	"logLevel" : {
		"ROOT" : "DEBUG"
	},
	"api_auth" : {
		"booking_dot_com" :{
		 "username": "username",
		 "password": "password"
	 },
		"google_place": "get-your-own-key"
	}
}
```

# Run

If you are in the working directory:

```shell
./OkGo-API
```

If not, use `-c` to specify the path of the configuration file
```shell
<PATH>/<TO>/OkGo-API -c <PATH>/<TO>/OkGo-API
```

# Test

For `/attractions`:

```shell
curl http://<Your IP>:<Your Port>/attractions\?place_id\=ChIJyWEHuEmuEmsRm9hTkapTCrk
```
For `/attractions/hotels`:

```shell
curl http://<Your IP>:<Your Port>/attractions/hotels\?checkin=2017-06-08\&checkout=2017-07-07\&place_ids=ChIJyWEHuEmuEmsRm9hTkapTCrk,ChIJLfySpTOuEmsRsc_JfJtljdc
```
