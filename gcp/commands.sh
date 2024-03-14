
## VARIABLES
CLUSTER_NAME="dwk-gke"
CLUSTER_ZONE="europe-north1-b"


# login to gcp account
gcloud auth login

# installing gcp kubernetes plugin
gcloud components install gke-gcloud-auth-plugin

# create a kubernetes cluster
gcloud container clusters create $CLUSTER_NAME --zone=$CLUSTER_ZONE

# get cluster credentials and apply to kubernetes
gcloud container clusters get-credentials $CLUSTER_NAME --zone=$CLUSTER_ZONE

# delete cluster
gcloud container clusters delete $CLUSTER_NAME --zone=$CLUSTER_ZONE

