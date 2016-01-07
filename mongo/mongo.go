package mongo

import (
	"time"
	"log"
	"encoding/base64"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"ProfileService/crypto"
)

// Profile - is the memory representation of one user profile
type Profile struct {
	Name 			string			`json: "username"`
	Password 		string			`json: "password"`
	Age 			int				`json: "age"`
	LastUpdated     time.Time		`json: "last_updated"`
}


func init() {
	tmp, _ := crypto.Encrypt([]byte( "test1234"))
	p1 := Profile{Name: "john@aol.com", Password: base64.StdEncoding.EncodeToString(tmp), Age: 23, LastUpdated: time.Now()}
	p1.CreateOrUpdateProfile()

	tmp, _ = crypto.Encrypt([]byte( "testing1"))
	p1 = Profile{Name: "harry@comcast.com", Password: base64.StdEncoding.EncodeToString(tmp), Age: 44, LastUpdated: time.Now()}
	p1.CreateOrUpdateProfile()

	tmp, _ = crypto.Encrypt([]byte( "pass4321"))
	p1 = Profile{Name: "sally@microsoft.com", Password: base64.StdEncoding.EncodeToString(tmp), Age: 63, LastUpdated: time.Now()}
	p1.CreateOrUpdateProfile()
}

func GetProfiles() []Profile {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("ProfileService").C("Profiles")
	var profiles []Profile
	err = c.Find(bson.M{}).All(&profiles)

	return profiles
}



// CreateOrUpdateProfile - Creates or Updates (Upsert) the profile in the Profiles Collection with id parameter
func (p *Profile) CreateOrUpdateProfile() bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return false
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("ProfileService").C("Profiles")
	_ , err = c.UpsertId( p.Name, p )
	if err != nil {
		log.Println("Error creating Profile: ", err.Error())
		return false
	}
	return true
}