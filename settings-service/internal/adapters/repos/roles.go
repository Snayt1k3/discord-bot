package repos

import (
	"gorm.io/gorm"
	"settings-service/internal/models"
)

type ReactionRolesRepoImpl struct {
	db *gorm.DB
}

func NewReactionRolesRepo(db *gorm.DB) *ReactionRolesRepoImpl {
	return &ReactionRolesRepoImpl{
		db: db,
	}
}

func (r *ReactionRolesRepoImpl) SetRoleMessageId(guildId, messageId string) error {
	return r.db.Model(&models.RolesSettings{}).
		Where("guild_id = ?", guildId).
		Update("message_id", messageId).Error
}

func (r *ReactionRolesRepoImpl) AddRole(guildId, roleId, emoji string) (models.Role, error) {
	var roleSetting models.RolesSettings

	if err := r.db.Where("guild_id = ?", guildId).First(&roleSetting).Error; err != nil {
		return models.Role{}, err
	}

	role := models.Role{
		RoleID: roleId,
		Emoji:  emoji,
	}

	if err := r.db.Create(&role).Error; err != nil {
		return models.Role{}, err
	}

	if err := r.db.Model(&roleSetting).Association("Roles").Append(&role); err != nil {
		return models.Role{}, err
	}

	return role, nil
}

func (r *ReactionRolesRepoImpl) DeleteRole(guildId, roleId string) error {
	var roleSetting models.RolesSettings

	if err := r.db.Where("guild_id = ?", guildId).First(&roleSetting).Error; err != nil {
		return err
	}

	var role models.Role
	if err := r.db.Where("role_id = ?", roleId).First(&role).Error; err != nil {
		return err
	}

	if err := r.db.Model(&roleSetting).Association("Roles").Delete(&role); err != nil {
		return err
	}

	var count int64
	r.db.Model(&models.RolesSettings{}).Where("id IN (SELECT roles_settings_id FROM roles_settings_roles WHERE role_id = ?)", role.ID).Count(&count)

	if count == 0 {
		return r.db.Delete(&role).Error
	}

	return nil
}
