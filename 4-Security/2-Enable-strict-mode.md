# Enabling Strict Mode


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


## apply strict mtls policy
```
oc apply  -f hello-istio-strict-mtls.yaml

apiVersion: "authentication.istio.io/v1alpha1"
kind: Policy
metadata:
  name: default
  namespace: hello-istio
spec:
  peers:
  - mtls:
      mode: STRICT

oc apply -f hello-istio-mtls-ports.yaml

oc apply -f hello-istio-destination-rules.yaml


while true; do echo; date; curl -kv  hello-istio:8080/api/hello  ; sleep 10; done

while true; do echo; date; curl -kv hello-istio.hello-istio.svc.cluster.local:8080/api/hello more ; sleep 10; done

```