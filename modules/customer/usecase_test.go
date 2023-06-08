package customer

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func TestUsecaseCustomer_CreateCustomer(t *testing.T) {
	type fields struct {
		customerRepo repositories.CustomerRepositoryInterface
	}
	type args struct {
		customer CustomerParam
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockCustomerRepositoryInterface(ctrl)

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
			fields: fields{customerRepo: mockRepo},
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
			uc := UsecaseCustomer{
				customerRepo: tt.fields.customerRepo,
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

func TestUsecaseCustomer_DeleteCustomerById(t *testing.T) {
	type fields struct {
		customerRepo repositories.CustomerRepositoryInterface
	}
	type args struct {
		id uint
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockCustomerRepositoryInterface(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Case 1",
			fields:  fields{customerRepo: mockRepo},
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseCustomer{
				customerRepo: tt.fields.customerRepo,
			}

			mockRepo.EXPECT().GetCustomerById(tt.args.id).Return(&entities.Customer{}, nil)

			mockRepo.EXPECT().DeleteCustomerById(tt.args.id, gomock.Any()).Return(nil).Times(1)

			if err := uc.DeleteCustomerById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCustomerById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecaseCustomer_GetCustomerById(t *testing.T) {
	type fields struct {
		customerRepo repositories.CustomerRepositoryInterface
	}
	type args struct {
		id uint
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockCustomerRepositoryInterface(ctrl)

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
			fields: fields{customerRepo: mockRepo},
			args:   args{id: 1},
			want: entities.Customer{
				ID:        1,
				FirstName: "dias",
				LastName:  "pangestu",
				Email:     "dias@gmail.com",
				Avatar:    "dias1.jpg",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseCustomer{
				customerRepo: tt.fields.customerRepo,
			}

			mockRepo.EXPECT().GetCustomerById(tt.args.id).Return(&tt.want, nil).Times(1)

			got, err := uc.GetCustomerById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomerById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomerById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseCustomer_UpdateCustomerById(t *testing.T) {
	type fields struct {
		customerRepo repositories.CustomerRepositoryInterface
	}
	type args struct {
		id       uint
		customer CustomerParam
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockCustomerRepositoryInterface(ctrl)

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
			fields: fields{customerRepo: mockRepo},
			args: args{
				id: 1,
				customer: CustomerParam{
					FirstName: "dias",
					LastName:  "pangestu",
					Email:     "dias@gmail.com",
					Avatar:    "dias.jpg",
				},
			},
			want: entities.Customer{
				ID:        1,
				FirstName: "dias updated",
				LastName:  "pangestu updated",
				Email:     "diasupdated@gmail.com",
				Avatar:    "diasupdated.jpg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseCustomer{
				customerRepo: tt.fields.customerRepo,
			}

			mockRepo.EXPECT().GetCustomerById(tt.args.id).Return(&entities.Customer{}, nil)
			mockRepo.EXPECT().UpdateCustomerById(gomock.Any(), gomock.Any()).Return(&tt.want, nil)

			got, err := uc.UpdateCustomerById(tt.args.id, tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCustomerById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateCustomerById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
