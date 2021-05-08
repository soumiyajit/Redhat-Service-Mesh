# Injecting HTTP Abort Faults

## Prerequisites

This work item assumes that you have running Istio installation on your OpenShift cluster.

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.


## Running

First, make sure everything is running correctly without any faults.

```
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

Next, configure a HTTP abort fault for any traffic to the hello-message v1 virtual service.

```yaml
- fault:
    abort:
      httpStatus: 500
      percentage:
        value: 100.0
  # optionally add header match here
  route:
  - destination:
      host: hello-message
      subset: v1
```

Apply the modified virtual service and check that the HTTP abort fault is configured correctly.

	oc apply -f hello-message-v1-abort.yaml

##### You should see an error message about the abort filter
	http get $INGRESS_HOST/api/hello 


