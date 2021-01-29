package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//建立数据库连接
func initDB() (err error) {
	connStr := "kumit:qaz78963@tcp(35.201.213.247)/learngo"
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	//设置数据库连接池最大连接数
	db.SetMaxOpenConns(10)
	//设置最大闲置连接数
	db.SetMaxIdleConns(10)
	return
}

//创建一个结构体来接受查询返回的数据
type user struct {
	id   int
	name string
	age  int
}

//单行查询
func queryRowSingle(n int) (err error) {
	sqlStr := "select id ,name ,age from user where id=?" //?是一个占位符
	// rowObj := db.QueryRow(sqlStr, n)                      //从连接池中取一个数据库连接来查询，1是参数

	//实例 测试超出连接数，scan不传值的情景
	// for i := 0; i < 11; i++ {
	// 	fmt.Println("开始第", i)
	// 	db.QueryRow(sqlStr, 1)
	// }
	var u user
	db.QueryRow(sqlStr, n).Scan(&u.id, &u.name, &u.age) // 常用写法
	// rowObj.Scan(&u.id, &u.name, &u.age) //必须使用scan方法来传值，调用该方法后才会释放连接，归还到连接池中，不然就卡住了
	fmt.Println(u)
	return
}

// 多行查询
func queryRows(n int) (err error) {
	sqlStr := "select id ,name,age from user where id<?"
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Printf("%#v\n", u)
	}
	return
}

// 插入数据
func insretRow() (err error) {
	sqlStr := `insert into user(name,age) values("唐五","40")`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println("insret err:", err)
		return
	}
	// 如果插入成功能拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(id)
	return
}

//更新数据
func updateRow(age, id int) (err error) {
	sqlStr := `update user set age=? where id=?`
	ret, err := db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Println("update err:", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("更新了：", n)
	return
}

//删除数据
func delectRow(id int) (err error) {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("删除了：", n)
	return
}

// 预处理插入,批量处理数据
func prepereInsertRow() (err error) {
	sqlStr := `insert into user(name,age) values(?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer stmt.Close()
	m := map[string]int{
		"刘五": 30,
		"马六": 50,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("mysql conn!")
	// queryRowSingle(2)

	// queryRows(5)
	// insretRow()
	// updateRow(22, 2)
	// delectRow(3)
	// prepereInsertRow()
	queryRows(10)

}
