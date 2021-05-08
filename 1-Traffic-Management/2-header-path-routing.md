# Path and Header based Routing

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

## Routing -

#apply the version subsets as destinations

```
oc apply -f hello-istio-destination.yaml

```

# apply path based routing
```
oc apply -f hello-istio-uri-match.yaml

http get $INGRESS_HOST/api/hello 

http get $INGRESS_HOST/api/v1/hello 

http get $INGRESS_HOST/api/v2/hello 

```

# apply header based routing
```
oc apply -f hello-istio-user-agent.yaml

http get $INGRESS_HOST/api/hello User-Agent:Chrome 

oc apply -f hello-istio-header-routing.yaml 

curl -kv $INGRESS_HOST/api/hello -H "v2-header: canary" 
```


