# MicroGo
Experimenting with micro-services in GO. 

MicroGo is an experimental project, definitely not a production server.<br>

Architecture:<br> 
<ul>
  <li><b>Product:</b> Product microservice expose REST APIs to retrieve and edit products</li>
  <li><b>Currency:</b> Currency microservice exposes live currencies rate (retrieved from an online repo) using a gRPC client</li>
  <li><b>Image:</b> Image microservice is used to upload product images onto the server</li>
 </ul>
