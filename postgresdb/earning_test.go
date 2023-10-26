package postgresdb

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/yigithancolak/custmate/graph/model"
)

type EarningTestSuite struct {
	StoreTestSuite
	Organization *model.Organization
	Group1       *model.Group
	Group2       *model.Group
	Customer     *model.Customer
}

func TestEarningSuite(t *testing.T) {
	suite.Run(t, new(EarningTestSuite))
}

func (s *EarningTestSuite) SetupTest() {
	s.Organization = s.createRandomOrganization()
	instructor1 := s.createRandomInstructor(s.Organization.ID)
	instructor2 := s.createRandomInstructor(s.Organization.ID)

	s.Group1 = s.createRandomGroup(instructor1.ID, s.Organization.ID)
	s.Group2 = s.createRandomGroup(instructor2.ID, s.Organization.ID)

	s.Customer = s.createRandomCustomer(s.Organization.ID, []string{s.Group1.ID, s.Group2.ID})

}
