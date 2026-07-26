apiVersion: apps/v1        # The API version of the Kubernetes object you’re defining (Deployment uses apps/v1)
kind: Deployment           # The type of object: Deployment manages Pods and ensures replicas are running
metadata:
  name: hello-gke          # The name of the Deployment (unique in the namespace)

spec:
  replicas: 1              # Number of Pods to run (here: 1 pod)
  selector:                # Defines how to find Pods managed by this Deployment
    matchLabels:
      app: hello-gke       # Must match labels in template.metadata.labels

  template:                # Pod template (blueprint for the Pods managed by Deployment)
    metadata:
      labels:
        app: hello-gke     # Labels applied to Pods (must match selector above)

    spec:
      containers:          # List of containers inside the Pod
      - name: hello-gke    # Name of the container
        image: gcr.io/mypoc-469116/hello-gke:v1   # Container image from Google Container Registry (GCR)
        ports:
        - containerPort: 8080   # The port the container listens on (your app runs on port 8080)
