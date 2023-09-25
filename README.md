# SHORTENED URL
<!-- ABOUT THE PROJECT -->
## About The Project

This is a simple shortened url exercise. 
For the exercise, use this documentation to build an API backend service, for simplyfy url link and use Base62 Encode method.
We use base62 method to encode shortened Url because base62 encoder allows us to use the combination of characters and numbers which contains A-Z, a-z, 0–9 total( 26 + 26 + 10 = 62). So for 7 characters short URL, we can serve 62^7 ~= 3500 billion URLs which is quite enough in comparison to base10 (base10 only contains numbers 0-9 so you will get only 10M combinations). If we use base62 making the assumption that the service is generating 1000 tiny URLs/sec then it will take 110 years to exhaust this 3500 billion combination.

### Built With

This section should list any major frameworks that you built your project using. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.
* [Golang](https://golang.com)
* [PostgreSQL](https://www.postgresql.org/)

<!-- GETTING STARTED -->
## Getting Started
Before we get started, it's important to know that that this code use a custom command to execute it with makefile to make more simple command like :
1. make update
2. make tidy
3. make start

So, let start it.
1. After clone this repository, just run `make update`.
2. Setup your `.env` file such as database connection and redis connection based on default setting on you device.
3. To make sure that all dependency is run well, than run `make tidy`.
4. Finally, you can run your project with command: `make start`.
5. Go to postman and set url like `http://localhost:8080/`, for information that port to run this project depend on configuratin on `.env`

And for additional information, i'm alredy setup unit-testing, just run `make test-service`.

## Documentation
Here is more information on the process flow and API Contract documentation:

* [Design system](https://docs.google.com/document/d/1PJ8AfhajSWCDgJUYy2092ffP2mMWHwsBfnMpVwGOE9c/edit?usp=sharing)
* [API Documentation](https://drive.google.com/file/d/1XeyXzBdDUOTNr0s_Tv283F80RnbPi6Ss/view?usp=sharing)

## Afterword
Hopefully, it can be easily understood and useful. Thank you~
