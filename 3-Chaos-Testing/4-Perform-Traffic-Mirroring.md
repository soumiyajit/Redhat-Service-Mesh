# Perform Traffic Mirroring

## Prerequisites
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

## Running

First, make sure everything is running correctly.

```
oc get all

http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
```

##### Check the container logs of the hello-message pod
	oc logs hello-message-v2-6dcc4fff9-hnbxs -c hello-message


Next, configure traffic mirroring for the hello-message v2 virtual service.

```yaml
  mirror:
    host: hello-message
    subset: v2
  mirror_percent: 100
```

Apply the changes to the virtual service, invoke the service and finally check the container logs.

```
oc  apply -f hello-message-v2-mirroring.yaml

http get $INGRESS_HOST/api/hello 

http get $INGRESS_HOST/api/hello 

oc logs hello-message-v2-6dcc4fff9-hnbxs -c hello-message

```