package director 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Agent struct {

	/*Blobstore - Descr: AWS credentials (static / env_or_profile) Default: static
*/
	Blobstore *Blobstore `yaml:"blobstore,omitempty"`

	/*Nats - Descr: Address for agent to connect to nats Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*Password - Descr: Password agent uses to connect to blobstore used by simple blobstore plugin Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*User - Descr: Username agent uses to connect to blobstore used by simple blobstore plugin Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

}