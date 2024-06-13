# Server
## gRPC-GateWay Server

### Usage
1. ```proto: ```
2. https://github.com/GrosbergKirr/proto_contracts

3. ```CLI: ```
4. https://github.com/GrosbergKirr/CLI

Данный gRPC-сервер имеет две ручки:
1. Изменение имени хоста
2. Изменение списка DNS-серверов

### 1. Эндпоинт для изменения имени хоста


```proto
service GatewayService{
  rpc ChangeHostName (HostRequest) returns (HostResponse){
    option (google.api.http) = {
      post: "/v1/changehost"
      body: "*"
    };
  }
}

message HostRequest{
  string NewHostName = 1;
  string Addr = 2;
  string Password = 3;
}

message HostResponse{
  string Result = 1;
}

```


### 1. Эндпоинт для изменения списка DNS-серверов


```proto
 rpc DNSChange (DNSRequest) returns (DNSResponse){
    option (google.api.http) = {
      post: "/v1/dnschange"
      body: "*"
    };
  }
}

message DNSRequest{
  string NewDNSName = 1;
  string Addr = 2;
  string Password = 3;
}

message DNSResponse{
  string Result = 1;
}
```
### Также к данному серверу написан CLI-клиент (Ссылка выше)


TODO:
1. Собрать Docker
2. Сделать Makefile
3. Дописать http клиента
4. Написать тесты
5. Улучшить CLI

