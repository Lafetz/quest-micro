# QUEST 

Go microservice, organized in a monorepo structure using gRPC, REST APIs, and kafka to facilitate interactions between services.The system is designed to demonstrate how these technologies can be integrated to build a scalable and efficient distributed architecture.
## Diagram

<img width="1028" alt="image" src="https://raw.githubusercontent.com/Lafetz/quest-demo/main/docs/diagram.png">

## Features

- gRPC for fast communication between services.
- RESTful APIs are available for easy connections with external systems.
- Kafka for messaging.
- Service Discovery using Consul.
- Resiliency Patterns like circuit breakers and retry mechanisms.
- Health Checks to monitor service status.
- Centralized Configuration Management to handle environment-specific settings.
- Load Balancing to evenly distribute traffic across service instances.
- Automated Deployment using Docker for containerization and orchestration.
- Scalability designed to handle increased load by adding more service instances.

## Monitoring

[x] structured Logging and Aggregation
[] Prometheus for performance monitoring and alerting
[] Tracing with Jaeger

## imgs

## Getting Started

### With Docker Compose

To run the application using Docker, you can follow these steps:

1. Navigate to the main directory.

2. Use the following command to build and start the containers:

   ```sh
   docker compose up --build
   ```

