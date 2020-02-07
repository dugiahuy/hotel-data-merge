package updater

import (
	"testing"

	"github.com/dugiahuy/hotel-data-merge/src/repository"
	"go.uber.org/zap"
)

func Test_usecase_CollectData(t *testing.T) {
	type fields struct {
		repo   repository.HotelStorage
		logger *zap.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			if err := u.CollectData(); (err != nil) != tt.wantErr {
				t.Errorf("usecase.CollectData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
