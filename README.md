/myapp
├── cmd
│   └── myapp               # Main application entry point (e.g., main.go)
├── pkg
│   ├── domain              # Domain layer
│   │   ├── entity          # Entities and domain models
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
