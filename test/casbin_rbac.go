package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	entAdapter "github.com/casbin/ent-adapter"
)

func Casbin() {
	// 定义模型
	text := `
    [request_definition]
    r = sub, obj, act

    [policy_definition]
    p = sub, obj, act

    [role_definition]
    g = _, _

    [policy_effect]
    e = some(where (p.eft == allow))

    [matchers]
    m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && r.act == p.act
    `

	// 创建模型
	m, err := model.NewModelFromString(text)
	if err != nil {
		fmt.Println("Model error:", err)
		return
	}

	// 创建适配器
	adapter, err := entAdapter.NewAdapter("mysql", "root:dachuang2574598.@tcp(117.72.65.213:13306)/gorm")
	if err != nil {
		fmt.Println("Adapter error:", err)
		return
	}

	// 创建 Enforcer
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		fmt.Println("Enforcer error:", err)
		return
	}

	// 加载策略
	err = enforcer.LoadPolicy()
	if err != nil {
		fmt.Println("Load policy error:", err)
		return
	}

	// 打印当前加载的策略
	fmt.Println("Current policies:")
	policies, err := enforcer.GetPolicy()
	if err != nil {
		fmt.Println("Get policy error:", err)
		return
	}
	for _, policy := range policies {
		fmt.Println(policy)
	}

	// 执行访问控制
	sub := "alice"  // 主体
	obj := "/data1" // 资源
	act := "read"   // 动作

	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		fmt.Println("Enforce error:", err)
		return
	}

	if ok {
		fmt.Println("Access granted")
	} else {
		fmt.Println("Access denied")
	}
}
