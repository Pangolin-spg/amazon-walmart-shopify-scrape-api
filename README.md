# amazon-walmart-shopify-scrape-api
Powerful Scrape API for Amazon, Walmart, Shopify, Shopee, eBay. Get product details, rankings, HTML, JSON, and Markdown. 专为电商卖家、数据服务商和工具开发者设计的数据采集API。
# 🚀 Pangolin Scrape API - 电商数据采集利器 

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Docs](https://img.shields.io/badge/API%20Docs-Read%20Here-blue)](https://docs.pangolinfo.com) [![Contact Us](https://img.shields.io/badge/Contact-Email%20Us-green)](mailto:your-contact-email@example.com) 
** Scrape API 是一款专为跨境电商卖家 (亚马逊、沃尔玛、Shopify、Shopee、eBay)、数据服务商 (如卖家精灵、SIF类似需求) 和跨境工具开发者设计的强大数据采集API。轻松获取商品详情、榜单排名、评论、价格等任何公开的电商平台数据。**

**支持的电商平台包括：亚马逊 (Amazon), 沃尔玛 (Walmart), Shopify, Shopee, eBay 等主流站点。**

---

## 核心功能 (Key Features)

* ✅ **广泛的平台支持**: 覆盖 Amazon (亚马逊), Walmart (沃尔玛), Shopify, Shopee, eBay 等。
* 📊 **多样化的数据类型**: 抓取商品详情 (Product Details), 搜索结果 (Search Results), 榜单排名 (Rankings/Best Sellers), 卖家信息 (Seller Information), 评论数据 (Reviews) 等。
* ⚙️ **灵活的数据格式**:
    * 原始 HTML 页面 (Raw HTML)
    * Markdown 格式 (Converted Markdown)
    * 结构化 JSON 数据 (Parsed Structured JSON Data)
* ⚡ **高效稳定**: 专为大规模、高并发的数据采集设计。
* 👨‍💻 **开发者友好**: 简洁易用的 API 接口，提供清晰的[API文档](https://docs.pangolinfo.com)。
* 🎯 **精准定位**: 特别适合需要进行**铺货 (Bulk Listing/Product Sourcing)** 和 **跟卖 (Price Monitoring/Competitor Analysis)** 的卖家。

---

## 为什么选择我们？(Why Choose Us?)

* **专注电商领域**: 深耕电商数据采集，更懂卖家和数据服务商的需求。
* **全面覆盖**: 支持您业务所需的大多数主流电商平台。
* **数据质量**: 提供高质量、准确的原始数据和解析后数据。
* **客户成功**: 助力亚马逊大卖、跨境数据服务商 、工具开发商提升效率，获取商业洞察。

---

## 快速上手 (Quick Start)

我们的 API 调用非常简单。以下是一个获取亚马逊商品详情的伪代码示例：

```http
GET [https://api.pangolinfo.com/scrape?api_key=YOUR_API_KEY&url=AMAZON_PRODUCT_URL&format=json](https://api.pangolinfo.com/scrape?api_key=YOUR_API_KEY&url=AMAZON_PRODUCT_URL&format=json)
