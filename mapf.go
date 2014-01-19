package mapf

import (
	"os"
	"fmt"
	"encoding/json"
)

type Mapf struct {
	filePath string
	writeMode int
	data map[string]interface{}
}

const (
	MODE_AUTOWRITE = iota
	MODE_NONWRITE
)

func New(path string)(m *Mapf, e error){

	f, oe := os.OpenFile(path, os.O_RDONLY, os.ModePerm)	
	defer f.Close()
	if oe != nil {
		fmt.Printf("%v", oe)	
		return nil, oe
	}
	
	fileinfo, _ := f.Stat()
	jsontext := make([]byte, fileinfo.Size())
	if _ , re := f.Read(jsontext); re != nil {
		fmt.Printf("%v", re)	
		return nil, re
	}

	m = &Mapf{path, MODE_AUTOWRITE, make(map[string]interface{})}

	if de := json.Unmarshal(jsontext, &m.data); de != nil {
		fmt.Printf("%v", de)
	}
	
	return
}

func (m *Mapf)Flush()(e error){

	b, ee := json.Marshal(m.data)
	if ee != nil{
		fmt.Printf("%v", ee)
		return ee
	}

	f, oe := os.OpenFile(m.filePath, os.O_WRONLY, os.ModePerm)
	defer f.Close()
	if oe != nil{
		fmt.Printf("%v", oe)	
		return oe
	}

	_, we := f.Write(b)
	if we != nil{
		fmt.Printf("%v",  we)
		return we
	}
	
	if te := f.Truncate(int64(len(b))); te != nil {
		fmt.Printf("%v",  te)
		return te
	}
	
	return nil
}

func (m *Mapf)SwitchMode(mode int){
	m.writeMode = mode
}

func (m *Mapf)Put(k string, v interface{}) error {
	m.data[k] = v
	if m.writeMode == MODE_AUTOWRITE{
		return m.Flush()
	}
	return nil
}

func (m *Mapf)Get(k string)(v interface{}, ok bool){
	v, ok = m.data[k]
	return 
}

func (m *Mapf)Delete(k string) error {
	delete(m.data, k)
	if m.writeMode == MODE_AUTOWRITE{
		return m.Flush()
	}
	return nil
}
