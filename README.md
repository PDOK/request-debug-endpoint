# request-debug-endpoint

Simple golang application that will log incoming request request

## What will it do

Have you ever found yourself in the situation that a 'customer' told you that the requests that they send to your API are not working, but you didn't change anything in your stack or infrastructure. But they are not in the position to check what kind of requests are sent your way, or even if there are any requests reaching your applications (9 out of 10 because of corporate infrastructures and weird proxies). Then this might help you and your customer in identifying what is wrong. It might be a missing header or a malformed XML body, letting them send their requests to an endpoint running a container with this simple Go application and you will know soon enough!

## Start

```docker
docker build -t pdok/request-debug-endpoint .
docker run --name request-debug-endpoint --rm -d -p 8001:80 pdok/request-debug-endpoint /request-debug-endpoint
```

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.