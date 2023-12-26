package repositories

import (
  "github.com/vietstars/postgres-api/dto"
  "github.com/vietstars/postgres-api/models"
  "golang.org/x/crypto/bcrypt"
)

//db.Unscoped().Where("id = 2").Find(&users)
//
// func (u *User) AfterFind(tx *gorm.DB) (err error) {
//   if u.Version == 0 {
//     u.Version = 1
//   }
//   return
// }

func NewUser(new dto.UserNew) (user *models.UserEntity, err error) {
  user = &models.UserEntity{
    UserName: new.UserName,
    Email: new.Email,
    Password: new.Password,
  }

  tx := models.DB.Table("users").Begin()

  if err := tx.Create(&user).Error; err != nil {
    tx.Rollback()
    return nil, err
  }

  tx.Commit()

  return user, nil
}


func GetAllUsers() (users *models.UserListEntity, err error) {
  if err = models.DB.Find(&users).Error; err != nil{

    return nil, err
  }

  return users, nil
}

func GetUserById(id uint) (user *models.UserEntity, err error) {
  if err = models.DB.First(&user, id).Error; err != nil{

    return user, nil
  }

  return nil, err
}

func GetUserByEmail(email string) (user *models.UserEntity, err error) {
  err = models.DB.First(&user, "email = ?", email).Error

  return user, err
}


func UpdateUserById(id uint, edit dto.UserEdit) (user *models.UserEntity, err error) {
  tx := models.DB.Table("users").Begin()

  if err := tx.Error; err != nil {

      return nil, err
  }

  if err := tx.Where("id = ? And version = ?", id, edit.Version).First(&user).Error; err != nil{
    tx.Rollback()

    return nil, err
  }

  user.UserName = edit.UserName
  user.Email = edit.Email
  if hash, err := bcrypt.GenerateFromPassword([]byte(edit.Password), 10);err == nil {
    user.Password = string(hash)
  }

  tx.Save(&user)
  tx.Commit()

  return user, nil
}
