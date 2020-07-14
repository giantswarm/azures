package label

// Cluster is the ID label put into all CRs to identify which Tenant Cluster the
// given CR is related to.
const Cluster = "giantswarm.io/cluster"

// ControlPlane is the ID label put into all Control Plane CRs to identify which
// Master Nodes the given CR groups together.
const ControlPlane = "giantswarm.io/control-plane"

// MachineDeployment is the ID label put into all Machine Deployment CRs to
// identify which Worker Nodes the given CR groups together.
const MachineDeployment = "giantswarm.io/machine-deployment"

// Organization is the ID label put into all CRs to identify which Organization
// the given CR is related to.
const Organization = "giantswarm.io/organization"