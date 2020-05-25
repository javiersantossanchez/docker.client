package main

/*
ContainerDto used to parse result
*/
type dockerCommand interface {
	execute() string
}
