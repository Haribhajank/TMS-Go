# Cloud-Based Task Management System

A microservices-based task management system built with **Go**, **Docker**, and **Kubernetes**, designed to provide scalable task management, user authentication, and notification services. The project integrates **Prometheus** and **Grafana** for real-time monitoring and observability.

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Technologies Used](#technologies-used)
- [Setup Instructions](#setup-instructions)
  - [1. Prerequisites](#1-prerequisites)
  - [2. Local Deployment](#2-local-deployment)
  - [3. Kubernetes Deployment](#3-kubernetes-deployment)
  - [4. Monitoring Setup](#4-monitoring-setup)
- [API Endpoints](#api-endpoints)
- [Future Improvements](#future-improvements)

---

## Features

- **User Management**: Register and authenticate users.
- **Task Management**: Create, update, delete, and view tasks.
- **Notifications**: Notify users of task updates or reminders.
- **Monitoring**: Real-time monitoring and alerts via Prometheus and Grafana.

## Architecture

The project uses a microservices-based architecture with the following components:

- **User Service**: Manages user authentication and profiles.
- **Task Service**: Handles task CRUD operations.
- **Notification Service**: Sends notifications to users.
- **API Gateway**: Routes client requests to the appropriate services.
  
Each service is containerized and orchestrated by **Kubernetes** for scalability and resilience.

## Technologies Used

- **Go**: Backend logic and microservices
- **Docker**: Containerization of each microservice
- **Kubernetes**: Orchestration and scaling
- **Nginx**: API Gateway for routing and load balancing
- **Prometheus & Grafana**: Monitoring and observability

---

## Setup Instructions

### 1. Prerequisites

- **Docker** and **Kubernetes** (with Minikube for local testing)
- **Helm** package manager
- **Go** installed locally (for development)

### 2. Local Deployment

To run each service individually using Docker:

1. **Build and Run Docker Containers**:
   
   ```bash
   docker build -t user-service ./user-service
   docker build -t task-service ./task-service
   docker build -t notification-service ./notification-service
2. **Run Containers**:

   ```bash
   docker run -d -p 8080:8080 user-service
   docker run -d -p 8081:8081 task-service
   docker run -d -p 8082:8082 notification-service

Each service will now be available on its respective port.

### 3. Kubernetes Deployment

To deploy the services in a Kubernetes cluster:

1. **Start Minikube**:
    
    ```bash
    minikube start
    ```

2. **Apply Kubernetes Manifests**:
    
    ```bash
    kubectl apply -f kubernetes/user-deployment.yaml
    kubectl apply -f kubernetes/task-deployment.yaml
    kubectl apply -f kubernetes/notification-deployment.yaml
    ```

3. **Access Services**:
    
    Use Minikube’s tunneling feature to access services locally:
    
    ```bash
    minikube service user-service
    ```

### 4. Monitoring Setup

To set up monitoring with Prometheus and Grafana:

1. **Install Prometheus**:
    
    ```bash
    helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
    helm install prometheus prometheus-community/prometheus -n monitoring
    ```

2. **Install Grafana**:
    
    ```bash
    helm repo add grafana https://grafana.github.io/helm-charts
    helm install grafana grafana/grafana -n monitoring
    ```

3. **Access Grafana**:
    
    ```bash
    kubectl port-forward -n monitoring deploy/grafana 3000
    ```

4. **Configure Prometheus in Grafana**:
    
    - Go to `http://localhost:3000`, add Prometheus as a data source, and create dashboards to monitor metrics.

---

## API Endpoints

### User Service

- `POST /register`: Register a new user
- `POST /login`: User login
- `GET /profile`: Get user profile

### Task Service

- `POST /tasks`: Create a new task
- `GET /tasks`: List all tasks
- `PUT /tasks/{id}`: Update a task
- `DELETE /tasks/{id}`: Delete a task

### Notification Service

- `POST /notify`: Send a notification

---

## Future Improvements

- **Cloud Deployment**: Deploy services on OCI, AWS, GCP, or Azure for scalability.
- **Authentication**: Enhance security with JWT-based authentication.
- **Advanced Monitoring**: Add alerting rules in Prometheus for improved observability.
