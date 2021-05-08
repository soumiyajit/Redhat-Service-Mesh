# Injecting HTTP Delay Faults

## Prerequisites

This work item assumes that you have running Istio installation on your OpenShift cluster.

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

## Running

Apply the modified virtual service and check that the HTTP delay is configured correctly.

	kubectl apply -f kubernetes/hello-message-v1-delay.yaml

##### You should see an error message after 3s delay -> timeout working

	http get $INGRESS_HOST/api/hello Host:hello-istio.cloud


## Running
```
oc new-project service-mesh-chaos-testing

oc label namespace service-mesh-chaos-testing istio-injection=enabled

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

### First, make sure everything is running correctly without delays.
```
oc get all
http get $INGRESS_HOST/api/hello
```

Next, configure a HTTP delay fault for any traffic to the hello-message v1 virtual service.

```yaml
- fault:
    delay:
      percentage:
        value: 100.0
      fixedDelay: 5s
  route:
  - destination:
      host: hello-message
      subset: v1
```
```
oc apply -f hello-message-v1-delay.yaml

curl -kv  $INGRESS_HOST/api/hello

oc apply -f hello-message-v1-delay-2s.yaml

curl -kv  $INGRESS_HOST/api/hello

oc apply -f hello-message-v1-delay-50-percent.yaml
```

##### To apply the virtual service without any delay:
	oc apply -f hello-message-virtual-service.yaml

## deletion method
```

oc delete -f hello-istio.yaml
oc delete -f hello-istio-gateway.yaml
oc delete -f hello-istio-v1-virtual-service.yaml
oc delete -f hello-istio-destination.yaml

oc delete -f hello-message-virtual-service.yaml
oc delete -f hello-message-destination.yaml


```
