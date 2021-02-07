# echopod
The minimal HTTP server that provides info about container/pod.

## Usage

### Docker
```bash
docker run --rm -d -p 80:8080 orginux/echopod
```

```bash
$ curl localhost:80
Hostname: 837b80954f04
IP: 172.17.0.2
URI: /
```

### Kubernetes

Create Deployment:
```bash
export DEPLOY_NAME="example"
kubectl create deployment $DEPLOY_NAME --image=orginux/echopod
```

Optional scaling:
```bash
kubectl scale deployment $DEPLOY_NAME --replicas=5
```


Ceate Service:
```bash
kubectl expose deployment $DEPLOY_NAME --port=80 --target-port=8080 --name=${DEPLOY_NAME}-service --type=LoadBalancer
```

Or forward port:
```bash
kubectl port-forward deployment/${DEPLOY_NAME} 8080:8080
```


Get content:
```
$ curl http://external-ip/debug
Hostname: deploy-name-5757fb5f64-k4jzv
IP: 10.0.8.7
URI: /debug
Namespace: default
```
