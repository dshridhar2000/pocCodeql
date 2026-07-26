## **Step 5 — Create GKE Cluster**

```sh
gcloud services enable container.googleapis.com

gcloud container clusters create hello-gke-cluster \
  --zone asia-south1-a \
  --num-nodes 1
```

---

## **Step 6 — Deploy to GKE**

```sh
sudo apt-get install google-cloud-sdk-gke-gcloud-auth-plugin

gcloud container clusters get-credentials hello-gke-cluster --zone asia-south1-a