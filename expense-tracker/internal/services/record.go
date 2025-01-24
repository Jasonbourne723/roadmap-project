package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"roadmap/expense-tracker/internal/models"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	fileName = "data.json"
)

var (
	ErrNotExist       = errors.New("record is not exist")
	ErrNegativeAmount = errors.New("amount not allow be less than 0")
)

var RecordSvc = new(recordSvc)

type recordSvc struct{}

func load() ([]models.Record, error) {
	f, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, fmt.Errorf("openfile failed:%w", err)
	}
	defer f.Close()
	var records []models.Record
	j := json.NewDecoder(f)
	err = j.Decode(&records)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("decode failed:%w", err)
	}
	return records, nil
}

func save(records []models.Record) error {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return fmt.Errorf("openfile failed:%w", err)
	}
	defer f.Close()

	j := json.NewEncoder(f)
	return j.Encode(records)
}

func (r *recordSvc) Add(description string, amount int) error {

	if amount <= 0 {
		return ErrNegativeAmount
	}

	list, err := load()
	if err != nil {
		return err
	}
	var id int32
	if len(list) == 0 {
		id = 1
	} else {
		id = list[len(list)-1].Id + 1
	}

	list = append(list, models.Record{
		Id:          id,
		Description: description,
		Amount:      amount,
		CreatedAt:   time.Now(),
	})
	return save(list)
}

func (r *recordSvc) Del(id int32) error {
	list, err := load()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return ErrNotExist
	}
	pos, exist := search(list, 0, len(list)-1, id)
	if !exist {
		return ErrNotExist
	}
	if pos == 0 {
		list = list[pos+1:]
	} else if pos == len(list)-1 {
		list = list[:pos]
	} else {
		list = append(list[:pos], list[pos+1:]...)
	}
	return save(list)
}

func search(list []models.Record, left int, right int, id int32) (int, bool) {
	if list[left].Id == id {
		return left, true
	}
	if list[right].Id == id {
		return right, true
	}
	if list[left].Id > id {
		return 0, false
	}
	if list[right].Id < id {
		return 0, false
	}
	if right <= left+1 {
		return 0, false
	}
	middle := (left+right)>>1 + left

	if list[middle].Id == id {
		return middle, true
	}
	if list[middle].Id > id {
		return search(list, left, middle, id)
	} else {
		return search(list, middle, right, id)
	}
}

func (r *recordSvc) List(year int, month int) ([]models.Record, error) {
	list, err := load()
	if err != nil {
		return nil, err
	}

	list = Filter(list, func(r models.Record) bool {
		if year == r.CreatedAt.Year() && month == int(r.CreatedAt.Month()) {
			return true
		}
		return false
	})
	return list, nil
}

func Filter(list []models.Record, f func(models.Record) bool) []models.Record {
	var newList []models.Record
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func Sum(list []models.Record, f func(models.Record) int) int {
	var sum int
	for _, v := range list {
		sum += f(v)
	}
	return sum
}

func (r *recordSvc) Summary(year int, month int) (int, error) {
	list, err := load()
	if err != nil {
		return 0, err
	}

	sum := Sum(Filter(list, func(r models.Record) bool {
		if year == r.CreatedAt.Year() && month == int(r.CreatedAt.Month()) {
			return true
		}
		return false
	}), func(r models.Record) int {
		return r.Amount
	})
	return sum, nil
}

func (r *recordSvc) Export() error {
	list, err := load()
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	buffer.WriteString("id,description,amount,createdAt\n")
	for _, v := range list {
		buffer.WriteString(strconv.Itoa(int(v.Id)))
		buffer.WriteString(",")
		buffer.WriteString(v.Description)
		buffer.WriteString(",")
		buffer.WriteString(strconv.Itoa(v.Amount))
		buffer.WriteString(",")
		buffer.WriteString(v.CreatedAt.Format("2006-01-02 15:04:05"))
		buffer.WriteString("\n")
	}
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	name := uid.String() + ".csv"
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(buffer.String())
	return nil
}
