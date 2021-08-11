// Code generated by MockGen. DO NOT EDIT.
// Source: src/repository/ingrediente_repository.go

// Package mock_repository is a generated GoMock package.
package mock

import (
	model "cookbook/src/model"
	"cookbook/src/repository"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockIngredienteRepository is a mock of IngredienteRepository interface.
type MockIngredienteRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIngredienteRepositoryMockRecorder
}

// MockIngredienteRepositoryMockRecorder is the mock recorder for MockIngredienteRepository.
type MockIngredienteRepositoryMockRecorder struct {
	mock *MockIngredienteRepository
}

// NewMockIngredienteRepository creates a new mock instance.
func NewMockIngredienteRepository(ctrl *gomock.Controller) *MockIngredienteRepository {
	mock := &MockIngredienteRepository{ctrl: ctrl}
	mock.recorder = &MockIngredienteRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIngredienteRepository) EXPECT() *MockIngredienteRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockIngredienteRepository) Delete(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIngredienteRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIngredienteRepository)(nil).Delete), arg0)
}

// FindById mocks base method.
func (m *MockIngredienteRepository) FindById(arg0 uint64) (model.Ingrediente, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(model.Ingrediente)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockIngredienteRepositoryMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockIngredienteRepository)(nil).FindById), arg0)
}

// GetAll mocks base method.
func (m *MockIngredienteRepository) GetAll(arg0 string) ([]model.Ingrediente, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]model.Ingrediente)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIngredienteRepositoryMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIngredienteRepository)(nil).GetAll), arg0)
}

// Save mocks base method.
func (m *MockIngredienteRepository) Save(arg0 model.Ingrediente) (model.Ingrediente, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(model.Ingrediente)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockIngredienteRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIngredienteRepository)(nil).Save), arg0)
}

// Update mocks base method.
func (m *MockIngredienteRepository) Update(arg0 model.Ingrediente) (model.Ingrediente, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(model.Ingrediente)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIngredienteRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIngredienteRepository)(nil).Update), arg0)
}

// WithTrx mocks base method.
func (m *MockIngredienteRepository) WithTrx(arg0 *gorm.DB) repository.IngredienteRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTrx", arg0)
	ret0, _ := ret[0].(repository.IngredienteRepository)
	return ret0
}

// WithTrx indicates an expected call of WithTrx.
func (mr *MockIngredienteRepositoryMockRecorder) WithTrx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTrx", reflect.TypeOf((*MockIngredienteRepository)(nil).WithTrx), arg0)
}
