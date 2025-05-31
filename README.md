# Talon.One DevOps Home Assignment

This repository contains a simple Go HTTP API deployed to Kubernetes using Helm, GitOps practices via ArgoCD, and TLS secured with Traefik.

---

## Tech Stack

* Go (HTTP API)
* Docker
* Kubernetes (Minikube)
* Helm
* ArgoCD (GitOps)
* Traefik (Ingress Controller)
* TLS (Self-signed certificate)
* GitHub Actions (CI/CD)

---

## Features

* Go web API
* Dockerized application
* Kubernetes deployment using Helm
* GitOps deployment with ArgoCD
* Traefik Ingress with TLS termination
* CI pipeline using GitHub Workflows
* Port-forwarding for local HTTPS access

---

## Setup Guide

### 1. Clone the repository

```bash
git clone https://github.com/your-username/talonone-home-task.git
cd talonone-home-task
```

### 2. Start Minikube

```bash
minikube start
```

### 3. Install Traefik (if not already installed)

```bash
helm repo add traefik https://helm.traefik.io/traefik
helm repo update
helm install traefik traefik/traefik
```

### 4. Create TLS Secret

```bash
kubectl create secret tls talonone-test-tls \
  --cert=certs/talonone-test.local.crt \
  --key=certs/talonone-test.local.key \
  --namespace default
```

### 5. Add local DNS record

Edit `/etc/hosts` and add:

```
127.0.0.1  talonone-test.local
```

### 6. Port Forward Traefik (HTTPS)

```bash
kubectl port-forward svc/traefik 9443:443
```

### 7. Access the Application

```bash
curl https://talonone-test.local:9443
```

---

## Deployment via ArgoCD

* Push changes to GitHub repository
* ArgoCD syncs changes automatically or manually
* App is deployed using the Helm chart in `charts/talonone-test`
* ArgoCD is accessed via port-forwarding:

```bash
kubectl port-forward svc/argocd-server -n argocd 8080:80
```

---

## Notes

* Certificates are stored in the `certs/` folder and ignored by Git via `.gitignore`
* Self-signed TLS certificates are used with a custom CA, so `curl` does not require the `-k` flag
* For production, certificates should be managed via `cert-manager` and secrets via Vault
* GitOps model keeps the desired state declarative and auditable

---

## Improvements

* Automate TLS with cert-manager
* Add observability (Prometheus, Grafana)
* Store secrets securely using Vault
* Improve GitHub Actions pipeline with testing, linting, and image publishing

---

## License

MIT

---

## Author

\[Your Name] - DevOps Engineer
