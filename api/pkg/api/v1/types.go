package v1

import (
	"encoding/json"

	"github.com/Masterminds/semver"
	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	ksemver "github.com/kubermatic/kubermatic/api/pkg/semver"

	cmdv1 "k8s.io/client-go/tools/clientcmd/api/v1"
	"sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

// LegacyObjectMeta is an object storing common metadata for persistable objects.
// Deprecated: LegacyObjectMeta is deprecated use ObjectMeta instead.
type LegacyObjectMeta struct {
	Name            string `json:"name"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
	UID             string `json:"uid,omitempty"`

	Annotations map[string]string `json:"annotations,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

// ObjectMeta defines the set of fields that objects returned from the API have
// swagger:model ObjectMeta
type ObjectMeta struct {
	// ID unique value that identifies the resource generated by the server
	ID string `json:"id,omitempty"`

	// Name represents human readable name for the resource
	Name string `json:"name"`

	// DeletionTimestamp is a timestamp representing the server time when this object was deleted.
	// swagger:strfmt date-time
	DeletionTimestamp *Time `json:"deletionTimestamp,omitempty"`

	// CreationTimestamp is a timestamp representing the server time when this object was created.
	// swagger:strfmt date-time
	CreationTimestamp Time `json:"creationTimestamp,omitempty"`
}

// DigitialoceanDatacenterSpec specifies a datacenter of DigitalOcean.
type DigitialoceanDatacenterSpec struct {
	Region string `json:"region"`
}

// HetznerDatacenterSpec specifies a datacenter of Hetzner.
type HetznerDatacenterSpec struct {
	Datacenter string `json:"datacenter"`
	Location   string `json:"location"`
}

// ImageList defines a map of operating system and the image to use
type ImageList map[string]string

// VSphereDatacenterSpec specifies a datacenter of VSphere.
type VSphereDatacenterSpec struct {
	Endpoint   string    `json:"endpoint"`
	Datacenter string    `json:"datacenter"`
	Datastore  string    `json:"datastore"`
	Cluster    string    `json:"cluster"`
	Templates  ImageList `json:"templates"`
}

// BringYourOwnDatacenterSpec specifies a data center with bring-your-own nodes.
type BringYourOwnDatacenterSpec struct{}

// AWSDatacenterSpec specifies a data center of Amazon Web Services.
type AWSDatacenterSpec struct {
	Region string `json:"region"`
}

// AzureDatacenterSpec specifies a datacenter of Azure.
type AzureDatacenterSpec struct {
	Location string `json:"location"`
}

// OpenstackDatacenterSpec specifies a generic bare metal datacenter.
type OpenstackDatacenterSpec struct {
	AvailabilityZone  string    `json:"availability_zone"`
	Region            string    `json:"region"`
	AuthURL           string    `json:"auth_url"`
	Images            ImageList `json:"images"`
	EnforceFloatingIP bool      `json:"enforce_floating_ip"`
}

// DatacenterSpec specifies the data for a datacenter.
type DatacenterSpec struct {
	Seed         string                       `json:"seed"`
	Country      string                       `json:"country,omitempty"`
	Location     string                       `json:"location,omitempty"`
	Provider     string                       `json:"provider,omitempty"`
	Digitalocean *DigitialoceanDatacenterSpec `json:"digitalocean,omitempty"`
	BringYourOwn *BringYourOwnDatacenterSpec  `json:"bringyourown,omitempty"`
	AWS          *AWSDatacenterSpec           `json:"aws,omitempty"`
	Azure        *AzureDatacenterSpec         `json:"azure,omitempty"`
	Openstack    *OpenstackDatacenterSpec     `json:"openstack,omitempty"`
	Hetzner      *HetznerDatacenterSpec       `json:"hetzner,omitempty"`
	VSphere      *VSphereDatacenterSpec       `json:"vsphere,omitempty"`
}

// DatacenterList represents a list of datacenters
// swagger:model DatacenterList
type DatacenterList []Datacenter

// Datacenter is the object representing a Kubernetes infra datacenter.
// swagger:model Datacenter
type Datacenter struct {
	Metadata LegacyObjectMeta `json:"metadata"`
	Spec     DatacenterSpec   `json:"spec"`
	Seed     bool             `json:"seed,omitempty"`
}

// DigitaloceanSizeList represents a object of digitalocean sizes.
// swagger:model DigitaloceanSizeList
type DigitaloceanSizeList struct {
	Standard  []DigitaloceanSize `json:"standard"`
	Optimized []DigitaloceanSize `json:"optimized"`
}

// DigitaloceanSize is the object representing digitalocean sizes.
// swagger:model DigitaloceanSize
type DigitaloceanSize struct {
	Slug         string   `json:"slug"`
	Available    bool     `json:"available"`
	Transfer     float64  `json:"transfer"`
	PriceMonthly float64  `json:"price_monthly"`
	PriceHourly  float64  `json:"price_hourly"`
	Memory       int      `json:"memory"`
	VCPUs        int      `json:"vcpus"`
	Disk         int      `json:"disk"`
	Regions      []string `json:"regions"`
}

// AzureSizeList represents an array of Azure VM sizes.
// swagger:model AzureSizeList
type AzureSizeList []AzureSize

// AzureSize is the object representing Azure VM sizes.
// swagger:model AzureSize
type AzureSize struct {
	Name                 string `json:"name"`
	NumberOfCores        int32  `json:"numberOfCores"`
	OsDiskSizeInMB       int32  `json:"osDiskSizeInMB"`
	ResourceDiskSizeInMB int32  `json:"resourceDiskSizeInMB"`
	MemoryInMB           int32  `json:"memoryInMB"`
	MaxDataDiskCount     int32  `json:"maxDataDiskCount"`
}

// SSHKey represents a ssh key
// swagger:model SSHKey
type SSHKey struct {
	ObjectMeta
	Spec SSHKeySpec `json:"spec"`
}

// SSHKeySpec represents the details of a ssh key
type SSHKeySpec struct {
	Fingerprint string `json:"fingerprint"`
	PublicKey   string `json:"publicKey"`
}

// User represent an API user
// swagger:model User
type User struct {
	ObjectMeta

	// Email an email address of the user
	Email string `json:"email"`

	// Projects holds the list of project the user belongs to
	// along with the group names
	Projects []ProjectGroup `json:"projects,omitempty"`
}

// ProjectGroup is a helper data structure that
// stores the information about a project and a group prefix that a user belongs to
type ProjectGroup struct {
	ID          string `json:"id"`
	GroupPrefix string `json:"group"`
}

// Project is a top-level container for a set of resources
// swagger:model Project
type Project struct {
	ObjectMeta
	Status string `json:"status"`
	// Owners an optional owners list for the given project
	Owners []User `json:"owners,omitempty"`
}

// Kubeconfig is a clusters kubeconfig
// swagger:model Kubeconfig
type Kubeconfig struct {
	cmdv1.Config
}

// OpenstackSize is the object representing openstack's sizes.
// swagger:model OpenstackSize
type OpenstackSize struct {
	// Slug holds  the name of the size
	Slug string `json:"slug"`
	// Memory is the amount of memory, measured in MB
	Memory int `json:"memory"`
	// VCPUs indicates how many (virtual) CPUs are available for this flavor
	VCPUs int `json:"vcpus"`
	// Disk is the amount of root disk, measured in GB
	Disk int `json:"disk"`
	// Swap is the amount of swap space, measured in MB
	Swap int `json:"swap"`
	// Region specifies the geographic region in which the size resides
	Region string `json:"region"`
	// IsPublic indicates whether the size is public (available to all projects) or scoped to a set of projects
	IsPublic bool `json:"isPublic"`
}

// OpenstackSubnet is the object representing a openstack subnet.
// swagger:model OpenstackSubnet
type OpenstackSubnet struct {
	// Id uniquely identifies the subnet
	ID string `json:"id"`
	// Name is human-readable name for the subnet
	Name string `json:"name"`
}

// OpenstackTenant is the object representing a openstack tenant.
// swagger:model OpenstackTenant
type OpenstackTenant struct {
	// Id uniquely identifies the current tenant
	ID string `json:"id"`
	// Name is the name of the tenant
	Name string `json:"name"`
}

// OpenstackNetwork is the object representing a openstack network.
// swagger:model OpenstackNetwork
type OpenstackNetwork struct {
	// Id uniquely identifies the current network
	ID string `json:"id"`
	// Name is the name of the network
	Name string `json:"name"`
	// External set if network is the external network
	External bool `json:"external"`
}

// OpenstackSecurityGroup is the object representing a openstack security group.
// swagger:model OpenstackSecurityGroup
type OpenstackSecurityGroup struct {
	// Id uniquely identifies the current security group
	ID string `json:"id"`
	// Name is the name of the security group
	Name string `json:"name"`
}

// VSphereNetwork is the object representing a vsphere network.
// swagger:model VSphereNetwork
type VSphereNetwork struct {
	// Name is the name of the network
	Name string `json:"name"`
}

// MasterVersion describes a version of the master components
// swagger:model MasterVersion
type MasterVersion struct {
	Version *semver.Version `json:"version"`
	Default bool            `json:"default,omitempty"`
}

// Cluster defines the cluster resource
//
// Note:
// Cluster has a custom MarshalJSON method defined
// and thus the output may vary
//
// swagger:model Cluster
type Cluster struct {
	ObjectMeta `json:",inline"`
	Spec       ClusterSpec   `json:"spec"`
	Status     ClusterStatus `json:"status"`
}

// ClusterSpec defines the cluster specification
type ClusterSpec struct {
	// Cloud specifies the cloud providers configuration
	Cloud kubermaticv1.CloudSpec `json:"cloud"`

	// MachineNetworks optionally specifies the parameters for IPAM.
	MachineNetworks []kubermaticv1.MachineNetworkingConfig `json:"machineNetworks,omitempty"`

	// Version desired version of the kubernetes master components
	Version ksemver.Semver `json:"version"`
}

// MarshalJSON marshals ClusterSpec object into JSON. It is overwritten to control data
// that will be returned in the API responses (see: PublicCloudSpec struct).
func (cs *ClusterSpec) MarshalJSON() ([]byte, error) {
	ret, err := json.Marshal(struct {
		Cloud           PublicCloudSpec                        `json:"cloud"`
		MachineNetworks []kubermaticv1.MachineNetworkingConfig `json:"machineNetworks,omitempty"`
		Version         ksemver.Semver                         `json:"version"`
	}{
		Cloud: PublicCloudSpec{
			DatacenterName: cs.Cloud.DatacenterName,
			Fake:           newPublicFakeCloudSpec(cs.Cloud.Fake),
			Digitalocean:   newPublicDigitaloceanCloudSpec(cs.Cloud.Digitalocean),
			BringYourOwn:   newPublicBringYourOwnCloudSpec(cs.Cloud.BringYourOwn),
			AWS:            newPublicAWSCloudSpec(cs.Cloud.AWS),
			Azure:          newPublicAzureCloudSpec(cs.Cloud.Azure),
			Openstack:      newPublicOpenstackCloudSpec(cs.Cloud.Openstack),
			Hetzner:        newPublicHetznerCloudSpec(cs.Cloud.Hetzner),
			VSphere:        newPublicVSphereCloudSpec(cs.Cloud.VSphere),
		},
		Version:         cs.Version,
		MachineNetworks: cs.MachineNetworks,
	})

	return ret, err
}

// PublicCloudSpec is a public counterpart of apiv1.CloudSpec.
type PublicCloudSpec struct {
	DatacenterName string                       `json:"dc"`
	Fake           *PublicFakeCloudSpec         `json:"fake,omitempty"`
	Digitalocean   *PublicDigitaloceanCloudSpec `json:"digitalocean,omitempty"`
	BringYourOwn   *PublicBringYourOwnCloudSpec `json:"bringyourown,omitempty"`
	AWS            *PublicAWSCloudSpec          `json:"aws,omitempty"`
	Azure          *PublicAzureCloudSpec        `json:"azure,omitempty"`
	Openstack      *PublicOpenstackCloudSpec    `json:"openstack,omitempty"`
	Hetzner        *PublicHetznerCloudSpec      `json:"hetzner,omitempty"`
	VSphere        *PublicVSphereCloudSpec      `json:"vsphere,omitempty"`
}

// PublicFakeCloudSpec is a public counterpart of apiv1.FakeCloudSpec.
type PublicFakeCloudSpec struct{}

func newPublicFakeCloudSpec(internal *kubermaticv1.FakeCloudSpec) (public *PublicFakeCloudSpec) {
	if internal == nil {
		return nil
	}

	return &PublicFakeCloudSpec{}
}

// PublicDigitaloceanCloudSpec is a public counterpart of apiv1.DigitaloceanCloudSpec.
type PublicDigitaloceanCloudSpec struct{}

func newPublicDigitaloceanCloudSpec(internal *kubermaticv1.DigitaloceanCloudSpec) (public *PublicDigitaloceanCloudSpec) {
	if internal == nil {
		return nil
	}

	return &PublicDigitaloceanCloudSpec{}
}

// PublicHetznerCloudSpec is a public counterpart of apiv1.HetznerCloudSpec.
type PublicHetznerCloudSpec struct{}

func newPublicHetznerCloudSpec(internal *kubermaticv1.HetznerCloudSpec) (public *PublicHetznerCloudSpec) {
	if internal == nil {
		return nil
	}

	return &PublicHetznerCloudSpec{}
}

// PublicAzureCloudSpec is a public counterpart of apiv1.AzureCloudSpec.
type PublicAzureCloudSpec struct{}

func newPublicAzureCloudSpec(internal *kubermaticv1.AzureCloudSpec) (public *PublicAzureCloudSpec) {
	if internal == nil {
		return nil
	}

	return &PublicAzureCloudSpec{}
}

// PublicVSphereCloudSpec is a public counterpart of apiv1.VSphereCloudSpec.
type PublicVSphereCloudSpec struct{}

func newPublicVSphereCloudSpec(internal *kubermaticv1.VSphereCloudSpec) (public *PublicVSphereCloudSpec) {
	if internal == nil {
		return nil
	}

	return &PublicVSphereCloudSpec{}
}

// PublicBringYourOwnCloudSpec is a public counterpart of apiv1.BringYourOwnCloudSpec.
type PublicBringYourOwnCloudSpec struct{}

func newPublicBringYourOwnCloudSpec(internal *kubermaticv1.BringYourOwnCloudSpec) (public *PublicBringYourOwnCloudSpec) {
	if internal == nil {
		return nil
	}

	return &PublicBringYourOwnCloudSpec{}
}

// PublicAWSCloudSpec is a public counterpart of apiv1.AWSCloudSpec.
type PublicAWSCloudSpec struct{}

func newPublicAWSCloudSpec(internal *kubermaticv1.AWSCloudSpec) (public *PublicAWSCloudSpec) {
	if internal == nil {
		return nil
	}

	return &PublicAWSCloudSpec{}
}

// PublicOpenstackCloudSpec is a public counterpart of apiv1.OpenstackCloudSpec.
type PublicOpenstackCloudSpec struct {
	FloatingIPPool string `json:"floatingIpPool"`
}

func newPublicOpenstackCloudSpec(internal *kubermaticv1.OpenstackCloudSpec) (public *PublicOpenstackCloudSpec) {
	if internal == nil {
		return nil
	}

	return &PublicOpenstackCloudSpec{
		FloatingIPPool: internal.FloatingIPPool,
	}
}

// ClusterStatus defines the cluster status
type ClusterStatus struct {
	// Version actual version of the kubernetes master components
	Version ksemver.Semver `json:"version"`

	// URL specifies the address at which the cluster is available
	URL string `json:"url"`
}

// ClusterHealth stores health information about the cluster's components.
// swagger:model ClusterHealth
type ClusterHealth struct {
	Apiserver         bool `json:"apiserver"`
	Scheduler         bool `json:"scheduler"`
	Controller        bool `json:"controller"`
	MachineController bool `json:"machineController"`
	Etcd              bool `json:"etcd"`
}

// ClusterList represents a list of clusters
// swagger:model ClusterList
type ClusterList []Cluster

// Node represents a worker node that is part of a cluster
// swagger:model Node
type Node struct {
	ObjectMeta `json:",inline"`
	Spec       NodeSpec   `json:"spec"`
	Status     NodeStatus `json:"status"`
}

// NodeCloudSpec represents the collection of cloud provider specific settings. Only one must be set at a time.
// swagger:model NodeCloudSpec
type NodeCloudSpec struct {
	Digitalocean *DigitaloceanNodeSpec `json:"digitalocean,omitempty"`
	AWS          *AWSNodeSpec          `json:"aws,omitempty"`
	Azure        *AzureNodeSpec        `json:"azure,omitempty"`
	Openstack    *OpenstackNodeSpec    `json:"openstack,omitempty"`
	Hetzner      *HetznerNodeSpec      `json:"hetzner,omitempty"`
	VSphere      *VSphereNodeSpec      `json:"vsphere,omitempty"`
}

// UbuntuSpec ubuntu specific settings
// swagger:model UbuntuSpec
type UbuntuSpec struct {
	// do a dist-upgrade on boot and reboot it required afterwards
	DistUpgradeOnBoot bool `json:"distUpgradeOnBoot"`
}

// CentOSSpec contains CentOS specific settings
type CentOSSpec struct {
	// do a dist-upgrade on boot and reboot it required afterwards
	DistUpgradeOnBoot bool `json:"distUpgradeOnBoot"`
}

// ContainerLinuxSpec ubuntu linux specific settings
// swagger:model ContainerLinuxSpec
type ContainerLinuxSpec struct {
	// disable container linux auto-update feature
	DisableAutoUpdate bool `json:"disableAutoUpdate"`
}

// OperatingSystemSpec represents the collection of os specific settings. Only one must be set at a time.
// swagger:model OperatingSystemSpec
type OperatingSystemSpec struct {
	Ubuntu         *UbuntuSpec         `json:"ubuntu,omitempty"`
	ContainerLinux *ContainerLinuxSpec `json:"containerLinux,omitempty"`
	CentOS         *CentOSSpec         `json:"centos,omitempty"`
}

// NodeVersionInfo node version information
// swagger:model NodeVersionInfo
type NodeVersionInfo struct {
	Kubelet string `json:"kubelet"`
}

// NodeSpec node specification
// swagger:model NodeSpec
type NodeSpec struct {
	// required: true
	Cloud NodeCloudSpec `json:"cloud"`
	// required: true
	OperatingSystem OperatingSystemSpec `json:"operatingSystem"`
	// required: true
	Versions NodeVersionInfo `json:"versions,omitempty"`
}

// DigitaloceanNodeSpec digitalocean node settings
// swagger:model DigitaloceanNodeSpec
type DigitaloceanNodeSpec struct {
	// droplet size slug
	// required: true
	Size string `json:"size"`
	// enable backups for the droplet
	Backups bool `json:"backups"`
	// enable ipv6 for the droplet
	IPv6 bool `json:"ipv6"`
	// enable monitoring for the droplet
	Monitoring bool `json:"monitoring"`
	// additional droplet tags
	Tags []string `json:"tags"`
}

// HetznerNodeSpec Hetzner node settings
// swagger:model HetznerNodeSpec
type HetznerNodeSpec struct {
	// server type
	// required: true
	Type string `json:"type"`
}

// AzureNodeSpec describes settings for an Azure node
// swagger:model AzureNodeSpec
type AzureNodeSpec struct {
	// VM size
	// required: true
	Size string `json:"size"`
	// should the machine have a publicly accessible IP address
	// required: false
	AssignPublicIP bool `json:"assignPublicIP"`
	// Additional metadata to set
	// required: false
	Tags map[string]string `json:"tags,omitempty"`
}

// VSphereNodeSpec VSphere node settings
// swagger:model VSphereNodeSpec
type VSphereNodeSpec struct {
	CPUs            int    `json:"cpus"`
	Memory          int    `json:"memory"`
	Template        string `json:"template"`
	TemplateNetName string `json:"templateNetName"`
}

// OpenstackNodeSpec openstack node settings
// swagger:model OpenstackNodeSpec
type OpenstackNodeSpec struct {
	// instance flavor
	// required: true
	Flavor string `json:"flavor"`
	// image to use
	// required: true
	Image string `json:"image"`
	// Additional metadata to set
	// required: false
	Tags map[string]string `json:"tags,omitempty"`
	// Defines whether floating ip should be used
	// required: false
	UseFloatingIP bool `json:"useFloatingIP,omitempty"`
}

// AWSNodeSpec aws specific node settings
// swagger:model AWSNodeSpec
type AWSNodeSpec struct {
	// instance type. for example: t2.micro
	// required: true
	InstanceType string `json:"instanceType"`
	// size of the volume in gb. Only one volume will be created
	// required: true
	VolumeSize int64 `json:"diskSize"`
	// type of the volume. for example: gp2, io1, st1, sc1, standard
	// required: true
	VolumeType string `json:"volumeType"`
	// ami to use. Will be defaulted to a ami for your selected operating system and region. Only set this when you know what you do.
	AMI string `json:"ami"`
	// additional instance tags
	Tags map[string]string `json:"tags"`
}

// NodeResources cpu and memory of a node
// swagger:model NodeResources
type NodeResources struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

// NodeStatus is information about the current status of a node.
// swagger:model NodeStatus
type NodeStatus struct {
	// name of the actual Machine object
	MachineName string `json:"machineName"`
	// resources in total
	Capacity NodeResources `json:"capacity,omitempty"`
	// allocatable resources
	Allocatable NodeResources `json:"allocatable,omitempty"`
	// different addresses of a node
	Addresses []NodeAddress `json:"addresses,omitempty"`
	// node versions and systems info
	NodeInfo NodeSystemInfo `json:"nodeInfo,omitempty"`

	// in case of a error this will contain a short error message
	ErrorReason string `json:"errorReason,omitempty"`
	// in case of a error this will contain a detailed error explanation
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// NodeAddress contains information for the node's address.
// swagger:model NodeAddress
type NodeAddress struct {
	// address type. for example: ExternalIP, InternalIP, InternalDNS, ExternalDNS
	Type string `json:"type"`
	// the actual address. for example: 192.168.1.1, node1.my.dns
	Address string `json:"address"`
}

// NodeSystemInfo is a set of versions/ids/uuids to uniquely identify the node.
// swagger:model NodeSystemInfo
type NodeSystemInfo struct {
	KernelVersion           string `json:"kernelVersion"`
	ContainerRuntime        string `json:"containerRuntime"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	OperatingSystem         string `json:"operatingSystem"`
	Architecture            string `json:"architecture"`
}

// ClusterMetric defines a metric for the given cluster
// swagger:model ClusterMetric
type ClusterMetric struct {
	Name   string    `json:"name"`
	Values []float64 `json:"values,omitempty"`
}

// NodeDeployment represents a set of worker nodes that is part of a cluster
// swagger:model NodeDeployment
type NodeDeployment struct {
	ObjectMeta `json:",inline"`

	Spec   NodeDeploymentSpec               `json:"spec"`
	Status v1alpha1.MachineDeploymentStatus `json:"status"`
}

// NodeDeploymentSpec node deployment specification
// swagger:model NodeDeploymentSpec
type NodeDeploymentSpec struct {
	// required: true
	Replicas int32 `json:"replicas,omitempty"`

	// required: true
	Template NodeSpec `json:"template"`

	Paused *bool `json:"paused,omitempty"`
}

// Event is a report of an event somewhere in the cluster.
type Event struct {
	ObjectMeta `json:",inline"`

	// A human-readable description of the status of this operation.
	Message string `json:"message,omitempty"`

	// Type of this event (i.e. normal or warning). New types could be added in the future.
	Type string `json:"type,omitempty"`

	// The object name that those events are about.
	InvolvedObjectName string `json:"involvedObjectName"`

	// The time at which the most recent occurrence of this event was recorded.
	// swagger:strfmt date-time
	LastTimestamp Time `json:"lastTimestamp,omitempty"`

	// The number of times this event has occurred.
	Count int32 `json:"count,omitempty"`
}

// KubermaticVersions describes the versions of running Kubermatic components.
// swagger:model KubermaticVersions
type KubermaticVersions struct {
	// Version of the Kubermatic API server.
	API string `json:"api"`
}
