package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nochzato/go-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	t.Helper()

	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordChangedAt, user2.PasswordChangedAt)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
}

func TestUpdateUserFullName(t *testing.T) {
	user := createRandomUser(t)

	newFullName := util.RandomOwner()

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		FullName: sql.NullString{
			String: newFullName,
			Valid:  true,
		},
		Username: user.Username,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.NotEqual(t, updatedUser.FullName, user.FullName)
	require.Equal(t, updatedUser.FullName, newFullName)
	require.Equal(t, updatedUser.Email, user.Email)
	require.Equal(t, updatedUser.HashedPassword, user.HashedPassword)
}

func TestUpdateUserEmail(t *testing.T) {
	user := createRandomUser(t)

	newEmail := util.RandomEmail()

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
		Username: user.Username,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.NotEqual(t, updatedUser.Email, user.Email)
	require.Equal(t, updatedUser.Email, newEmail)
	require.Equal(t, updatedUser.FullName, user.FullName)
	require.Equal(t, updatedUser.HashedPassword, user.HashedPassword)
}

func TestUpdateUserPassword(t *testing.T) {
	user := createRandomUser(t)

	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		HashedPassword: sql.NullString{
			String: newHashedPassword,
			Valid:  true,
		},
		Username: user.Username,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.NotEqual(t, updatedUser.HashedPassword, user.HashedPassword)
	require.Equal(t, updatedUser.HashedPassword, newHashedPassword)
	require.Equal(t, updatedUser.FullName, user.FullName)
	require.Equal(t, updatedUser.Email, user.Email)
}

func TestUpdateUserAllFields(t *testing.T) {
	user := createRandomUser(t)

	newFullName := util.RandomOwner()
	newEmail := util.RandomEmail()
	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		FullName: sql.NullString{
			String: newFullName,
			Valid:  true,
		},
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
		HashedPassword: sql.NullString{
			String: newHashedPassword,
			Valid:  true,
		},
		Username: user.Username,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.NotEqual(t, updatedUser.FullName, user.FullName)
	require.NotEqual(t, updatedUser.HashedPassword, user.HashedPassword)
	require.NotEqual(t, updatedUser.Email, user.Email)
	require.Equal(t, updatedUser.FullName, newFullName)
	require.Equal(t, updatedUser.HashedPassword, newHashedPassword)
	require.Equal(t, updatedUser.Email, newEmail)
}
