package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindIgnition = "Ignition"
)

func NewIgnitionTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindIgnition,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=common;giantswarm
// +k8s:openapi-gen=true

// Ignition is a Kubernetes resource (CR) which is based on the Ignition CRD defined above.
//
// An example Ignition resource can be viewed here
// https://github.com/giantswarm/apiextensions/blob/master/docs/cr/core.giantswarm.io_v1alpha1_ignition.yaml
type Ignition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              IgnitionSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status IgnitionStatus `json:"status"`
}

// IgnitionSpec is the interface which defines the input parameters for
// a newly rendered g8s ignition template.
// +k8s:openapi-gen=true
type IgnitionSpec struct {
	// APIServerEncryptionKey is used in EncryptionConfiguration to encrypt Kubernetes secrets at rest.
	APIServerEncryptionKey string `json:"apiServerEncryptionKey"`
	// BaseDomain is the base domain for all cluster services.
	// For test installations, this may be in the form
	// <clusterId>.k8s.<installation>.<region>.<provider>.gigantic.io.
	BaseDomain string `json:"baseDomain"`
	// Calico provides configuration for all calico-related services.
	Calico IgnitionSpecCalico `json:"calico"`
	// ClusterID is the name of the workload cluster to be created.
	ClusterID string `json:"clusterID"`
	// DisableEncryptionAtRest will disable secret encryption at rest when set to true.
	DisableEncryptionAtRest bool `json:"disableEncryptionAtRest"`
	// Docker provides configuration for all calico-related services.
	Docker IgnitionSpecDocker `json:"docker"`
	// Etcd provides configuration for all etcd-related services.
	Etcd IgnitionSpecEtcd `json:"etcd"`
	// Extension can be used to extend an ignition with extra configuration provided by the provider operator.
	Extension IgnitionSpecExtension `json:"extension"`
	// Ingress provides configuration for all ingress-related services.
	Ingress IgnitionSpecIngress `json:"ingress"`
	// IsMaster determines if the rendered ignition should contain master-specific configuration.
	IsMaster bool `json:"isMaster"`
	// Kubernetes provides configuration for all Kubernetes-related services.
	Kubernetes IgnitionSpecKubernetes `json:"kubernetes"`
	// Defines the provider which should be rendered.
	Provider string `json:"provider"`
	// Registry provides configuration for the docker registry used for core component images.
	Registry IgnitionSpecRegistry `json:"registry"`
	// SSO provides configuration for all SSO-related services.
	SSO IgnitionSpecSSO `json:"sso"`
}

// +k8s:openapi-gen=true
type IgnitionSpecCalico struct {
	// CIDR is the CIDR-component of the IPv4 overlay subnetwork. Combined with Subnet below.
	CIDR string `json:"cidr"`
	// Disable can be set to true to disable Calico setup.
	Disable bool `json:"disable"`
	// MTU is the maximum size of packets sent over Calico in bytes.
	MTU string `json:"mtu"`
	// Subnet is the IP-component of the IPv4 overlay subnetwork. Combined with CIDR above.
	Subnet string `json:"subnet"`
}

// +k8s:openapi-gen=true
type IgnitionSpecDocker struct {
	// Daemon provides information about the Docker daemon running on TC nodes.
	Daemon IgnitionSpecDockerDaemon `json:"daemon"`
	// NetworkSetup provides the Docker image to be used for network environment setup.
	NetworkSetup IgnitionSpecDockerNetworkSetup `json:"networkSetup"`
}

// +k8s:openapi-gen=true
type IgnitionSpecDockerDaemon struct {
	// CIDR is the fully specified subnet used for DOCKER_OPT_BIP.
	CIDR string `json:"cidr"`
}

// +k8s:openapi-gen=true
type IgnitionSpecDockerNetworkSetup struct {
	// Image provides the Docker image to be used for network environment setup.
	Image string `json:"image"`
}

// +k8s:openapi-gen=true
type IgnitionSpecEtcd struct {
	// Domain is the domain of the etcd service.
	Domain string `json:"domain"`
	// Port is the port of the etcd service, usually 2379.
	Port int `json:"port"`
	// Prefix is the prefix to add to all etcd keys created by Kubernetes.
	Prefix string `json:"prefix"`
}

// +k8s:openapi-gen=true
type IgnitionSpecExtension struct {
	// Files is an optional array of files which will be rendered and added to the final node ignition.
	Files []IgnitionSpecExtensionFile `json:"files,omitempty"`
	// Files is an optional array of systemd units which will be rendered and added to the final node ignition.
	Units []IgnitionSpecExtensionUnit `json:"units,omitempty"`
	// Files is an optional array of users which will be added to the final node ignition.
	Users []IgnitionSpecExtensionUser `json:"users,omitempty"`
}

// +k8s:openapi-gen=true
type IgnitionSpecExtensionFile struct {
	// Content is the string containing a file with optional go-template-style replacements.
	Content string `json:"content"`
	// Metadata is the filesystem metadata of the given file.
	Metadata IgnitionSpecExtensionFileMetadata `json:"metadata"`
}

// +k8s:openapi-gen=true
type IgnitionSpecExtensionFileMetadata struct {
	// Compression allows a file to be passed in as a base64-encoded compressed string.
	Compression bool `json:"compression"`
	// Owner is the owner of the file.
	Owner IgnitionSpecExtensionFileMetadataOwner `json:"owner"`
	// Path is the path of the file.
	Path string `json:"path"`
	// Permissions is the numeric permissions applied to the file.
	Permissions int `json:"permissions"`
}

// +k8s:openapi-gen=true
type IgnitionSpecExtensionFileMetadataOwner struct {
	// Group is the group which owns the file.
	Group IgnitionSpecExtensionFileMetadataOwnerGroup `json:"group"`
	// User is the user which owns the file.
	User IgnitionSpecExtensionFileMetadataOwnerUser `json:"user"`
}

// +k8s:openapi-gen=true
type IgnitionSpecExtensionFileMetadataOwnerUser struct {
	// ID is the UID of the user.
	ID string `json:"id"`
	// Name is the name of the user.
	Name string `json:"name"`
}

// +k8s:openapi-gen=true
type IgnitionSpecExtensionFileMetadataOwnerGroup struct {
	// ID is the GID of the group.
	ID string `json:"id"`
	// Name is the name of the group.
	Name string `json:"name"`
}

// +k8s:openapi-gen=true
type IgnitionSpecExtensionUnit struct {
	// Content is the string containing a systemd unit with optional go-template-style replacements.
	Content string `json:"content"`
	// Metadata is the filesystem metadata of the given file.
	Metadata IgnitionSpecExtensionUnitMetadata `json:"metadata"`
}

// +k8s:openapi-gen=true
type IgnitionSpecExtensionUnitMetadata struct {
	// Enabled indicates that the unit should be enabled by default.
	Enabled bool `json:"enabled"`
	// Name is the name of the unit on the filesystem and used in systemctl commands.
	Name string `json:"name"`
}

// +k8s:openapi-gen=true
type IgnitionSpecExtensionUser struct {
	// Name is the name of the user to be added to the node via ignition.
	Name string `json:"name"`
	// PublicKey is the public key of the user to be added to the node via ignition.
	PublicKey string `json:"publicKey"`
}

// +k8s:openapi-gen=true
type IgnitionSpecIngress struct {
	// Disable will disable the ingress controller in the TC when true.
	Disable bool `json:"disable"`
}

// +k8s:openapi-gen=true
type IgnitionSpecKubernetes struct {
	// API holds information about the desired TC Kubernetes API.
	API IgnitionSpecKubernetesAPI `json:"api"`
	// CloudProvider is the provider upon which the cluster is running. It is passed to API server as a flag.
	CloudProvider string `json:"cloudProvider"`
	// DNS hold information about the in-cluster DNS service.
	DNS IgnitionSpecKubernetesDNS `json:"dns"`
	// Domain is the domain used for services running in the cluster. Usually this is "cluster.local".
	Domain string `json:"domain"`
	// Kubelet holds information about the kubelet running on nodes.
	Kubelet IgnitionSpecKubernetesKubelet `json:"kubelet"`
	// IPRange is the range of IPs used for pod networking.
	IPRange string `json:"ipRange"`
	// OIDC hold configuration which will be applied to the apiserver OIDC flags.
	OIDC IgnitionSpecOIDC `json:"oidc"`
}

// +k8s:openapi-gen=true
type IgnitionSpecKubernetesAPI struct {
	// Domain is the domain of the API server.
	Domain string `json:"domain"`
	// Secure port is the port on which the API will listen.
	SecurePort int `json:"securePort"`
}

// +k8s:openapi-gen=true
type IgnitionSpecKubernetesDNS struct {
	// IP is the IP of the in-cluster DNS service. Usually this is
	// the same as the API server IP with the final component replaced with .10.
	IP string `json:"ip"`
}

// +k8s:openapi-gen=true
type IgnitionSpecKubernetesKubelet struct {
	// Domain is the domain of the network.
	Domain string `json:"domain"`
}

// +k8s:openapi-gen=true
type IgnitionSpecRegistry struct {
	// Domain is the domain of the registry to be used for pulling core component images.
	Domain string `json:"domain"`
	// Pull progress deadline is a string representing a duration to be used as a deadline
	// for pulling images.
	PullProgressDeadline string `json:"pullProgressDeadline"`
}

// +k8s:openapi-gen=true
type IgnitionSpecSSO struct {
	// PublicKey is the public key of the SSO service.
	PublicKey string `json:"publicKey"`
}

// +k8s:openapi-gen=true
type IgnitionSpecOIDC struct {
	// Enabled indicates that the OIDC settings should be applied when true.
	Enabled bool `json:"enabled"`
	// The client ID for the OpenID Connect client, must be set if IssuerURL is set.
	ClientID string `json:"clientID"`
	// The URL of the OpenID issuer, only HTTPS scheme will be accepted.
	// If set, it will be used to verify the OIDC JSON Web Token (JWT).
	IssuerURL string `json:"issuerUrl"`
	// The OpenID claim to use as the user name. Note that claims other
	// than the default ('sub') is not guaranteed to be unique and immutable.
	UsernameClaim string `json:"usernameClaim"`
	// If provided, all usernames will be prefixed with this value. If not provided, username
	// claims other than 'email' are prefixed by the issuer URL to avoid clashes. To skip any
	// prefixing, provide the value '-'.
	UsernamePrefix string `json:"usernamePrefix"`
	// If provided, the name of a custom OpenID Connect claim for specifying
	// user groups. The claim value is expected to be a string or JSON encoded array of strings.
	GroupsClaim string `json:"groupsClaim"`
	// If provided, all groups will be prefixed with this value to prevent conflicts with other
	// authentication strategies.
	GroupsPrefix string `json:"groupsPrefix"`
}

// IgnitionStatus holds the rendering result.
// +k8s:openapi-gen=true
type IgnitionStatus struct {
	// DataSecret is a reference to the secret containing the rendered ignition once created.
	DataSecret IgnitionStatusSecret `json:"dataSecretName"`
	// FailureReason is a short string indicating the reason rendering failed (if it did).
	FailureReason string `json:"failureReason"`
	// FailureMessage is a longer message indicating the reason rendering failed (if it did).
	FailureMessage string `json:"failureMessage"`
	// Ready will be true when the referenced secret contains the rendered ignition and can be used for creating nodes.
	Ready bool `json:"ready"`
	// Verification is a hash of the rendered ignition to ensure that it has
	// not been changed when loaded as a remote file by the bootstrap ignition.
	// See https://coreos.com/ignition/docs/latest/configuration-v2_2.html
	Verification IgnitionStatusVerification `json:"verification"`
}

// +k8s:openapi-gen=true
type IgnitionStatusVerification struct {
	// The content of the full rendered ignition hashed by the corresponding algorithm.
	Hash string `json:"hash"`
	// The algorithm used for hashing. Must be sha512 for now.
	Algorithm string `json:"algorithm"`
}

// +k8s:openapi-gen=true
type IgnitionStatusSecret struct {
	// Name is the name of the secret containing the rendered ignition.
	Name string `json:"name"`
	// Namespace is the namespace of the secret containing the rendered ignition.
	Namespace string `json:"namespace"`
	// ResourceVersion is the Kubernetes resource version of the secret.
	// Used to detect if the secret has changed, e.g. 12345.
	ResourceVersion string `json:"resourceVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IgnitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Ignition `json:"items"`
}
