package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

var ErrTaskNoExist = errors.New("task does not exist")

type ResponseError struct {
	Err error
}

type Task struct {
	title string
	done  bool
	due   *time.Time
}

type ID string

type DataAccess interface {
	Get(id ID) (Task, error)
	Put(id ID, t Task) error
	Post(t Task) (ID, error)
	Delete(id ID) error
}

type Response struct {
	ID    ID            `json:"id,omitempty"`
	Task  Task          `json:"task"`
	Error ResponseError `json:"error"`
}

type MemoryDataAccess struct {
	tasks  map[ID]Task
	nextID int64
}

func (m *MemoryDataAccess) Get(id ID) (Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return Task{}, ErrTaskNoExist
	}
	return t, nil
}

func (m *MemoryDataAccess) Put(id ID, t Task) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNoExist
	}
	m.tasks[id] = t
	return nil
}

func (m *MemoryDataAccess) Post(t Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t

	return id, nil
}

func (m *MemoryDataAccess) Delete(id ID) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNoExist
	}
	delete(m.tasks, id)
	return nil
}

func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		tasks:  map[ID]Task{},
		nextID: int64(1),
	}
}

var m = NewMemoryDataAccess()

const pathPrefix = "/api/v1/task/"

func apiHandler(response http.ResponseWriter, request *http.Request) {
	getID := func() (ID, error) {
		id := ID(request.URL.Path[len(pathPrefix):])
		if id == "" {
			return id, errors.New("apiHandler: ID is empty")
		}
		return id, nil
	}

	getTasks := func() ([]Task, error) {
		var result []Task
		if err := request.ParseForm(); err != nil {
			return nil, err
		}
		encodedTasks, ok := request.PostForm["task"]
		if !ok {
			return nil, errors.New("task paameter expected")
		}
		for _, encodedTask := range encodedTasks {
			var t Task
			if err := json.Unmarshal([]byte(encodedTask), &t); err != nil {
				return nil, err
			}
			result = append(result, t)
		}
		return result, nil
	}

	switch request.Method {
	case "GET":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		t, err := m.Get(id)

		err = json.NewEncoder(response).Encode(Response{
			ID:    id,
			Task:  t,
			Error: ResponseError{err},
		})
		if err != nil {
			log.Println(err)
		}
	case "PUT":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		tasks, err := getTasks()
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range tasks {
			err = m.Put(id, t)
			err = json.NewEncoder(response).Encode(Response{
				ID:    id,
				Task:  t,
				Error: ResponseError{err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "POST":
		tasks, err := getTasks()
		if err != nil {
			log.Println(err)
			return
		}

		for _, t := range tasks {
			id, err := m.Post(t)
			err = json.NewEncoder(response).Encode(Response{
				ID:    id,
				Task:  t,
				Error: ResponseError{err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "DELETE":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		err = m.Delete(id)
		err = json.NewEncoder(response).Encode(Response{
			ID:    id,
			Error: ResponseError{err},
		})

		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", apiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
