# Controlling Ingress Traffic

## Prerequisites

This working exercise assumes that you have running Service Mesh installed in your OpenShift Cluster

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

```
oc label namespace default istio-injection=enabled

oc label namespace istio-demo istio-injection=enabled

```

## Running

```
export INGRESS_HOST=$(oc -n istio-system get route istio-ingressgateway -o jsonpath='{.spec.host}')

echo $INGRESS_HOST

```

# deploy sample application

```
oc apply -f hello-istio.yaml

oc get all
```

# create ingress gateway and route traffic to microservices

```
oc  apply -f hello-istio-gateway.yaml
oc apply -f hello-istio-virtual-service.yaml

http get $INGRESS_HOST/api/hello 

```

