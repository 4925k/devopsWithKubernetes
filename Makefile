# start k3d cluster
# hash -> listening @localhost:4444/hash/status
# pingpong -> listening @locahost:4445/pingpong
# image -> localhost

# kubectl apply -f ./image/manifests
# kubectl apply -f ./pingpong/manifests
create:
	k3d cluster create -p 8081:80@loadbalancer --agents 2
	kubectl create namespace ns-sideproject
	kubectl create namespace ns-mainproject
	kubectl apply -f ./manifests/
	kubectl apply -f ./manifests/hash
	kubectl apply -f ./manifests/pingpong
	kubectl apply -f ./manifests/postgres
	kubectl apply -f ./manifests/todolist
	


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