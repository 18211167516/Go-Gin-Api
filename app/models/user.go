package models

type User struct {
    Model

    Name string `json:"name"`
    CreatedBy string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State int `json:"state"`
}

func GetUsers(pageNum int, pageSize int, maps interface {}) (users []User) {
    db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users)
    
    return
}

func GetUserTotal(maps interface {}) (count int){
    db.Model(&User{}).Where(maps).Count(&count)

    return
}

func ExistUserByMaps(maps interface{}) bool {
    var user User
    db.Select("id").Where(maps).First(&user)
    if user.ID > 0 {
        return true
    }

    return false
}

func AddUser(Users map[string]interface{}) bool{
    user := User {
        Name : Users["Name"].(string),
        State : Users["State"].(int),
        CreatedBy : Users["CreatedBy"].(string),
    }
    db.Create(&user)
    return !db.NewRecord(user)
}

func ExistTagByID(id int) bool {
    var user User
    db.Select("id").Where("id = ?", id).First(&user)
    if user.ID > 0 {
        return true
    }

    return false
}

func DeleteUser(maps interface{}) bool {
    db.Where(maps).Delete(&User{})

    return true
}

func EditUser(id int, data interface {}) bool {
    db.Model(&User{}).Where("id = ?", id).Updates(data)

    return true
}