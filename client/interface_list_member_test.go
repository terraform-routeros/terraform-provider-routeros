package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateInterfaceListMemberTestObjects() (*InterfaceListMember, *InterfaceList, error) {
	c := NewClient(GetCredentialsFromEnvVar())
	newList := new(InterfaceList)
	newList.Name = "List_TEST"
	res, err := c.CreateInterfaceList(newList)
	if err != nil {
		return nil, nil, err
	}
	newMember := new(InterfaceListMember)
	newMember.List = newList.Name
	newMember.Interface = "ether1"
	res_m, err := c.CreateInterfaceListMember(newMember)
	if err != nil {
		return nil, nil, err
	}
	return res_m, res, nil
}

func TestCreateInterfaceListMember(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	newList := new(InterfaceList)
	newList.Name = "List_TEST"
	res, err := c.CreateInterfaceList(newList)
	newMember := new(InterfaceListMember)
	newMember.List = newList.Name
	newMember.Interface = "ether1"
	res_m, err := c.CreateInterfaceListMember(newMember)
	assert.Nil(t, err, "expecting a nil error")
	assert.NotNil(t, res_m, "expecting a non-nil result")
	assert.NotNil(t, res_m.ID, "expecting the Member to have an id")
	assert.Equal(t, newMember.List, res_m.List)
	err = c.DeleteInterfaceListMember(res_m.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
	err = c.DeleteInterfaceList(res.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestReadInterfaceListMember(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	list_member, list, err := CreateInterfaceListMemberTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	member_get, err := c.ReadInterfaceListMember(list_member.ID)
	assert.Equal(t, list_member.Interface, member_get.Interface)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, member_get, "expecting non-nil result")
	err = c.DeleteInterfaceListMember(list_member.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
	err = c.DeleteInterfaceList(list.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
func TestUpdateInterfaceListMember(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	list_member, list, err := CreateInterfaceListMemberTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	newMember := InterfaceListMember{}
	newMember.Interface = "bridge"
	newMember.List = list_member.List
	resp, err := c.UpdateInterfaceListMember(list_member.ID, &newMember)
	assert.Nil(t, err, "expecting a nil error")
	assert.Equal(t, newMember.Interface, resp.Interface)
	err = c.DeleteInterfaceListMember(list_member.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
	err = c.DeleteInterfaceList(list.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}

func TestDeleteInterfaceListMember(t *testing.T) {
	c := NewClient(GetCredentialsFromEnvVar())
	list_member, list, err := CreateInterfaceListMemberTestObjects()
	assert.Nil(t, err, "expecting a nil error")
	err = c.DeleteInterfaceListMember(list_member.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
	err = c.DeleteInterfaceList(list.ID)
	assert.Nil(t, err, "expecting a nil error on delete")
}
