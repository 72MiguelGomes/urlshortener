package model

import "gopkg.in/couchbase/gocb.v1"

var bucket *gocb.Bucket

func InitDB() {

	cluster, _ := gocb.Connect("couchbase://localhost")

	bucket, _ = cluster.OpenBucket("default", "")
}
