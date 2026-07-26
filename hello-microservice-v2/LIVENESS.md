Liveness Probe
Tells K8s if the container is alive.
If it fails → Kubernetes kills the Pod and restarts it.
Think of it as a health heartbeat.

Readiness Probe
Tells K8s if the Pod is ready to serve traffic.
If it fails → Pod stays running, but Service won’t send traffic to it.
Think of it as a “ready for requests” gate.

Watch Pod status live
kubectl get pods -w

If liveness fails → Pod will restart (you’ll see RESTARTS counter increase).
If readiness fails → Pod stays Running but won’t become Ready.

