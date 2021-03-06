# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api.proto](#api.proto)
    - [LoginInfoReq](#sso.LoginInfoReq)
    - [LoginInfoRsp](#sso.LoginInfoRsp)
    - [Resource](#sso.Resource)
    - [User](#sso.User)
    - [UserGroup](#sso.UserGroup)
  
  
  
    - [SsoService](#sso.SsoService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api.proto
指定proto版本


<a name="sso.LoginInfoReq"></a>

### LoginInfoReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basicRawToken | [string](#string) |  |  |






<a name="sso.LoginInfoRsp"></a>

### LoginInfoRsp
HelloRequest 请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#sso.User) |  |  |
| resources | [Resource](#sso.Resource) | repeated |  |
| userGroupes | [UserGroup](#sso.UserGroup) | repeated |  |






<a name="sso.Resource"></a>

### Resource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| pid | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| sysName | [string](#string) |  |  |
| priority | [int32](#int32) |  |  |
| path | [string](#string) |  |  |
| type | [string](#string) |  | 菜单、按钮 |
| permission | [string](#string) |  |  |
| available | [bool](#bool) |  |  |






<a name="sso.User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| gids | [int32](#int32) | repeated |  |
| available | [bool](#bool) |  |  |






<a name="sso.UserGroup"></a>

### UserGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| resources | [int32](#int32) | repeated |  |
| available | [bool](#bool) |  |  |





 

 

 


<a name="sso.SsoService"></a>

### SsoService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetLoginInfo | [LoginInfoReq](#sso.LoginInfoReq) | [LoginInfoRsp](#sso.LoginInfoRsp) |  |
| SaveUserInfo | [User](#sso.User) | [User](#sso.User) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

