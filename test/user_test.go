package test

import (
	"context"
	"fmt"
	db "parcel-management/db/sqlc"
	"parcel-management/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetListUsers(t *testing.T) {

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Millisecond)
		createRandomUser(t)
	}

	listUsers, err := testQueries.GetListUsers(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, listUsers)

	for _, u := range listUsers {
		fmt.Printf("u: %v\n", u)
	}
}

func createRandomUser(t *testing.T) db.User {
	role := createRandomRole(t)
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := db.CreateUserParams{
		Name:     util.RandomString(10),
		Email:    util.RandomEmail(),
		Password: hashedPassword,
		RoleID:   role.ID,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.RoleID, user.RoleID)

	return user
}
