# Setting request timeouts

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

http get $INGRESS_HOST/api/hello sleep==3

```



Next, edit the virtual service definitions for `hello-istio` and the `hello-message` service to configure the timeouts.

```yaml
    # configure a 2s timeout
    timeout: 2s
```

Issue the following commands to apply and see the timeouts in action.

```bash
oc apply -f hello-istio-virtual-service-timeout.yaml

oc apply -f hello-message-virtual-service-timeout.yaml

http get $INGRESS_HOST/api/hello  sleep==3

http get $INGRESS_HOST/api/hello sleep==1

```