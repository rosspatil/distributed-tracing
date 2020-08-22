# A full flegde example for Distributed Tracing

This Repository contains two micro-services in which tracing agent is integrated.
I am using Apache ZipKin, a Open Source tracer for demonstration.

## Services -
 1. Employee Service - This is a demo service performs CRUD operation for employee.
 2. Auth Service - This service creates and checks Auth for employee service APIs.

## Employee flows

<img  src="https://raw.githubusercontent.com/rosspatil/distributed-tracing/master/reg.png" title="Employee registration flow" width="50">
<br> <br>
<img  src="https://raw.githubusercontent.com/rosspatil/distributed-tracing/master/get.png" title="Employee registration flow" width="50">

In both flows employee service interacts with auth service using inter-service communication.