package mongo

import (
	"time"
)

// Profile - is the memory representation of one user profile
type Profile struct {
	Name 			string			`json: "username"`
	Password 		string			`json: "password"`
	Age 			int				`json: "age"`
	LastUpdated     time.Time
}

var profiles []Profile

func init() {
	profiles = make([]Profile, 3)

	profiles[0] = Profile{Name: "john@aol.com", Password: "test1234", Age: 23, LastUpdated: time.Now()}
	profiles[1] = Profile{Name: "harry@comcast.com", Password: "test1234", Age: 44, LastUpdated: time.Now()}
	profiles[2] = Profile{Name: "sally@microsoft.com", Password: "test1234", Age: 63, LastUpdated: time.Now()}
}

func GetProfiles() []Profile {
	return profiles
}
