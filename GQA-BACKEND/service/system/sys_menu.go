package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
)

type ServiceMenu struct {
}

func (s *ServiceMenu) GetMenuList(pageInfo global.RequestPage) (err error, menu interface{}, total int64) {
	pageSize := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GqaDb.Model(&system.SysMenu{})
	var menuList []system.SysMenu
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(pageInfo.SortBy, pageInfo.Desc)).Find(&menuList).Error
	return err, menuList, total
}

func (s *ServiceMenu) EditMenu(toEditMenu system.SysMenu) (err error) {
	var sysMenu system.SysMenu
	if err = global.GqaDb.Where("id = ?", toEditMenu.Id).First(&sysMenu).Error; err != nil {
		return err
	}
	if sysMenu.Stable == "yes" {
		return errors.New("系统内置不允许编辑：" + toEditMenu.Title)
	}
	err = global.GqaDb.Updates(&toEditMenu).Error
	return err
}

func (s *ServiceMenu) AddMenu(toAddMenu system.SysMenu) (err error) {
	err = global.GqaDb.Create(&toAddMenu).Error
	return err
}

func (s *ServiceMenu) DeleteMenu(id uint) (err error) {
	var sysMenu system.SysMenu
	if err = global.GqaDb.Where("id = ?", id).First(&sysMenu).Error; err != nil {
		return err
	}
	if sysMenu.Stable == "yes" {
		return errors.New("系统内置不允许删除！" + sysMenu.Title)
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysMenu).Error
	return err
}

func (s *ServiceMenu) QueryMenuById(id uint) (err error, menuInfo system.SysMenu) {
	var menu system.SysMenu
	err = global.GqaDb.First(&menu, "id = ?", id).Error
	return err, menu
}
