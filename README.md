# 🚀 电商数据采集 API - 专业的跨境电商数据解决方案

[![API Status](https://img.shields.io/badge/API-在线-brightgreen)](https://docs.pangolinfo.com)
[![支持平台](https://img.shields.io/badge/支持平台-Amazon%20|%20Walmart%20|%20eBay%20|%20Shopify-blue)](https://www.pangolinfo.com)
[![数据格式](https://img.shields.io/badge/数据格式-JSON%20|%20HTML%20|%20Markdown-orange)](https://docs.pangolinfo.com)

## 📖 项目简介

**Pangolin Scrape API by PANGOLIN INFO TECH PTE. LTD.是一款专为跨境电商卖家、数据服务商和工具开发者设计的强大、稳定、高效的数据采集API。**
Scrape API 可动态兼容Amazon等各类电商页面结构变化。该接口通过智能识别算法自动识别并提取相关产品数据，如标题、折扣、价格、可用性和描述等。开发者无需关注目标页面DOM结构变更，系统将持续维护数据解析逻辑，显著降低电商数据集成与维护成本，支持通过API密钥快速调用并获取实时数据。 [cite: 1]

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

### 1. 获取 API 密钥

```bash
curl -X POST http://scrapeapi.pangolinfo.com/api/v1/auth \
  -H 'Content-Type: application/json' \
  -d '{"email": "your-email@example.com", "password": "your-password"}'
```

### 2. 基础调用示例

#### Amazon 商品详情抓取

```bash
curl -X POST http://scrapeapi.pangolinfo.com/api/v1/scrape \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer YOUR_TOKEN' \
  -d '{
    "url": "https://www.amazon.com/dp/B0DYTF8L2W",
    "parserName": "amzProductDetail",
    "formats": ["json"],
    "bizContext": {
      "zipcode": "10041"
    }
  }'
```

#### 批量数据抓取

```bash
curl -X POST http://scrapeapi.pangolinfo.com/api/v1/scrape/batch \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer YOUR_TOKEN' \
  -d '{
    "urls": [
      "https://www.amazon.com/dp/PRODUCT1",
      "https://www.amazon.com/dp/PRODUCT2"
    ],
    "formats": ["markdown"]
  }'
```

### 3. Python SDK 示例

```python
import requests
import json

class ScrapeAPI:
    def __init__(self, token):
        self.token = token
        self.base_url = "http://scrapeapi.pangolinfo.com/api/v1/scrape"
        self.headers = {
            'Content-Type': 'application/json',
            'Authorization': f'Bearer {token}'
        }
    
    def scrape_amazon_product(self, asin, zipcode="10041"):
        """抓取 Amazon 商品详情"""
        url = f"https://www.amazon.com/dp/{asin}"
        payload = {
            "url": url,
            "parserName": "amzProductDetail", 
            "formats": ["json"],
            "bizContext": {"zipcode": zipcode}
        }
        
        response = requests.post(self.base_url, 
                               headers=self.headers, 
                               data=json.dumps(payload))
        return response.json()
    
    def search_amazon_keyword(self, keyword, zipcode="10041"):
        """Amazon 关键词搜索"""
        url = f"https://www.amazon.com/s?k={keyword}"
        payload = {
            "url": url,
            "parserName": "amzKeyword",
            "formats": ["json"],
            "bizContext": {"zipcode": zipcode}
        }
        
        response = requests.post(self.base_url,
                               headers=self.headers,
                               data=json.dumps(payload))
        return response.json()

# 使用示例
api = ScrapeAPI("YOUR_TOKEN")

# 获取商品详情
product_data = api.scrape_amazon_product("B0DYTF8L2W")
print(json.dumps(product_data, indent=2, ensure_ascii=False))

# 关键词搜索
search_results = api.search_amazon_keyword("wireless headphones")
print(json.dumps(search_results, indent=2, ensure_ascii=False))
```

### 4. Node.js 示例

```javascript
const axios = require('axios');

class ScrapeAPI {
    constructor(token) {
        this.token = token;
        this.baseURL = 'http://scrapeapi.pangolinfo.com/api/v1/scrape';
        this.headers = {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        };
    }

    async scrapeAmazonProduct(asin, zipcode = '10041') {
        const payload = {
            url: `https://www.amazon.com/dp/${asin}`,
            parserName: 'amzProductDetail',
            formats: ['json'],
            bizContext: { zipcode }
        };

        try {
            const response = await axios.post(this.baseURL, payload, {
                headers: this.headers
            });
            return response.data;
        } catch (error) {
            throw new Error(`API 调用失败: ${error.message}`);
        }
    }

    async batchScrape(urls, formats = ['json']) {
        const payload = {
            urls,
            formats
        };

        try {
            const response = await axios.post(`${this.baseURL}/batch`, payload, {
                headers: this.headers
            });
            return response.data;
        } catch (error) {
            throw new Error(`批量抓取失败: ${error.message}`);
        }
    }
}

// 使用示例
(async () => {
    const api = new ScrapeAPI('YOUR_TOKEN');
    
    try {
        // 单个商品抓取
        const productData = await api.scrapeAmazonProduct('B0DYTF8L2W');
        console.log('商品数据:', JSON.stringify(productData, null, 2));
        
        // 批量抓取
        const batchData = await api.batchScrape([
            'https://www.amazon.com/dp/PRODUCT1',
            'https://www.amazon.com/dp/PRODUCT2'
        ]);
        console.log('批量数据:', JSON.stringify(batchData, null, 2));
        
    } catch (error) {
        console.error('错误:', error.message);
    }
})();
```

## 📊 API 解析器说明

### Amazon 解析器

| 解析器名称 | 用途 | 返回字段 |
|-----------|------|----------|
| `amzProductDetail` | 商品详情页 | asin、title、price、star、rating、image、sales、seller、brand、description 等 30+ 字段 |
| `amzKeyword` | 关键词搜索 | asin、title、price、star、rating、image、sales |
| `amzProductOfCategory` | 分类商品列表 | asin、title、price、star、rating、image |
| `amzProductOfSeller` | 卖家商品列表 | asin、title、price、star、rating、image |
| `amzBestSellers` | 热销榜 | rank、asin、title、price、star、rating、image |
| `amzNewReleases` | 新品榜 | rank、asin、title、price、star、rating、image |

### Walmart 解析器

| 解析器名称 | 用途 | 返回字段 |
|-----------|------|----------|
| `walmProductDetail` | 商品详情页 | productId、img、title、star、rating、size、color、desc、price、hasCart |
| `walmKeyword` | 关键词搜索 | productId、image、title、star、rating、price |

## 🌍 支持的地区和邮编

| 国家 | 支持邮编 |
|------|----------|
| 美国 | 10041, 90001, 60601, 84104 |
| 英国 | W1S 3AS, EH15 1LR, M13 9PL, M2 5BQ |
| 法国 | 75000, 69001, 06000, 13000 |
| 德国 | 80331, 10115, 20095, 60306 |

## 💰 计费说明

| 数据格式 | 消耗积分 |
|----------|----------|
| markdown | 0.75 积分/次 |
| rawHtml | 0.75 积分/次 |
| json | 1 积分/次 |

## 🔧 高级功能

### 异步处理

对于大批量数据处理，我们提供异步 API：

```bash
curl -X POST https://extapi.pangolinfo.com/api/v1 \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer YOUR_TOKEN' \
  -d '{
    "url": "https://www.amazon.com/dp/B0DYTF8L2W",
    "callbackUrl": "https://your-domain.com/callback",
    "bizKey": "amzProduct",
    "zipcode": "10041"
  }'
```

### 自定义解析需求

我们支持根据您的业务需求定制解析字段，提交需求后我们将在每周的迭代中更新解析引擎。

## 🛠️ 技术特点

- **智能反爬**: 内建多种反反爬策略，确保抓取成功率
- **分布式架构**: 全球多节点部署，就近提供服务
- **实时监控**: 24/7 系统监控，确保服务稳定性
- **数据一致性**: 多重校验机制，保证数据准确性
- **弹性扩容**: 自动扩容机制，应对流量高峰

## 📚 完整文档

- **API 文档**: [docs.pangolinfo.com](https://docs.pangolinfo.com)
- **用户指南**: [Amazon Data API 用户指南](https://www.pangolinfo.com/amazon-data-api-user-guide/)
- **官方网站**: [www.pangolinfo.com](https://www.pangolinfo.com)

## 🤝 适用场景

### 🎯 跨境电商卖家
- **选品分析**: 批量获取竞品数据，分析市场趋势
- **价格监控**: 实时监控竞争对手价格变化
- **库存管理**: 跟踪热销商品库存状态
- **评论分析**: 获取用户评价数据，优化产品

### 🔧 工具开发商
- **数据服务**: 为卖家精灵、SIF 等工具提供数据支撑
- **市场分析**: 构建商品分析和趋势预测工具
- **自动化运营**: 开发自动定价、库存管理系统
- **竞品监控**: 构建竞争对手监控平台

### 📊 数据服务商
- **数据清洗**: 获取原始数据进行二次加工
- **报告生成**: 生成行业分析报告
- **API 集成**: 集成到现有数据平台
- **定制化服务**: 为企业客户提供定制数据服务

## 🆚 与其他方案对比

| 特性 | 我们的 API | 自建爬虫 | 其他 API 服务 |
|------|-----------|----------|---------------|
| **开发成本** | ⭐⭐⭐⭐⭐ 低 | ⭐⭐ 高 | ⭐⭐⭐ 中 |
| **维护成本** | ⭐⭐⭐⭐⭐ 零维护 | ⭐ 持续维护 | ⭐⭐⭐ 中等 |
| **稳定性** | ⭐⭐⭐⭐⭐ 99.9% | ⭐⭐ 不稳定 | ⭐⭐⭐⭐ 较好 |
| **数据质量** | ⭐⭐⭐⭐⭐ 结构化 | ⭐⭐⭐ 需处理 | ⭐⭐⭐⭐ 较好 |
| **反爬能力** | ⭐⭐⭐⭐⭐ 专业 | ⭐⭐ 需自建 | ⭐⭐⭐ 一般 |
| **扩展性** | ⭐⭐⭐⭐⭐ 弹性 | ⭐⭐ 受限 | ⭐⭐⭐ 中等 |

## 📞 联系我们

### 商务合作
- **邮箱**: charon@pangolinfo.com
- **公司**: PANGOLIN INFO TECH PTE. LTD.
- ### 微信咨询
- **微信号**: Pangolin-Scraper
- **添加时请备注**: GitHub-API咨询

### 技术支持
- **邮箱**: shiyang@pangolinfo.com
- **响应时间**: 工作日 24 小时内回复

### 在线支持
- **官网**: [www.pangolinfo.com](https://www.pangolinfo.com)
- **文档**: [docs.pangolinfo.com](https://docs.pangolinfo.com)

## 📝 开源贡献

我们欢迎社区贡献和反馈：

1. **Issues**: 提交 bug 报告或功能请求
2. **文档改进**: 帮助完善文档和示例代码
3. **SDK 开发**: 贡献不同语言的 SDK
4. **使用案例**: 分享您的使用经验和最佳实践

## 📄 许可证

本项目仅用于展示 API 功能和使用方法，具体服务条款请参考官方协议。

---

⭐ **如果这个项目对您有帮助，请给我们一个 Star！**

**让数据采集变得简单，让跨境电商更高效！** 🚀
