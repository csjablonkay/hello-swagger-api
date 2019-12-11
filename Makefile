gen-server: validate
	swagger generate server \
		-t ./ \
		-f ./swagger.yaml \
		--exclude-main \
		-A hello-swagger-api

gen-client: validate
	swagger generate client \
		-t ./ \
		-f ./swagger.yaml \
		-A hello-swagger-api

gen: validate gen-server gen-client

validate:
	swagger validate ./swagger.yaml