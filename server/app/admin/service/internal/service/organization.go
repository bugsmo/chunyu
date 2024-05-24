package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/bugsmo/chunyu/app/admin/service/internal/data"

	v1 "github.com/bugsmo/chunyu/api/gen/go/user/service/v1"

	pagination "github.com/bugsmo/cy/contrib/kratos/bootstrap/gen/api/go/pagination/v1"
)

type OrganizationService struct {
	v1.UnimplementedOrganizationServiceServer

	log *log.Helper

	uc *data.OrganizationRepo
}

func NewOrganizationService(uc *data.OrganizationRepo, logger log.Logger) *OrganizationService {
	l := log.NewHelper(log.With(logger, "module", "organization/service/admin-service"))
	return &OrganizationService{
		log: l,
		uc:  uc,
	}
}

func (s *OrganizationService) ListOrganization(ctx context.Context, req *pagination.PagingRequest) (*v1.ListOrganizationResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *OrganizationService) GetOrganization(ctx context.Context, req *v1.GetOrganizationRequest) (*v1.Organization, error) {
	return s.uc.Get(ctx, req)
}

func (s *OrganizationService) CreateOrganization(ctx context.Context, req *v1.CreateOrganizationRequest) (*emptypb.Empty, error) {
	err := s.uc.Create(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *OrganizationService) UpdateOrganization(ctx context.Context, req *v1.UpdateOrganizationRequest) (*emptypb.Empty, error) {
	err := s.uc.Update(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *OrganizationService) DeleteOrganization(ctx context.Context, req *v1.DeleteOrganizationRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
