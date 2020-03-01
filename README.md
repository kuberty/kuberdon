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

## Documentation
This is a specification of the inner workings of kuberdon, you should not worry about it.
### Collission-avoidance
To avoid collissions, kuberdon prefixes all deployed secrets with 'kuberdon-'. 

Kuberdon also sets the ownerReferences to the Kuberdon Registry. If a kuberdon- prefixed secret already exists with a dfferent owner, kuberdon will display this in the status

### Garbage collection
For this to work, Registry objects have to be cluster scoped.
```yaml
ownerReferences:
  - apiVersion: kuberdon.kuberty.io/v1
    blockOwnerDeletion: true
    controller: true
    kind: Registry
    name: kuberty-gitlab-read
    uid: 24c17568-daa9-4cbb-b121-f5bd42dc703a
```
