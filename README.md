# A full flegde example for Distributed Tracing

This Repository contains two micro-services in which tracing agent is integrated.
I am using Apache ZipKin, a Open Source tracer for demonstration.

## Prerequisites
    1. GO 1.13+
    2. Docker

## Services -
 1. Employee Service - This is a demo service performs CRUD operation for employee.
 2. Auth Service - This service creates and checks Auth for employee service APIs.

## Employee flows

<img  src="https://raw.githubusercontent.com/rosspatil/distributed-tracing/master/reg.png" title="Employee registration flow" width="500">
<br> <br>
<img  src="https://raw.githubusercontent.com/rosspatil/distributed-tracing/master/get.png" title="Employee registration flow" width="500">

In both of the flows employee service interacts with auth service using REST APIs.

## Demo

## Perform below actions to run this demo -<br>
    docker run -d -p 9411:9411 openzipkin/zipkin 
    cd auth
    go run main.go
    cd ..
    cd employee
    go run main.go

## Blow are the cURL to perform employee action

### 1. Employee Registration

#### Request -
    curl --location --request POST 'localhost:8080/employee' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name":"<your name>"
    }'

#### Response - 
    {
    "id": "<Employee Id>"
    }

    

### 2. Get Employee details
#### Request -
    curl --location --request GET 'localhost:8080/employee?id=<Employee Id>' \
    --header 'Authorization: <Employee Id>' \
    --header 'Content-Type: application/json'

#### Response - 
    {
        "employee": {
            "id": "<Employee Id>",
            "name": "your name"
        }
    }

### 3. Open ZipKin dashboard using below URL to see trace generated
    http://localhost:9411/zipkin