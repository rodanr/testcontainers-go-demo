package customer

import (
	"context"
	"log"
	"testing"

	"testcontainers-go-demo/testhelpers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CustomerRepoTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repository  *Repository
	ctx         context.Context
}

func (s *CustomerRepoTestSuite) SetupSuite() {
	s.ctx = context.Background()
	pgContainer, err := testhelpers.CreatePostgresContainer(s.ctx)
	if err != nil {
		log.Fatal(err)
	}
	s.pgContainer = pgContainer
	repository, err := NewRepository(s.ctx, s.pgContainer.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	s.repository = repository
}

func (s *CustomerRepoTestSuite) TearDownSuite() {
	if err := s.pgContainer.Terminate(s.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (s *CustomerRepoTestSuite) TestCreateCustomer() {
	t := s.T()

	customer, err := s.repository.CreateCustomer(s.ctx, Customer{
		Name:  "Henry",
		Email: "henry@gmail.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, customer.Id)
}

func (s *CustomerRepoTestSuite) TestGetCustomerByEmail() {
	t := s.T()

	customer, err := s.repository.GetCustomerByEmail(s.ctx, "john@gmail.com")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "John", customer.Name)
	assert.Equal(t, "john@gmail.com", customer.Email)
}

func TestCustomerRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepoTestSuite))
}
