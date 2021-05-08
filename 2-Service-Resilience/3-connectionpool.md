# Connection Pools and Bulk Heading

## Prerequisites

```
oc label namespace default istio-injection=enabled

oc get svc istio-ingressgateway -n istio-system

export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')

```

## Running

```

oc apply -f hello-istio.yaml

oc apply -f hello-istio-gateway.yaml

oc apply -f hello-istio-virtual-service.yaml

oc apply -f hello-istio-destination.yaml

oc apply -f hello-message-virtual-service.yaml

oc apply -f hello-message-destination.yaml

oc get all
curl -kv  $INGRESS_HOST/api/hello 

```

Next, edit the destination rule definitions for `hello-istio` and  `hello-message` to configure the connection pools and bulk heads.

```yaml
  trafficPolicy:
    connectionPool:
      tcp:
        maxConnections: 100             # maximum number of TCP conns
        connectTimeout: 5s              # TCP connection timeout
        tcpKeepalive:                   # keep alive settings
          time: 3600s
          interval: 60s
      http:
        maxRequestsPerConnection: 25    # max request per keep-alive
        http2MaxRequests: 5             # max number of HTTP2 conns
        http1MaxPendingRequests: 5      # max number of pending reqs
        maxRetries: 3                   # max number of retries
        idleTimeout: 60s                # idle timeout for connection
```

Now apply the modified destination rule definitions for the two services.

```bash
oc apply -f hello-istio-connection-pool.yaml

oc apply -f hello-message-connection-pool.yaml

http get $INGRESS_HOST/api/hello 
```
