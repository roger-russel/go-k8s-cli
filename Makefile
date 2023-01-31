version=v0.0.2

.PHONY: helm
helm:
	@helm create go-k8s-cli

.PHONY: build
build:
	@docker build -t go-k8s-cli:v0.0.1 -f k8s/Dockerfile .

.PHONY: build/minikube
build/minikube:
	bash -c "$(cat <<EOF
		minikube docker-env;
		eval $(minikube -p minikube docker-env);
		docker build -t go-k8s-cli:v0.0.2 -f k8s/Dockerfile .
	EOF
	)"

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

.PHONY: logs
logs:
	kubectl logs deploy/go-k8s-cli -f

.PHONY: minikube
minikube:
	minikube start --memory=4096 --cpus=4