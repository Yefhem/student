package main

import (
	"github.com/Yefhem/student-syllabus/configs"
	"github.com/Yefhem/student-syllabus/repository"
	"github.com/Yefhem/student-syllabus/service"
	"github.com/Yefhem/student-syllabus/site/controller"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	// ----------> DB Connection
	db *gorm.DB = configs.InitDB()

	// ----------> Store
	store *sessions.CookieStore = sessions.NewCookieStore([]byte("secret"))

	// ----------> Repository Layer
	userRepo repository.UserRepository = repository.NewUserRepository(db)
	taskRepo repository.TaskRepository = repository.NewTaskRepository(db)

	// ----------> Service Layer
	userService    service.UserService    = service.NewUserService(userRepo)
	taskService    service.TaskService    = service.NewTaskService(taskRepo)
	sessionService service.SessionService = service.NewSessionService(userRepo, store)
	alertService   service.AlertService   = service.NewAlertService(store)

	// ----------> Controller Layer
	authController controller.AuthController = controller.NewAuthController(userService, sessionService, alertService)
	taskController controller.TaskController = controller.NewTaskController(taskService, alertService, sessionService)
)

func main() {
	e := echo.New()

	e.Static("/site/assets", "site/assets")

	e.GET("/login", authController.LoginPage)
	e.GET("/register", authController.RegisterPage)

	e.POST("/user-login", authController.Login)

	e.GET("/dashboard", taskController.DashboardPage)
	e.GET("/tasks", taskController.TasksPage)
	e.GET("/new-task", taskController.NewTaskPage)
	e.GET("/edit-task/:id", taskController.EditTaskPage)

	// ----------> Methods
	e.POST("/create-account", authController.CreateAccount)
	e.POST("/add-new-task", taskController.TaskAdd)
	e.POST("/update-task/:id", taskController.TaskUpdate)
	e.GET("/delete-task/:id", taskController.TaskDelete)

	e.GET("/status/:status/:id", taskController.Status)

	e.Logger.Fatal(e.Start(":8080"))

}
