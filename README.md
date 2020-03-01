# Kuberdon
Stop copying your docker registry secrets to each namespace, use Kuberdon instead.

Kuberdon is a dead simple controller. It copies your secret to all desired namespaces and optionally adds it automatically as an image-pull-secret.

```yaml
apiVersion: kuberdon.kuberty.io/v1
kind: Registry
metadata:
  name: kuberty-gitlab-read
  namespace: kuberdon
spec:
  secret: kuberty-gitlab-read-secret
  namespaces:
  - name: "*"
    add-automatically: true
```

Want to exclude a namespace (rules follow the priority of the namespace list):
```yaml
apiVersion: kuberdon.kuberty.io/v1
kind: Registry
metadata:
  name: kuberty-gitlab-read
  namespace: kuberdon
spec:
  secret: kuberty-gitlab-read-secret
  namespaces:
  - name: kube-system
    exclude: true
  - name: "*"
    add-automatically: true
```
