package service

import (
	"errors"
	"go-im/model"
	"time"
)

type ContactService struct {
}

func (service *ContactService) AddFriend(userid, dstid int64) error {
	if userid == dstid {
		return errors.New("不能添加自己")
	}
	tmp := model.Contact{}
	DbEngin.Where("owner_id=?", userid).And("dst_obj = ?", dstid).And("cate = ?", model.CONTACT_CATE_USER).Get(&tmp)
	if tmp.Id > 0 {
		return errors.New("该用户已经添加")
	}

	session := DbEngin.NewSession()
	session.Begin()
	_, e := session.InsertOne(model.Contact{OwnerId: userid, DstObj: dstid, Cate: model.CONTACT_CATE_USER, CreateAt: time.Now()})
	_, e2 := session.InsertOne(model.Contact{OwnerId: dstid, DstObj: userid, Cate: model.CONTACT_CATE_USER, CreateAt: time.Now()})
	if e == nil && e2 == nil {
		session.Commit()
		return nil
	} else {
		session.Rollback()
		if e != nil {
			return e
		} else {
			return e2
		}
	}
}

func (service *ContactService) CreateCommunity(comm model.Community) (ret model.Community, err error) {
	if len(comm.Name) == 0 {
		err = errors.New("缺少群名称")
		return ret, err
	}
	if comm.OwnerId == 0 {
		err = errors.New("请登录")
		return ret, err
	}
	com := model.Community{OwnerId: comm.OwnerId}
	num, err := DbEngin.Count(&com)
	if num > 5 {
		err = errors.New("用户只能建5个群")
		return com, err
	}

	session := DbEngin.NewSession()
	session.Begin()

	comm.CreateAt = time.Now()
	_, err = session.InsertOne(&comm)
	if err != nil {
		session.Rollback()
		return com, err
	}
	_, err = session.InsertOne(model.Contact{OwnerId: comm.OwnerId, DstObj: comm.Id, Cate: model.CONTACT_CATE_COMUNITY, CreateAt: time.Now()})
	if err != nil {
		session.Rollback()
	} else {
		session.Commit()
	}
	return com, err
}

func (service *ContactService) JoinCommunity(userId, comId int64) error {
	cot := model.Contact{OwnerId: userId, DstObj: comId, Cate: model.CONTACT_CATE_COMUNITY}
	DbEngin.Get(&cot)
	if cot.Id == 0 { //不在群里面
		cot.CreateAt = time.Now()
		_, err := DbEngin.InsertOne(cot)
		return err
	} else {
		return nil
	}
}

func (service *ContactService) SearchComunity(userId int64) []model.Community {
	contacts := make([]model.Contact, 0)
	comIds := make([]int64, 0)
	DbEngin.Where("owner_id = ? and cate = ?", userId, model.CONTACT_CATE_COMUNITY).Find(&contacts)
	for _, v := range contacts {
		comIds = append(comIds, v.DstObj)
	}
	coms := make([]model.Community, 0)
	if len(comIds) == 0 {
		return coms
	}
	DbEngin.In("id", comIds).Find(&coms) //查找群信息
	return coms
}

func (service *ContactService) SearchComunityIds(userId int64) (comIds []int64) {
	contacts := make([]model.Contact, 0)
	comIds = make([]int64, 0)
	DbEngin.Where("owner_id = ? and cate = ?", userId, model.CONTACT_CATE_COMUNITY).Find(&contacts)
	for _, v := range contacts {
		comIds = append(comIds, v.DstObj)
	}
	return comIds
}

func (service *ContactService) SearchFriend(userId int64) []model.User {
	contacts := make([]model.Contact, 0)
	objIds := make([]int64, 0)
	DbEngin.Where("owner_id = ? and cate = ?", userId, model.CONTACT_CATE_USER).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, v.DstObj)
	}
	coms := make([]model.User, 0)
	if len(objIds) == 0 {
		return coms
	}
	DbEngin.In("id", objIds).Find(&coms)
	return coms
}
