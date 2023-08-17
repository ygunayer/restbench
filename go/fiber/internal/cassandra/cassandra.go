package cassandra

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/ygunayer/restbench/internal/config"
	"github.com/ygunayer/restbench/internal/logger"
)

var session *gocql.Session

func Open(cfg *config.CassandraConfig) error {
	cluster := gocql.NewCluster(cfg.Hosts...)
	cluster.Keyspace = cfg.Keyspace
	cluster.ProtoVersion = 4
	cluster.Consistency = gocql.Quorum

	sess, err := cluster.CreateSession()
	if err != nil {
		return err
	}

	ctx := context.Background()

	if err := sess.AwaitSchemaAgreement(ctx); err != nil {
		return err
	}

	logger.Infof("Successfully connected to Cassandra cluster at %s", cfg.Hosts)

	session = sess
	return nil
}

func Get() *gocql.Session {
	return session
}
