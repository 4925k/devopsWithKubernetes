# start k3d cluster
# hash -> listening @localhost:4444/hash/status
# pingpong -> listening @locahost:4445/pingpong
# image -> localhost

# kubectl apply -f ./image/manifests
# kubectl apply -f ./pingpong/manifests
create: 
	k3d cluster create -p 8081:80@loadbalancer --agents 2

setup:
	k3d cluster create -p 8081:80@loadbalancer --agents 2
	kubectl apply -f ./manifests/
	kubectl apply -f ./manifests/hash
	kubectl apply -f ./manifests/pingpong
	kubectl apply -f ./manifests/postgres
	kubectl apply -f ./manifests/todolist
	kubectl apply -f ./manifests/ubuntu
	
helm: 
	brew install helm
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo add stable https://charts.helm.sh/stable


prometheus:
	kubectl create namespace prometheus
	helm install prometheus-community/kube-prometheus-stack --generate-name --namespace prometheus


loki:
	helm repo add grafana https://grafana.github.io/helm-charts
	helm repo update
	kubectl create namespace loki-stack
	helm upgrade --install loki --namespace=loki-stack grafana/loki-stack
	

apply: 
	kubectl apply -f ./manifests/
	kubectl apply -f ./manifests/hash
	kubectl apply -f ./manifests/pingpong
	kubectl apply -f ./manifests/postgres
	kubectl apply -f ./manifests/todolist

remove: 
	kubectl delete -f ./manifests

reapply: remove
	kubectl apply -f ./manifests

stats:
	kubectl get deployments
	kubectl get pods 

down:
	k3d cluster delete



### DOCKER COMMANDS

dockerhashfind:
	docker build -t 4925k/hashfind:v1 ./app-hash/hashfind
	docker push 4925k/hashfind:v1

dockertodolistfe:
	docker build -t 4925k/todolistfe:v1 ./app-todolist/fe
	docker push 4925k/todolistfe:v1

dockertodolistbe:
	docker build -t 4925k/todolistbe:v1 ./app-todolist/be
	docker push 4925k/todolistbe:v1

dockerpingpong:
	docker build -t 4925k/pingpong:v1 ./app-pingpong
	docker push 4925k/pingpong:v1

grafana:
	brew install helm
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo add stable https://charts.helm.sh/stable
	kubectl create namespace logging
	helm install prometheus-community/kube-prometheus-stack --generate-name --namespace logging
	helm repo add grafana https://grafana.github.io/helm-charts
	helm repo update
	helm upgrade --install loki --namespace=logging grafana/loki-stack
	kubectl get pods -n logging
	echo "kubectl -n logging port-forward <grafana pod> 3000"