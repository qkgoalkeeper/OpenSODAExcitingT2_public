

![Untitled](./assets/fig/Untitled.png)

## 1.Introduction

### 1.1 项目背景

本项目的目标是开发一个命令行工具，为用户提供一种交互式的方式，用于查询和获取 OpenDigger 平台上的数据。OpenDigger 是一个提供各种统计型和网络型指标的平台，用户可以通过修改 HTTPS URL 来获取特定仓库或开发者在各项指标上的数据结果。为了更方便地查询和浏览这些数据，我们希望开发一种新型的指标结果查询方式，即通过命令行实现可交互的指标结果查询。

本命令行工具的核心功能包括查询特定仓库在特定指标上的数据、查询特定仓库在特定自然月上在特定指标上的数据、查询特定仓库在特定自然月份上的整体报告、以及特定仓库在全部自然月份上的报告，除此以外本工具还支持对github上的火热项目进行批量分析生成报告。同时，该工具还支持将查询结果导出到本地文件。

通过该工具，用户可以方便地在终端中进行各项指标的查询，探索仓库或开发者的数据情况。用户可以根据自己的需求进行定制化的查询，例如只查询特定指标或只查询特定类型的数据。

该命令行工具的设计旨在提供简单易用、高效便捷的查询功能，帮助用户更好地了解和分析 GitHub 上的开源项目和开发者的数据表现。通过这种交互式的查询方式，用户可以快速获取所需的数据结果，并进行进一步的分析和研究。

该项目的atomgit仓库见[此](https://atomgit.com/wsytcl/2023-opensoda-final-t2).

### 1.2 命令行框架 Cobra

本次项目基于Cobra框架进行开发，Cobra 是一个 Go 语言开发的命令行框架，它提供了**简洁、灵活且强大**的方式来创建命令行程序。它包含一个用于创建命令行程序的库，以及一个用于快速生成基于 Cobra 库的命令行程序工具。Cobra具有以下特性和优势：

- 简单易用：Cobra提供了直观的API和清晰的设计，使得构建命令行应用程序变得简单而直观。开发者可以快速定义命令、子命令和相关的标志和参数。
- 子命令支持：Cobra支持创建具有层次结构的命令行应用程序，即主命令下可以有多个子命令。这种结构使得应用程序的命令组织更清晰、更易扩展。
- 命令行补全：Cobra支持命令行补全功能，用户可以通过按Tab键自动补全命令、子命令、标志和参数，提高了交互性和用户体验

### 1.3 运行准备

本项目采用golang开发，需要运行环境具有golang 1.18 及以上的版本 / Docker部署。

1. 编译原代码获得可执行文件。

   ```bash
   go build
   ```

   可以进一步将可执行文件安装进系统中。
   ```bash
   go install
   ```
2. Docker部署
   拉取docker image。
   ```bash
   docker pull yinzhengsun/exciting-opendigger:v1.0
   ```

   进入容器
   ```bash
   docker run -it yinzhengsun/exciting-opendigger:v1.0 /bin/bash
   ```

3. 执行可执行文件

   ```bash
   ./exciting-opendigger commands
   ```
   在将文件安装进系统后，可以在任意位置直接运行该文件
   ```bash
   exciting-opendigger commands
   ```
   
   具体的执行命令见下，其中均假定未将文件安装入系统中：

## 2. Usage

Exciting-Opendigger提供了多样的查询功能。为了更好地满足精细化的查询需求，我们给出了一系列的命令选项，其整体语法如下：

![USAGESTMT.svg](./assets/fig/USAGESTMT.svg)

通过这一系列命令，Exciting-Opendigger满足了开源爱好者在查询贡献指标上的不同需求。如下表所示：

| 模块 | 命令 | 结果 | 需求 |
| --- | --- | --- | --- |
| 单点查询 | SHOW | 单个仓库或开发者在特定指标、特定月份的数据或整体报告 | 对开源指标的简单浏览 |
| 比较查询 | COMPARE | 对两个仓库、开发者、月份的比较结果 | 对开源指标的差异分析 |
| 文件下载 | DOWNLOAD | 查询结果的文件和可视化分析 | 对开源指标的详细分析和可视化 |
| 批量分析 | BATCH | 批量对象的数据 | 对Github社区的宏观分析 |
| 日志查询 | LOG | 历史查询日志的查看和回放 | 连续使用本工具的便捷性 |
| 帮助菜单 | HELP | 使用说明 | 辅助工具的便捷使用 |

### 2.0 视频教程

我们提供了一个[视频教程](https://www.bilibili.com/video/BV1mH4y1f7EM/?vd_source=6034ef2a6f692353baf55e8cbcd00df3)，便于使用者快速上手该工具。

### 2.1 **基础查询**

**2.1.1 单点查询**

其中，`SHOW`命令提供了基础的查询功能，即对特定对象的指标查询。此处我们提供一系列的查询参数于`SearchOpt`如下：

![SearchOpt.svg](./assets/fig/SearchOpt.svg)

在基础查询中，我们支持对仓库和开发者两个维度的指标查询，并且可以使用两个参数，`month`和`metric`，约束查询对应的月份和指标。用户需要至少提供一个参数以约束查询结果。选择不同参数约束的结果如下：

| month | metric | 结果 |
| --- | --- | --- |
| 填写 | 填写 | 特定指标在特定月份上的结果 |
| 填写 | 忽略 | 特定月份上的整体报告 |
| 忽略 | 填写 | 特定指标的长期趋势 |

用例：

![1689325868868](./assets/fig/showCase.png)

> 开发日志：该模块需要与网络交互，访问OpenDigger提供的网络API接口。这部分已经基本完成。

### 2.2 指标比较

除了单个仓库的查询，Exciting-Opendigger也提供多个对象间的比较。当启用`COMPARE`子命令时，开源爱好者可以通过追加新的对象或月份的方式比较两个指标数据的差异。

用例：

![1689326073589](./assets/fig/compareCase.png)


### 2.3 文件下载

Exciting-Opendigger支持使用`DOWNLOAD`命令下载基础查询的结果。考虑到数据下载往往用于详细分析，我们为下载文件进行可视化的数据分析，将结果转化为html输出。我们同样提供`SearchOpt`用于约束查询的内容，由于文件下载具有更强的表达能力，其支持不使用`SearchOpt`，直接输出所有的结果。`DownloadClause`的具体语法如下：

![DownloadClause.svg](./assets/fig/DownloadClause.svg)

用例：

![1689322132927](./assets/fig/downloadCase.png)

执行结果（部分）

![1689322084150](./assets/fig/downloadReport.png)


### 2.4 批量分析

Exciting-Opendigger为开源爱好者进一步提供了批量分析的功能，从而鸟瞰Github社区的整体情况。`BatchClause`的具体语法如下：

![BatchClause.svg](./assets/fig/BatchClause.svg)

当使用`BATCH`命令进行批量分析时，用户可以请求两种不同的数据来源。首先，可以使用`TOP`参数请求查询当前Github最活跃仓库；其次，也可以采用提供文件的方式，重点关注特定的仓库。由于批量分析的数据量往往较多，直接输出到屏幕不利于分析数据。我们提供了文件下载的方式提供查询的结果。Exciting-Opendigger支持查询不同语言、编程语言、时间周期中的最活跃仓库。`BatchOpt`的具体语法如下，这部分约束只会在数据来源为最活跃仓库时起效，此时可约束查询的最活跃仓库的范围：

![BatchOpt.svg](./assets/fig/BatchOpt.svg)



用例：

![1689390060134](./assets/fig/batchCase.png)

执行结果如下：

![1689323036671](./assets/fig/batchResult.png)

### 2.5 日志查询

为了提供更为连续的操作体验，Exciting-Opendigger默认在查询时缓存查询命令。随后，用户可以使用`LOG`命令查询之前的指标查询历史和结果。

用例：

![1689248846120](./assets/fig/logCase.png)


### 2.6 帮助菜单

由于Exciting-Opendigger提供了多样化的查询服务，其查询命令和参数较为详细。为便于用户使用这一工具，我们提供了全面的帮助菜单以展示可选操作和参数，解除了用户的记忆压力，降低使用门槛。

用例：

![1689249550000](./assets/fig/helpCase.png)


## 3. Demo

命令行接口展示，此处展示了工具入口的帮助菜单，显示了Exciting-Opendigger提供的命令类型和用法说明。

![1689249550000](./assets/fig/helpCase.png)

html版本报告demo：

输出的报告与用户需求有关，将查询的数据进行格式整理以html格式输出，查询的指标和参数可以灵活设置，并且支持搜索具体指标。

下图的demo展示了X-lab2017/open-digger项目的一个分析报告，这里以openrank指标为例，输出了X-lab2017/open-digger项目近五个月的openrank值，并以柱状图的形式输出了X-lab2017/open-digger项目立项以来openrank的月份变化趋势。

![report_html.jpg](./assets/fig/report_pdf.jpg)

## 4.项目成员和分工

- 翁思扬   交互模块，开发管理
- 钱堃       输出模块
- 梁辉       查询模块
- 孙印政   查询模块

