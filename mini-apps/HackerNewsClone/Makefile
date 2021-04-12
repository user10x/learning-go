version:=latest
APP_NAME:=hn
GCP_PROJECT=test-nchawla-proj
PORT = 4000
IMAGE := gcr.io/$(GCP_PROJECT)/$(version)

.PHONY: build
all:
	docker build -t hn:latest .
	docker tag hn:latest gcr.io/$(GCP_PROJECT)/hn:latest
	docker push gcr.io/$(GCP_PROJECT)/hn:latest
	#kubectl run hn-app --image=gcr.io/$(GCP_PROJECT)/hn:latest --port=4000

.PHONY: deploy
deploy:
	kubectl apply -f deploy/app-deployment.yaml
	kubectl apply -f deploy/app-svc.yaml



.PHONY: all
all:  deploy build

#docker build -t hn:latest .
#docker tag go-app:latest  gcr.io/$GCP_PROJECT/go-app:latest
#"docker push gcr.io/$GCP_PROJECT/go-app"