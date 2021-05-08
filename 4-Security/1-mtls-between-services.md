# Mutual TLS between services

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

