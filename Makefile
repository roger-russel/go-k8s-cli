version=v0.0.4

.PHONY: helm
helm:
	@helm create go-k8s-cli

.PHONY: build
build:
	@docker build -t go-k8s-cli:$(version) -f k8s/Dockerfile .

.PHONY: build/minikube
build/minikube:
	@./scripts/registry-image.sh $(version)

.PHONY: helm/install
helm/install:
	@helm install go-k8s-cli ./k8s/helm -f ./k8s/helm/values.yaml -n default

.PHONY: helm/upgrade
helm/upgrade:
	@helm upgrade go-k8s-cli ./k8s/helm -f ./k8s/helm/values.yaml -n default

.PHONY: helm/list
helm/list:
	@helm list

.PHONY: port-forward
port-forward:
	@kubectl port-forward svc/go-k8s-cli 8080:80

.PHONY: pod/logs
pod/logs:
	kubectl logs deploy/go-k8s-cli -f

.PHONY: pod/delete
pod/delete:
	kubectl delete pod $(shell kubectl get pod | grep go | awk -F" " '{print $$1}')

.PHONY: minikube
minikube:
	minikube start --memory=4096 --cpus=4

.PHONY: test
test:
	go test ./... -coverpkg=./... -race 
