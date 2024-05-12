package grpc

import (
	"context"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/api/grpc/promotion"
	"github.com/minhvuongrbs/financial-service-example/config"
	"github.com/minhvuongrbs/financial-service-example/internal/common/utils"
	"github.com/minhvuongrbs/financial-service-example/internal/promotion/adapters/repository/campaign"
	"github.com/minhvuongrbs/financial-service-example/internal/promotion/app"
	entitycampaign "github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/campaign"
	"google.golang.org/grpc"
)

type AdminService struct {
	promotion.UnimplementedPromotionAdminServer
	App app.App
}

func (s AdminService) DefineCampaign(ctx context.Context, req *promotion.DefineCampaignRequest) (*promotion.DefineCampaignReply, error) {
	md, err := convertCampaignMetadata(req)
	if err != nil {
		return nil, fmt.Errorf("invalid campaign metadata: %w", err)
	}
	appReq := app.DefineCampaignRequest{
		Key:      req.CampaignKey,
		Name:     req.Name,
		Metadata: md,
	}
	err = s.App.DefineCampaignHandler.Handle(ctx, appReq)
	if err != nil {
		return nil, fmt.Errorf("handle define campaign failed: %w", err)
	}
	return &promotion.DefineCampaignReply{
		Code:    utils.CodeSuccess,
		Message: utils.MessageSuccess,
	}, nil
}

func convertCampaignMetadata(req *promotion.DefineCampaignRequest) (entitycampaign.Metadata, error) {
	var strategy entitycampaign.Strategy
	switch v := req.GetMetadata().GetStrategy().GetMetadata().(type) {
	case *promotion.Strategy_StrategyTopRegister:
		var vouchers []entitycampaign.VoucherApplication
		for _, voucher := range v.StrategyTopRegister.Vouchers {
			vouchers = append(vouchers, entitycampaign.NewVoucher(voucher.Id, voucher.TotalVouchers))
		}
		strategy = entitycampaign.NewStrategy(entitycampaign.NewStrategyDetailTopRegister(v.StrategyTopRegister.TopUsers, vouchers))
	default:
		return entitycampaign.Metadata{}, fmt.Errorf("invalid campaign strategy: %v", v)
	}

	md := entitycampaign.NewMetadata(req.Metadata.StartAt.AsTime(), req.Metadata.EndAt.AsTime(), strategy)
	return md, nil
}

func (s AdminService) DefineVoucher(ctx context.Context, request *promotion.DefineVoucherRequest) (*promotion.DefineVoucherReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s AdminService) mustEmbedUnimplementedPromotionAdminServer() {
	//TODO implement me
	panic("implement me")
}

func NewPromotionAdminService(conf config.Config) (AdminService, error) {
	// load config
	infraDependencies, err := initInfrastructure(conf)
	if err != nil {
		return AdminService{}, fmt.Errorf("failed to init infra dependcies: %w", err)
	}
	campaignRepo := campaign.NewRepository(infraDependencies.db)

	return AdminService{
		App: app.App{
			DefineCampaignHandler: app.NewDefineCampaignHandler(campaignRepo),
		},
	}, nil
}

func (s AdminService) RegisterService(grpcServiceRegistrar grpc.ServiceRegistrar) {
	promotion.RegisterPromotionAdminServer(grpcServiceRegistrar, s)
}
