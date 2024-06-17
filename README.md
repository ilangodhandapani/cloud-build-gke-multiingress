# cloud-build-gke-multiingress 

gcloud container clusters get-credentials my-auto-cluster-east \
    --region=us-east1 \
    --project=flawless-lacing-392113

gcloud container clusters get-credentials my-auto-cluster-central \
    --region=us-central1 \
    --project=flawless-lacing-392113

gcloud container fleet memberships list --project=flawless-lacing-392113

gcloud container fleet ingress enable \
    --config-membership=my-auto-cluster-east \
    --location=us-east1 \
    --project=flawless-lacing-392113

gcloud container fleet ingress describe \
    --project=flawless-lacing-392113
    
kubectl config use-context gke_flawless-lacing-392113_us-east1_my-auto-cluster-east

apiVersion: networking.gke.io/v1
kind: MultiClusterService
metadata:
  name: mcs
spec:
  template:
    spec:
      selector:
        app: myapp-blue
      ports:
      - name: web
        protocol: TCP
        port: 80
        targetPort: 8080
