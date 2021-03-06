package vsphere_cpi 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Vcenter struct {

	/*ConnectionOptions - Descr: All required custom CA certificates Default: <nil>
*/
	ConnectionOptions *ConnectionOptions `yaml:"connection_options,omitempty"`

	/*Address - Descr: Address of vCenter server used by vsphere cpi Default: <nil>
*/
	Address interface{} `yaml:"address,omitempty"`

	/*Datacenters - Descr: Datacenters in vCenter to use (value is an array of Hashes representing datacenters and clusters, See director.yml.erb.erb) Default: <nil>
*/
	Datacenters interface{} `yaml:"datacenters,omitempty"`

	/*Password - Descr: Password to connect to vCenter server used by vspher cpi Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*User - Descr: User to connect to vCenter server used by vsphere cpi Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

}