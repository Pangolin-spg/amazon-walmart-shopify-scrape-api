# 🚀 电商数据采集 API - 免费Amazon/Walmart数据提取工具

[![API Status](https://img.shields.io/badge/API-在线-brightgreen)](https://docs.pangolinfo.com)
[![支持平台](https://img.shields.io/badge/支持平台-Amazon%20|%20Walmart%20|%20eBay%20|%20Shopify-blue)](https://www.pangolinfo.com)
[![数据格式](https://img.shields.io/badge/数据格式-JSON%20|%20HTML%20|%20Markdown-orange)](https://docs.pangolinfo.com)

## 🎁 **免费福利 - 限时提供**

🔥 **为开源社区提供特别优惠：**
- ✅ **免费获取完整解析模板代码**
- ✅ **200次免费API调用额度**  
- ✅ **技术支持和使用指导**
- ✅ **无需信用卡，即刻开始**

### 📞 **获取免费资源 - 联系我们**

#### 🔥 快速咨询（推荐）
- **微信**: `Pangolin-Scraper` 
  - 添加备注：`GitHub免费额度`
  - 工作时间内秒回，提供技术支持

<div>
  <img src="wechat-qrcode.png" alt="微信二维码" width="200"/>
  <br>
  <em>扫码添加微信获取免费资源</em>
</div>

#### 📧 邮件联系
- **技术支持**: shiyang@pangolinfo.com
- **商务合作**: charon@pangolinfo.com
- **响应时间**: 24小时内回复

#### 🌐 在线资源
- **项目文档**: [docs.pangolinfo.com](https://docs.pangolinfo.com)
- **官方网站**: [www.pangolinfo.com](https://www.pangolinfo.com)

---

## 📖 项目简介

**Pangolin Scrape API by PANGOLIN INFO TECH PTE. LTD.是一款专为跨境电商卖家、数据服务商和工具开发者设计的强大、稳定、高效的数据采集API。**
Scrape API 可动态兼容Amazon等各类电商页面结构变化。该接口通过智能识别算法自动识别并提取相关产品数据，如标题、折扣、价格、可用性和描述等。开发者无需关注目标页面DOM结构变更，系统将持续维护数据解析逻辑，显著降低电商数据集成与维护成本，支持通过API密钥快速调用并获取实时数据。

**主要支持平台：亚马逊 (Amazon), 沃尔玛 (Walmart)。同时也支持 Shopify, Shopee, eBay 等其他主流站点 (更多详情请查阅官方文档)。**

### 🎯 核心优势

- **🔄 动态适配**: 智能识别算法自动适配页面结构变化，无需担心DOM更新
- **⚡ 高效稳定**: 99.9% 可用性保证，毫秒级响应时间
- **🌍 多平台支持**: 覆盖 Amazon、Walmart、eBay、Shopify、Shopee 等主流电商平台
- **📊 多格式输出**: 支持 JSON、HTML、Markdown 等多种数据格式
- **🔧 零维护成本**: 系统自动维护解析逻辑，开发者专注业务开发

## 🛍️ 支持的电商平台

| 平台 | 支持页面类型 | 数据字段 |
|------|-------------|----------|
| **Amazon** | 商品详情、关键词搜索、分类列表、卖家商品、热销榜、新品榜 | ASIN码、标题、价格、评分、评论数、图片、销量、卖家信息、商品描述等 30+ 字段 |
| **Walmart** | 商品详情、关键词搜索 | 商品ID、标题、价格、评分、评论数、图片、尺寸、颜色等 |
| **eBay** | 商品详情、搜索结果 | 商品ID、标题、价格、卖家信息、图片等 |
| **Shopify** | 商品详情、店铺数据 | 商品信息、库存、价格、变体等 |

---

## 最好的开源亚马逊解析器 (Open-Source Amazon Keyword Page Parser)

我们很高兴与社区分享一部分针对 **Amazon 关键词搜索结果页面** 的解析代码。这部分开源代码旨在帮助开发者理解和快速上手如何从 Amazon 关键词搜索结果的 HTML 页面中提取结构化数据。

**请注意：**
* 此开源解析器专注于**解析本地已有的 HTML 源文件**。
* 它本身**不包含**网络请求、IP代理、验证码处理等数据采集功能。
* 此解析器是根据特定时间的 Amazon 页面结构编写的，虽然我们努力维护，但 Amazon 页面结构可能会发生变化，届时您可能需要自行调整代码。具体用法，参考parser_test.go

## 🚀 快速开始

### 方式一：使用开源解析模板（学习参考）

```bash
# 克隆仓库
git clone https://github.com/your-repo/ecommerce-parser.git
cd ecommerce-parser

# 安装依赖
go mod tidy

# 运行测试
go test ./parser_test.go
```

### 方式二：使用免费API（推荐生产）

#### 1. 获取免费API密钥
```bash
# 联系微信: Pangolin-Scraper 获取免费token
# 或者访问: https://docs.pangolinfo.com/free-signup
```

#### 2. Python 快速示例
```python
import requests
import json

# 使用您的免费token
TOKEN = "your_free_token_here"

def get_amazon_product(asin):
    """免费获取Amazon商品数据"""
    url = "http://scrapeapi.pangolinfo.com/api/v1/scrape"
    headers = {
        'Authorization': f'Bearer {TOKEN}',
        'Content-Type': 'application/json'
    }
    
    payload = {
        "url": f"https://www.amazon.com/dp/{asin}",
        "parserName": "amzProductDetail",
        "formats": ["json"]
    }
    
    response = requests.post(url, headers=headers, json=payload)
    return response.json()

# 测试使用
if __name__ == "__main__":
    # 免费测试ASIN
    result = get_amazon_product("B0DYTF8L2W")
    print(json.dumps(result, indent=2, ensure_ascii=False))
```

#### 3. Node.js 示例
```javascript
const axios = require('axios');

class FreeEcommerceAPI {
    constructor(token) {
        this.token = token;
        this.baseURL = 'http://scrapeapi.pangolinfo.com/api/v1/scrape';
    }

    async getAmazonProduct(asin) {
        const payload = {
            url: `https://www.amazon.com/dp/${asin}`,
            parserName: 'amzProductDetail',
            formats: ['json']
        };

        const response = await axios.post(this.baseURL, payload, {
            headers: {
                'Authorization': `Bearer ${this.token}`,
                'Content-Type': 'application/json'
            }
        });

        return response.data;
    }
}

// 使用免费API
const api = new FreeEcommerceAPI('your_free_token');
api.getAmazonProduct('B0DYTF8L2W')
   .then(data => console.log(data))
   .catch(err => console.error(err));
```

## 🔧 开源解析模板详解

### Amazon关键词页面解析器（开源模板）

```go
// 这是我们开源的Amazon关键词搜索解析模板
package parser

import (
    "github.com/PuerkitoBio/goquery"
    "strings"
)

type AmazonProduct struct {
    ASIN      string  `json:"asin"`
    Title     string  `json:"title"`
    Price     float64 `json:"price"`
    Rating    float64 `json:"rating"`
    Reviews   int     `json:"reviews"`
    ImageURL  string  `json:"image_url"`
}

func ParseAmazonSearchPage(html string) ([]AmazonProduct, error) {
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
    if err != nil {
        return nil, err
    }
    
    var products []AmazonProduct
    
    // 解析商品列表
    doc.Find("[data-asin]").Each(func(i int, s *goquery.Selection) {
        asin, _ := s.Attr("data-asin")
        if asin == "" {
            return
        }
        
        product := AmazonProduct{
            ASIN:     asin,
            Title:    s.Find("h2 a span").Text(),
            ImageURL: s.Find("img").AttrOr("src", ""),
        }
        
        // 解析价格
        priceText := s.Find(".a-price-whole").First().Text()
        // 价格解析逻辑...
        
        // 解析评分
        ratingText := s.Find(".a-icon-alt").First().Text()
        // 评分解析逻辑...
        
        products = append(products, product)
    })
    
    return products, nil
}
```

## 📊 功能对比

| 功能特性 | 开源模板 | 免费API | 付费API |
|----------|----------|---------|---------|
| 解析模板获取 | ✅ 关键词搜索 | ❌ | ❌ |
| 免费额度 | ✅ 无限制 | ✅ 200次/月 | ❌ |
| 技术支持 | ✅ 社区支持 | ✅ 官方支持 | ✅ VIP支持 |
| 自定义修改 | ✅ 完全自由 | ❌ | ✅ 定制服务 |
| 反爬处理 | 🔧 需自建 | ✅ 内置 | ✅ 高级 |
| 稳定性 | 🔧 自维护 | ✅ 99% | ✅ 99.9% |

## 🤝 开源贡献指南

我们热烈欢迎社区贡献！

### 🌟 如何贡献
1. **Fork** 本仓库
2. **创建** 您的功能分支 (`git checkout -b feature/AmazingFeature`)
3. **提交** 您的更改 (`git commit -m 'Add some AmazingFeature'`)
4. **推送** 到分支 (`git push origin feature/AmazingFeature`)
5. **提交** Pull Request

### 🏆 贡献者福利
- ✅ **专属徽章** - GitHub profile展示
- ✅ **免费API额度** - 贡献者专享
- ✅ **技术交流群** - 与核心开发者直接沟通
- ✅ **优先支持** - 问题优先解决

### 🎯 我们需要的贡献
- 🐛 **Bug修复** - 发现和修复解析问题
- 🆕 **新平台支持** - 增加新的电商平台解析器
- 📚 **文档完善** - 改进使用文档和示例
- 🧪 **测试用例** - 添加更多测试覆盖
- 🌍 **国际化** - 支持更多国家和语言

## 📈 使用场景

### 🎓 **学习研究**
- 📖 理解网页解析原理
- 🔬 数据科学项目
- 🎯 爬虫技术学习
- 📊 市场分析研究

### 🚀 **个人项目**
- 🛒 价格监控工具
- 📈 商品趋势分析
- 🔍 选品助手开发
- 📱 个人数据看板

### 🏢 **商业应用**
- 💼 竞品分析平台
- 📊 电商数据服务
- 🤖 自动化运营工具
- 📈 市场研究报告

## 📚 完整文档

- **📖 快速开始指南**: [docs.pangolinfo.com/quick-start](https://docs.pangolinfo.com/quick-start)
- **🔧 API参考文档**: [docs.pangolinfo.com/api](https://docs.pangolinfo.com/api)
- **💡 最佳实践**: [docs.pangolinfo.com/best-practices](https://docs.pangolinfo.com/best-practices)
- **❓ 常见问题**: [docs.pangolinfo.com/faq](https://docs.pangolinfo.com/faq)

## 🗺️ 开发路线图

### 🔥 近期计划 (Q2 2025)
- [ ] 完善Amazon解析器文档
- [ ] 增加Walmart解析器开源代码
- [ ] 提供Docker化部署方案
- [ ] 添加更多编程语言SDK

### 🚀 长期规划 (2025年)
- [ ] 支持更多电商平台
- [ ] 机器学习解析优化
- [ ] 云端部署方案
- [ ] 可视化配置界面

## ⭐ Star History

[![Star History Chart](https://api.star-history.com/svg?repos=your-repo/ecommerce-parser&type=Date)](https://star-history.com/#your-repo/ecommerce-parser&Date)

## 📄 开源协议

本项目采用 [MIT License](LICENSE) 开源协议。

---

## 🎉 立即开始

**不要犹豫，立即获取您的免费资源！**

1. ⭐ **给项目点个Star**
2. 📱 **添加微信**: `Pangolin-Scraper` (备注：GitHub免费)
3. 🚀 **开始您的数据采集之旅**

**让数据采集变得简单，让开源精神传递！** 🌟

---

*如果这个项目对您有帮助，请不要忘记给我们一个 ⭐Star！这是对我们最大的鼓励！*
