package admin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func TestUsecaseAdmin_CreateCustomer(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
	}
	type args struct {
		customer CustomerParam
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockAdminRepositoryInterface(ctrl)

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
			fields: fields{adminRepo: mockRepo},
			args: args{customer: CustomerParam{
				FirstName: "dias",
				LastName:  "pangestu",
				Email:     "dias@gmail.com",
				Avatar:    "dias.jpg",
			}},
			want: entities.Customer{
				FirstName: "dias",
				LastName:  "pangestu",
				Email:     "dias@gmail.com",
				Avatar:    "dias.jpg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}

			mockRepo.EXPECT().CreateCustomer(gomock.Any()).Return(&tt.want, nil)

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

func TestUsecaseAdmin_DeleteAdminById(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
	}
	type args struct {
		id uint
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Case 1",
			fields:  fields{adminRepo: mockRepo},
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}

			mockRepo.EXPECT().GetAdminById(gomock.Any()).Return(&entities.Actor{}, nil)

			mockRepo.EXPECT().DeleteAdminById(gomock.Any(), gomock.Any()).Return(nil).Times(1)

			if err := uc.DeleteAdminById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteAdminById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecaseAdmin_DeleteCustomerById(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
	}
	type args struct {
		id uint
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Case 1",
			fields:  fields{adminRepo: mockRepo},
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}

			mockRepo.EXPECT().GetCustomerById(gomock.Any()).Return(&entities.Customer{}, nil)

			mockRepo.EXPECT().DeleteCustomerById(gomock.Any(), gomock.Any()).Return(nil).Times(1)

			if err := uc.DeleteCustomerById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCustomerById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecaseAdmin_GetAdminById(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
	}
	type args struct {
		id uint
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			fields: fields{adminRepo: mockRepo},
			args:   args{id: 1},
			want: entities.Actor{
				ID:         1,
				Username:   "diaspangestu",
				Password:   "dias123",
				RoleID:     2,
				IsVerified: "true",
				IsActived:  "true",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}

			mockRepo.EXPECT().GetAdminById(tt.args.id).Return(&tt.want, nil).Times(1)

			got, err := uc.GetAdminById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdminById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdminById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseAdmin_GetAllCustomers(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
	}
	type args struct {
		first_name string
		last_name  string
		email      string
		page       int
		pageSize   int
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			fields: fields{adminRepo: mockRepo},
			args: args{
				first_name: "dias",
				last_name:  "pangestu",
				email:      "dias@gmail.com",
				page:       2,
				pageSize:   6,
			},
			want:    []*entities.Customer{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}

			mockRepo.EXPECT().GetAllCustomers(tt.args.first_name, tt.args.last_name, tt.args.email, tt.args.page, tt.args.pageSize).Return(tt.want, nil)

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

func TestUsecaseAdmin_LoginAdmin(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
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
			uc := UsecaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}
			got, got1, err := uc.LoginAdmin(tt.args.id, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginAdmin() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LoginAdmin() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUsecaseAdmin_RegisterAdmin(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
	}
	type args struct {
		admin LoginAdminParam
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			fields: fields{adminRepo: mockRepo},
			args: args{admin: LoginAdminParam{
				ID:       1,
				Username: "dias",
				Password: "dias123",
			}},
			want: &entities.Actor{
				ID:       1,
				Username: "dias",
				Password: "dias123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}

			mockRepo.EXPECT().RegisterAdmin(gomock.Any()).Return(tt.want, nil)

			got, err := uc.RegisterAdmin(tt.args.admin)
			got.Password = "dias123"
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseAdmin_SaveCustomersFromAPI(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Case 1",
			fields:  fields{adminRepo: mockRepo},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}

			mockRepo.EXPECT().SaveCustomersFromAPI(gomock.Any()).Return(nil)

			if err := uc.SaveCustomersFromAPI(); (err != nil) != tt.wantErr {
				t.Errorf("SaveCustomersFromAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecaseAdmin_UpdateAdminById(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminRepositoryInterface
	}
	type args struct {
		id    uint
		admin AdminParam
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockAdminRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			fields: fields{adminRepo: mockRepo},
			args: args{
				id: 1,
				admin: AdminParam{
					Username: "dias",
					Password: "dias123",
				},
			},
			want: entities.Actor{
				ID:       2,
				Username: "diass",
				Password: "diass123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}

			mockRepo.EXPECT().GetAdminById(tt.args.id).Return(&entities.Actor{}, nil)

			mockRepo.EXPECT().UpdateAdminById(gomock.Any(), gomock.Any()).Return(&tt.want, nil)

			got, err := uc.UpdateAdminById(tt.args.id, tt.args.admin)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAdminById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAdminById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
