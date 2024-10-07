import requests
import json

# 服务器的 URL，假设你使用的是前面 Go 服务器的 "/hello" 路径
url = "http://localhost:1234/hello"

# JSON-RPC 请求数据
json_rpc_data = {
    "method": "HelloService.Hello",  # 方法名
    "params": ["John"],  # 参数
    "id": 1  # 请求ID，用于匹配请求和响应
}

# 将数据转换为 JSON 格式
headers = {'Content-Type': 'application/json'}
response = requests.post(url, data=json.dumps(json_rpc_data), headers=headers)

# 输出服务器响应内容
print("Response status code:", response.status_code)
print("Response body:", response.json())
