import requests
import json

# --- 用户配置 START ---
# 请根据您的实际API信息修改以下URL
PANGOLIN_AUTH_URL = "http://scrapeapi.pangolinfo.com/api/v1/auth" # [cite: 11]
PANGOLIN_SYNC_API_URL = "http://scrapeapi.pangolinfo.com/api/v1" # [cite: 17]

# 请替换为您的真实注册邮箱和密码
USER_EMAIL = "YOUR_REGISTERED_EMAIL"  # 示例: "your_email@example.com" [cite: 12]
USER_PASSWORD = "YOUR_PASSWORD"    # 示例: "your_secure_password" [cite: 12]

# 您想要抓取的Amazon商品详情页面的完整URL
TARGET_PRODUCT_URL = "https://www.amazon.com/dp/B0DYTF8L2W" # 这是一个示例URL，请替换为您想测试的实际商品URL [cite: 18, 35]
# 目标地区的邮编，对于Amazon站点通常是必需的
TARGET_ZIPCODE = "10041" # 示例美国邮编，请根据目标区域修改 [cite: 31, 35]
# --- 用户配置 END ---

def get_auth_token(email, password):
    """
    使用邮箱和密码获取认证Token。
    Token具有较长有效期。 [cite: 11]
    """
    headers = {
        "Content-Type": "application/json" # [cite: 11]
    }
    payload = {
        "email": email, # [cite: 12]
        "password": password # [cite: 12]
    }
    print(f"向 {PANGOLIN_AUTH_URL} 发送认证请求...")
    try:
        response = requests.post(PANGOLIN_AUTH_URL, headers=headers, json=payload, timeout=10) # 添加超时设置
        response.raise_for_status()  # 如果HTTP状态码是4xx或5xx，则抛出异常

        response_json = response.json()
        if response_json.get("code") == 0 and response_json.get("data"): # [cite: 10, 16]
            print("认证成功！")
            return response_json.get("data") # [cite: 16]
        else:
            print(f"认证失败: {response_json.get('message', '未知错误')}")
            print(f"完整响应: {response_json}")
            return None
    except requests.exceptions.Timeout:
        print("认证请求超时。")
        return None
    except requests.exceptions.HTTPError as http_err:
        print(f"认证请求发生HTTP错误: {http_err}")
        print(f"响应内容: {response.text}")
        return None
    except requests.exceptions.RequestException as e:
        print(f"认证请求发生异常: {e}")
        return None
    except json.JSONDecodeError:
        print(f"认证响应不是有效的JSON格式: {response.text}")
        return None


def get_product_details_sync(token, product_url, parser_name, formats, zipcode, timeout_ms=30000): # [cite: 33]
    """
    使用Token同步获取商品详情。
    """
    headers = {
        "Content-Type": "application/json", # [cite: 17]
        "Authorization": f"Bearer {token}" # [cite: 8, 17]
    }
    payload = {
        "url": product_url, # [cite: 18]
        "formats": formats, # [cite: 20]
        "parserName": parser_name, # [cite: 24]
        "bizContext": { # [cite: 29]
            "zipcode": zipcode # [cite: 31]
        },
        "timeout": timeout_ms # [cite: 33]
    }
    print(f"\n向 {PANGOLIN_SYNC_API_URL} 发送同步任务请求...")
    print(f"请求体: {json.dumps(payload, indent=2)}")
    try:
        response = requests.post(PANGOLIN_SYNC_API_URL, headers=headers, json=payload, timeout=(timeout_ms / 1000) + 5) # API超时+网络buffer
        response.raise_for_status()

        response_json = response.json()
        if response_json.get("code") == 0: # [cite: 10, 37]
            print("同步任务成功！")
            return response_json.get("data") # [cite: 37]
        else:
            print(f"同步任务失败: {response_json.get('message', '未知错误')}")
            print(f"完整响应: {response_json}")
            return None
    except requests.exceptions.Timeout:
        print("同步任务请求超时。")
        return None
    except requests.exceptions.HTTPError as http_err:
        print(f"同步任务请求发生HTTP错误: {http_err}")
        print(f"响应内容: {response.text}")
        return None
    except requests.exceptions.RequestException as e:
        print(f"同步任务请求发生异常: {e}")
        return None
    except json.JSONDecodeError:
        print(f"同步任务响应不是有效的JSON格式: {response.text}")
        return None

if __name__ == "__main__":
    print("--- Pangolin Scrape API Python 同步调用示例 ---")

    # 检查用户是否已填写凭证
    if USER_EMAIL == "YOUR_REGISTERED_EMAIL" or USER_PASSWORD == "YOUR_PASSWORD":
        print("\n错误：请先在脚本顶部的 '用户配置' 部分填写您的真实邮箱和密码。")
        exit()

    print("\n步骤 1: 获取认证Token...")
    api_token = get_auth_token(USER_EMAIL, USER_PASSWORD)

    if api_token:
        print(f"获取到的Token (部分): {api_token[:20]}...") # 仅显示部分token以保护隐私

        print("\n步骤 2: 使用Token同步获取Amazon商品详情...")
        # 根据文档，获取JSON格式数据时，parserName 是必需的 [cite: 23]
        # Amazon商品详情对应的 parserName 是 "amzProductDetail" [cite: 26]
        # 返回格式选择 ["json"] [cite: 20, 22]
        product_data = get_product_details_sync(
            token=api_token,
            product_url=TARGET_PRODUCT_URL,
            parser_name="amzProductDetail", # [cite: 26]
            formats=["json"], # [cite: 20, 22]
            zipcode=TARGET_ZIPCODE # [cite: 31]
        )

        if product_data:
            print("\n--- 商品数据 ---")
            # API 返回的 JSON 数据在 'json' 键下，且是一个包含单个字符串元素的列表 [cite: 37]
            if "json" in product_data and isinstance(product_data["json"], list) and len(product_data["json"]) > 0:
                try:
                    # 解析这个字符串为Python字典
                    parsed_json_content = json.loads(product_data["json"][0])
                    print(json.dumps(parsed_json_content, indent=4, ensure_ascii=False))
                except json.JSONDecodeError:
                    print("解析返回的JSON内容时出错。原始字符串内容:")
                    print(product_data["json"][0])
            elif "rawHtml" in product_data and isinstance(product_data["rawHtml"], list) and len(product_data["rawHtml"]) > 0:
                print("获取到原始HTML内容 (部分显示):")
                print(product_data["rawHtml"][0][:500] + "...") # 显示前500个字符
            elif "markdown" in product_data and isinstance(product_data["markdown"], list) and len(product_data["markdown"]) > 0:
                print("获取到Markdown内容 (部分显示):")
                print(product_data["markdown"][0][:500] + "...") # 显示前500个字符
            else:
                print("获取到的数据格式未知或内容为空。完整原始数据:")
                print(json.dumps(product_data, indent=4, ensure_ascii=False))
        else:
            print("\n未能获取到商品数据。")
    else:
        print("\n获取认证Token失败，无法继续执行任务。请检查您的邮箱、密码以及网络连接。")

    print("\n--- 示例执行完毕 ---")
