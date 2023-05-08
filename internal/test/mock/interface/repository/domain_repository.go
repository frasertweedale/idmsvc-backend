// Code generated by mockery v2.26.1. DO NOT EDIT.

package repository

import (
	model "github.com/hmsidm/internal/domain/model"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// DomainRepository is an autogenerated mock type for the DomainRepository type
type DomainRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: db, orgID, data
func (_m *DomainRepository) Create(db *gorm.DB, orgID string, data *model.Domain) error {
	ret := _m.Called(db, orgID, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, *model.Domain) error); ok {
		r0 = rf(db, orgID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteById provides a mock function with given fields: db, orgID, uuid
func (_m *DomainRepository) DeleteById(db *gorm.DB, orgID string, uuid string) error {
	ret := _m.Called(db, orgID, uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, string) error); ok {
		r0 = rf(db, orgID, uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: db, orgID, uuid
func (_m *DomainRepository) FindById(db *gorm.DB, orgID string, uuid string) (*model.Domain, error) {
	ret := _m.Called(db, orgID, uuid)

	var r0 *model.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, string) (*model.Domain, error)); ok {
		return rf(db, orgID, uuid)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, string) *model.Domain); ok {
		r0 = rf(db, orgID, uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, string, string) error); ok {
		r1 = rf(db, orgID, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: db, orgID, offset, limit
func (_m *DomainRepository) List(db *gorm.DB, orgID string, offset int, limit int) ([]model.Domain, int64, error) {
	ret := _m.Called(db, orgID, offset, limit)

	var r0 []model.Domain
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, int, int) ([]model.Domain, int64, error)); ok {
		return rf(db, orgID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, int, int) []model.Domain); ok {
		r0 = rf(db, orgID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, string, int, int) int64); ok {
		r1 = rf(db, orgID, offset, limit)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(*gorm.DB, string, int, int) error); ok {
		r2 = rf(db, orgID, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RhelIdmClearToken provides a mock function with given fields: db, orgID, uuid
func (_m *DomainRepository) RhelIdmClearToken(db *gorm.DB, orgID string, uuid string) error {
	ret := _m.Called(db, orgID, uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, string) error); ok {
		r0 = rf(db, orgID, uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: db, orgID, data
func (_m *DomainRepository) Update(db *gorm.DB, orgID string, data *model.Domain) error {
	ret := _m.Called(db, orgID, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, *model.Domain) error); ok {
		r0 = rf(db, orgID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDomainRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewDomainRepository creates a new instance of DomainRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDomainRepository(t mockConstructorTestingTNewDomainRepository) *DomainRepository {
	mock := &DomainRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
