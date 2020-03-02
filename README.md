# Kuberdon
**Stop copying your docker registry secrets to each namespace, use Kuberdon instead.**

Kuberdon is a dead simple controller. It copies your secret to all desired namespaces and automatically binds them to the default serviceaccount. Pods will then automatically use them.

```yaml
apiVersion: kuberdon.kuberty.io/v1
kind: Registry
metadata:
  name: kuberty-gitlab-read
spec:
  secret: kuberty-gitlab-read-secret
  namespaces:
  - name: "*"
```

Do you want to exclude a namespace?
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
Note that the higher the namespace rule, the higher its priority.

## Similar projects
[titansoft-pte-ltd/imagepullsecret-patcher](https://github.com/titansoft-pte-ltd/imagepullsecret-patcher): Very similar, though not kubernetes native (does not use the kubectl api)

To read more about Kuberdon please see our [Documentation](docs/main.md).
