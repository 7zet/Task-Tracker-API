package internal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	input.ID = len(Users) + 1
	Users = append(Users, input)
	if err := SaveToFile(); err != nil {
		c.JSON(500, gin.H{
			"error":"Faylga yozib bolmadi!",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Registered",
		"id":      input.ID,
	})
}

func Login(c *gin.Context) {
	var input User

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var foundUser *User
	found := false

	for _, u := range Users {

		if u.Username == input.Username && u.Password == input.Password {
			found = true
			foundUser = &u
			break

		}
	}
	if !found {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Username yoki Password Xato!!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tizimga muvafaqqiyatli kirdingiz",
		"data": gin.H{
			"id":       foundUser.ID,
			"username": foundUser.Username,
		},
	})

}

func CreateTask(c *gin.Context) {
	// TASK structi bilan yaxshiroq ishlay olishimiz uchun yangi ozgaruvchi yaratib olamiz
	var newTask Task

	// 1. agar bizga kelgan POST json bolmasa yoki jsonda xatolik bolsa ishlaydigan IF
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		// IF xato berganda toxtashi uchun return
		return
	}

	// ID'ni bittaga qoshib boramiz
	newTask.ID = len(Tasks) + 1
	// newTaskni Tasks Slicega qoshib qoyamiz
	Tasks = append(Tasks, newTask)

	// hamma jarayon togri bolganda
	c.JSON(http.StatusCreated, gin.H{
		"message": "Task Created",
		"data":    newTask,
	})
}

func GetTasks(c *gin.Context) {
	// 1. Urldan "user_id"ni string korinishiga olib kelamiz
	// misol: /tasks?user_id=5
	idParam := c.Query("user_id")

	// 2. Stringni INTga ogiramiz
	userID, err := strconv.Atoi(idParam)

	// 3. Agar foydalanuvchi son orniga harf yozgan bolsa yoki
	// bosh bolsa
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user_id faqat son bolishi kerak",
		})
		// return berib funksiyani toxtatamiz
		return
	}

	// 4. endi userID ozgaruvchisi bilan xuddi sondek ishlay olamiz
	var userTasks []Task
	for _, t := range Tasks {
		if t.UserID == userID { //endi ikkalasi ham int!
			userTasks = append(userTasks, t)
		}
	}

	c.JSON(http.StatusOK, userTasks)
}
