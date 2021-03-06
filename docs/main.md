
## Documentation
This is a specification of the inner workings of kuberdon, you should not worry about it.
### CRD
- [Openshift guide to creating a CRD with kubernetes code-generator](https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/)
- [Guide to creating a Custom Resource Definition](https://github.com/jinghzhu/KubernetesCRD)
- [Code-generator](https://github.com/kubernetes/code-generator)

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
