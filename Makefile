run:
	GOOGLE_CLOUD_PROJECT=$$(gcloud config get-value project) API_KEY=dev go run *.go

open:
	open "http://localhost:8080/swagger-ui/?url=http://localhost:8080/api-docs.json"

dbuild:
	docker build -t landskape .

# for local run we need a service account and some env set
drun:
	docker run -it -p8080:8080 \
	-eAPI_KEY=dev \
	-eGOOGLE_CLOUD_PROJECT=$$(gcloud config get-value project) \
	-v $$(pwd):/config \
	-eGOOGLE_APPLICATION_CREDENTIALS=/config/landskape.json \
	landskape