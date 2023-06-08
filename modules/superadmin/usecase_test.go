package superadmin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestUsecaseSuperadmin_ApprovedAdminRegister(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}
			if err := uc.ApprovedAdminRegister(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ApprovedAdminRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecaseSuperadmin_CreateCustomer(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		customer CustomerParam
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			fields: fields{superadminRepo: mockRepo},
			args: args{customer: CustomerParam{
				FirstName: "dias",
				LastName:  "pangestu",
				Email:     "dias@gmail.com",
				Avatar:    "dias.jpg",
			}},
			want: entities.Customer{
				ID:        1,
				FirstName: "dias",
				LastName:  "pangestu",
				Email:     "dias@gmail.com",
				Avatar:    "dias.jpg",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().CreateSuperadmin(gomock.Any()).Return(&tt.want, nil)

			got, err := uc.CreateCustomer(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_CreateSuperadmin(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		superadmin SuperAdminParam
	}

	//ctrl := gomock.NewController(t)
	//mockRepo := mocks.NewMockSuperAdminRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}
			got, err := uc.CreateSuperadmin(tt.args.superadmin)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSuperadmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSuperadmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_DeleteCustomerById(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		id uint
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "asdasd",
			fields:  fields{superadminRepo: mockRepo},
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}
			mockRepo.EXPECT().GetCustomerById(gomock.Any()).Return(&entities.Customer{}, nil)
			mockRepo.EXPECT().DeleteCustomerById(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			//if err := uc.DeleteCustomerById(tt.args.id); (err != nil) != tt.wantErr {
			//	t.Errorf("DeleteCustomerById() error = %v, wantErr %v", err, tt.wantErr)
			//}
			err := uc.DeleteCustomerById(tt.args.id)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestUsecaseSuperadmin_GetAllAdmins(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		username string
		page     int
		pageSize int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}
			got, err := uc.GetAllAdmins(tt.args.username, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllAdmins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllAdmins() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_GetAllCustomers(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		first_name string
		last_name  string
		email      string
		page       int
		pageSize   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}
			got, err := uc.GetAllCustomers(tt.args.first_name, tt.args.last_name, tt.args.email, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCustomers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_GetApprovalRequest(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entities.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}
			got, err := uc.GetApprovalRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApprovalRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApprovalRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_LoginSuperadmin(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		id       uint
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Actor
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}
			got, got1, err := uc.LoginSuperadmin(tt.args.id, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginSuperadmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginSuperadmin() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LoginSuperadmin() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUsecaseSuperadmin_RejectedAdminRegister(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}
			if err := uc.RejectedAdminRegister(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("RejectedAdminRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecaseSuperadmin_UpdateActivedAdmin(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}
			if err := uc.UpdateActivedAdmin(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateActivedAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecaseSuperadmin_UpdateDeadactivedAdmin(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}
			if err := uc.UpdateDeadactivedAdmin(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDeadactivedAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
