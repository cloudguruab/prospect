package src

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func showLoginPage(c *gin.Context) {}

func performLogin(c *gin.Context) {}

func logout(c *gin.Context) {}

func CheckEmpty(name string) (bool, error) {
	/*
		Handler for checking /tmp for newly uploaded files
	*/

	f, err := os.Open(name)

	if err != nil {
		return false, err
	}

	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true, nil
	}

	return false, err
}

func Check(e error) {
	/*
		Custom error handler for instances
	*/
	if e != nil {
		panic(e)
	}
}
