# Quasar Fire Operation - MELI

![enter image description here](https://i.ibb.co/fVWmnTN/images.jpg)

Challenge project "Fuego de Quasar"

# Quick Start 

Requires  **GO 1.5**

Run locally with: 
> $ go run main.go

Run all tests from the root folder with:
> $ go test ./...

## Endpoints

**Topsecret**
Get the position and complete message by posting three satellites with them distance and incomplete message. 
Some positions are not possible to calculate.
`POST -> /topsecret`
Example request body:

    {
    	"satellites": [
    		{
    			"name": "kenobi",
    			"distance": 989.945,
    			"message": ["este","","","mensaje",""]
    		},
    		{
    			"name": "skywalker",
    			"distance": 608.276,
    			"message": ["","es","","","secreto"]
    		},
    		{
    			"name": "sato",
    			"distance": 500,
    			"message": ["este","","un","",""]
    		}
    	]
    }
Example succesful 200 response:

    {
      "position": {
        "x": 199.99,
        "y": 500
      },
      "message": "este es un mensaje secreto"
    }

Example failed 400 response

    {
      "error": "could not calculate position or message with provided info"
    }
<br>
<br>

**Topsecret Split**
Posts a satellite's distance and an incomplete message that will be saved in system memory. It will receive the satellite's name as a path parameter.
`POST -> /topsecret_split/{satellite_name}`

Example request body

    {
	    "distance": 989.945,
	    "message": ["este", "", "", "mensaje", ""]
    }
Example 200 response

    "saved"

`GET -> /topsecret_split/`

Example 200 response:

    {
      "position": {
        "x": 199.99,
        "y": 500
      },
      "message": "este es un mensaje secreto"
    }





## Miscellaneous

 - The app runs with a cache-like memory to store which flushes every 10 minutes and it's values get stored for 8 minutes. It's data will be lost if the app's process is interrupted.
 - It would be nice to add more tests
 - You can test it online [https://quasar-fire-meli.herokuapp.com](https://quasar-fire-meli.herokuapp.com)