package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/seregaa020292/capitalhub/infrastructure/session/mock"
	"github.com/seregaa020292/capitalhub/internal/models"
)

func TestSessionUC_CreateSession(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSessRepo := mock.NewMockSessRepository(ctrl)
	sessUC := NewSessionUseCase(mockSessRepo, nil)

	ctx := context.Background()
	sess := &models.Session{}
	sid := "session id"

	mockSessRepo.EXPECT().CreateSession(gomock.Any(), gomock.Eq(sess), 10).Return(sid, nil)

	createdSess, err := sessUC.CreateSession(ctx, sess, 10)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotEqual(t, createdSess, "")
}

func TestSessionUC_GetSessionByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSessRepo := mock.NewMockSessRepository(ctrl)
	sessUC := NewSessionUseCase(mockSessRepo, nil)

	ctx := context.Background()
	sess := &models.Session{}
	uid := uuid.New()
	sid := "session id"

	mockSessRepo.EXPECT().GetSessionByID(gomock.Any(), gomock.Eq(uid), gomock.Eq(sid)).Return(sess, nil)

	session, err := sessUC.GetSessionByID(ctx, uid, sid)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, session)
}

func TestSessionUC_DeleteByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSessRepo := mock.NewMockSessRepository(ctrl)
	sessUC := NewSessionUseCase(mockSessRepo, nil)

	ctx := context.Background()
	uid := uuid.New()
	sid := "session id"

	mockSessRepo.EXPECT().DeleteByID(gomock.Any(), gomock.Eq(uid), gomock.Eq(sid)).Return(nil)

	err := sessUC.DeleteByID(ctx, uid, sid)
	require.NoError(t, err)
	require.Nil(t, err)
}
