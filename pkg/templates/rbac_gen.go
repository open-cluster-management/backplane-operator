// Code generated by go generate; DO NOT EDIT.

package main

//+kubebuilder:rbac:groups="",resources=configmaps,verbs=create
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=create;get;list;watch
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=create;update;get;delete
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;create;update;patch
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;update;delete
//+kubebuilder:rbac:groups="",resources=configmaps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups="",resources=configmaps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups="",resources=configmaps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups="",resources=configmaps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups="",resources=configmaps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups="",resources=configmaps;jobs;namespaces;pods;secrets,verbs=list;watch
//+kubebuilder:rbac:groups="",resources=configmaps;namespaces;serviceaccounts;services;secrets,verbs=create;get;list;update;watch;patch;delete
//+kubebuilder:rbac:groups="",resources=configmaps;secrets,verbs=*
//+kubebuilder:rbac:groups="",resources=configmaps;secrets,verbs=*
//+kubebuilder:rbac:groups="",resources=configmaps;secrets,verbs=*
//+kubebuilder:rbac:groups="",resources=events,verbs=create
//+kubebuilder:rbac:groups="",resources=events,verbs=create
//+kubebuilder:rbac:groups="",resources=events,verbs=create
//+kubebuilder:rbac:groups="",resources=events,verbs=create;patch
//+kubebuilder:rbac:groups="",resources=events,verbs=create;patch
//+kubebuilder:rbac:groups="",resources=events,verbs=create;patch
//+kubebuilder:rbac:groups="",resources=events,verbs=create;patch
//+kubebuilder:rbac:groups="",resources=namespaces,verbs=delete
//+kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=namespaces;secrets;pods;pods/portforward,verbs=*
//+kubebuilder:rbac:groups="",resources=nodes,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=nodes;pods;endpoints;services;secrets,verbs=get;watch;list
//+kubebuilder:rbac:groups="",resources=persistentvolumeclaims,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups="",resources=pods,verbs=list
//+kubebuilder:rbac:groups="",resources=pods;nodes,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=pods;pods/log,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=pods;services;endpoints,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=pods;services;endpoints,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=secrets,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups="",resources=secrets,verbs=create;get;list;update;watch
//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;update;watch;patch
//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;watch;list;create
//+kubebuilder:rbac:groups="",resources=secrets,verbs=list;get;watch
//+kubebuilder:rbac:groups="",resources=secrets;configmaps;events;services,verbs=get;list;watch;create;update;delete;deletecollection;patch
//+kubebuilder:rbac:groups="",resources=secrets;events,verbs=get;list;watch;create;update;patch
//+kubebuilder:rbac:groups="",resources=secrets;namespaces,verbs=list;get;watch;delete
//+kubebuilder:rbac:groups="",resources=serviceaccounts,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups="",resources=serviceaccounts,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups="",resources=serviceaccounts/token,verbs=create
//+kubebuilder:rbac:groups="",resources=serviceaccounts;serviceaccounts/finalizers;secrets;secrets/finalizers;services;services/finalizers;endpoints;events;configmaps;namespaces;persistentvolumeclaims,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=serviceaccounts;services,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=services,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups="",resources=services,verbs=list
//+kubebuilder:rbac:groups="",resources=services;events;serviceaccounts,verbs=*
//+kubebuilder:rbac:groups="",resources=services;services/finalizers;events;configmaps;secrets;serviceaccounts;namespaces,verbs=list;create;update;get;watch;patch;delete
//+kubebuilder:rbac:groups="";coordination.k8s.io,resources=configmaps;leases,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="";coordination.k8s.io,resources=configmaps;leases,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="";events.k8s.io,resources=configmaps;namespaces;events;endpoints;secrets,verbs=get;list;watch;create;update;delete;deletecollection;patch
//+kubebuilder:rbac:groups="";events.k8s.io,resources=events,verbs=create;patch;update
//+kubebuilder:rbac:groups="";events.k8s.io,resources=events,verbs=create;patch;update
//+kubebuilder:rbac:groups="";events.k8s.io,resources=events,verbs=create;update;patch
//+kubebuilder:rbac:groups=action.open-cluster-management.io,resources=managedclusteractions,verbs=get;create;update;delete
//+kubebuilder:rbac:groups=action.open-cluster-management.io,resources=managedclusteractions,verbs=get;list;watch
//+kubebuilder:rbac:groups=action.open-cluster-management.io,resources=managedclusteractions/status,verbs=update;patch
//+kubebuilder:rbac:groups=action.open-cluster-management.io,resources=managedclusteractions;managedclusteractions/status,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=addondeploymentconfigs,verbs=get;list;watch
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=addondeploymentconfigs,verbs=get;list;watch
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=addondeploymentconfigs,verbs=get;list;watch
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=addondeploymentconfigs,verbs=get;list;watch
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=clustermanagementaddons,verbs=get;list;watch
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=clustermanagementaddons,verbs=get;list;watch
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=clustermanagementaddons/finalizers,verbs=update
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=clustermanagementaddons/finalizers,verbs=update
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=clustermanagementaddons;managedclusteraddons,verbs=list;watch
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=clustermanagementaddons;managedclusteraddons;clustermanagementaddons/status;clustermanagementaddons/finalizers;managedclusteraddons/status,verbs=*
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=managedclusteraddons,verbs=get;list;watch
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=managedclusteraddons,verbs=get;list;watch;create;update;delete
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=managedclusteraddons,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=managedclusteraddons/finalizers,verbs=*
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=managedclusteraddons/finalizers,verbs=update
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=managedclusteraddons/finalizers,verbs=update
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=managedclusteraddons/status,verbs=update;patch
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=managedclusteraddons/status,verbs=update;patch
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=managedclusteraddons;managedclusteraddons/finalizers,verbs=get;list;watch;patch;update;delete;deletecollection
//+kubebuilder:rbac:groups=addon.open-cluster-management.io,resources=managedclusteraddons;managedclusteraddons/finalizers;managedclusteraddons/status;clustermanagementaddons;clustermanagementaddons/finalizers,verbs=create;delete;get;list;watch;patch;update
//+kubebuilder:rbac:groups=admission.hive.openshift.io,resources=clusterdeployments;clusterimagesets;clusterprovisions;dnszones;machinepools;selectorsyncsets;syncsets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=admission.hive.openshift.io,resources=dnszones,verbs=get;list;watch
//+kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations;mutatingwebhookconfigurations,verbs=create;get;list;update;watch;patch;delete
//+kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations;mutatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=agentclassifications,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=agentclassifications/finalizers,verbs=update
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=agentclassifications/status,verbs=get;patch;update
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=agents,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=agents/ai-deprovision,verbs=update
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=agents/status,verbs=get;patch;update
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=agents;infraenvs,verbs=get;list;watch;update
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=agents;infraenvs;nmstateconfigs;agentserviceconfigs,verbs=list;watch
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=agentserviceconfigs,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=agentserviceconfigs/finalizers,verbs=update
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=agentserviceconfigs/status,verbs=get;patch;update
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=hypershiftagentserviceconfigs,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=hypershiftagentserviceconfigs/finalizers,verbs=update
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=hypershiftagentserviceconfigs/status,verbs=get;patch;update
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=infraenvs,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=infraenvs,verbs=get;list;watch
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=infraenvs/status,verbs=get;patch;update
//+kubebuilder:rbac:groups=agent-install.openshift.io,resources=nmstateconfigs,verbs=get;list;watch
//+kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=create;get;list;update;watch;patch;delete
//+kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;create;update;patch
//+kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch
//+kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch
//+kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=list;watch
//+kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions/status,verbs=update;patch
//+kubebuilder:rbac:groups=apiregistration.k8s.io,resources=apiservices,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=apiregistration.k8s.io,resources=apiservices,verbs=create;get;list;update;watch;patch;delete
//+kubebuilder:rbac:groups=apiregistration.k8s.io,resources=apiservices;apiservices/finalizers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=create;get;list;update;watch;patch;delete
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch
//+kubebuilder:rbac:groups=apps,resources=deployments/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments;deployments/finalizers,verbs=get;list;create;delete;update;patch
//+kubebuilder:rbac:groups=apps,resources=deployments;deployments/finalizers;daemonsets;daemonsets/finalizers;statefulsets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=deployments;deployments/scale,verbs=*
//+kubebuilder:rbac:groups=apps,resources=replicasets,verbs=get
//+kubebuilder:rbac:groups=apps,resources=replicasets;deployments,verbs=get;list;watch;create;update;patch
//+kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=apps.open-cluster-management.io,resources=deployables;deployables/status,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=apps.openshift.io,resources=deploymentconfigs,verbs=get;list;watch
//+kubebuilder:rbac:groups=authentication.k8s.io,resources=tokenrequests,verbs=create
//+kubebuilder:rbac:groups=authentication.k8s.io,resources=tokenreviews,verbs=create
//+kubebuilder:rbac:groups=authentication.k8s.io,resources=tokenreviews,verbs=create
//+kubebuilder:rbac:groups=authentication.open-cluster-management.io,resources=managedserviceaccounts;managedserviceaccounts/status,verbs=get;list;watch;update;patch;delete
//+kubebuilder:rbac:groups=authorization.k8s.io,resources=subjectaccessreviews,verbs=create
//+kubebuilder:rbac:groups=authorization.k8s.io,resources=subjectaccessreviews,verbs=create
//+kubebuilder:rbac:groups=authorization.k8s.io,resources=subjectaccessreviews,verbs=create
//+kubebuilder:rbac:groups=authorization.k8s.io,resources=subjectaccessreviews,verbs=create
//+kubebuilder:rbac:groups=authorization.k8s.io,resources=subjectaccessreviews,verbs=get;create
//+kubebuilder:rbac:groups=authorization.k8s.io,resources=subjectaccessreviews,verbs=get;create
//+kubebuilder:rbac:groups=authorization.openshift.io,resources=clusterroles;clusterrolebindings,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;watch;list;delete;create
//+kubebuilder:rbac:groups=batch;hive.openshift.io;tower.ansible.com;"",resources=ansiblejobs;jobs;clusterdeployments;serviceaccounts;machinepools,verbs=get
//+kubebuilder:rbac:groups=capi-provider.agent-install.openshift.io,resources=agentmachines,verbs=list;watch
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests,verbs=create;get;list;watch
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests,verbs=list;watch
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests/approval,verbs=update
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests;certificatesigningrequests/approval,verbs=create;get;list;watch;patch;update
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests;certificatesigningrequests/approval,verbs=get;list;watch;create;update
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests;certificatesigningrequests/approval;certificatesigningrequests/status,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests;certificatesigningrequests/approval;certificatesigningrequests/status,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=signers,verbs=*
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=signers,verbs=*
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=signers,verbs=approve
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=signers,verbs=approve
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=signers,verbs=approve
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=addonplacementscores/status,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=addonplacementscores;addonplacementscores/status,verbs=create;delete;deletecollection;get;list;patch;update;watch
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=clustercurators,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=clustercurators/status,verbs=get;patch;update
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclusters,verbs=get;list;create;delete;watch;update;patch
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclusters,verbs=get;list;watch
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclusters,verbs=get;list;watch
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclusters,verbs=get;list;watch
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclusters/status,verbs=patch;update
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclusters;managedclusters/status;managedclusters/finalizers,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclusters;managedclustersets,verbs=get;list;watch
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclusters;managedclustersets;managedclustersetbindings,verbs=get;list;watch;patch;update
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclusters;managedclustersets;managedclustersetbindings;clustercurators;placements;placementdecisions,verbs=list;watch
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclusters;managedclustersets;placementdecisions;placementdecisions/status,verbs=get;list;watch;update
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclustersetbindings,verbs=create;update;get;list;watch;delete;deletecollection;patch
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclustersets/bind,verbs=create
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclustersets/join,verbs=create
//+kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=managedclustersets/join,verbs=create
//+kubebuilder:rbac:groups=clusterview.open-cluster-management.io,resources=managedclusters;managedclustersets,verbs=list;get;watch
//+kubebuilder:rbac:groups=config.openshift.io,resources=apiservers,verbs=get;list;watch
//+kubebuilder:rbac:groups=config.openshift.io,resources=clusteroperators,verbs=get;list;watch
//+kubebuilder:rbac:groups=config.openshift.io,resources=clusterversions,verbs=get
//+kubebuilder:rbac:groups=config.openshift.io,resources=clusterversions,verbs=get;list
//+kubebuilder:rbac:groups=config.openshift.io,resources=infrastructures,verbs=get;list;watch
//+kubebuilder:rbac:groups=config.openshift.io,resources=infrastructures,verbs=get;list;watch;patch;update
//+kubebuilder:rbac:groups=config.openshift.io;console.openshift.io;project.openshift.io;tower.ansible.com,resources=infrastructures;consolelinks;projects;featuregates;ansiblejobs;clusterversions,verbs=list;get;watch
//+kubebuilder:rbac:groups=console.open-cluster-management.io,resources=userpreferences,verbs=list;watch
//+kubebuilder:rbac:groups=console.openshift.io,resources=consoleclidownloads,verbs=get;list;create;delete;update;patch
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=*
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=*
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=create;get;list;update;watch;patch
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=create;get;update;delete
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;create;update
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;create;update
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;create;update
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;create;update
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;create;update;patch
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;watch;update;create;patch
//+kubebuilder:rbac:groups=discovery.open-cluster-management.io,resources=discoveredclusters,verbs=create;delete;deletecollection;get;list;patch;update;watch
//+kubebuilder:rbac:groups=discovery.open-cluster-management.io,resources=discoveredclusters/finalizers,verbs=get;patch;update
//+kubebuilder:rbac:groups=discovery.open-cluster-management.io,resources=discoveredclusters/status,verbs=get;patch;update
//+kubebuilder:rbac:groups=discovery.open-cluster-management.io,resources=discoveryconfigs,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=discovery.open-cluster-management.io,resources=discoveryconfigs/finalizers,verbs=get;patch;update
//+kubebuilder:rbac:groups=discovery.open-cluster-management.io,resources=discoveryconfigs/status,verbs=get;patch;update
//+kubebuilder:rbac:groups=discovery.open-cluster-management.io,resources=discoveryconfigs;discoveredclusters,verbs=list;watch
//+kubebuilder:rbac:groups=extensions.hive.openshift.io,resources=*,verbs=*
//+kubebuilder:rbac:groups=extensions.hive.openshift.io,resources=agentclusterinstalls,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=extensions.hive.openshift.io,resources=agentclusterinstalls,verbs=get;list;watch;update
//+kubebuilder:rbac:groups=extensions.hive.openshift.io,resources=agentclusterinstalls,verbs=list;watch
//+kubebuilder:rbac:groups=extensions.hive.openshift.io,resources=agentclusterinstalls/finalizers,verbs=update
//+kubebuilder:rbac:groups=extensions.hive.openshift.io,resources=agentclusterinstalls/status,verbs=get;patch;update
//+kubebuilder:rbac:groups=flowcontrol.apiserver.k8s.io,resources=prioritylevelconfigurations;flowschemas,verbs=get;list;watch
//+kubebuilder:rbac:groups=hive.openshift.io,resources=*,verbs=*
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterclaims;clusterdeployments;clusterpools;clusterimagesets;clusterprovisions;clusterdeprovisions;machinepools,verbs=list;watch
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterclaims;clusterpools,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterdeployments,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterdeployments,verbs=get
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterdeployments,verbs=get;list;watch
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterdeployments,verbs=patch;delete;update
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterdeployments/finalizers,verbs=update
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterdeployments/status,verbs=get;patch;update
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterdeployments;clusterpools;clusterclaims;machinepools,verbs=*
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterdeployments;clusterpools;clusterclaims;machinepools,verbs=get;list;watch
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterdeployments;clusterpools;clusterclaims;machinepools,verbs=get;list;watch;update
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterdeployments;syncsets;selectorsyncsets,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterimagesets,verbs=get;list;watch
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterimagesets,verbs=get;list;watch
//+kubebuilder:rbac:groups=hive.openshift.io,resources=clusterimagesets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hiveinternal.openshift.io,resources=*,verbs=*
//+kubebuilder:rbac:groups=hiveinternal.openshift.io,resources=clustersyncs,verbs=get;list;watch
//+kubebuilder:rbac:groups=hypershift.openshift.io,resources=hostedclusters;nodepools,verbs=list;watch
//+kubebuilder:rbac:groups=imageregistry.open-cluster-management.io,resources=managedclusterimageregistries;managedclusterimageregistries,verbs=get;list;watch
//+kubebuilder:rbac:groups=imageregistry.open-cluster-management.io,resources=managedclusterimageregistries;managedclusterimageregistries/status,verbs=get;list;watch;update
//+kubebuilder:rbac:groups=internal.open-cluster-management.io,resources=managedclusterinfos,verbs=get;list;watch
//+kubebuilder:rbac:groups=internal.open-cluster-management.io,resources=managedclusterinfos,verbs=get;list;watch
//+kubebuilder:rbac:groups=internal.open-cluster-management.io,resources=managedclusterinfos,verbs=list;watch
//+kubebuilder:rbac:groups=internal.open-cluster-management.io,resources=managedclusterinfos/status,verbs=update;patch
//+kubebuilder:rbac:groups=internal.open-cluster-management.io,resources=managedclusterinfos;managedclusterinfos/status,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=internal.open-cluster-management.io;"",resources=managedclusterinfos;pods;secrets,verbs=get
//+kubebuilder:rbac:groups=metal3.io,resources=baremetalhosts,verbs=get;list;patch;update;watch
//+kubebuilder:rbac:groups=metal3.io,resources=baremetalhosts,verbs=get;list;watch;update
//+kubebuilder:rbac:groups=metal3.io,resources=baremetalhosts;provisionings,verbs=list;watch
//+kubebuilder:rbac:groups=metal3.io,resources=preprovisioningimages,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=metal3.io,resources=preprovisioningimages/status,verbs=get;patch;update
//+kubebuilder:rbac:groups=metal3.io,resources=provisionings,verbs=get
//+kubebuilder:rbac:groups=migration.k8s.io,resources=storageversionmigrations,verbs=create;get;list;update;watch;patch;delete
//+kubebuilder:rbac:groups=monitoring.coreos.com,resources=prometheusrules;servicemonitors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=monitoring.coreos.com,resources=servicemonitors,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=monitoring.coreos.com,resources=servicemonitors,verbs=get;create
//+kubebuilder:rbac:groups=multicluster.openshift.io,resources=multiclusterengines,verbs=list;watch
//+kubebuilder:rbac:groups=operator.open-cluster-management.io,resources=clustermanagers,verbs=get;list;watch;update;delete
//+kubebuilder:rbac:groups=operator.open-cluster-management.io,resources=clustermanagers/status,verbs=update;patch
//+kubebuilder:rbac:groups=operator.open-cluster-management.io,resources=klusterlets,verbs=create;delete;deletecollection;get;list;patch;update;watch;escalate
//+kubebuilder:rbac:groups=operator.open-cluster-management.io,resources=multiclusterhubs,verbs=get;list;watch
//+kubebuilder:rbac:groups=operators.coreos.com,resources=clusterserviceversions,verbs=get;list
//+kubebuilder:rbac:groups=operators.coreos.com,resources=subscriptions,verbs=get;list;watch
//+kubebuilder:rbac:groups=policy.open-cluster-management.io;app.k8s.io;apps.open-cluster-management.io;argoproj.io,resources=applications;applicationsets;appprojects;argocds;channels;gitopsclusters;helmreleases;placementrules;placementbindings;policies;policyautomations;policysets;subscriptions;subscriptionreports,verbs=list;watch
//+kubebuilder:rbac:groups=proxy.open-cluster-management.io,resources=clusterstatuses/aggregator,verbs=get;create
//+kubebuilder:rbac:groups=proxy.open-cluster-management.io,resources=clusterstatuses/aggregator,verbs=get;create
//+kubebuilder:rbac:groups=proxy.open-cluster-management.io,resources=managedproxyconfigurations;managedproxyconfigurations/status;managedproxyconfigurations/finalizers;managedproxyserviceresolvers;managedproxyserviceresolvers/status;managedproxyserviceresolvers/finalizers,verbs=*
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings;rolebindings,verbs=create;get;list;update;watch;patch;delete
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles/finalizers,verbs=get;update
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles;clusterrolebindings,verbs=create;delete;get;list;patch;update;watch;bind;escalate
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles;clusterrolebindings,verbs=get;list;watch;create;update;delete;escalate;bind
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles;clusterrolebindings;roles;rolebindings,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles;roles,verbs=create;get;list;update;watch;patch;delete;escalate;bind
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles;rolebindings,verbs=create;delete;get;list;patch;update;watch;escalate;bind
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles;rolebindings,verbs=get;list;watch;create;update;delete
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles;rolebindings,verbs=get;list;watch;create;update;delete
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles;rolebindings,verbs=get;list;watch;create;update;patch
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles;rolebindings,verbs=get;list;watch;create;update;patch
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io;"",resources=roles;rolebindings,verbs=create;get
//+kubebuilder:rbac:groups=register.open-cluster-management.io,resources=managedclusters/accept,verbs=update
//+kubebuilder:rbac:groups=route.openshift.io,resources=routes,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=route.openshift.io,resources=routes,verbs=get;list;create;delete;update;patch
//+kubebuilder:rbac:groups=route.openshift.io,resources=routes/custom-host,verbs=create;delete;get;list;patch;update;watch
//+kubebuilder:rbac:groups=route.openshift.io,resources=routes;routes/custom-host,verbs=create;get
//+kubebuilder:rbac:groups=storage.k8s.io,resources=storageclasses,verbs=list;watch
//+kubebuilder:rbac:groups=submarineraddon.open-cluster-management.io,resources=submarinerconfigs,verbs=list;watch
//+kubebuilder:rbac:groups=tower.ansible.com;batch;"",resources=ansiblejobs;jobs;secrets;serviceaccounts,verbs=create
//+kubebuilder:rbac:groups=velero.io,resources=backups,verbs=create
//+kubebuilder:rbac:groups=view.open-cluster-management.io,resources=managedclusterviews,verbs=get;create;update;delete
//+kubebuilder:rbac:groups=view.open-cluster-management.io,resources=managedclusterviews,verbs=get;list;watch
//+kubebuilder:rbac:groups=view.open-cluster-management.io,resources=managedclusterviews/status,verbs=update;patch
//+kubebuilder:rbac:groups=view.open-cluster-management.io,resources=managedclusterviews;managedclusterviews/status,verbs=get;list;watch;create;update;patch
//+kubebuilder:rbac:groups=wgpolicyk8s.io,resources=policyreports,verbs=list;watch
//+kubebuilder:rbac:groups=work.open-cluster-management.io,resources=manifestworks,verbs=create;update;get;list;watch;delete;deletecollection;patch
//+kubebuilder:rbac:groups=work.open-cluster-management.io,resources=manifestworks,verbs=create;update;get;list;watch;delete;deletecollection;patch
//+kubebuilder:rbac:groups=work.open-cluster-management.io,resources=manifestworks,verbs=get;list;watch
//+kubebuilder:rbac:groups=work.open-cluster-management.io,resources=manifestworks,verbs=get;list;watch;create;update;patch
//+kubebuilder:rbac:groups=work.open-cluster-management.io,resources=manifestworks,verbs=get;list;watch;create;update;patch
//+kubebuilder:rbac:groups=work.open-cluster-management.io,resources=manifestworks,verbs=get;list;watch;update
//+kubebuilder:rbac:groups=work.open-cluster-management.io,resources=manifestworks;manifestworks/finalizers,verbs=create;delete;get;list;patch;update;watch
