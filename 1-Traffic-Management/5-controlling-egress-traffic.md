# Controlling Egress Traffic

## Prerequisites

This working exercise assumes that you have running Service Mesh installed in your OpenShift Cluster

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

```
oc label namespace default istio-injection=enabled

oc label namespace istio-demo istio-injection=enabled

export INGRESS_HOST=$(oc -n istio-system get route istio-ingressgateway -o jsonpath='{.spec.host}')
echo $INGRESS_HOST

```

### Deploy sample application

```
oc apply -f hello-istio_with_console.yaml

oc get all

```

### Create ingress gateway and route traffic to microservices

```
oc  apply -f hello-istio-gateway.yaml

oc apply -f hello-istio-virtual-service.yaml

http get $INGRESS_HOST/api/hello 

```

## Running

### Get the current egress mode

```

oc get configmap istio -n istio-system -o yaml | grep -o "mode: ALLOW_ANY"

export SOURCE_POD=$(kubectl get pod -l app=hello-istio-console -o jsonpath={.items..metadata.name})

oc rsh $SOURCE_POD 
$ wget -S -q https://www.google.com

```

### Disable ALLOW_ANY egress mode

```
oc get configmap istio -n istio-system -o yaml | sed 's/mode: ALLOW_ANY/mode: REGISTRY_ONLY/g' | oc replace -n istio-system -f -

oc get configmap istio -n istio-system -o yaml | grep -o "mode: REGISTRY_ONLY"

oc rsh $SOURCE_POD 

wget -S -q https://www.google.com

oc apply -f hello-istio-egress.yaml

oc rsh $SOURCE_POD 

wget -S -q https://www.google.com
```

### To enable the ALLOW_ANY mode

```
oc get configmap istio -n istio-system -o yaml | sed 's/mode: REGISTRY_ONLY/mode: ALLOW_ANY/g' | oc replace -n istio-system -f -

oc get configmap istio -n istio-system -o yaml | grep -o "mode: REGISTRY_ONLY"

oc get configmap istio -n istio-system -o yaml | grep -o "mode: ALLOW_ANY"

```

