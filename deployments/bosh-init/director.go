package boshinit

import "github.com/enaml-ops/omg-cli/deployments/bosh-init/enaml-gen/director"

func NewDirectorProperty(name, cpijob string, db *director.Db) DirectorProperty {
	return DirectorProperty{
		Address: "127.0.0.1",
		Director: director.Director{
			Name:       name,
			CpiJob:     cpijob,
			MaxThreads: 10,
			Db:         db,
			UserManagement: &director.UserManagement{
				Provider: "local",
				Local: &director.Local{
					Users: []user{
						user{Name: "admin", Password: "admin"},
						user{Name: "hm", Password: "hm-password"},
					},
				},
			},
		},
	}
}
