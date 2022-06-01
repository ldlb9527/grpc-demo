![](https://img.shields.io/badge/-ceph-green)
# proto3语法
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