# WebSub Hub

A [WebSub](https://www.w3.org/TR/websub/) hub implemented in with Go. Includes a prebuilt subscriber service and MySQL database for storing subscriber data. Uses a Docker compose file for running all applications.

The hub mainly fulfills the following functionalities:

1. The server verifies the subscriber and its intent of subscribing to a specific topic. The subscriber data (its callback URL, topic and secret) is persisted in a MySQL database.
2. Messages are signed, meaning an [HMAC](https://en.wikipedia.org/wiki/HMAC) signature header is added to each POST requests sent to the subscribers.

## Run locally

Run the three services:

`docker compose up -d`

## Walkthrough

The hub takes ~5-10 seconds to start since it waits on the MySQL server to start. On a successfull start, the hub logs something like this:

```
2024-11-13 18:22:13 subscribers table created successfully
2024-11-13 18:22:13 2024/11/13 17:22:13 Successfully connected to MySQL database
```

Until it is sucessfull, the subscriber sends a new subscribe request to the hub every 10 seconds. Hopefully it will work as soon as the hub is up and running:

```
// logs from the hub
2024-11-13 18:22:16 2024/11/13 17:22:16 Verifying intent of subscriber...
2024-11-13 18:22:16 2024/11/13 17:22:16 Intent verified
2024-11-13 18:22:16 2024/11/13 17:22:16 Subscriber added successfully
```

The logs verifies that the hub has a new subscriber to some topic. Looking at the subscribers logs, it seems to be happy as well:

```
// logs from the subscriber
2024-11-13 18:22:16 time="2024-11-13T17:22:16Z" level=info msg="Sending subscribe request"
2024-11-13 18:22:16 time="2024-11-13T17:22:16Z" level=info msg="Got subscribe to /a/topic and challenge IcUVOxwkVR"
2024-11-13 18:22:16 time="2024-11-13T17:22:16Z" level=info msg="Verifying intent for /MDQQkuvQfU"
```

Lastly, try to generate a publication by sending a GET request to the hubs `/generate` endpoint, including a `topic` parameter. It will try to generate a random user and post it to the subscribers of the topic.

```
// Generate a publication. Data will be posted to all subscribers of the topic "/a/topic"
curl -X GET http://localhost:8080/generate?topic=/a/topic
```

```
// A random user is returned
{"firstname":"Katie","lastname":"Smith","age":53}
```

The hub logs verifies that it generated a new user:

```
// logs from the hub
2024-11-13 18:27:35 2024/11/13 17:27:35 Generate request for topic: /a/topic
2024-11-13 18:27:35 2024/11/13 17:27:35 Successfully sent data to subscriber http://web-sub-client:8080/MDQQkuvQfU, status: 200 OK
```

And the subscriber verifies that it recieved a post:

```
// logs from the subscriber. Happy once again
2024-11-13 18:27:35 time="2024-11-13T17:27:35Z" level=info msg="Got a post, and all is good"
```

Lastly, we can verify that the subscriber actually recieved the generated data. Visit http://localhost:8081/log and you should see something like

```
{
    "Callback": "/MDQQkuvQfU",
    "Timestamp": "2024-11-13T17:27:35.387392591Z",
    "Payload": {
        "age": 53,
        "firstname": "Katie",
        "lastname": "Smith"
    }
}
```

## Dev commands

````

// Rebuild Hub image and run with docker compose
docker compose up -d --build

```

```
````
