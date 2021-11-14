package domain

import (
	"golvl2/app/domain/db/environments"
	"golvl2/app/domain/db/users"
)

type Repositories interface {
	Users() users.Repository
	Environment() environments.Repository
}
