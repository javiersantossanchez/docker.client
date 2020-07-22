package dto

/*
ContainerDto used to parse result
*/
type ImageDetailDto struct {
	ID string

	Size int

	Created string

	RepoTags []string

	Tag string

	Repository string

	Author string
}
