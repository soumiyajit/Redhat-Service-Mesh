# Red Hat OpenShift Service Mesh

## Understanding Red Hat OpenShift Service Mesh

A service mesh is the network of microservices that make up applications in a distributed microservice architecture and the interactions between those microservices. When a Service Mesh grows in size and complexity, it can become harder to understand and manage.

Based on the open source Istio project, Red Hat OpenShift Service Mesh adds a transparent layer on existing distributed applications without requiring any changes to the service code. You add Red Hat OpenShift Service Mesh support to services by deploying a special sidecar proxy to relevant services in the mesh that intercepts all network communication between microservices. You configure and manage the Service Mesh using the control plane features.

Red Hat OpenShift Service Mesh gives us an easy way to create a network of deployed services that provide: 

```
 - Discovery 
 - Load balancing 
 - Service-to-service authentication 
 - Failure recovery 
 - Metrics 
```


Service Mesh also provides more complex operational functions including:

```
 - A/B testing
 - Canary releases
 - Rate limiting
 - Access control
 - End-to-end authentication
```

## Red Hat OpenShift Service Mesh Architecture

![alt text](Images/servicemesh-architecture.png?raw=true "Service Mesh Architecture")

Red Hat OpenShift Service Mesh is logically split into a data plane and a control plane:

**Data Plane** is a set of intelligent proxies deployed as sidecars. These proxies intercept and control all inbound and outbound network communication between microservices in the service mesh. Sidecar proxies also communicate with Mixer, the general-purpose policy and telemetry hub.

**Envoy Proxy** intercepts all inbound and outbound traffic for all services in the service mesh. Envoy is deployed as a sidecar to the relevant service in the same pod.

**Control Plane** manages and configures proxies to route traffic, and configures Mixers to enforce policies and collect telemetry.

**Mixer** enforces access control and usage policies (such as authorization, rate limits, quotas, authentication, and request tracing) and collects telemetry data from the Envoy proxy and other services.

**Pilot** configures the proxies at runtime. Pilot provides service discovery for the Envoy sidecars, traffic management capabilities for intelligent routing (for example, A/B tests or canary deployments), and resiliency (timeouts, retries, and circuit breakers).

**Citadel** issues and rotates certificates. Citadel provides strong service-to-service and end-user authentication with built-in identity and credential management. You can use Citadel to upgrade unencrypted traffic in the service mesh. Operators can enforce policies based on service identity rather than on network controls using Citadel.

**Galley** ingests the service mesh configuration, then validates, processes, and distributes the configuration. Galley protects the other service mesh components from obtaining user configuration details from OpenShift Container Platform.

Red Hat OpenShift Service Mesh also uses the istio-operator to manage the installation of the control plane. An Operator is a piece of software that enables you to implement and automate common activities in your OpenShift cluster. It acts as a controller, allowing you to set or change the desired state of objects in your cluster.


### Multi-tenancy in Red Hat OpenShift Service Mesh versus cluster-wide installations

The main difference between a multi-tenant installation and a cluster-wide installation is the scope of privileges used by the control plane deployments, for example, Galley and Pilot. The components no longer use cluster-scoped Role Based Access Control (RBAC) resource ClusterRoleBinding, but rely on project-scoped RoleBinding.

Every project in the members list will have a RoleBinding for each service account associated with a control plane deployment and each control plane deployment will only watch those member projects. Each member project has a maistra.io/member-of label added to it, where the member-of value is the project containing the control plane installation.

Red Hat OpenShift Service Mesh configures each member project to ensure network access between itself, the control plane, and other member projects. The exact configuration differs depending on how OpenShift software-defined networking (SDN) is configured. See About OpenShift SDN for additional details.

If the OpenShift Container Platform cluster is configured to use the SDN plug-in:

- NetworkPolicy: Red Hat OpenShift Service Mesh creates a NetworkPolicy resource in each member project allowing ingress to all pods from the other members and the control plane. If you remove a member from Service Mesh, this NetworkPolicy resource is deleted from the project.

- **These network policies  also restricts ingress to only member projects. If ingress from non-member projects is required, you need to create a NetworkPolicy to allow that traffic through.** 
Example: if we add the member project to the service mesh we may find the service mesh applications functioning normally as it is using the ingress gateway however to make the other non istio applications reachable we may require to add the network policy to add allow all the ingress and egress traffic for non-istio applications.
Sample network policy to allow all traffic:

```
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-all-traffic
spec:
  podSelector: {}
  ingress:
  - {}
  egress:
  - {}
  policyTypes:
  - Ingress
  - Egress
```

* Multitenant: Red Hat OpenShift Service Mesh joins the NetNamespace for each member project to the NetNamespace of the control plane project (the equivalent of running *oc adm pod-network join-projects --to control-plane-project member-project*). If you remove a member from the Service Mesh, its NetNamespace is isolated from the control plane (the equivalent of running oc adm pod-network isolate-projects member-project).

### Automatic injection
The upstream Istio community installation automatically injects the sidecar into pods within the projects you have labeled.

Red Hat OpenShift Service Mesh does not automatically inject the sidecar to any pods, but requires you to specify the ***sidecar.istio.io/inject*** annotation.