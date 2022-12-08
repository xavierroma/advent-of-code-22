package day7

type Directory struct {
	name            string
	size            int
	parentDirectory *Directory
	directories     []*Directory
	files           []*File
}

type File struct {
	name string
	size int
}
