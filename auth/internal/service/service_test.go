package service

import (
	"auth/internal/mocks"
	"auth/internal/model"
	"auth/pkg/transport/medicalpb"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type ServiceSuite struct {
	suite.Suite
	repo       *mocks.MockRepository
	blockchain *mocks.MockBlockchainProcessor
	service    *Service
	ctrl       *gomock.Controller
}

func (s *ServiceSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())

	conf := ServiceConf{
		SecretKey:         "secret",
		ExpiredAfterHours: 24,
	}

	s.repo = mocks.NewMockRepository(s.ctrl)
	s.blockchain = mocks.NewMockBlockchainProcessor(s.ctrl)

	s.service = NewService(s.repo, conf, s.blockchain)
}

func (s *ServiceSuite) TearDownSuite() {
	s.ctrl.Finish()
}

func (s *ServiceSuite) TestRegister() {
	req := model.RegisterRequest{
		Username: "testuser",
		Password: "testpassword",
		Email:    "email@m.com",
		Phone:    "1234567890",
		ID:       uuid.New(),
	}
	user := model.RequestToUser(req)
	user.ID = req.ID
	user.PKey = "pkey"
	user.Password, _ = model.HashPassword(user.Password)

	gomock.InOrder(
		s.blockchain.EXPECT().GenerateNewAccount(gomock.Any(), &medicalpb.Empty{}).Return(&medicalpb.AccountResponse{PrivateKey: "pkey"}, nil),
		s.repo.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(nil),
	)

	token, err := s.service.Register(context.Background(), req)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), token)
}

func (s *ServiceSuite) TestAuthorize() {
	request := model.AuthRequest{
		Username: "testuser",
		Password: "123123",
	}

	gomock.InOrder(
		s.repo.EXPECT().GetUserByUsername(gomock.Any(), "testuser").Return(model.User{Password: "$2a$10$.z28atrV.5o6XJQsGlfnTeNslfxhkm/UxBabqDRKO4jU7RC.0OJWG"}, nil),
	)

	token, err := s.service.Authorize(context.Background(), request)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), token)
}

func (s *ServiceSuite) TestValidateToken() {
	gomock.InOrder()

	token, err := s.service.ValidateToken(context.Background(), "token")
	assert.Error(s.T(), err)
	assert.NotNil(s.T(), token)
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
