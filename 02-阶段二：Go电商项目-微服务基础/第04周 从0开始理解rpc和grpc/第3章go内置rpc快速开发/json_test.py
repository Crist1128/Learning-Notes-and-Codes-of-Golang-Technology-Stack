import socket
import json

# JSON-RPC 请求的数据结构
request_data = {
    "jsonrpc": "2.0",  # JSON-RPC 版本号
    "method": "HelloService.Hello",  # 远程调用的方法名：服务名.方法名
    "params": ["cc"],  # 方法的参数，这里是一个字符串数组 ["cc"]
    "id": 0  # 客户端生成的唯一请求 ID，用于匹配响应
}

# 将 Python 字典转化为 JSON 格式的字符串
json_data = json.dumps(request_data)

# 创建 TCP socket
client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# 连接到服务器，IP 和端口需要根据服务端实际情况修改
server_address = ('localhost', 1234)  # 服务端 IP 和端口
client_socket.connect(server_address)

try:
    # 发送 JSON-RPC 请求数据
    # 需要将字符串转换为字节数据发送
    client_socket.sendall(json_data.encode('utf-8'))

    # 接收响应数据，缓冲区大小 1024 字节
    response = client_socket.recv(1024)

    # 将接收到的字节数据转换为字符串
    response_data = response.decode('utf-8')

    # 将字符串转换为 Python 字典，方便处理
    response_json = json.loads(response_data)

    # 输出响应内容
    print("Received response:", response_json)

finally:
    # 关闭连接
    client_socket.close()
