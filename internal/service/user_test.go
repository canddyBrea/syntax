package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordEncrypt(t *testing.T) {
	password := []byte("1649564Dl`")
	encrypted, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	assert.NoError(t, err)
	fmt.Println(string(encrypted))

	err = bcrypt.CompareHashAndPassword(encrypted, []byte("1649564DL~"))
	fmt.Println(err)
	assert.NotNil(t, err)
}
