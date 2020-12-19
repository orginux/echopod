# pod-info
The minimal image with web-server for debugging network in Docker or Kubernetes

## Usage

### Docker
```bash
docker run --rm -d -p 80:80 orginux/pod-info
```

```bash
$ curl localhost:80
Hostname: 837b80954f04
IP: 172.17.0.2
URI: /
```

### Kubernetes

Create Deployment
```bash
export DEPLOY_NAME="example"
kubectl create deployment $DEPLOY_NAME --image=orginux/pod-info
```

Optional scaling
```bash
kubectl scale deployment $DEPLOY_NAME --replicas=5
```


Ceate Service
```bash
kubectl expose deployment $DEPLOY_NAME --name ${DEPLOY_NAME}-service --target-port 80 --port 80
```

and add rule Ingress
```yaml
spec:
  rules:
    - host: my.host.name
      http:
        paths:
          - path: /debug
            backend:
              serviceName: example-service
              servicePort: 80
```


Or forward port
```bash
kubectl port-forward deployment/${DEPLOY_NAME} 8080:80
```


Get content
```
$ curl http://my.host.name/debug
Hostname: deploy-name-5757fb5f64-k4jzv
IP: 10.233.64.51
Namespace: default
URI: /debug
```
