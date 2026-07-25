package golang_gorm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	dialect := mysql.Open("root:password@tcp(localhost:3306)/belajar_golang_gorm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("insert into sample(id, name) values (?, ?)", "1", "Eko").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values (?, ?)", "2", "Budi").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values (?, ?)", "3", "Joko").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values (?, ?)", "4", "Rully").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id   string
	Name string
}

func TestRawSQL(t *testing.T) {
	var sample Sample
	err := db.Raw("select id, name from sample where id = ?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "Eko", sample.Name)

	var samples []Sample
	err = db.Raw("select id, name from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}

// Cara ribet
func TestSqlRow(t *testing.T) {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		var id string
		var name string

		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			Id:   id,
			Name: name,
		})
	}
	assert.Equal(t, 4, len(samples))
}

// Cara simpel
func TestScanRow(t *testing.T) {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}
	assert.Equal(t, 4, len(samples))
}

// Create
func TestCreateUser(t *testing.T) {
	user := User{
		ID:       "1",
		Password: "rahasia",
		Name: Name{
			FirstName:  "Eko",
			MiddleName: "Kurniawan",
			LastName:   "Khannedy",
		},
		Information: "ini akan di ignore",
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 2; i < 10; i++ { // karena sudah ada data pertama/ke 1, jadi mulai dari 2 sampai 9, total 8 data
		users = append(users, User{
			ID:       strconv.Itoa(i),
			Password: "rahasia",
			Name: Name{
				FirstName: "User " + strconv.Itoa(i),
			},
		})
	}

	result := db.Create(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 8, int(result.RowsAffected))
}

// Transaction
// - Kita tidak perlu Begin, Commit, Rollback, karena gorm sudah menghandle itu semua
func TestTransactionSuccess(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error { // Dijamin hanya menggunakan 1 koneksi database
		err := tx.Create(&User{ID: "10", Password: "rahasia", Name: Name{FirstName: "User 10"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{ID: "11", Password: "rahasia", Name: Name{FirstName: "User 11"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{ID: "12", Password: "rahasia", Name: Name{FirstName: "User 12"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

func TestTransactionRollback(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{ID: "13", Password: "rahasia", Name: Name{FirstName: "User 13"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{ID: "11", Password: "rahasia", Name: Name{FirstName: "User 11"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.NotNil(t, err)
}

func TestManualTransactionSuccess(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "13", Password: "rahasia", Name: Name{FirstName: "User 13"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{ID: "14", Password: "rahasia", Name: Name{FirstName: "User 14"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestManualTransactionError(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "15", Password: "rahasia", Name: Name{FirstName: "User 15"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{ID: "14", Password: "rahasia", Name: Name{FirstName: "User 14"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

// Query
func TestQuerySingleObject(t *testing.T) {
	user := User{}
	err := db.First(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.ID)

	user = User{}
	err = db.Last(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "9", user.ID)
}

func TestQuerySingleObjectInlineCondition(t *testing.T) {
	user := User{}
	err := db.Take(&user, "id = ?", "5").Error
	assert.Nil(t, err)
	assert.Equal(t, "5", user.ID)
	assert.Equal(t, "User 5", user.Name.FirstName)
}

func TestQueryAllObjects(t *testing.T) {
	var users []User
	err := db.Find(&users, "id in ?", []string{"1", "2", "3", "4"}).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))

	for _, user := range users {
		fmt.Printf("User ID: %s | First Name: %s | Password: %s\n", user.ID, user.Name.FirstName, user.Password)
	}
}

// Advance Query
func TestQueryCondition(t *testing.T) {
	var users []User
	err := db.Where("first_name like ?", "%User%").Where("password = ?", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 13, len(users))

	for _, user := range users {
		fmt.Printf("User ID: %s | First Name: %s | Password: %s\n", user.ID, user.Name.FirstName, user.Password)
	}
}

func TestOrOperator(t *testing.T) {
	var users []User
	err := db.Where("first_name like ?", "%User%").Or("password = ?", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 14, len(users))

	for _, user := range users {
		fmt.Printf("User ID: %s | First Name: %s | Password: %s\n", user.ID, user.Name.FirstName, user.Password)
	}
}

func TestNotOperator(t *testing.T) {
	var users []User
	err := db.Not("first_name like ?", "%User%").Where("password = ?", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))

	for _, user := range users {
		fmt.Printf("User ID: %s | First Name: %s | Password: %s\n", user.ID, user.Name.FirstName, user.Password)
	}
}

func TestSelectFields(t *testing.T) {
	var users []User
	err := db.Select("id", "first_name").Find(&users).Error
	assert.Nil(t, err)

	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)

		fmt.Printf("User ID: %s | First Name: %s\n", user.ID, user.Name.FirstName)
	}

	assert.Equal(t, 14, len(users))
}

func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "User 5",
			LastName:  "", // tidak bisa, karena dianggap default value
		},
		Password: "rahasia",
	}

	var users []User
	err := db.Where(userCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))

	for _, user := range users {
		fmt.Printf("User ID: %s | First Name: %s | Password: %s\n", user.ID, user.Name.FirstName, user.Password)
	}
}

func TestMapCondition(t *testing.T) {
	mapCondition := map[string]interface{}{
		"middle_name": "",
		"last_name":   "",
	}

	var users []User
	err := db.Where(mapCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 13, len(users))

	for _, user := range users {
		fmt.Printf("User ID: %s | First Name: %s | Password: %s\n", user.ID, user.Name.FirstName, user.Password)
	}
}

func TestOrderLimitOffset(t *testing.T) {
	var users []User
	err := db.Order("id asc, first_name desc").Limit(5).Offset(5).Find(&users).Error // biasanya digunakan untuk pagination
	assert.Nil(t, err)
	assert.Equal(t, 5, len(users))

	for _, user := range users {
		fmt.Printf("User ID: %s | First Name: %s | Password: %s\n", user.ID, user.Name.FirstName, user.Password)
	}
}

// Query Non Model
type UserResponse struct {
	ID        string
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse
	err := db.Model(&User{}).Select("id", "first_name", "last_name").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 14, len(users))

	for _, user := range users {
		fmt.Printf("User ID: %s | First Name: %s | Last Name: %s\n", user.ID, user.FirstName, user.LastName)
	}
}

// Update
func TestUpdate(t *testing.T) {
	user := User{}
	err := db.Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)

	user.Name.FirstName = "Budi"
	user.Name.MiddleName = ""
	user.Name.LastName = "Nugraha"
	user.Password = "rahasia123"

	err = db.Save(&user).Error
	assert.Nil(t, err)
}

func TestUpdateSelectedColumns(t *testing.T) {
	err := db.Model(&User{}).Where("id = ?", "1").Updates(map[string]interface{}{
		"middle_name": "",
		"last_name":   "Morro",
	}).Error
	assert.Nil(t, err)

	err = db.Model(&User{}).Where("id = ?", "1").Update("password", "diubahlagi").Error
	assert.Nil(t, err)

	err = db.Where("id = ?", "1").Updates(User{
		Name: Name{
			FirstName: "Eko",
			LastName:  "Khannedy",
		},
	}).Error
	assert.Nil(t, err)
}

// Auto Increment
func TestAutoIncrement(t *testing.T) {
	for i := range 10 {
		userLog := UserLog{
			UserId: "1",
			Action: "Test Action",
		}

		err := db.Create(&userLog).Error
		assert.Nil(t, err)

		assert.NotEqual(t, 0, userLog.ID)
		fmt.Println(userLog.ID)
		_ = i
	}
}

// Upsert
func TestSaveOrUpdate(t *testing.T) { // cocok untuk kasus auto increment
	userLog := UserLog{
		UserId: "1",
		Action: "Test Action",
	}

	err := db.Save(&userLog).Error // insert
	assert.Nil(t, err)

	userLog.UserId = "2"
	err = db.Save(&userLog).Error // update
	assert.Nil(t, err)
}

func TestSaveOrUpdateNonAutoIncrement(t *testing.T) {
	user := User{
		ID: "100",
		Name: Name{
			FirstName: "User 100",
		},
	}

	err := db.Save(&user).Error // insert
	assert.Nil(t, err)

	user.Name.FirstName = "User 100 Updated"
	err = db.Save(&user).Error // update
	assert.Nil(t, err)
}

func TestConflict(t *testing.T) {
	user := User{
		ID: "88",
		Name: Name{
			FirstName: "User 88",
		},
	}

	err := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&user).Error // insert
	assert.Nil(t, err)
}

// Delete
func TestDelete(t *testing.T) {
	var user User
	err := db.Take(&user, "id = ?", "88").Error
	assert.Nil(t, err)

	err = db.Delete(&user).Error
	assert.Nil(t, err)

	err = db.Delete(&User{}, "id = ?", "99").Error
	assert.Nil(t, err)

	err = db.Where("id = ?", "100").Delete(&User{}).Error
	assert.Nil(t, err)
}

// Soft Delete
func TestSoftDelete(t *testing.T) {
	todo := Todo{
		UserId:      "1",
		Title:       "Todo 1",
		Description: "Description 1",
	}
	err := db.Create(&todo).Error
	assert.Nil(t, err)

	err = db.Delete(&todo).Error
	assert.Nil(t, err)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	err = db.Find(&todos).Error
	assert.Nil(t, err)
	assert.Equal(t, 0, len(todos))
}

func TestUnscoped(t *testing.T) {
	var todo Todo
	err := db.Unscoped().First(&todo, "id = ?", 3).Error
	assert.Nil(t, err)
	fmt.Println(todo)

	err = db.Unscoped().Delete(&todo).Error
	assert.Nil(t, err)

	var todos []Todo
	err = db.Unscoped().Find(&todos).Error
	assert.Nil(t, err)
}

// Lock
func TestLock(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&user, "id = ?", "1").Error
		if err != nil {
			return err
		}

		user.Name.FirstName = "Joko"
		user.Name.LastName = "Morro"
		err = tx.Save(&user).Error
		return err
	})
	assert.Nil(t, err)
}

// One to One
func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		ID:      "1",
		UserId:  "1",
		Balance: 1000000,
	}

	err := db.Create(&wallet).Error
	assert.Nil(t, err)
}

func TestRetrieveRelation(t *testing.T) {
	var user User
	err := db.Model(&User{}).Preload("Wallet").Take(&user, "id = ?", "1").Error // bersifat EAGER, .Preload("NamaField")
	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
}

func TestRetrieveRelationJoin(t *testing.T) {
	var user User
	err := db.Model(&User{}).Joins("Wallet").Take(&user, "users.id = ?", "1").Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)

	fmt.Printf("%+v\n", user)
}

// Auto Upsert Relation
func TestAutoCreateUpdate(t *testing.T) {
	user := User{
		ID:       "20",
		Password: "rahasia",
		Name: Name{
			FirstName: "User 20",
		},
		Wallet: Wallet{
			ID:      "20",
			UserId:  "20",
			Balance: 1000000,
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
}

func TestSkipAutoCreateUpdate(t *testing.T) {
	user := User{
		ID:       "21",
		Password: "rahasia",
		Name: Name{
			FirstName: "User 21",
		},
		Wallet: Wallet{
			ID:      "21",
			UserId:  "21",
			Balance: 1000000,
		},
	}

	err := db.Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)
}

// One to Many
func TestUserAndAddresses(t *testing.T) {
	user := User{
		ID:       "2",
		Password: "rahasia",
		Name: Name{
			FirstName: "User 50",
		},
		Wallet: Wallet{
			ID:      "2",
			UserId:  "2",
			Balance: 1000000,
		},
		Addresses: []Address{
			{
				UserId:  "2",
				Address: "Jalan A",
			},
			{
				UserId:  "2",
				Address: "Jalan B",
			},
		},
	}

	err := db.Save(&user).Error
	assert.Nil(t, err)
}

func TestUserAndAddresses2(t *testing.T) {
	// 1. Buat user dulu
	user := User{
		ID:       "60",
		Password: "rahasia",
		Name:     Name{FirstName: "User 60"},
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)

	// 2. Buat wallet
	wallet := Wallet{
		ID:      "60",
		UserId:  user.ID,
		Balance: 1000000,
	}
	err = db.Create(&wallet).Error
	assert.Nil(t, err)

	// 3. Batch insert addresses (tanpa loop!)
	addresses := []Address{
		{UserId: user.ID, Address: "Jalan A"},
		{UserId: user.ID, Address: "Jalan B"},
	}
	err = db.Create(&addresses).Error // Batch insert
	assert.Nil(t, err)
}

func TestPreloadJoinOneToMany(t *testing.T) {
	var users []User
	err := db.Model(&User{}).Preload("Addresses").Joins("Wallet").Find(&users).Error
	assert.Nil(t, err)
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}

func TestTakePreloadJoinOneToMany(t *testing.T) {
	var user User
	err := db.Model(&User{}).Preload("Addresses").Joins("Wallet").
		Take(&user, "users.id = ?", "2").Error
	assert.Nil(t, err)
	fmt.Println("User :", user)
}

// Belongs To / Many to One
func TestBelongsTo(t *testing.T) {
	// 1. Select all addresses, setelah dapa semua address,
	// 2. Select all addresses, lalu ambil user_id dari address, lalu select user by id, lalu set ke address.User
	fmt.Println("Preload")
	var addresses []Address
	err := db.Model(&Address{}).Preload("User").Find(&addresses).Error
	assert.Nil(t, err)
	assert.Equal(t, 6, len(addresses))
	for _, address := range addresses {
		fmt.Printf("%+v\n\n", address)
	}

	fmt.Println("Joins")
	addresses = []Address{}
	err = db.Model(&Address{}).Joins("User").Find(&addresses).Error
	assert.Nil(t, err)
	assert.Equal(t, 6, len(addresses))
	for _, address := range addresses {
		fmt.Printf("%+v\n\n", address)
	}
}

func TestBelongsToWallet(t *testing.T) {
	fmt.Println("Preload")
	var wallets []Wallet
	err := db.Model(&Wallet{}).Preload("User").Find(&wallets).Error
	assert.Nil(t, err)

	fmt.Println("Joins")
	wallets = []Wallet{}
	err = db.Model(&Wallet{}).Joins("User").Find(&wallets).Error
	assert.Nil(t, err)
}
