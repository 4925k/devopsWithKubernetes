# start k3d cluster
start:
	k3d cluster create -p 8081:80@loadbalancer --agents 2
	kubectl apply -f ./hash-gen/manifests
	kubectl apply -f ./pingpong/manifests

stats:
	kubectl get deployments
	kubectl get pods 

down:
	k3d cluster delete