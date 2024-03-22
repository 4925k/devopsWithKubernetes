
export CLUSTER_NApostgreME:="dwk-gke"
export CLUSTER_ZONE:="asia-south2-a"

create:
	gcloud container clusters create ${CLUSTER_NAME} --zone=${CLUSTER_ZONE}

delete:
	gcloud container clusters delete ${CLUSTER_NAME} --zone=${CLUSTER_ZONE}

postgres:
	kubectl apply -f manifests/pvc.yaml
	kubectl apply -f manifests/postgres

postgresrm:
	kubectl delete -f manifests/pvc.yaml
	kubectl delete -f manifests/postgres


pingpong:
	kubectl apply -f manifests/pingpong

pingpongrm:
	kubectl delete -f manifests/pingpong

ppwithdb:
	kubectl apply -f manifests/pvc.yaml
	kubectl apply -f manifests/postgres
	kubectl apply -f manifests/pingpong

ppwithdbrm:
	kubectl delete -f manifests/pvc.yaml
	kubectl delete  -f manifests/pingpong
	kubectl delete -f manifests/postgres



stats:
	kubectl get pv
	kubectl get pvc
	kubectl get deployments
	kubectl get pods 
	