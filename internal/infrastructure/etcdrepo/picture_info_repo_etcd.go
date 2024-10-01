package etcdrepo

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kipitix/picture_spawn/internal/domain/pictureinfo"
	"github.com/rs/zerolog/log"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type PictureInfoRepoEtcd struct {
	client *clientv3.Client
}

func NewPictureInfoRepoEtcd(client *clientv3.Client) *PictureInfoRepoEtcd {
	return &PictureInfoRepoEtcd{
		client: client,
	}
}

var _ pictureinfo.PictureInfoRepo = (*PictureInfoRepoEtcd)(nil)

func (r *PictureInfoRepoEtcd) Put(ctx context.Context, picInfo pictureinfo.PictureInfo) error {
	const errT = "failed to store picture info in repo: %w"

	pij := pictureinfo.NewPictureInfoJson(picInfo)
	data, _ := json.Marshal(pij)

	// Put data
	if _, err := r.client.Put(ctx, picInfo.ID(), string(data)); err != nil {
		return fmt.Errorf(errT, err)
	}

	log.Debug().
		Str("ID", picInfo.ID()).
		Str("name", picInfo.Name()).
		Str("URL", picInfo.URL()).
		Str("tags", strings.Join(picInfo.Tags(), ",")).
		Str("resolution", picInfo.Resolution()).
		Send()

	return nil
}

func (r *PictureInfoRepoEtcd) GetRandom(ctx context.Context) (pictureinfo.PictureInfo, error) {
	const errT = "failed to get random picture info from repo: %w"

	resp, err := r.client.Get(ctx, "", clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf(errT, err)
	}

	for i, k := range resp.Kvs {

		log.Debug().Any("index", i).
			Any("key", string(k.Key)).
			Send()

		// pij := pictureinfo.NewPictureInfoJson(nil)
		// if err := json.Unmarshal(v.Value, &pij); err != nil {
		// 	return nil, fmt.Errorf(errT, err)
		// }

		// return pij, nil
	}

	return nil, nil
}

// func (p *PictureInfoRepoEtcd) SearchPictureInfo(ctx context.Context, picInfoRequest pictureinfo.PictureInfo) ([]pictureinfo.PictureInfo, error) {
// 	// p.client.Get()

// 	return nil, nil
// }
