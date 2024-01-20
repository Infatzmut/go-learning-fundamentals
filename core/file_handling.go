package core

import (
	"fmt"
	"os"
	"path/filepath"
)

func join_method() {
	path := filepath.Join("dir1", "dir2", "text.txt")
	fmt.Println(path)
	fmt.Println(filepath.Dir(path))
	// return the last element fo the path
	fmt.Println(filepath.Base(path))
	// check if the path is an absolute path, returns a boolean
	fmt.Println(filepath.IsAbs(path))
	fmt.Println(filepath.IsAbs("/dir/file"))

	// Returns the file name extension used by path. Ex: .txt
	fmt.Println(filepath.Ext(path))
}

func os_method() {
	fileInfo, err := os.Stat("./temp.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.IsDir())
}

func read_file() {
	path := "./temp.txt"
	// ReadFile reads the named file and returns the contents
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	// returns as bytes
	//fmt.Println(data)
	fmt.Println(string(data))
}

func open_file() {
	path := "./temp.txt"
	// Open opens the named file for reading, If successful, methods on the returned file can be used for reading. If there is an error, it will be of type *PathError
	file, _ := os.Open(path)
	b := make([]byte, 4)
	for {
		n, err := file.Read(b)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		fmt.Println(string(b[0:n]))
	}
}

func write_file() {

	// OpenFile is the gerealized open call. It opens the named file with specified flag
	// WriteString is like Write, but writes the contents of string s rather than a slice of bytes
	file, err := os.OpenFile("./temp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.WriteString("Hope you had a good day!")
	if err != nil {
		fmt.Println(err)
	}

}
