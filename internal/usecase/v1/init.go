package v1

import "awesomeProject/internal/repository"

func GetApi(repoDbLayer repository.DBRepository) ApiV1 {
	return ApiV1{db: repoDbLayer}
}
