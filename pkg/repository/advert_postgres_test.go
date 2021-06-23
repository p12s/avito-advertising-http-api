package repository

import (
	"fmt"
	"github.com/p12s/avito-advertising-http-api"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
	"time"
)


func TestAdvertPostgres_Get(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		closeErr := db.Close()
		if closeErr != nil {
			panic(closeErr)
		}
	}()

	r := NewAdvertPostgres(db)

	type args struct {
		advertId int
		params common.AdvertFieldParams
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    common.AdvertWithPhoto
		wantErr bool
	}{
		{
			// expected: common.AdvertWithPhoto{Id:1, Title:"Title1", Description:"Description1", Price:123.5, CreatedAt:time.Time{wall:0xc02ce48a70dca6d8, ext:880629, loc:(*time.Location)(0x13a7920)}, Photos:[]common.Photo{}}
			// actual  : common.AdvertWithPhoto{Id:0, Title:"", Description:"", Price:0, CreatedAt:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}, Photos:[]common.Photo(nil)}



			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "price", "created_at"}).
					AddRow(1, "Title1", "Description1", 123.50, time.Now())

				mock.ExpectQuery("SELECT (.+) FROM advert a LEFT JOIN photo ph ON (.+) WHERE (.+)").
					WithArgs(1, common.AdvertFieldParams{Fields: common.Fields{}}).WillReturnRows(rows)
			},
			input: args{
				advertId: 1,
				params: common.AdvertFieldParams{Fields: common.Fields{}},
			},
			want: common.AdvertWithPhoto{1, "Title1", "Description1", 123.50, time.Now(), []common.Photo{}},
		},/*
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "done"})

				mock.ExpectQuery("SELECT (.+) FROM todo_items ti INNER JOIN lists_items li on (.+) INNER JOIN users_lists ul on (.+) WHERE (.+)").
					WithArgs(404, 1).WillReturnRows(rows)
			},
			input: args{
				itemId: 404,
				userId: 1,
			},
			wantErr: true,
		},*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			fmt.Println("tt.input")
			fmt.Println(tt.input)

			// actual  : common.AdvertWithPhoto{Id:0, Title:"", Description:"", Price:0, CreatedAt:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}, Photos:[]common.Photo(nil)}

			got, err := r.Get(tt.input.advertId, tt.input.params)
			fmt.Println("!!! got !!!")
			fmt.Println(got)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				//fmt.Println("tt.want")
				//fmt.Println(tt.want)
				//fmt.Println("got")
				//fmt.Println(got)

				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

