# grpc-saga-orchestrator

## Overview

This project demonstrates a **distributed microservices system** built using **gRPC in Go**, implementing the **Saga pattern (Orchestrator-based)** for handling distributed transactions.

The system simulates a real-world order processing workflow across multiple services and ensures consistency using **compensation logic** instead of traditional database transactions.

---

## Architecture

The system consists of three independent services:

* **Order Service (Orchestrator)**
  Controls the workflow and manages the Saga execution

* **Payment Service**
  Handles payment processing and refunds

* **Shipping Service**
  Simulates shipment processing and failure scenarios

---

## Saga Workflow

### Success Flow

1. Order is created
2. Payment is processed
3. Shipping is initiated
4. Order is marked as successful

---

### Failure Flow (Compensation)

1. Order is created
2. Payment is processed
3. Shipping fails
4. Payment is refunded
5. Order is marked as failed

---

## Key Concepts

* gRPC-based inter-service communication
* Protocol Buffers for service contracts
* Orchestrator-based Saga pattern
* Compensation transactions
* Distributed system design

---

## Running the Application (Without Docker)

### Start Services

Run each service in separate terminals:

```bash
go run payment/main.go
go run shipping/main.go
go run order/main.go
```

---

### Run Client

```bash
go run client/main.go success
```

---

## Running the Application (Docker)

### Build and Start All Services

```bash
docker-compose up --build
```

This will:

* Build all services
* Start containers
* Create internal Docker network

---

### Run Client (outside Docker)

```bash
go run client/main.go success
```

---

## Controlled Simulation (IMPORTANT)

The system supports controlled execution modes:

```bash
go run client/main.go success   # Full success flow
go run client/main.go fail      # Force Saga failure (compensation)
go run client/main.go timeout   # Simulate timeout scenario
```

---

## Expected Output

### Success Case

```text
Order Created
Payment processed
Shipping SUCCESS
Order Completed
Final Result: SUCCESS
```

---

### Failure Case (Saga Compensation)

```text
Order Created
Payment processed
Shipping FAILED
Payment refunded
Order FAILED
Final Result: FAILED
```

---

### Timeout Case

```text
Request FAILED: context deadline exceeded
```

---

## Monitoring and Logs

You can observe system behavior through logs:

### Docker Logs

```bash
docker-compose logs -f
```

---

### Check Running Containers

```bash
docker ps
```

---

### Observe Individual Service Logs

Each service prints:

* Order lifecycle
* Payment status
* Shipping result
* Compensation execution

---

## Code Generation

Protobuf files are pre-generated and included.

To regenerate:

```bash
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

---

## Project Highlights

* Distributed transaction handling using Saga pattern
* gRPC-based microservice communication
* Failure handling with compensation logic
* Controlled simulation for testing scenarios
* Fully containerized using Docker

---

## Future Improvements

* Add database (persistent storage)
* Add structured logging (Zap / Logrus)
* Add tracing (OpenTelemetry)
* Add retry and circuit breaker
* Deploy using Kubernetes

---

## Learning Outcomes

This project demonstrates:

* Designing distributed systems
* Handling failures in microservices
* Implementing Saga pattern
* Building production-style backend systems

---

