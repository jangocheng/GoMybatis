package example

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
	"github.com/zhuxiujia/GoMybatis/lib/github.com/go-xorm/xorm"
	"os"
	"fmt"
	"io/ioutil"
	"github.com/zhuxiujia/GoMybatis"
)
//定义mapper文件的接口和结构体
type ExampleActivityMapper interface {
	SelectAll(result *[]Activity) error
	SelectByCondition(Name string, StartTime time.Time, EndTime time.Time, Page int, Size int, result *[]Activity) error
	UpdateById(arg Activity, result *int64) error
	Insert(arg Activity, result *int64) error
	CountByCondition(name string, startTime time.Time, endTime time.Time, result *int) error
}
//定义mapper文件的接口和结构体，也可以只定义结构体就行
//mapper.go文件 函数必须为2个参数（第一个为自定义结构体参数（属性必须大写），第二个为指针类型的返回数据） error 为返回错误
type ExampleActivityMapperImpl struct {
	ExampleActivityMapper
	SelectAll         func(result *[]Activity) error
	SelectByCondition func(name string, startTime time.Time, endTime time.Time, page int, size int, result *[]Activity) error `mapperParams:"name,startTime,endTime,page,size"`
	UpdateById        func(arg Activity, result *int64) error
	Insert            func(arg Activity, result *int64) error
	CountByCondition  func(name string, startTime time.Time, endTime time.Time, result *int) error                            `mapperParams:"name,startTime,endTime"`
}

func Test_main(t *testing.T) {
	var err error
	//mysql链接格式为         用户名:密码@(数据库链接地址:端口)/数据库名称   例如root:123456@(***.mysql.rds.aliyuncs.com:3306)/test
	engine, err := xorm.NewEngine("mysql", "*?charset=utf8&parseTime=True&loc=Local") //此处请按格式填写你的mysql链接，这里用*号代替
	if err != nil {
		panic(err.Error())
	}
	engine.ShowSQL()

	file, err := os.Open("Example_ActivityMapper.xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)
	var exampleActivityMapperImpl ExampleActivityMapperImpl
	//设置对应的mapper xml文件
	GoMybatis.UseProxyMapper(&exampleActivityMapperImpl, bytes, engine)

	//使用mapper
	var result []Activity
	exampleActivityMapperImpl.SelectByCondition("",time.Time{},time.Time{},0,2000,&result)

	fmt.Println(result)

	//result[0].Name="rs-446"
	//err=exampleActivityMapperImpl.UpdateById(result[0],nil)
	//fmt.Println(err)
	//
	//result[0].Id=result[0].Id+"2"
	//err=exampleActivityMapperImpl.Insert(result[0],nil)
	//fmt.Println(err)
}