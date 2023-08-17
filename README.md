# Notification Service

The Notification Service is a lightweight and efficient system designed to accept messages via API requests and then asynchronously dispatch them to various channels such as Slack and email. This README provides a comprehensive guide on how to set up and use this service.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [Getting available channels](#getting-available-channels)
  - [Sending Notifications](#sending-notifications)
- [Configuration](#configuration)

## Getting Started

Follow these steps to set up and run the Notification Service on your machine.

### Prerequisites
- Docker
- Docker Compose

### Installation

1. Clone this repository:

   ```sh
   git clone https://github.com/phgermanov/notification-service
   cd notification-service
   ```

2. Build and start the service using Docker Compose:
    ```sh
    docker-compose up --build -d
    ```
3. The service is now running and ready to accept notifications.

# Usage

## Getting available channels
To get all available channels, make a GET request to `/channels`. Here's an example using curl:
```sh
curl -X GET http://localhost:8087/channels
```

## Sending Notifications
To send a notification, make a POST request to the `/notifications` endpoint with a JSON payload containing the necessary information. Here's an example using curl:
```sh
curl -X POST \
    -H "Content-Type: application/json" \
    -d '{
        "message": "Hello, world!",
        "channels": ["Slack", "Email"]
    }' \
    http://localhost:8087/notifications
```

## Configuration
The Notification Service can be configured using settings.yml and environment variables. Example settings.yml file:
```yml
PORT: "8087"
SLACK_WEBHOOK_URL: "<SLACK_TOKEN>"
RETRY_DURATION: "5s"
```

### Getting Slack Webhook URL
Slack sender is using Incomming Webhooks for simplicity. 
You can see how to set it up [here](https://api.slack.com/messaging/webhooks).