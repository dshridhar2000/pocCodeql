apiVersion: v1                    # API version for Service (core group, stable v1)
kind: Service                     # We are creating a Service resource
metadata:
  name: hello-gke-service         # Name of the Service (must be unique in namespace)
spec:
  type: LoadBalancer              # Exposes Service externally using a cloud load balancer (GCP)
  selector:                       # Selects Pods to forward traffic to
    app: hello-gke                # Matches Pods with label "app=hello-gke"
  ports:
  - protocol: TCP                 # Protocol used for communication (default: TCP)
    port: 80                      # Port exposed by Service (clients connect here)
    targetPort: 8080              # Port inside the Pod/container (Go app is listening here)
