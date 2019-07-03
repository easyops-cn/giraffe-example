package main

import (
	"context"
	"fmt"
	"time"

	"github.com/easyops-cn/giraffe-micro/hack"
	"github.com/easyops-cn/giraffe-micro/plugins/easyopsrest"
	"github.com/easyops-cn/giraffe-micro/plugins/easyopsrest/auth"
	"github.com/gogo/protobuf/types"

	logic_cmdb_service "github.com/easyops-cn/giraffe-example/protorepo_logic_cmdb_service"
	"github.com/easyops-cn/giraffe-example/protorepo_logic_cmdb_service/instances"
)

var instance = &instances.Instance{
	ObjectId: "APP",
	Data: &types.Struct{
		Fields: map[string]*types.Value{
			"name": {Kind: &types.Value_StringValue{StringValue: "GiraffeApp7"}},
		},
	},
}

func main() {
	// 初始化 RESTFul 客户端, 在这里可以对客户端进行简单的设置
	rest, err := easyopsrest.NewClient(
		// 设置名字服务始终返回静态地址, 方便开发时测试
		easyopsrest.WithNameService(hack.StaticAddress("192.168.100.162", 8079)),
		easyopsrest.WithTimeout(time.Second*60),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 用刚才创建的 RESTFul 客户端, 生成 CMDB 服务 Client, 稍后可以通过该 Client 对象调用 CMDB 服务接口
	client := logic_cmdb_service.NewClient(rest)

	// 在 context 中注入 CMDB 服务所需要的用户身份信息
	ctx := auth.WithUserInfo(context.Background(), auth.UserInfo{"easyops", 8888})

	// 接下来尝试利用 CMDB 服务 Client, 在 CMDB 上创建应用实例
	instanceID, err := client.Instances.CreateInstance(ctx, instance) // 调用 CMDB 服务实例创建接口
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("CreateInstance succeed: object_id: %s, instance_id: %s\n", instanceID.GetXObjectId(), instanceID.GetInstanceId())

	st, err := client.Instances.GetInstance(ctx, instanceID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("GetInstance succeed: instance name: %s\n", st.Fields["name"].GetStringValue())

	_, err = client.Instances.DeleteInstance(ctx, instanceID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("DeleteInstance succeed: object_id: %s, instance_id: %s\n", instanceID.GetXObjectId(), instanceID.GetInstanceId())
}
