package banyan_api

import (
	"strconv"
)

type basicsFunc interface {
	Set(field string, val interface{}) error
	Get(field string) (string, error)
	Incr(field string, num int) (int, error)
	Exists(field string) (bool, error)
	Size() (int, error)
	Clear() error
	Scan(member string, start, end, limit int) (map[string]string, error)
	Rscan(member string, start, end, limit int) (map[string]string, error)
	Mget(fields []string) (map[string]string, error)
	GetAll() (map[string]string, error)
}

type OperationSet struct {
	Db *BanyanClient
}

func (this OperationSet) Set(field string, val interface{}) error {
	return this.Db.Set(field, val.(string))
}

func (this OperationSet) Get(field string) (string, error) {
	return this.Db.Get(field)
}

func (this OperationSet) Incr(field string, num int) (int, error) {
	return this.Db.Incr(field, num)
}

func (this OperationSet) Exists(field string) (bool, error) {
	return this.Db.Exists(field)
}

func (this OperationSet) Size() (int, error) {
	return 0, nil
}

func (this OperationSet) Clear() error {
	return nil
}

func (this OperationSet) Scan(member string, start, end, limit int) (map[string]string, error) {
	return nil, nil
}

func (this OperationSet) Rscan(member string, start, end, limit int) (map[string]string, error) {
	return nil, nil
}

func (this OperationSet) Mget(fields []string) (map[string]string, error) {
	return nil, nil
}

func (this OperationSet) GetAll() (map[string]string, error) {
	return nil, nil
}

type OperationHset struct {
	Db  *BanyanClient
	Key string
}

func (this OperationHset) Set(field string, val interface{}) error {
	return this.Db.Hset(this.Key, field, val.(string))
}

func (this OperationHset) Get(field string) (string, error) {
	return this.Db.Hget(this.Key, field)
}

func (this OperationHset) Incr(field string, num int) (int, error) {
	return this.Db.Hincr(this.Key, field, num)
}

func (this OperationHset) Exists(field string) (bool, error) {
	return this.Db.Hexists(this.Key, field)
}

func (this OperationHset) Size() (int, error) {
	return this.Db.Hsize(this.Key)
}

func (this OperationHset) Clear() error {
	return this.Db.Hclear(this.Key)
}

func (this OperationHset) Scan(member string, start, end, limit int) (map[string]string, error) {
	startStr := strconv.Itoa(start)
	endStr := strconv.Itoa(end)
	return this.Db.Hscan(this.Key, startStr, endStr, limit)
}

func (this OperationHset) Rscan(member string, start, end, limit int) (map[string]string, error) {
	KeyInt, _ := strconv.Atoi(this.Key)
	return this.Db.Hrscan(KeyInt, start, end, limit)
}

func (this OperationHset) Mget(fields []string) (map[string]string, error) {
	return this.Db.MultiHget(this.Key, fields...)
}

func (this OperationHset) GetAll() (map[string]string, error) {
	return this.Db.Hgetall(this.Key)
}

type OperationZset struct {
	Db  *BanyanClient
	Key string
}

func (this OperationZset) Set(field string, val interface{}) error {
	_, vErr := this.Db.Zset(this.Key, field, val.(int))
	return vErr
}

func (this OperationZset) Get(field string) (string, error) {
	return this.Db.Zget(this.Key, field)
}

func (this OperationZset) Incr(field string, num int) (int, error) {
	return this.Db.Zincr(this.Key, field, num)
}

func (this OperationZset) Exists(field string) (bool, error) {
	return this.Db.Zexists(this.Key, field)
}

func (this OperationZset) Size() (int, error) {
	return this.Db.Zsize(this.Key)
}

func (this OperationZset) Clear() error {
	return this.Db.Zclear(this.Key)
}

func (this OperationZset) Scan(member string, start, end, limit int) (map[string]string, error) {
	data, err := this.Db.Zscan(this.Key, member, start, end, limit)
	if err != nil {
		return nil, err
	}
	var result = make(map[string]string)
	for k, v := range data {
		value := strconv.Itoa(v)
		result[k] = value
	}
	return result, nil
}

func (this OperationZset) Rscan(member string, start, end, limit int) (map[string]string, error) {
	data, err := this.Db.Zrscan(this.Key, member, start, end, limit)
	if err != nil {
		return nil, err
	}
	var result = make(map[string]string)
	for k, v := range data {
		value := strconv.Itoa(v)
		result[k] = value
	}
	return result, nil
}

func (this OperationZset) Mget(fields []string) (map[string]string, error) {
	return this.Db.MultiHget(this.Key, fields...)
}

func (this OperationZset) GetAll() (map[string]string, error) {
	data, err := this.Db.Zgetall(this.Key)
	if err != nil {
		return nil, err
	}
	var result = make(map[string]string)
	for k, v := range data {
		value := strconv.Itoa(v)
		result[k] = value
	}
	return result, nil
}

type UseFunc struct {
	basicsFunc
}

func NewBanyanUse(b basicsFunc) UseFunc {
	return UseFunc{b}
}
