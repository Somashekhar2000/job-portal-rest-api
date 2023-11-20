package service

import (
	"errors"
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/model"
	"job-portal-api/internal/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestService_UserSignup(t *testing.T) {
	type args struct {
		userData model.UserSignup
	}
	tests := []struct {
		name    string
		args    args
		want    model.User
		wantErr bool
		mockUserResponse func ()(model.User,error)
	}{
		{
			name: "failure",
			args: args{userData: model.UserSignup{}},
			want: model.User{},
			wantErr: true,
			mockUserResponse: func() (model.User, error) {
				return model.User{},errors.New("invalid user")
			},
		},
		{
			name: "success",
			args: args{userData: model.UserSignup{}},
			want: model.User{},
			wantErr: false,
			mockUserResponse: func() (model.User, error) {
				return model.User{},nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			ms := repository.NewMockUserRepository(mc)
			ma := authentication.NewMockAuthenticaton(mc)
			s,_:=NewUserService(ms,ma)
			if tt.mockUserResponse != nil {
				ms.EXPECT().CreateUser(gomock.Any()).Return(tt.mockUserResponse()).AnyTimes()
			}
			got, err := s.UserSignup(tt.args.userData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UserSignup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Userlogin(t *testing.T) {
	type args struct {
		userSignin model.UserLogin
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		mockUserResponse func()(model.User,error)
	}{
		
		{
			name: "faile",
			args: args{userSignin: model.UserLogin{}},
			want: "",
			wantErr: true,
			mockUserResponse:func() (model.User, error) {
				return model.User{},errors.New("error")
			},
		},
		{
			name: "success",
			args: args{userSignin: model.UserLogin{EmailID: "abc@gmail.com",Password: "12345678"}},
			want: "",
			wantErr: false,
			mockUserResponse: func() (model.User, error) {
				return model.User{
					EmailID: "abc@gmail.com",
					Password: "$2a$10$hNkswO/Wr.gDQJPnaYqvoOh0oQSnw8PkNm6Ipj6CVEYTpNetUPabC",
				},nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			ms := repository.NewMockUserRepository(mc)
			ma := authentication.NewMockAuthenticaton(mc)
			s,_:=NewUserService(ms,ma)
			if tt.mockUserResponse != nil {
				ms.EXPECT().CheckUser(gomock.Any()).Return(tt.mockUserResponse()).AnyTimes()
			}
			got, err := s.Userlogin(tt.args.userSignin)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UserSignup() = %v, want %v", got, tt.want)
			}
		})
	}
}
