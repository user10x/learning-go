version:=latest
APP_NAME:=hn
GCP_PROJECT=test-nchawla-proj
PORT = 4000
IMAGE := gcr.io/$(GCP_PROJECT)/$(version)

.PHONY: bin
bin:
	go build -o hnc main.go

.PHONY: run
run:
	go run -race main.go

.PHONY: clean
clean:
	rm hnc

.PHONY: build-remote
build-remote:
	docker build -t hn:latest .
	docker tag hn:latest gcr.io/$(GCP_PROJECT)/hn:latest
	docker push gcr.io/$(GCP_PROJECT)/hn:latest
	#kubectl run hn-app --image=gcr.io/$(GCP_PROJECT)/hn:latest --port=4000

.PHONY: deploy-remote
deploy-remote:
	kubectl apply -f deploy/app-deployment.yaml
	kubectl apply -f deploy/app-svc.yaml

.PHONY: all-remote
all-remote:  deploy-remote build-remote

.PHONY: clean-remote
clean-remote:
	kubectl delete -f deploy/
