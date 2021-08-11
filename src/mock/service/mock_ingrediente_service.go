// Code generated by MockGen. DO NOT EDIT.
// Source: src/service/ingrediente_service.go

// Package mock is a generated GoMock package.
package mock

import (
	model "cookbook/src/model"
	service "cookbook/src/service"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockIngredienteService is a mock of IngredienteService interface.
type MockIngredienteService struct {
	ctrl     *gomock.Controller
	recorder *MockIngredienteServiceMockRecorder
}

// MockIngredienteServiceMockRecorder is the mock recorder for MockIngredienteService.
type MockIngredienteServiceMockRecorder struct {
	mock *MockIngredienteService
}

// NewMockIngredienteService creates a new mock instance.
func NewMockIngredienteService(ctrl *gomock.Controller) *MockIngredienteService {
	mock := &MockIngredienteService{ctrl: ctrl}
	mock.recorder = &MockIngredienteServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIngredienteService) EXPECT() *MockIngredienteServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockIngredienteService) Delete(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIngredienteServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIngredienteService)(nil).Delete), arg0)
}

// FindById mocks base method.
func (m *MockIngredienteService) FindById(arg0 uint64) (model.Ingrediente, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(model.Ingrediente)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockIngredienteServiceMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockIngredienteService)(nil).FindById), arg0)
}

// GetAll mocks base method.
func (m *MockIngredienteService) GetAll(arg0 string) ([]model.Ingrediente, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]model.Ingrediente)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIngredienteServiceMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIngredienteService)(nil).GetAll), arg0)
}

// Save mocks base method.
func (m *MockIngredienteService) Save(arg0 model.Ingrediente) (model.Ingrediente, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(model.Ingrediente)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockIngredienteServiceMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIngredienteService)(nil).Save), arg0)
}

// Update mocks base method.
func (m *MockIngredienteService) Update(arg0 model.Ingrediente) (model.Ingrediente, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(model.Ingrediente)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIngredienteServiceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIngredienteService)(nil).Update), arg0)
}

// WithTrx mocks base method.
func (m *MockIngredienteService) WithTrx(arg0 *gorm.DB) service.IngredienteService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTrx", arg0)
	ret0, _ := ret[0].(service.IngredienteService)
	return ret0
}

// WithTrx indicates an expected call of WithTrx.
func (mr *MockIngredienteServiceMockRecorder) WithTrx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTrx", reflect.TypeOf((*MockIngredienteService)(nil).WithTrx), arg0)
}
