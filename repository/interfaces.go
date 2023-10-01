// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
	RegisterUser(ctx context.Context, phoneNumber, fullName, passwordHash string) (int, error)
	LoginUser(ctx context.Context, phoneNumber, passwordHash string) (int, error)
	GetMyProfile(ctx context.Context, userID int, anotherAttr ...string) (string, string, error)
	UpdateMyProfile(ctx context.Context, userID int, phoneNumber, fullName string) error
}
