package infrastructure

import (
	"context"
	"github.com/ssoql/faq-service/internal/app/entities"
	"github.com/ssoql/faq-service/internal/infrastructure/db"
	"github.com/ssoql/faq-service/utils/apiErrors"
	"reflect"
	"testing"
)

type writeRepositoryTest struct{}

func TestNewFaqWriteRepository(t *testing.T) {

}

func Test_faqWriteRepository_Delete(t *testing.T) {
	type fields struct {
		db *db.ClientDB
	}
	type args struct {
		ctx context.Context
		faq *entities.Faq
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   apiErrors.ApiError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &faqWriteRepository{
				db: tt.fields.db,
			}
			if got := r.Delete(tt.args.ctx, tt.args.faq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_faqWriteRepository_Insert(t *testing.T) {
	type fields struct {
		db *db.ClientDB
	}
	type args struct {
		ctx context.Context
		faq *entities.Faq
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   apiErrors.ApiError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &faqWriteRepository{
				db: tt.fields.db,
			}
			if got := r.Insert(tt.args.ctx, tt.args.faq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_faqWriteRepository_Update(t *testing.T) {
	type fields struct {
		db *db.ClientDB
	}
	type args struct {
		ctx context.Context
		faq *entities.Faq
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   apiErrors.ApiError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &faqWriteRepository{
				db: tt.fields.db,
			}
			if got := r.Update(tt.args.ctx, tt.args.faq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
