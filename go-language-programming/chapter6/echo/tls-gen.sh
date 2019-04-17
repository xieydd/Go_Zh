#! /bin/nash

rm -rf client* server* ca*
# self 
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -days 10000 -out ca.crt -subj "/IP=127.0.0.1"
#server
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/IP=127.0.0.1"
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365
#client
openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr -subj "/IP=127.0.0.1"
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365
rm -rf ./ca.key ./ca.srl  && rm -rf ./server.csr && rm -rf ./client.csr
