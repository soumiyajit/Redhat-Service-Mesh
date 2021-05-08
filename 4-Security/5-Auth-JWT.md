# Authorization with JWT

## Prerequisites

```
cd httpbin

oc apply -f httpbin.yaml

oc apply -f sleep.yaml

oc apply -f httpbin-gateway.yaml

oc apply -f httpbin-virtualservice.yaml
```

cat httpbin-virtualservice.yaml

	apiVersion: networking.istio.io/v1alpha3
	kind: VirtualService
	metadata:
	  name: httpbin
	  namespace: hello-istio
	spec:
	  hosts:
	  - "*"
	  gateways:
	  - httpbin-gateway
	  http:
	  - route:
	    - destination:
	        port:
	          number: 8000
	        host: httpbin.hello-istio.svc.cluster.local
 
Now export the INGRESS_HOST path and check the return type for the http request to the httpbin application:

```

export INGRESS_HOST=$(oc -n istio-system get route istio-ingressgateway -o jsonpath='{.spec.host}')

echo $INGRESS_HOST

curl $INGRESS_HOST -s -o /dev/null -w "%{http_code}\n"

```

Now, add a policy that requires end-user JWT for httpbin.hello-istio. 
The next command assumes there is no service-specific policy for httpbin.foo . 
You can run "oc get policies.authentication.istio.io -n hello-istio" to confirm.

cat auth.yaml

	apiVersion: "authentication.istio.io/v1alpha1"
	kind: "Policy"
	metadata:
	  name: "jwt-example"
	spec:
	  targets:
	  - name: httpbin
	  origins:
	  - jwt:
	      issuer: "testing@secure.istio.io"
	      jwksUri: "https://raw.githubusercontent.com/istio/istio/release-1.4/security/tools/jwt/samples/jwks.json"
	  principalBinding: USE_ORIGIN

Now generate the token for the JWT access:

```
TOKEN=$(curl https://raw.githubusercontent.com/istio/istio/release-1.4/security/tools/jwt/samples/demo.jwt -s)

curl --header "Authorization: Bearer $TOKEN" $INGRESS_HOST/headers -s -o /dev/null -w "%{http_code}\n"
```
