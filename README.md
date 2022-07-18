This is a simple bank backend web services developed from scratch right to deployment.

Create and manage bank accounts.
Record all balance changes to each of the accounts.
Perform a money transfer between 2 accounts.
The programming language used to develop the service is Golang, but the course is not just about coding in Go. 

The database was created using sqlc, generate codes to talk to the DB in a consistent and reliable way using transactions, understand the DB isolation levels, and how to use it correctly in production. Besides the database, you will also learn how to use docker for local development, how to use Git to manage your codes, and how to use Github Action to run unit tests automatically.

RESTful HTTP APIs using Gin - one of the most popular Golang frameworks for building web services. This includes everything from loading app configs, mocking DB for more robust unit tests, handling errors, authenticating users, and securing the APIs with JWT and PASETO access tokens.

Docker is used to deploy it to a production Kubernetes cluster on AWS. The lectures are very detailed with a step-by-step guide, from how to build a minimal docker image, set up a free-tier AWS account, create a production database, store and retrieve production secrets, create a Kubernetes cluster with EKS, use Github Action to automatically build and deploy the image to the EKS cluster, buy a domain name and route the traffics to the service, secure the connection with HTTPs and auto-renew TLS certificate.
 
Future work-in-progress, where we will do advanced backend topics such as managing user sessions, building gRPC APIs, using gRPC gateway to serve both gRPC and HTTP with 1 single implementation of the handler, and embedding Swagger documentation as part of the backend service, etc. We will keep updating the project.
