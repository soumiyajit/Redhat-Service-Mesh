# Using Envoy Filters

## Prerequisites

This work item assumes that you have running Istio installation on your OpenShift cluster.

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

## Running

First, make sure everything is running correctly.

```
oc get all

oc apply -f hello-istio.yaml
oc apply -f hello-istio-gateway.yaml
oc apply -f hello-istio-v1-virtual-service.yaml
oc apply -f hello-istio-destination.yaml

oc apply -f hello-message-virtual-service.yaml
oc apply -f hello-message-destination.yaml

oc get svc istio-ingressgateway -n istio-system


export INGRESS_HOST=$(oc -n istio-system get route istio-ingressgateway -o jsonpath='{.spec.host}')
echo $INGRESS_HOST

```

Apply the prepared envoy filter manifest and check that everything is working as expected. It does take a while for the sidecar to pick up the new filter configuration.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: hello-message-sam
spec:
  workloadSelector:
    labels:
      app: hello-message
  configPatches:
  - applyTo: HTTP_FILTER
    match:
      context: SIDECAR_INBOUND
      listener:
        filterChain:
          filter:
            name: "envoy.http_connection_manager"
            subFilter:
              name: "envoy.router"
    patch:
      operation: INSERT_BEFORE
      value:
       name: envoy.lua
       config:
         inlineCode: |
           function envoy_on_request(request_handle)
             -- send back static response and do not continue
             request_handle:respond({[":status"] = "200"}, "MSS Envoy Filtered Message")
           end

           function envoy_on_response(request_handle)
             -- add response specific logic here
           end
```

```
oc apply -f hello-message-v1-filter.yaml

watch -n 1 -d http get $INGRESS_HOST/api/hello 

oc logs hello-message-v1-6dcc4fff9-hnbxs -c hello-message

```