# Adding a Circuit Breaker

## Prerequisites

```
oc label namespace istio-demo istio-injection=enabled

export INGRESS_HOST=$(oc -n istio-system get route istio-ingressgateway -o jsonpath='{.spec.host}')

echo $INGRESS_HOST

```

## Running

```bash
oc apply -f hello-istio.yaml
oc apply -f hello-istio-gateway.yaml
oc apply -f hello-istio-virtual-service.yaml
oc apply -f hello-istio-destination.yaml

oc get all
curl -kv  $INGRESS_HOST/api/hello 

```

Next, edit the `hello-message-destination.yaml` and add the traffic policy and circuit breaker definitions.

```yaml
  trafficPolicy:
    outlierDetection:
      consecutiveErrors: 5    # 5 upstream errors (502, 503, 504)
      interval: 30s           # sliding window of 30s
      baseEjectionTime: 1m    # eject upstream for 1 minute
      maxEjectionPercent: 50  # max 50% of upstream hosts ejected
```

Now apply the virtual service and destination definitions for the `hello-message` service.

```bash
oc apply -f hello-message-virtual-service.yaml
oc apply -f hello-message-destination.yaml

curl -kv  $INGRESS_HOST/api/hello 