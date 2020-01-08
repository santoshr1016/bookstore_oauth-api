package cassandra

import (
	"pkg/mod/github.com/gocql/gocql@v0.0.0-20200103014340-68f928edb90a"
	//"github.com/gocql/gocql"
)


var (
	cfg  *gocql.ClusterConfig
)

func init(){
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	//cfg.Port = 9042
	return cfg.CreateSession()
	//session, err := cluster.CreateSession()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Connection created")
	//defer session.Close()
	//return session, nil
}