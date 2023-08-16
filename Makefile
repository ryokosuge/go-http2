.PHONY: nginx/certs

nginx/certs:
	rm -rf nginx/certs
	mkdir -p nginx/certs
	openssl genrsa -out nginx/certs/server.key 2048
	openssl req -new -key nginx/certs/server.key -out nginx/certs/server.csr
	openssl x509 -req -days 3650 -signkey nginx/certs/server.key < nginx/certs/server.csr > nginx/certs/server.crt