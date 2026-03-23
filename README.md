# grpc-saga-orchestrator

## Overview

This project demonstrates a distributed microservices system built using **gRPC** in Go, implementing the **Saga pattern (Orchestrator-based)** for handling distributed transactions.

The system simulates a real-world order processing workflow involving multiple services and ensures consistency using compensation logic instead of traditional transactions.

---

## Architecture

The system consists of three independent services:

* **Order Service (Orchestrator)**
  Controls the workflow and manages the Saga execution

* **Payment Service**
  Handles payment processing and refunds

* **Shipping Service**
  Simulates order shipment and failure scenarios

---

## Workflow

### Success Flow

1. Order is created
2. Payment is processed
3. Shipping is initiated
4. Order is marked as successful

---

### Failure Flow

1. Order is created
2. Payment is processed
3. Shipping fails
4. Payment is refunded (compensation)
5. Order is marked as failed

---

## Key Concepts

* gRPC-based inter-service communication
* Protocol Buffers for service contracts
* Orchestrator-based Saga pattern
* Compensation transactions for failure handling
* Distributed system design principles

---

## Running the Application

### Start Services

Run each service in a separate terminal:

```bash
go run payment/main.go
go run shipping/main.go
go run order/main.go
```

---

### Run Client

```bash
go run client/main.go
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

### Failure Case

```text
Order Created
Payment processed
Shipping FAILED
Payment refunded
Final Result: FAILED
```

---

## Code Generation

Protobuf files are already generated and included in the repository.

To regenerate:

```bash
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

---

## Learning Outcomes

This project demonstrates:

* Designing distributed systems using microservices
* Handling failures using Saga pattern
* Implementing inter-service communication with gRPC
* Understanding real-world backend architecture patterns

---

## Future Improvements

* Add persistent storage (database)
* Introduce service discovery
* Add logging and tracing
* Containerize using Docker
* Deploy using Kubernetes

---

