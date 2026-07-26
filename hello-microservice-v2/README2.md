Got it 👍
Here’s an updated **README.md** for your new **Hello Microservice with liveness & readiness probe break simulation**:

---

# Hello Microservice on GKE

This is a simple **Go microservice** deployed on **Google Kubernetes Engine (GKE)**.
It demonstrates **Kubernetes liveness and readiness probes** with the ability to simulate failures for testing.

---

## Features

* `/` → Returns **Hello, <name>!**
* `/healthz` → Endpoint used by Kubernetes for **liveness** and **readiness probes**

  * Accepts query parameters to simulate probe failures:

    * `break=liveness` → Breaks **liveness probe**
    * `break=readiness` → Breaks **readiness probe**
    * `time=N` → Breaks the probe for **N seconds**

---

## Example

### Normal Hello

```bash
curl "http://<EXTERNAL-IP>/?name=Shridhar"
```

**Response:**

```
Hello, Shridhar!
```

---

### Break Readiness for 20s

```bash
curl "http://<EXTERNAL-IP>/healthz?break=readiness&time=20"
```

* Readiness will fail for **20 seconds**.
* During this time, Kubernetes will **stop sending traffic** to this pod.

---

### Break Liveness for 15s

```bash
curl "http://<EXTERNAL-IP>/healthz?break=liveness&time=15"
```

* Liveness will fail for **15 seconds**.
* If failure persists beyond the configured threshold, Kubernetes will **restart the pod**.

---

## Deploy to GKE

### 1. Build & Push Docker Image

```bash
docker build -t gcr.io/<PROJECT_ID>/hello-gke:v2 .
docker push gcr.io/<PROJECT_ID>/hello-gke:v2
```

### 2. Update Deployment

Edit your `deployment.yaml` to use the new image:

```yaml
containers:
- name: hello-gke
  image: gcr.io/<PROJECT_ID>/hello-gke:v2
  ports:
  - containerPort: 8080
  livenessProbe:
    httpGet:
      path: /healthz
      port: 8080
    initialDelaySeconds: 5
    periodSeconds: 5
  readinessProbe:
    httpGet:
      path: /healthz
      port: 8080
    initialDelaySeconds: 5
    periodSeconds: 5
```

Apply changes:

```bash
kubectl apply -f deployment.yaml
```

### 3. Get External IP

```bash
kubectl get svc hello-gke-service
```

---

## Testing in GCP Console

1. Go to **GCP Console → Kubernetes Engine → Workloads**
2. Select your microservice pod
3. Under **Probes tab** you can see liveness & readiness status
4. Call `/healthz` with `break` params to watch how probes fail and recover

---

✅ With this setup, you can visually test how **Kubernetes health checks behave** when services become unhealthy and then recover.

---

Do you want me to also add a **step-by-step "How to observe pod restart/readiness change in GKE console"** guide inside the README so your teammates can follow along?
