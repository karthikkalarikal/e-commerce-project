package repository_test

import (
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

func TestUserSignUp(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		input models.UserDetails
	}

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
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`
							INSERT INTO users (name,email,phone,password)
							VALUES ($1,$2,$3,$4);	
						`)).
					WithArgs(models.UserDetails{}).
					WillReturnResult(errors.New("the values are null"))
			},
			wantErr: errors.New("the values are null"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			u := repository.NewUserRepository(db)

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got, err := u.UserSignUp(tt.args.input)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("repository.UserSignUp() = %v, want %v", got, tt.want)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.UserSignUp() = %v, want %v", got, tt.want)
			}

		})
	}
}
