package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/kipitix/picture_spawn/internal/application/pipelines"
	"github.com/kipitix/picture_spawn/internal/infrastructure/etcdrepo"
	"github.com/kipitix/picture_spawn/internal/infrastructure/sourceparser"
	"github.com/rs/zerolog/log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Application which parse given site (or another source) and puts data into storage

func main() {

	parser := sourceparser.NewWallpapersWideSourceParser()

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		log.Fatal().Err(err)
	}

	repo := etcdrepo.NewPictureInfoRepoEtcd(etcdClient)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	pipelines.ParsePictureInfoFromSourceAndStoreInRepo(ctx, parser, repo)
}
