Alright — you want the **most basic Golang “Hello World” microservice** but **deployed on GKE** instead of Cloud Run.
We’ll keep it as barebones as possible while touching all GKE essentials.

---

## **Step 1 — Minimal Go Microservice**

**main.go**

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go on GKE! 🚀")
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", helloHandler)
    log.Printf("Server listening on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
```

**go.mod**

```sh
go mod init github.com/yourusername/hello-gke
go mod tidy
```

---

## **Step 2 — Dockerfile**

```dockerfile
# Stage 1: Build
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o server .

# Stage 2: Run
FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/server /server
CMD ["/server"]
```

---

## **Step 3 — Build & Push Image**

```sh
# Enable Container Registry or Artifact Registry
gcloud services enable containerregistry.googleapis.com

gcloud services enable artifactregistry.googleapis.com containerregistry.googleapis.com
gcloud auth configure-docker
gcloud auth configure-docker us-docker.pkg.dev

# Build and push to GCP Container Registry
docker build -t gcr.io/<PROJECT_ID>/hello-gke:v1 .
docker push gcr.io/<PROJECT_ID>/hello-gke:v1
```

---

## **Step 4 — Kubernetes Manifests**

**deployment.yaml**

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-gke
spec:
  replicas: 1
  selector:
    matchLabels:
      app: hello-gke
  template:
    metadata:
      labels:
        app: hello-gke
    spec:
      containers:
      - name: hello-gke
        image: gcr.io/PROJECT_ID/hello-gke:v1
        ports:
        - containerPort: 8080
```

**service.yaml**

```yaml
apiVersion: v1
kind: Service
metadata:
  name: hello-gke-service
spec:
  type: LoadBalancer
  selector:
    app: hello-gke
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
```

---

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

kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

---

## **Step 7 — Get External IP**

```sh
kubectl get service hello-gke-service
```

Open the `EXTERNAL-IP` in your browser to see:

```
Hello from Go on GKE! 🚀
```

---

If you want, I can make the **next level of this microservice** —
still minimal but with **health checks**, **JSON output**, and a **configurable replica count** so it’s a proper production-ready GKE microservice skeleton.

Do you want me to prepare that next?
