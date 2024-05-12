package grpc

import (
	"context"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/api/grpc/promotion"
	"github.com/minhvuongrbs/financial-service-example/config"
	"github.com/minhvuongrbs/financial-service-example/internal/common/utils"
	"github.com/minhvuongrbs/financial-service-example/internal/promotion/adapters/repository/campaignuser"
	"github.com/minhvuongrbs/financial-service-example/internal/promotion/app/service"
	"google.golang.org/grpc"
)

type PromotionService struct {
	promotion.UnimplementedPromotionServer
	App service.App
}

func NewPromotionService(conf config.Config) (PromotionService, error) {
	// load config
	infraDependencies, err := initInfrastructure(conf)
	if err != nil {
		return PromotionService{}, fmt.Errorf("failed to init infra dependcies: %w", err)
	}
	campaignUserRepo := campaignuser.NewRepository(infraDependencies.db)

	return PromotionService{
		App: service.App{
			JoinCampaignHandler: service.NewJoinCampaignHandler(campaignUserRepo),
		},
	}, nil
}

func (s PromotionService) RegisterService(grpcServiceRegistrar grpc.ServiceRegistrar) {
	promotion.RegisterPromotionServer(grpcServiceRegistrar, s)
}

func (s PromotionService) JoinCampaign(ctx context.Context, req *promotion.JoinCampaignRequest) (*promotion.JoinCampaignReply, error) {
	err := s.App.JoinCampaignHandler.Handle(ctx, service.JoinCampaignRequest{
		CampaignId: req.CampaignId,
		UserId:     1,
	})
	if err != nil {
		return nil, fmt.Errorf("join campaign handler failed: %w", err)
	}
	return &promotion.JoinCampaignReply{
		Code:    utils.CodeSuccess,
		Message: utils.MessageSuccess,
	}, nil
}
