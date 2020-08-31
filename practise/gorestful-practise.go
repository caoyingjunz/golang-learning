package main

// repo: https://github.com/emicklei/go-restful
// Refer to https://www.kubernetes.org.cn/1788.html
// https://github.com/emicklei/go-restful/tree/master/examples

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

type User struct {
	Id, Name, Age string
}

type UserResource struct {
	// normally one would use DAO (data access object)
	users map[string]User
}

type RequestBody struct {
	Name string
	Age string
}

func (u UserResource) Register(container *restful.Container) {
	// 创建新的WebService
	ws := new(restful.WebService)

	// 设定WebService对应的路径("/users")和支持的MIME类型(restful.MIME_XML/ restful.MIME_JSON)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	// 添加路由： GET /{user-id} --> u.findUser
	// 带括号的表示变量，不带的是固定路径，如果不写 则404
	//ws.Route(ws.GET("/{user-id}/test/{user-name}").To(u.findUser))
	ws.Route(ws.GET("/{user-id}").To(u.findUser))


	// 添加路由： POST / --> u.updateUser
	ws.Route(ws.POST("").To(u.createUser))

	// 添加路由： PUT /{user-id} --> u.createUser
	ws.Route(ws.PUT("/{user-id}").To(u.updateUser))

	// 添加路由： DELETE /{user-id} --> u.removeUser
	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser))

	// 将初始化好的WebService添加到Container中
	container.Add(ws)
}

// GET http://127.0.0.1:8080/users/test?k1=111&k2=222
// GET http://127.0.0.1:8080/users/id2/test/user-name?k1=111&k2=222
func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	// 获取user-id 和 url 中的变量方法
	userId := request.PathParameter("user-id")
	userName := request.PathParameter("user-name")
	fmt.Println(userId, userName)
	request.Request.ParseForm()
	k1 := request.Request.Form.Get("k1")
	k2 := request.Request.Form.Get("k2")
	fmt.Println(k1, k2)

	usr := u.users[userId]
	if len(usr.Id) == 0 {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(usr)
	}
}

//PUT http://127.0.0.1:8080/users/test
//Content-Type: application/json
//
//{
//"name": "caoyingjun",
//"age": "18"
//}
func (u *UserResource) updateUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	userId := request.PathParameter("user-id")
	// 获取 request 里面的 body
	body := request.Request.Body
	b3, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
		response.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	// 转换给结构体，以便后续使用
    var rb RequestBody
	err = json.Unmarshal(b3, &rb)
	if err != nil {
		fmt.Println(err.Error())
		response.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	//err = request.ReadEntity(&usr)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	response.WriteErrorString(http.StatusOK, err.Error())
	//	return
	//}

	usr.Id = userId
	usr.Name = rb.Name
	usr.Age = rb.Age

	u.users[usr.Id] = *usr
	//response.WriteHeader(http.StatusOK)
    response.WriteEntity(usr)
}

//POST http://127.0.0.1:8080/users
//Content-Type: application/json
//
//{
//"name": "caoyingjun",
//"age": "18"
//}
func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	usr := User{
		Id: request.PathParameter("user-id"),
	}

	body := request.Request.Body
	b3, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
		response.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	// 转换给结构体，以便后续使用
	var rb RequestBody
	err = json.Unmarshal(b3, &rb)
	if err != nil {
		fmt.Println(err.Error())
		response.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	usr.Id = "test"
	usr.Name = rb.Name
	usr.Age = rb.Age

	u.users[usr.Id] = usr
	// superfluous  多余的设置
	//response.WriteHeader(200)
	response.WriteEntity(usr)
}

//DELETE  http://127.0.0.1:8080/users/test
func (u *UserResource) removeUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	delete(u.users, id)
	// 设置返回的格式
	response.WriteEntity(map[string]string{"delete": "true"})
}

func main() {
	// 创建一个空的Container
	wsContainer := restful.NewContainer()

	// 设定路由为CurlyRouter(快速路由)
	wsContainer.Router(restful.CurlyRouter{})

	// 创建自定义的Resource Handle(此处为UserResource)
	u := UserResource{map[string]User{}}

	// 创建WebService，并将WebService加入到Container中
	u.Register(wsContainer)

	log.Printf("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}

	// 启动服务
	log.Fatal(server.ListenAndServe())
}
