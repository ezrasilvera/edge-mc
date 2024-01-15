# User Guide

## KubeStellar Primer

This is a description of the concepts behind KubeStellar.

## Glossary

**Downsynced Object** - One of two categories of workload object, complementary to "upsynced object"**.  In KubeStellar, a downsynced object first appears in a Workload Description Space and the object's desired state propagates from there through Mailbox Spaces to Workload Execution Clusters and that object's reported state originates in the Workload Execution Clusters and propagates back to the Mailbox Spaces and in the future will be summarized into the Workload Description Space.  
****Note:** Upsync is currently not yet supported and may be added to the plugable transport in the future.

**EdgePlacement** - `TBD`

**Inventory Space (IS)** - `TBD` 

**Location** - `TBD`

**Mailbox Namespace** - `TBD`

**Control Plane** - `TBD`

**Status Summarizer** - A planned central KubeStellar controller that will maintain the status summary objects in the Workload Description Spaces 

**SyncTarget** - `TBD`

**Workload Description Space (WDS)** - Holds workload objects and the adjacent KubeStellar control objects, which are `TBD`

**Workload Execution Cluster** - A Kubernetes cluster which can execute a workload. In the examples on this website, we use [Kind](https://kind.sigs.k8s.io/) clusters.

