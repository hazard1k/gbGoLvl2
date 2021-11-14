package models

type EnvironmentType string

type Environment struct {
	Type  EnvironmentType
	Users []*User
}
