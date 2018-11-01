package demo

import (
	"fmt"
	"reflect"
)

type Address struct {
	City string
	Area string
}

type Student struct {
	Address
	Name string
	Age  int
}

func (this Student) Say() {
	fmt.Println("hello, i am ", this.Name, "and i am ", this.Age)
}

func (this Student) Hello(word string) {
	fmt.Println("hello", word, ". i am ", this.Name)
}

/*
  ��ȡ�������Ϣ
*/
func StructInfo(o interface{}) {
	//��ȡ���������
	t := reflect.TypeOf(o)
	fmt.Println(t.Name(), "object type: ", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("the object is not a struct, but it is", t.Kind())
		return
	}

	//��ȡ�����ֵ
	v := reflect.ValueOf(o)
	fmt.Println(t.Name(), "object value: ", v)

	//��ȡ������ֶ�
	fmt.Println(t.Name(), "fields: ")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v = %v \n", f.Name, f.Type, val)
		//ͨ���ݹ���û�ȡ�����͵���Ϣ
		t1 := reflect.TypeOf(val)
		if k := t1.Kind(); k == reflect.Struct {
			StructInfo(val)
		}
	}
	//��ȡ����ĺ���
	fmt.Println(t.Name(), "methods: ", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%10s:%v \n", m.Name, m.Type)
	}
}

/*
  �����ֶεķ���
*/
func Annoy(o interface{}) {
	t := reflect.TypeOf(o)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%10s:%#v \n", f.Name, f)
	}
}

/*
  ͨ�����������ֶ�
*/
func ReflectSet(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("�޸�ʧ��")
		return
	}
	v = v.Elem()
	//��ȡ�ֶ�
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("�޸�ʧ��")
		return
	}
	//����ֵ
	if f.Kind() == reflect.String {
		f.SetString("chairis")
	}
}

/*
  ͨ��������ú���
*/
func ReflectMethod(o interface{}) {
	v := reflect.ValueOf(o)
	//�޲κ�������
	m1 := v.MethodByName("Say")
	m1.Call([]reflect.Value{})

	//�вκ�������
	m2 := v.MethodByName("Hello")
	m2.Call([]reflect.Value{reflect.ValueOf("iris")})
}
