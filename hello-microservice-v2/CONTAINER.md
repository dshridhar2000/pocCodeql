Verify container is actually listening inside pod
Exec into pod:
kubectl exec -it hello-gke-deployment-7745b86f9d-r75g5 -- sh

curl localhost:8080


exec: "sh": executable file not found in $PATH
means that the container image you built doesn’t include /bin/sh (or any shell).

kubectl exec -it <pod-name> -- /hello-gke

------------

kubectl rollout restart deployment hello-gke-deployment
kubectl logs -f deployment/hello-gke-deployment
kubectl get svc
curl <EXTERNAL-IP>:<PORT>
