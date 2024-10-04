package etcdrepo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kipitix/picture_spawn/internal/domain/imginfo"
	"github.com/rs/zerolog/log"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type ImageRepoEtcd struct {
	client *clientv3.Client
}

func NewImageRepoEtcd(client *clientv3.Client) *ImageRepoEtcd {
	return &ImageRepoEtcd{
		client: client,
	}
}

var _ imginfo.ImageRepo = (*ImageRepoEtcd)(nil)

func (r *ImageRepoEtcd) Put(ctx context.Context, img imginfo.Image) error {
	const errT = "failed to store image info in repo: %w"

	ij := imginfo.NewImageJSON(img)
	data, _ := json.Marshal(ij)

	// Put data
	if _, err := r.client.Put(ctx, img.ID(), string(data)); err != nil {
		return fmt.Errorf(errT, err)
	}

	log.Debug().Interface("image JSON", ij).Send()

	return nil
}

func (r *ImageRepoEtcd) Get(ctx context.Context, reqImg imginfo.Image) (imginfo.Image, error) {
	//const errT = "failed to get random picture info from repo: %w"

	// resp, err := r.client.Get(ctx, "", clientv3.WithPrefix())
	// if err != nil {
	// 	return nil, fmt.Errorf(errT, err)
	// }

	// for i, k := range resp.Kvs {

	// 	log.Debug().Any("index", i).
	// 		Any("key", string(k.Key)).
	// 		Send()

	// 	// pij := imginfo.NewPictureInfoJson(nil)
	// 	// if err := json.Unmarshal(v.Value, &pij); err != nil {
	// 	// 	return nil, fmt.Errorf(errT, err)
	// 	// }

	// 	// return pij, nil
	// }

	return nil, nil
}

// func (p *PictureInfoRepoEtcd) SearchPictureInfo(ctx context.Context, imginfoRequest imginfo.PictureInfo) ([]imginfo.PictureInfo, error) {
// 	// p.client.Get()

// 	return nil, nil
// }
