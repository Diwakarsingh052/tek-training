package cmd

import (
	"go.uber.org/mock/gomock"
	core "mocking-cod/stores"
	"mocking-cod/stores/mockcore"
	"testing"
)

func TestCreate(t *testing.T) {
	u := core.User{Name: "diwakar"}
	ctrl := gomock.NewController(t)
	mStore := mockcore.NewMockStorer(ctrl)
	mStore.EXPECT().Create(u).Return(nil).Times(0)

	s := core.NewStore(mStore)
	err := s.Create(u)
	if err != nil {
		t.Fatal("problem in creating user")
	}

}
