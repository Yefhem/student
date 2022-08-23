package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Yefhem/student-syllabus/helpers"
	"github.com/Yefhem/student-syllabus/model"
	"github.com/Yefhem/student-syllabus/service"
	"github.com/labstack/echo/v4"
)

type TaskController interface {
	DashboardPage(c echo.Context) error
	NewTaskPage(c echo.Context) error
	TasksPage(c echo.Context) error
	EditTaskPage(c echo.Context) error

	TaskAdd(c echo.Context) error
	TaskUpdate(c echo.Context) error
	TaskDelete(c echo.Context) error
	Status(c echo.Context) error
}

type taskController struct {
	taskService    service.TaskService
	alertService   service.AlertService
	sessionService service.SessionService
}

func NewTaskController(taskService service.TaskService, alertService service.AlertService, sessionService service.SessionService) TaskController {
	return &taskController{
		taskService:    taskService,
		alertService:   alertService,
		sessionService: sessionService,
	}
}

var dateNumber = model.DateNumber{
	Day:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
	Month: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	Year:  []int{2022, 2023, 2024, 2025, 2026, 2027, 2028, 2029, 2030},
	Hour:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
}

var taskMap = make(map[string]interface{})

// --------------------> Pages

// ----------> Dashboard Page
func (cont *taskController) DashboardPage(c echo.Context) error {
	if !cont.sessionService.CheckUser(c) {
		cont.alertService.SetAlert(c, "Lütfen Giriş Yapınız!")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	results, err := helpers.Include("pages/dashboard")
	if err != nil {
		log.Fatal(err)
	}
	view, err := template.ParseFiles(results...)
	if err != nil {
		fmt.Println(err)
	}

	totalTasks, err := cont.taskService.TotalTask()
	if err != nil {
		log.Println(err)
	}

	lastTask, err := cont.taskService.LastTask()
	if err != nil {
		log.Println(err)
	}

	taskMap["TotalTasks"] = totalTasks
	taskMap["LastTask"] = lastTask
	taskMap["Alert"] = cont.alertService.GetAlert(c)

	return view.ExecuteTemplate(c.Response(), "index", taskMap)
}

// ----------> Tasks Page
func (cont *taskController) TasksPage(c echo.Context) error {

	if !cont.sessionService.CheckUser(c) {
		cont.alertService.SetAlert(c, "Lütfen Giriş Yapınız!")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	results, err := helpers.Include("pages/tasks")
	if err != nil {
		log.Fatal(err)
	}
	view, err := template.ParseFiles(results...)
	if err != nil {
		fmt.Println(err)
	}

	result, err := cont.taskService.GetAllTasks()
	if err != nil {
		log.Println(err)
	}

	taskMap["Tasks"] = result
	taskMap["Alert"] = cont.alertService.GetAlert(c)

	return view.ExecuteTemplate(c.Response(), "index", taskMap)
}

// ----------> New Task Page
func (cont *taskController) NewTaskPage(c echo.Context) error {

	if !cont.sessionService.CheckUser(c) {
		cont.alertService.SetAlert(c, "Lütfen Giriş Yapınız!")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	results, err := helpers.Include("pages/taskadd")
	if err != nil {
		log.Fatal(err)
	}
	view, err := template.ParseFiles(results...)
	if err != nil {
		fmt.Println(err)
	}

	log.Println(helpers.CurrentTime())

	taskMap["DateNumber"] = dateNumber
	taskMap["CurrentTime"] = helpers.CurrentTime()
	taskMap["Alert"] = cont.alertService.GetAlert(c)

	return view.ExecuteTemplate(c.Response(), "index", taskMap)
}

// ----------> Edit Task Page
func (cont *taskController) EditTaskPage(c echo.Context) error {

	if !cont.sessionService.CheckUser(c) {
		cont.alertService.SetAlert(c, "Lütfen Giriş Yapınız!")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	results, err := helpers.Include("pages/taskedit")
	if err != nil {
		log.Fatal(err)
	}
	view, err := template.ParseFiles(results...)
	if err != nil {
		fmt.Println(err)
	}

	id := c.Param("id")

	result, err := cont.taskService.GetTask(id)
	if err != nil {
		log.Println(err)
	}

	log.Println(result)

	taskMap["DateNumber"] = dateNumber
	taskMap["Task"] = result
	taskMap["Alert"] = cont.alertService.GetAlert(c)

	return view.ExecuteTemplate(c.Response(), "index", taskMap)
}

// --------------------> Methods

func (cont *taskController) TaskAdd(c echo.Context) error {

	if !cont.sessionService.CheckUser(c) {
		cont.alertService.SetAlert(c, "Lütfen Giriş Yapınız!")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	year, _ := strconv.Atoi(c.FormValue("year"))
	month, _ := strconv.Atoi(c.FormValue("month"))
	day, _ := strconv.Atoi(c.FormValue("day"))
	starthour, _ := strconv.Atoi(c.FormValue("starthour"))
	endhour, _ := strconv.Atoi(c.FormValue("endhour"))

	date := model.Date{
		Year:      year,
		Month:     month,
		Day:       day,
		StartHour: starthour,
		EndHour:   endhour,
	}

	taskDTO := model.TaskDTO{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Date:        date,
		Rate:        c.FormValue("rate"),
	}

	if err := cont.taskService.CreateTask(taskDTO); err != nil {
		log.Println(err)
	}

	return c.Redirect(http.StatusSeeOther, "/tasks")
}

func (cont *taskController) TaskUpdate(c echo.Context) error {

	if !cont.sessionService.CheckUser(c) {
		cont.alertService.SetAlert(c, "Lütfen Giriş Yapınız!")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	id := c.Param("id")

	year, _ := strconv.Atoi(c.FormValue("year"))
	month, _ := strconv.Atoi(c.FormValue("month"))
	day, _ := strconv.Atoi(c.FormValue("day"))
	starthour, _ := strconv.Atoi(c.FormValue("starthour"))
	endhour, _ := strconv.Atoi(c.FormValue("endhour"))

	date := model.Date{
		Year:      year,
		Month:     month,
		Day:       day,
		StartHour: starthour,
		EndHour:   endhour,
	}

	taskDTO := model.TaskDTO{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Date:        date,
		Rate:        c.FormValue("rate"),
	}

	if err := cont.taskService.UpdateTask(taskDTO, id); err != nil {
		log.Println(err)
	}

	return c.Redirect(http.StatusSeeOther, "/tasks")
}

func (cont *taskController) TaskDelete(c echo.Context) error {

	if !cont.sessionService.CheckUser(c) {
		cont.alertService.SetAlert(c, "Lütfen Giriş Yapınız!")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	id := c.Param("id")

	if err := cont.taskService.DeleteTask(id); err != nil {
		log.Println(err)
		cont.alertService.SetAlert(c, "Silme İşlemi Başarısız! Lütfen Tekrar Deneyin")
		return c.Redirect(http.StatusSeeOther, "/tasks")
	}

	cont.alertService.SetAlert(c, "Silme İşlemi Başarılı")
	return c.Redirect(http.StatusSeeOther, "/tasks")
}

func (cont *taskController) Status(c echo.Context) error {

	if !cont.sessionService.CheckUser(c) {
		cont.alertService.SetAlert(c, "Lütfen Giriş Yapınız!")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	id := c.Param("id")
	status := c.Param("status")

	if err := cont.taskService.StatusService(status, id); err != nil {
		log.Println(err)
		return c.Redirect(http.StatusSeeOther, "/tasks")
	}

	return c.Redirect(http.StatusSeeOther, "/tasks")

}
