package dao

import (
	"database/sql"
	"go-IM/logic/db"
	"go-IM/logic/model"
)

type groupDao struct{}

var GroupDao = new(groupDao)

// 获取群组信息
func (*groupDao) Get(groupId int64) (*model.Group, error) {
	row := db.Cli.QueryRow("SELECT name,avatar_url,introduction,user_num,type,extra,create_time,update_time FROM `group` WHERE group_id = ?", groupId)
	group := model.Group{
		GroupId: groupId,
	}
	err := row.Scan(&group.Name, &group.AvatarUrl, &group.Introduction, &group.UserNum, &group.Type, &group.Extra, &group.CreateTime, &group.UpdateTime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &group, nil
}

// 插入一条群组
func (*groupDao) Add(group *model.Group) (int64, error) {
	result, err := db.Cli.Exec("INSERT IGNORE INTO `group`(app_id,group_id,name,avatar_url,introduction,type,extra) VALUES (?,?,?,?,?,?,?)",
		group.AppId, group.GroupId, group.Name, group.AvatarUrl, group.Introduction, group.Type, group.Extra)
	if err != nil {
		return 0, err
	}
	num, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return num, nil
}

// 更新群组信息
func (*groupDao) Update(appId, groupId int64, name, avatarUrl, introduction, extra string) error {
	_, err := db.Cli.Exec("UPDATE `group` SET name = ?, avatar_url = ?, introduction = ?,extra = ? WHERE app_id = ? AND group_id = ?",
		name, avatarUrl, introduction, extra, appId, groupId)
	return err
}

// 更新群组信息
func (*groupDao) AddUserNum(appId, groupId int64, userNum int) error {
	_, err := db.Cli.Exec("UPDATE `group` SET user_num = user_num + ? WHERE app_id = ? AND group_id = ?",
		userNum, appId, groupId)
	return err
}

// 更新群组群成员人数
func (*groupDao) UpdateUserNum(groupId, userNum int64) error {
	_, err := db.Cli.Exec("UPDATE `group` SET user_num = user_num + ? WHERE group_id = ?",
		userNum, groupId)
	return err
}
