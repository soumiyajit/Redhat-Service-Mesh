# Authorization on Ingress Gateway


## Prerequisites

export INGRESS_HOST=$(oc -n istio-system get route istio-ingressgateway -o jsonpath='{.spec.host}')

```bash
oc create namespace hello-istio

oc label namespace hello-istio istio-injection=enabled

```

##### The default namespace will not have the sidecar injection
oc label namespace default istio-injection=disabled

```

oc apply -f demo

oc apply -f hello-istio-secure.yaml

oc apply -f hello-istio-insecure.yaml

while true; do echo; date; curl -kv  hello-istio:8080/api/hello  ; sleep 10; done

while true; do echo; date; curl -kv hello-istio.hello-istio.svc.cluster.local:8080/api/hello ; sleep 10; done

```

First, make sure that you can call the services without any applied `AuthorizationPolicy` and
that you have configured to the gateway to forward source IP addresses.

```bash

export INGRESS_HOST=$(oc -n istio-system get route istio-ingressgateway -o jsonpath='{.spec.host}')

http get $INGRESS_HOST/api/hello 

```

Next, find out your client IP address and issue the following commands to `DENY` or `ALLOW` any traffic entering the ingress gateway from your IP.

```bash
curl -s 'https://api.ipify.org?format=json'

export CLIENT_IP=$(curl -s 'https://api.ipify.org?format=text')

sed "s/<<CLIENT_IP>>/$CLIENT_IP/" kubernetes/hello-istio-gateway-policy-deny.yaml | kubectl apply -f -

http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

sed "s/<<CLIENT_IP>>/$CLIENT_IP/" kubernetes/hello-istio-gateway-policy-allow.yaml | kubectl apply -f -

http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

#####cleanup
$ kubectl delete AuthorizationPolicy -n istio-system --all

http get $INGRESS_HOST/api/hello

curl -s 'https://api.ipify.org?format=json'
export CLIENT_IP=$(curl -s 'https://api.ipify.org?format=text')

sed "s/<<CLIENT_IP>>/$CLIENT_IP/" hello-istio-gateway-policy-deny.yaml | oc apply -f -

http get $INGRESS_HOST/api/hello

sed "s/<<CLIENT_IP>>/$CLIENT_IP/" hello-istio-gateway-policy-allow.yaml | oc apply -f -

http get $INGRESS_HOST/api/hello

kubectl delete AuthorizationPolicy -n istio-system --all

http get $INGRESS_HOST/api/hello

```
