package ch32

import (
	"errors"
	"time"
)

type ReusableObj struct {
}
type ObjPool struct {
	bufChan chan *ReusableObj // 用于缓存可重用对象
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		// 往池中放对象  对象可以是具体的
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

// 需要超时很重要 高可用性 sony

//func (p *ObjPool) GetObj(timeout time.Duration) (obj interface, error) {
//建议  使用不同的池   缓存不同的对象 具体的比较好 别obj interface 得多判断对象啥的
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil

	case <-time.After(timeout):
		return nil, errors.New("time out")
	}

}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")

	}
}
