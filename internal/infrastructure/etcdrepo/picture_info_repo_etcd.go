package etcdrepo

import (
	"context"
	"encoding/json"
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

func (p *PictureInfoRepoEtcd) StorePictureInfo(ctx context.Context, picInfo pictureinfo.PictureInfo) error {

	pij := pictureinfo.NewPictureInfoJson(picInfo)
	data, _ := json.Marshal(pij)

	p.client.Put(ctx, picInfo.URL(), string(data))

	log.Debug().
		Str("name", picInfo.Name()).
		Str("URL", picInfo.URL()).
		Str("tags", strings.Join(picInfo.Tags(), ",")).
		Str("resolution", picInfo.Resolution()).
		Send()

	return nil
}

func (p *PictureInfoRepoEtcd) SearchPictureInfo(ctx context.Context, picInfoRequest pictureinfo.PictureInfo) ([]pictureinfo.PictureInfo, error) {
	return nil, nil
}
