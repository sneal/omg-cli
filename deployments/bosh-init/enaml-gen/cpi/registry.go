package cpi 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Registry struct {

	/*Username - Descr: User to access the Registry Default: <nil>
*/
	Username interface{} `yaml:"username,omitempty"`

	/*Password - Descr: Password to access the Registry Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*Endpoint - Descr: Full URL for the registry endpoint that may include basic auth credentials Default: <nil>
*/
	Endpoint interface{} `yaml:"endpoint,omitempty"`

	/*Port - Descr: Port of the Registry to connect to Default: 25777
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Host - Descr: Address of the Registry to connect to Default: <nil>
*/
	Host interface{} `yaml:"host,omitempty"`

}