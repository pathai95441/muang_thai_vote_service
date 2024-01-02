package vote_config_domain

import (
	"context"

	"github.com/pathai95441/muang_thai_vote_service/src/redis_cache"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/vote_config"
)

const (
	VOTE_CLOSED_CONFIG = "VOTE_CLOSED_CONFIG"
)

//go:generate mockgen -source=./vote_config_domain.go -destination=./mock/mock_vote_config_domain.go -package=mock_vote_config_domain
type IVoteConfigDomain interface {
	GetClosedConfig(ctx context.Context) (bool, error)
}

type VoteConfigDomain struct {
	voteConfigRepo vote_config.IRepository
	cache          redis_cache.IRedisCache
}

func NewVoteConfigDomain(voteConfigRepo vote_config.IRepository, cache redis_cache.IRedisCache) VoteConfigDomain {
	return VoteConfigDomain{voteConfigRepo, cache}
}

func (d VoteConfigDomain) GetClosedConfig(ctx context.Context) (bool, error) {
	val, err := d.cache.GetValue(ctx, VOTE_CLOSED_CONFIG)
	if err != nil {
		cacheVal := "0"
		isClosed, err := d.voteConfigRepo.Get(ctx)
		if err != nil {
			return false, err
		}

		if isClosed {
			cacheVal = "1"
		}

		err = d.cache.SetKeyValue(ctx, VOTE_CLOSED_CONFIG, cacheVal)
		if err != nil {
			return false, err
		}

		return isClosed, nil
	}

	return val == "1", nil
}
