package repository_test

import (
	"errors"
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
					ExpectQuery(`insert into users\(name, email, password, phone\) values \(\$1,\$2,\$3,\$4\) returning \*`).
					WithArgs(det.Name, det.Email, det.Password, det.Phone).
					WillReturnError(errors.New("insert into users table violates not null constraints"))
			},
			wantErr: errors.New("insert into users table violates not null constraints"),
			want:    models.UserDetailsResponse{},
		},
		{
			name: "success signup user",
			args: args{
				input: models.UserDetails{Name: "John", Email: "john00@gmail.com", Phone: "123456789", Password: "password", ConfirmPassword: "password"},
			},
			beforeTest: func(s sqlmock.Sqlmock) {
				s.
					ExpectQuery(`insert into users\(name, email, password, phone\) values \(\$1,\$2,\$3,\$4\) returning \*`).
					WithArgs("John", "john00@gmail.com", "password", "123456789").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
						AddRow(1, "John", "john00@gmail.com", "123456789"))
			},
			want:    models.UserDetailsResponse{Id: 1, Name: "John", Email: "john00@gmail.com", Phone: "123456789"},
			wantErr: nil,
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

			// assert.Equal(t, tt.wantErr, err)
			// fmt.Println("gggg")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.UserSignUp() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestCheckUserAvailability(t *testing.T) {
	tests := []struct {
		name string
		args string
		stub func(sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "user available",
			args: "john00@gmail.com",
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `^SELECT COUNT\(\*\) from users where email= \$1`
				s.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"COUNT"}).AddRow(5))
			},

			want: true,
		},
		{
			name: "error from database",
			args: "john@gmail.com",
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `^SELECT COUNT\(\*\) from users where email= \$1`

				s.ExpectQuery(expectedQuery).WithArgs("john@gmail.com").
					WillReturnError(errors.New("text string"))
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := repository.NewUserRepository(gormDB)

			result := u.CheckUserAvailability(tt.args)
			assert.Equal(t, tt.want, result)
		})
	}
}
