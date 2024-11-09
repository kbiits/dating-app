# Requirements
- PostgreSQL
- Redis stack (or any redis with bloom filters module)
- Go >= ^1.22.2
- Make tools to run the app using Makefile

# How to run
1. Copy the `.env.example` to `.env` and adjust the value
2. Migrate the schema using `make migrate-up`
3. Run the http server with `make run-http`
4. You can also run the cron with `make run-cron`

# The interesting parts in this app
To make this app scalable, we have to make sure the API should have low latencies
when it comes to high traffic.
To make it possible, I'm using a bloom filter to achieve low latencies when it
comes to checking if a user already swipe other user.
But, still lot of things that I still can't finish because 2 days is really not
enough for me. I do write a little unit tests for several function that shouldn't
be changed, specially in the domain layer. 

# API Documentation
Please find the `Dealls Test.postman_collection.json`, and import it to your postman
to test the API.

# Directory Structure
This app follows the Uncle Bob Clean Architecture in terms of the dependencies rules.
This spec allows the app to be more easy to maintain because of the dependencies rules.
This app also following the port-adapter pattern as you can see we can implement
the entire application by just implementing the transport layer whether it's for
the incoming or outbound connection such as switching the database or switching the
tranport for receiving requests.

```
dating-app  
├── cmd  
│   └── cron               # Main application entry point for cron transport  
│   └── grpc               # Main application entry point for grpc transport (not used)  
│   └── http               # Main application entry point for http transport  
├── adapters
│   └── cron               # Main application entry point for cron transport  
│   └── grpc               # Main application entry point for grpc transport (not used)  
│   └── http               # Main application entry point for http transport  
|
├── domain                  # Domain layer  
│   │   ├── entity          # Entities and domain models  
│   │   ├── errors          # Business-specific error  
│   │   ├── repository      # Interfaces for persistence  
│   │   └── service         # Domain-specific services  
│   ├── application         # Application layer (use cases)  
│   ├── interface           # Interface adapters  
│   │   ├── http            # HTTP handlers/controllers  
│   │   ├── grpc            # gRPC handlers (if needed)  
│   │   └── repository      # Implementation of repository interfaces  
│   └── infrastructure      # Infrastructure layer  
│       ├── database        # Database connection and setup  
│       ├── repository      # Concrete repository implementations  
│       ├── config          # Configuration and environment variables  
│       └── logger          # Logger setup  
```

## Infrastructure directory
Infrastructure layer is a gateway to external world of the application, such as
connecting to rest api, postgresql, etc. All outbound connection should use this
as a gateway

## Adapters directory
Adapters is an implementation of each transport the app used to communicate.
This is one to one to interface adapter in terms of clean architecture by Uncle Bob.

## Usecase directory
Usecase layer is application-specific logic, which involving user interactions.
Ex: User Swiping Left/Right Other Profile

## Domain directory
Domain is used to put all of enterprise and business rules. This is the most inward layer
in terms of clean architecture by Uncle Bob.
All of the business rules should be put here.
I'm also using domain service terms which is comes from DDD (domain design driven)
which is used to put business logic that not fit into a domain entity

## Repositories directory
This is the concrete implementations of domain repository.
This should be specific and can be dependent and coupled to other package
(such as postgres, redis, etc)

## Migrations directory
This directory contains migration files used to build the schema of the database

## Config directory
This directory contains application config

## Utils directory
Contains all of helpers and utilities used by the application
