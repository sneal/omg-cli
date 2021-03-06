package director 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Nats struct {

	/*Port - Descr: Port that the nats server listens on Default: 4222
*/
	Port interface{} `yaml:"port,omitempty"`

	/*User - Descr: Username to connect to nats with Default: nats
*/
	User interface{} `yaml:"user,omitempty"`

	/*Address - Descr: Address for agent to connect to nats Default: <nil>
*/
	Address interface{} `yaml:"address,omitempty"`

	/*Password - Descr: Password to connect to nats with Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

}