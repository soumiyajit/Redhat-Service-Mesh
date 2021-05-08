#Install the Service Mesh Operators

OpenShift uses an Operator to install the Red Hat Service Mesh. There are also separate Operators for Elasticsearch, Jaeger, and Kiali. We need all 4 and install them in sequence.

In the Web Console, go to Catalog, OperatorHub and search for Elasticsearch:

![alt text](Images/elasticsearch-operator-install.png?raw=true "Elasticsearch Operator")

Click on the Elasticsearch (provided by Red Hat, Inc.) tile, click “Install”, accept all defaults for “Create Operator Subscription”, and click “Subscribe”. Make sure to select the correct elasticsearch version appropriate for your openshift release. 

In the “Subscription Overview” wait for “UPGRADE STATUS” to be “Up to date”, then check section “Installed Operators” for “STATUS: InstallSucceeded”:

![alt text](Images/elasticsearch-install-success.png?raw=true "Elasticsearch Operator Success")

Repeat these steps for Jaeger. There are Community and Red Hat provided Operators, make sure to use the Red Hat provided ones!

![alt text](Images/elasticsearch-operator-install.png?raw=true "Jaeger Operator")

Confirm from the installed operator tab that the Jaeger installation is success.

![alt text](Images/elasticsearch-install-success.png?raw=true "Jaeger Operator Success")

Next we can install the Kiali operator from the operator hub:

![alt text](Images/kiali-operator-install.png?raw=true "Kiali Operator")

Confirm from the installed operator tab that the Kiali installation is success.

![alt text](Images/kiali-install-success.png?raw=true "Kiali Operator Success")

Next we can install the ServiceMesh operator from the operator hub:

![alt text](Images/servicemesh-operator-install.png?raw=true "ServiceMesh Operator")

Confirm from the installed operator tab that the ServiceMesh installation is success.

![alt text](Images/servicemesh-install-success.png?raw=true "ServiceMesh Operator Success")

### Create the Service Mesh Control Plane

The Service Mesh Control Plane is the actual installation of all Istio components into OpenShift.

We begin with creating a project ‘istio-system’, either in the Web Console or via command line (‘oc new-project istio-system‘) You can actually name the project whatever you like, in fact you can have more than one service mesh in a single OpenShift instance. But to be consistent with Istio I like to stay with ‘istio-system’ as name.

In the Web Console in project: ‘istio-system’ click on “Installed Operators”. You should see all 4 Operators in status “Copied”. The Operators are installed in project ‘openshift-operators’ but we will create the Control Plane in ‘istio-system’. Click on “Red Hat OpenShift Service Mesh”. This Operator provides 2 APIs: ‘Member Role’ and ‘Control Plane’:

![alt text](Images/servicemesh-controlplane.png?raw=true "Control Plane")

Click on “Create New” Control Plane. This opens an editor with a YAML file of kind “ServiceMeshControlPlane”. Look at it but accept it as is. It will create a Control Plane of name ‘basic-install’ with Kiali, Grafana, and Tracing (Jaeger) enabled, Jaeger will use an ‘all-in-one’ template (without Elasticsearch). Click “Create”.

You will now see “basic-install” in the list of Service Mesh Control Planes. Click on “basic-install” and “Resources”. This will display a list of objects that belong to the control plane and this list will grow in the next minutes as more objects are created:

![alt text](Images/controlplane-resources.png?raw=true "ControlPlane Resources")

A good way to check if the installation is complete is by looking into Networking – Routes. You should see 5 routes:

![alt text](Images/servicemesh-routes.png?raw=true "ControlPlane Routes")

Click on the Routes for grafana, jaeger, prometheus, and kiali. Accept the security settings. I click on Kiali last because Kiali is using the other services and in that way all the security settings for those are in place already.

One last thing to do: you need to specify which projects are managed by your Service Mesh Control Plane and this is done by creating a Service Mesh Member Role.

In your project ‘istio-system’ go to “Installed Operator” and click on the “OpenShift Service Mesh” operator. In the Overview, create a new ‘Member Roll’:

![alt text](Images/servicemesh-member-roll.png?raw=true "ServiceMesh Member Roll")

In the YAML file make sure that namespace is indeed ‘istio-system’ and then add all projects to the ‘members’ section that you want to be managed.

![alt text](Images/member-roll-projects.png?raw=true "Member Roll Projects")

You may also add the names of the projects that do not exists. 






