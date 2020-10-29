# MicroGo
Experimenting with micro-services in GO. 

MicroGo is an experimental project, definitely not a production server. 

Architecture: 
  .Product: Product microservice expose REST APIs to manipulate products in a data set 
  .Currency: Currency microservice exposes live currencies rate (retrieved from an online repo) using a gRPC client
  .Image: image microservice is used to upload product images onto the server
