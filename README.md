![](https://img.shields.io/badge/grpc-proto3-green)
# 一、grpc usage (完整demo,可直接从第五步开始运行)
* **1.安装`protoc`**
```
protoc --version  #确保版本在3.0+
```
***
* 2.**创建项目,gomod引入对应包**
```
go get -u github.com/golang/protobuf/protoc-gen-go
```
***
* 3.**编写user.proto文件,前往[proto3语法学习](#proto)**
```
protoc -I . users/user.proto --go_out=plugins=grpc:.  //会在users包下生成user.pb.go文件
```
***
* 4.**编写服务端代码,`server/main.go`,编写客户端代码(也可通过postman直接访问服务端,新版本已支持),`client/main.go`**
***
* 5.**先启动服务端，再启动客户端访问**
```
go run server/main.go
go run client/main.go
```
***
* **6.控制台日志**
```
//server log

2022/05/31 15:17:32 receive users index request:page 1 page_size 12
2022/05/31 15:17:32 receive users uid request:uid 1
2022/05/31 15:17:32 receive users uid request:name big_cat password:123456,age:29
2022/05/31 15:17:32 receive users uid request:uid 1

====================================================================

//client log

aaaa 28
bbbb 1
aaaa 28
2022/05/31 15:30:53 user index success: success
2022/05/31 15:30:53 user view success: success
2022/05/31 15:30:53 user post success: success
2022/05/31 15:30:53 user delete success: success
```
***
# 二、<a id="proto"></a>proto3语法学习
## 1.定义一个 Message
```protobuf
syntax = "proto3"; //指定使用proto3，如果不指定的话，编译器会使用proto2去编译

message SearchRequests {
    // 定义SearchRequests的成员变量，需要指定：变量类型、变量名、变量Tag
    string query = 1;
    int32 page_number = 2;
    int32 result_per_page = 3;
}
```
***
## 2.定义多个message类型
```protobuf
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}

message SearchResponse {
  repeated string result = 1;
}
```
***
## 3.message可以嵌套定义，比如message可以定义在另一个message内部
```protobuf
syntax = "proto3";
message SearchResponse {
    message Result {
        string url = 1;
        string title = 2;
        repeated string snippets = 3;
    }
    repeated Result results = 1;
}

message SomeOtherMessage {
  SearchResponse.Result result = 1; //定义在内部的message的使用
}
```
***
## 4.定义变量类型
|.proto | 说明 | C++ | Java | Python | Go | Ruby | C# | PHP |
|  ---- | ----| ----| ----| ----| ----| ----| ----| ----|
| double |  | double | double | float | float64 | Float | double | float |
|float |  | float | float | float | float32 | Float | float | float |
| int32 | 使用变长编码，对负数编码效率<br>低，如果你的变量可能是负数，可以使用sint32 | int32 | int | int | int32 | Fixnum or Bignum (as required) | int | integer |
| int64 | 使用变长编码，对负数编码效率<br>低，如果你的变量可能是负数，可以使用sint64 | int64 | long | int/long | int64 | Bignum | long | integer/string|
|uint32 | 使用变长编码 | uint32 | int | int/long | uint32 | Fixnum or Bignum (as required) | uint | integer|
|uint64 | 使用变长编码 | uint64 | long | int/long | uint64 | Bignum | ulong | integer/string|
|sint32 | 使用变长编码，带符号的int类型，对负数<br>编码比int32高效 | int32 | int | int | int32 | Fixnum or Bignum (as required) | int | integer|
|sint64 | 使用变长编码，带符号的int类型，对负数<br>编码比int64高效 | int64 | long | int/long | int64 | Bignum | long | integer/string|
|fixed32 | 4字节编码， 如果变量经常大于的话，<br>会比uint64高效 | uint64 | long | int/long | uint64 | Bignum | ulong | integer/string|
|sfixed32 | 4字节编码 | int32 | int | int | int32 | Fixnum or Bignum (as required) | int | integer|
|sfixed64 | 8字节编码 | int64 | long | int/long | int64 | Bignum | long | integer/string|
|bool |  | bool | boolean | bool | bool | TrueClass/FalseClass | bool | boolean|
|string | 必须包含utf-8编码或者7-bit ASCII text | string | String | str/unicode | string | String (UTF-8) | string | string|
|bytes | 任意的字节序列 | string | ByteString | str | []byte | String (ASCII-8BIT) | ByteString | string|
***
## 5.分配tag
* 每一个变量在message内都需要自定义一个唯一的数字Tag，protobuf会根据Tag从数据中查找变量对应的位置，Tag一旦指定，以后更新协议的时候也不能修改，否则无法对旧版本兼容
* Tag的取值范围最小是1，最大是-1，但 19000~19999 是 protobuf 预留的，用户不能使用。
* Tag的编码：1 ~ 15：单字节编码，16 ~ 2047：双字节编码
***
## 6.指定变量规则
* 在proto3中，可以给变量指定两个规则：`singular`:0或者1个，但不能多于1个;`repeated`:任意数量（包括0）
* 在proto2中，规则为：`required`：必须有一个;`optional`：0或者1个;`repeated`：任意数量（包括0）
***
## 7.保留变量不被使用
* 一旦 Tag 指定后就不能变更，这就会带来一个问题，假如在版本1的协议中，我们有个变量：`int32 number = 1;`<br>在版本2中，我们决定废弃对它的使用,那我们应该如何修改协议呢？注释掉它？删除掉它？如果把它删除了，后来<br>者很可能在定义新变量的时候，使新的变量 Tag = 1,这样会导致协议不兼容,我们可以用 reserved 关键字，当一个变<br>量不再使用的时候，我们可以把它的变量名或 Tag 用 reserved 标注，这样，当这个Tag或者变量名字被重新使用的时候，编译器会报错
```protobuf
message Foo {
    // 注意，同一个 reserved 语句不能同时包含变量名和 Tag 
    reserved 2, 15, 9 to 11;
    reserved "foo", "bar";
}
```
***
## 8.默认值
* 解析 message 时，如果被编码的 message 里没有包含某些变量，那么根据类型不同，他们会有不同的默认值
* `string`：默认是空的字符串,`byte`：默认是空的bytes,`bool`：默认为false,`numeric`：默认为0,
* `enums`：定义在第一位的枚举值，也就是0,`messages`：根据生成的不同语言有不同的表现
## 9.定义枚举
```protobuf
syntax = "proto3";
message SearchRequest {
    string query = 1;
    int32 page_number = 2; // 页码
    int32 page_size = 3; // 每页条数
    enum Corpus {
        UNIVERSAL = 0;
        WEB = 1;
        IMAGES = 2;
        LOCAL = 3;
        NEWS = 4;
        PRODUCTS = 5;
        VIDEO = 6;
    }
    Corpus corpus = 4;
}
```
* 枚举定义在一个消息内部或消息外部都是可以的，如果枚举是 定义在 message 内部，而其他 message 又想使用，那么可以通过 `MessageType.EnumType` 的方式引用。定义枚举的时候，我们要保证第一个枚举值必须是0，枚举值不能重复，除非使用 `option allow_alias = true` 选项来开启别名。
```protobuf
syntax = "proto3";
enum EnumAllowingAlias {
    option allow_alias = true;
    UNKNOWN = 0;
    STARTED = 1;
    RUNNING = 1;
}
```
* 枚举值的范围是32-bit integer，但因为枚举值使用变长编码，所以不推荐使用负数作为枚举值，因为这会带来效率问题
***
## 10.引用其他 proto 文件
```protobuf
// 在1.proto中
//import "2.proto";

// 在2.proto中
//import "3.proto";   //1中导入2，2中导入3，此时1无法使用3中定义的内容

// 在1.proto中
//import "2.proto";

// 在2.proto中
//import public "3.proto";   //1中导入2，2中导入3，此时1可以使用3中定义的内容
//2都能使用3中定义的内容
```
***
## 11.正确升级proto文件
* 不要修改任何已存在的变量的 Tag
* 如果你新增了变量，新生成的代码依然能解析旧的数据，但新增的变量将会变成默认值。相应的，新代码序列化的数据也能被旧的代码解析，但旧代码会自动忽略新增的变量。
* 废弃不用的变量用 reserved 标注
* int32、 uint32、 int64、 uint64 和 bool 是相互兼容的，这意味你可以更改这些变量的类型而不会影响兼容性
* sint32 和 sint64 是兼容的，但跟其他类型不兼容
* string 和 bytes 可以兼容，前提是他们都是UTF-8编码的数据
* fixed32 和 sfixed32 是兼容的, fixed64 和 sfixed64是兼容的
***
## 12.使用`Any`
* Any可以让你在proto文件中使用未定义的类型，具体里面保存什么数据，是在上层业务代码使用的时候决定的，使用`Any`必须导入`import google/protobuf/any.proto`
```protobuf
syntax = "proto3";
import "google/protobuf/any.proto";

message ErrorStatus {
    string message = 1;
    repeated google.protobuf.Any details = 2;
}
```
***
## 13.使用`oneof`
* Oneof 类似union，如果你的消息中有很多可选字段，而同一个时刻最多仅有其中的一个字段被设置的话，你可以使用oneof来强化这个特性并且节约存储空间
```protobuf
syntax = "proto3";
message LoginReply {
    oneof test_oneof {
        string name = 3;  //name 和 age 都是 LoginReply 的成员，但不能给他们同时设置值（设置一个oneof字段会自动清理其他的oneof字段）
        string age = 4;
    }
   repeated string status = 1;
   repeated string token = 2;
}
```
***
## 14.使用`map`
* map<key_type, value_type> map_field = N;
* key_type:必须是string或者int，value_type：任意类型，例如：`map<string, Project> projects = 3;`
* Map 类型不能使`repeated`，Map 是无序的，以文本格式展示时，Map以key来排序，如果有相同的键会导致解析失败
***
## 15.使用`package`
* 防止不同消息之间的命名冲突，可以对特定的.proto文件提指定package名字。在定义消息的成员的时候，可以指定包的名字
```protobuf
syntax = "proto3";
package foo.bar;
message Open {
}

message Foo {
  // 带上包名
  foo.bar.Open open = 1;
}
```
***
## 16.Options
* Options 分为 file-level options（只能出现在最顶层，不能在消息、枚举、服务内部使用）、 message-level options（只能在消息内部使用）、field-level options（只能在变量定义时使用）
* java_package (file option)：指定生成类的包名，如果没有指定此选项，将由关键字package指定包名。此选项只在生成 java 代码时有效
* java_multiple_files (file option)：如果为 true， 定义在最外层的 message 、enum、service 将作为单独的类存在
* java_outer_classname (file option)：指定最外层class的类名，如果不指定，将会以文件名作为类名
* optimize_for (file option)：可选有 `SPEED`,`CODE_SIZE`,`LITE_RUNTIME` ，分别是效率优先、空间优先，第三个lite是兼顾效率和代码大小，但是运行时需要依赖 libprotobuf-lite
* cc_enable_arenas (file option):启动arena allocation，c++代码使用
* objc_class_prefix (file option)：Objective-C使用
* deprecated (field option)：提示变量已废弃、不建议使用
```protobuf
syntax = "proto3";
option java_package = "com.example.foo";
option java_multiple_files = true;
option java_outer_classname = "Ponycopter";
option optimize_for = CODE_SIZE;

message Foo {
  int32 old_field = 6 [deprecated=true];
}
```
***
## 17.定义services
* 定义一个服务，在 .proto文件中指定service
```protobuf
syntax = "proto3";

service User {
  rpc UserIndex(UserIndexRequest) returns (UserIndexResponse) {}
}

message UserIndexRequest {
  int32 page = 1;
  int32 page_size = 2;
}

message UserIndexResponse {
  int32 err = 1;
  string msg = 2;
  // 返回一个 UserEntity 对象的列表数据
  repeated UserEntity data = 3;
}

message UserEntity {
  string name = 1;
  int32 age = 2;
}
```
***
## 18.代码生成(不同版本的protoc的命令不一样，请查看[gRPC官方文档](https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code))
* 使用 `protoc`工具可以把编写好的proto文件“编译”为Java, Python, C++, Go, Ruby, JavaNano, Objective-C,或C#代码
```shell
protoc --proto_path=IMPORT_PATH --cpp_out=DST_DIR --java_out=DST_DIR --python_out=DST_DIR --go_out=DST_DIR path/to/file.proto
```
* `IMPORT_PATH`：指定proto文件的路径，如果没有指定， protoc会从当前目录搜索对应的proto文件，如果有多个路径，那么可以指定多次`--proto_path`
* 指定各语言代码的输出路径,–-cpp_out：生成c++代码  --java_out ：生成java代码







