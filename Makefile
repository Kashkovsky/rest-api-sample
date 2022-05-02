release:
	docker build --tag rest-api-sample .
	docker tag rest-api-sample bourneagain/rest-api-sample:latest
	docker push bourneagain/rest-api-sample:latest

