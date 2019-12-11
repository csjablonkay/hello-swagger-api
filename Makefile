gen-server: validate
	swagger generate server \
		-t ./ \
		-f ./swagger.yaml \
		--exclude-main \
		-A github.com/csjablonkay/hello-swagger-api

gen-client: validate
	swagger generate client \
		-t ./ \
		-f ./swagger.yaml \
		-A github.com/csjablonkay/hello-swagger-api

gen: validate gen-server gen-client

validate:
	swagger validate ./swagger.yaml