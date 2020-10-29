# MicroGo
Experimenting with micro-services in GO. 

MicroGo is an experimental project, definitely not a production server.<br>

Architecture:<br> 
<ul>
  <li>Product: Product microservice expose REST APIs to manipulate products in a data set </li>
  <li>Currency: Currency microservice exposes live currencies rate (retrieved from an online repo) using a gRPC client</li>
  <li>Image: image microservice is used to upload product images onto the server</li>
 </ul>
