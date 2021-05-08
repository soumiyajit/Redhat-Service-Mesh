# Authorization for HTTP traffic


## Prerequisites

export INGRESS_HOST=$(oc -n istio-system get route istio-ingressgateway -o jsonpath='{.spec.host}')

```bash
oc create namespace hello-istio

oc label namespace hello-istio istio-injection=enabled

```

# The default namespace will not have the sidecar injection
oc label namespace default istio-injection=disabled

```

oc apply -f demo

oc apply -f hello-istio-secure.yaml

oc apply -f hello-istio-insecure.yaml

while true; do echo; date; curl -kv  hello-istio:8080/api/hello  ; sleep 10; done

while true; do echo; date; curl -kv hello-istio.hello-istio.svc.cluster.local:8080/api/hello ; sleep 10; done

```

## Running

Any HTTP traffic inside the service mesh can be enabled to disabled using an `AuthorizationPolicy`.

```yaml
apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
  name: hello-message-http-policy
  namespace: hello-istio
spec:
  selector:
    matchLabels:
      app: hello-message
  # toggle between DENY and ALLOW
  action: DENY
  rules:
  - to:
      - operation:
          methods: ["GET"]
    from:
      - source:
          namespaces:
            - "hello-istio"
```

Apply the above policy and check that the HTTP communication within the `hello-istio` namespace is not possible anymore but it is possible from outside the namespace.

```bash
http get $INGRESS_HOST/api/hello
oc apply -f hello-message-http-policy.yaml

oc get AuthorizationPolicy -n hello-istio

$ wget hello-istio:8080/api/hello -S -O - | more
$ wget hello-istio.hello-istio.svc.cluster.local:8080/api/hello -S -O - | more

while true; do echo; date; wget hello-istio:8080/api/hello -S -O - | more ; sleep 10; done

while true; do echo; date; wget hello-istio.hello-istio.svc.cluster.local:8080/api/hello -S -O - | more ; sleep 10; done

oc delete -f hello-message-http-policy.yaml; oc get authorizationpolicy

```

















