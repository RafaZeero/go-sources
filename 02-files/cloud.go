package main

import (
	"errors"
	"fmt"
	"os"
)

// CLIENT CLOUD SDK

// fake cloud system
type MyCloud struct {
	*CloudFile
}

// fake cloud opts
type CloudOpts struct {
	ApiToken    string
	Credentials struct {
		Username string
		Password string
	}
	// ... others stuff that are necessary
}

// should check for apitoken or credentials
func NewCloudInstance(opts CloudOpts) *MyCloud {
	return &MyCloud{}
}

// should ping server for connecttion check
func (c *MyCloud) Ping() error {
	return nil
}

// maybe this could be the file management or just get info about the storage
// like space, total files, owners etc.
func (c *MyCloud) Storage() error {
	return nil
}

type CloudFile struct {
	*os.File
}

// should have a proper error handling for the system
var (
	// error X due to Y
	ErrFailedToCreate = errors.New("err_msg_x")
)

func manageFileError(fileErr error) error { return nil }
func veryGoodLogSystem(log any)           {}

func (cf *CloudFile) CreateFile(filename string, content string) (*CloudFile, error) {
	// should validate filename + error handling
	f, err := os.Create(filename)
	if err != nil {
		veryGoodLogSystem(err)
		return nil, manageFileError(err)
	}
	veryGoodLogSystem(f)

	// should validate file content
	n, err := f.Write([]byte(content))
	if err != nil {
		veryGoodLogSystem(err)
		return nil, manageFileError(err)
	}
	veryGoodLogSystem(n)

	return &CloudFile{f}, nil
}

func (cf *CloudFile) RemoveFile(filename string) error {
	// should validate filename + error handling
	err := os.Remove(filename)
	manageFileError(err)
	veryGoodLogSystem("msg_rm_ok")

	return nil
}

func (cf *CloudFile) ReadFile(filename string) ([]byte, error) {
	// should validate filename + error handling
	data, err := os.ReadFile(filename)
	manageFileError(err)
	veryGoodLogSystem(data)

	return data, nil
}

func main() {
	// setup
	cloud := NewCloudInstance(CloudOpts{})

	filename := "foo.txt"
	content := "asdjkaspokd"

	if err := cloud.Ping(); err != nil {
		// ping failed
		return
	}

	// exec create file
	f, err := cloud.CreateFile(filename, content)
	if err != nil {
		// yea...
		return
	}

	// do things with file
	_ = f

	// then close it
	defer f.Close()

	// exec read file content
	b, err := cloud.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// do things with file content
	_ = b

	// removing file created
	if err := cloud.RemoveFile(filename); err != nil {
		// yea 2.0 ...
		return
	}
}
