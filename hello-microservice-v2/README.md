Nice 👍 let’s extend your **Hello Microservice** with **Level 1 improvements**:

* Add a **second endpoint** `/healthz` (for Kubernetes readiness/liveness probes)
* Return **JSON responses** instead of plain text
* Add **query parameter support** (`/hello?name=Shridhar`)

---

## **1. Updated Go Code**

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Response struct for JSON responses
type Response struct {
	Message string `json:"message"`
	Status  string `json:"status,omitempty"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	resp := Response{Message: fmt.Sprintf("Hello, %s! 🚀", name)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "Service is healthy", Status: "ok"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/healthz", healthHandler)

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
```

---

## **2. Test Locally**

```sh
go run main.go
```

* Visit:

  * `http://localhost:8080/hello` → `{"message":"Hello, World! 🚀"}`
  * `http://localhost:8080/hello?name=Shridhar` → `{"message":"Hello, Shridhar! 🚀"}`
  * `http://localhost:8080/healthz` → `{"message":"Service is healthy","status":"ok"}`

---

## **3. Update Deployment (for probes)**

Add liveness & readiness probes using `/healthz`:

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
        image: gcr.io/PROJECT_ID/hello-gke:v2   # build and push new version
        ports:
        - containerPort: 8080
        livenessProbe:
          httpGet:
            path: /healthz
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /healthz
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 10
```

---

## **4. Build & Push New Image**

```sh
docker build -t gcr.io/<PROJECT_ID>/hello-gke:v2 .
docker push gcr.io/<PROJECT_ID>/hello-gke:v2
kubectl apply -f deployment.yaml
```

---

✅ Now your service is more “microservice-like”:

* JSON-based API
* Query param support
* Health checks for Kubernetes

---

👉 Next logical step (Level 2) could be adding **structured logging + /metrics** for observability.

Do you want me to show you how to integrate **Prometheus metrics (/metrics endpoint)** next, so your microservice can be monitored inside GKE?


http://34.100.194.76/healthz
http://34.100.194.76/hello?name=shridhar