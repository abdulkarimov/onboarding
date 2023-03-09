package repositories


import (
    "github.com/gofiber/fiber/v2"
    "github.com/abdulkarimov/onboarding/database"
    "github.com/abdulkarimov/onboarding/models"
    gf "github.com/shareed2k/goth_fiber"
    "fmt"
    "net/http"
    "net/url"
)

func RegisterUser(c *fiber.Ctx) error {
   return c.Render("anketa", fiber.Map{
            "Email":  c.Params("email"),
    })  
}

func ShowIMG(c *fiber.Ctx) error {
    return c.Render("IMG",fiber.Map{}) 
 }

func GetUsers(c *fiber.Ctx) error {

    user := []models.User{}
    database.DB.Db.Preload("Status").Find(&user)

    return c.Status(200).JSON(user)
}

func AddUser(c *fiber.Ctx) error {

    user := new(models.User)
    c.BodyParser(user)
    database.DB.Db.Create(user)
    
    return c.Status(200).JSON(user)
}

func UpdateUser( c *fiber.Ctx) error {

    oldUser := models.User{}
    user := models.User{}
    c.BodyParser(&user)

    result :=  database.DB.Db.First(&oldUser, c.Params("id")).Updates(user)

    if result.RowsAffected != 0 {
        return c.Status(200).JSON(oldUser)
    }

    return c.SendStatus(204)
}

func DeleteUser( c *fiber.Ctx) error {

    result := database.DB.Db.Delete(&models.User{}, c.Params("id"))

    if result.RowsAffected != 0 {
        return c.SendStatus(200)
    }
    
    return c.SendStatus(204)
}


func GoogleCallback( c *fiber.Ctx) error {

    user, err := gf.CompleteUserAuth(c)

    if err != nil {
        return err
    }

    test := models.User{}
    result := database.DB.Db.Where("email = ?",user.Email).First(&test)

    if result.RowsAffected == 0 {

        database.DB.Db.Create(&models.User{
            Email : user.Email,
            StatusId : 2,
        })
    } else if test.StatusId == 2 {

        data := url.Values{
            "chat_id": {"861921150"},
            "text": {user.Email + " Ждет вашего одобрения."},
        }
        http.PostForm("https://api.telegram.org/bot6140313805:AAEO0wiwir08eYD4VcZAuxRKnc6PBwMJzs4/sendMessage", data)
    } else if test.StatusId == 1 {

        fmt.Println("next")

        c.Redirect("/api/registerUser/"+user.Email)

     
    }
   
    return nil
}

func GoogleAuth(c *fiber.Ctx) error {
    if gothUser, err := gf.CompleteUserAuth(c); err == nil {
        c.JSON(gothUser)
    } else {
        gf.BeginAuthHandler(c)
    }
    return nil
}