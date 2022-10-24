package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInterfaceList(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newInterfaceList := new(InterfaceList)
	newInterfaceList.Name = "InterfaceList_TEST"
	res, err := c.CreateInterfaceList(newInterfaceList)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res, "expecting a non-nil result")
	assert.NotNil(t, res.ID, "expecting the InterfaceList to have an id")
	assert.Equal(t, newInterfaceList.Name, res.Name)
	err = c.DeleteInterfaceList(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadInterfaceList(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newInterfaceList := new(InterfaceList)
	newInterfaceList.Name = "InterfaceList_TEST"
	res, err := c.CreateInterfaceList(newInterfaceList)
	assert.Nil(t, err, "expecting a nil error")
	list_get, err := c.ReadInterfaceList(res.ID)
	assert.Equal(t, newInterfaceList.Name, list_get.Name)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	err = c.DeleteInterfaceList(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
func TestUpdateInterfaceList(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newInterfaceList := new(InterfaceList)
	newInterfaceList.Name = "InterfaceList_TEST"
	res, err := c.CreateInterfaceList(newInterfaceList)
	assert.Nil(t, err, "expecting a nil error")
	newInterfaceList.Name = "InterfaceList_TEST2"
	resp, err := c.UpdateInterfaceList(res.ID, newInterfaceList)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, newInterfaceList.Name, resp.Name)
	err = c.DeleteInterfaceList(resp.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestDeleteInterfaceList(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newInterfaceList := new(InterfaceList)
	newInterfaceList.Name = "InterfaceList_TEST"
	res, err := c.CreateInterfaceList(newInterfaceList)
	assert.Nil(t, err, "expecting a nil error")
	err = c.DeleteInterfaceList(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
