package paginator

import (
	"github.com/dyoshikawa/paginate-gorm/paginator/mock_paginator"
	"github.com/golang/mock/gomock"
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestExampleSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_paginator.NewMockDBIface(ctrl)
	m.
		EXPECT().
		New()
}
