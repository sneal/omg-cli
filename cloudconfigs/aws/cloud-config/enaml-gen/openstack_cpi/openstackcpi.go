package openstack_cpi 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type OpenstackCpi struct {

	/*Blobstore - Descr: Username agent uses to connect to blobstore used by 'simple' blobstore plugin Default: <nil>
*/
	Blobstore *Blobstore `yaml:"blobstore,omitempty"`

	/*Env - Descr: No proxy environment variable Default: <nil>
*/
	Env *Env `yaml:"env,omitempty"`

	/*Openstack - Descr: OpenStack region (optional) Default: <nil>
*/
	Openstack *Openstack `yaml:"openstack,omitempty"`

	/*Registry - Descr: Full URL for the registry endpoint that may include basic auth credentials Default: <nil>
*/
	Registry *Registry `yaml:"registry,omitempty"`

	/*Agent - Descr: Message bus endpoint for the agent to start accepting agent requests Default: <nil>
*/
	Agent *Agent `yaml:"agent,omitempty"`

	/*Nats - Descr: NATS address used by agent to subscribe to agent requests Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*Ntp - Descr: List of NTP servers Default: []
*/
	Ntp interface{} `yaml:"ntp,omitempty"`

}