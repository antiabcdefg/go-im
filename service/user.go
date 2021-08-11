package service

import (
	"errors"
	"fmt"
	"go-im/model"
	"go-im/util"
	"math/rand"
	"time"
)

type UserService struct {
}

func (service *UserService) Login(name, plainpwd string) (user model.User, err error) {
	tmp := model.User{}
	DbEngin.Where("name=?", name).Get(&tmp)
	if tmp.Id == 0 {
		return tmp, errors.New("该用户不存在")
	}
	if !util.ValidatePasswd(plainpwd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码不正确")
	}

	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token = token

	DbEngin.ID(tmp.Id).Cols("token").Update(&tmp)
	return tmp, nil
}

func (service *UserService) Register(name, plainpwd, telephone, avatar, sex string) (user model.User, err error) {
	tmp := model.User{}
	_, err = DbEngin.Where("name=?", name).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	if tmp.Id > 0 {
		return tmp, errors.New("该用户已经注册")
	}

	tmp.Name = name
	tmp.Telephone = telephone
	tmp.Avatar = avatar
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
	tmp.CreateAt = time.Now()
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())

	_, err = DbEngin.InsertOne(&tmp)
	return tmp, err
}

func (service *UserService) FindById(userId int64) (user model.User) {
	tmp := model.User{}
	DbEngin.ID(userId).Get(&tmp)
	return tmp
}
