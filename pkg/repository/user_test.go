package repository_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserSignUp(t *testing.T) {

	type args struct {
		input models.UserDetails
	}

	var det models.UserDetails

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       models.UserDetailsResponse
		wantErr    error
	}{
		{
			name: "null date in user signup",
			args: args{
				input: models.UserDetails{},
			},
			beforeTest: func(s sqlmock.Sqlmock) {
				s.
					ExpectQuery(`insert into users\(name, email, password, phone\) values \(\$1,\$2,\$3,\$4\) returning \*`).WithArgs(det.Name, det.Email, det.Password, det.Phone).WillReturnError(errors.New("insert into users table violates not null constraints"))
			},
			wantErr: errors.New("insert into users table violates not null constraints"),
			want:    models.UserDetailsResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New()
			assert.NoError(t, err)
			defer mockDB.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			assert.NoError(t, err)

			// db := sqlx.NewDb(mockDB, "sqlmock")

			u := repository.NewUserRepository(gormDB)

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got, err := u.UserSignUp(tt.args.input)
			if (err != nil) && err == tt.wantErr {
				t.Errorf("repository.UserSignUp() = %v, want %v", got, tt.want)
				return
			}
			fmt.Println("gggg")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.UserSignUp() = %v, want %v", got, tt.want)
			}

		})
	}
}
