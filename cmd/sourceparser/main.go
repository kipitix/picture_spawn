// Application which parse given site (or another source) and puts data into storage

package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/kipitix/picture_spawn/internal/application/pipelines"
	"github.com/kipitix/picture_spawn/internal/infrastructure/etcdrepo"
	"github.com/kipitix/picture_spawn/internal/infrastructure/sourceparser"
	"github.com/kipitix/picture_spawn/internal/tools/arguments"
	"github.com/kipitix/picture_spawn/internal/tools/logger"
	"github.com/rs/zerolog/log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var Version = "0.0.1"

var args struct {
	EtcdEndpoints string `arg:"-e,--etcd-endpoints" default:"http://localhost:2379" help:"Endpoints to connect to etcd (comma separated)"`
}

func main() {
	logger.SetupZerolog()

	log.Info().Msgf("sourceparser v%s", Version)

	log.Info().Msg("parsing args...")

	arg.MustParse(&args)

	log.Debug().Interface("args", args).Send()

	log.Info().Msgf("connection to etcd...")

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   arguments.ParseEtcdEndpoints(args.EtcdEndpoints),
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		log.Fatal().Err(err)
	}

	repo := etcdrepo.NewImageRepoEtcd(etcdClient)

	// repo2 := imginfo.NewImageRepoLog()

	parser := sourceparser.NewWallpapersWideSourceParser()

	pipeline := pipelines.NewParseImagesFromSourceAndPutInRepo(parser, repo, 5*time.Second)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	pipeline.Do(ctx)
}
