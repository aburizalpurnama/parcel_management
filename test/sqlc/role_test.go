package db

import (
	"context"
	db "parcel-management/db/sqlc"
	"parcel-management/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateRandomRole(t *testing.T) {
	createRandomRole(t)
}

func createRandomRole(t *testing.T) db.Role {
	name := util.RandomString(6)
	role, err := testQueries.CreateRole(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, role)

	require.Equal(t, name, role.Name)
	require.NotZero(t, role.ID)
	require.NotZero(t, role.CreatedAt)

	return role
}
