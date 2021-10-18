# qiyu

*代办*

- [x] 基本能运行起来
- [ ] list接口
    - [x] 分页
    - [x] 搜索
        - [ ] 拼接到db查询方法上
    - [ ] 排序
    - [x] 过滤
        - [x] 返回结果的时间等格式化(重新定义一个结构体)
        - [x] 数据表中字段可读，改为了time.Time
    - [x] 数据表迁移
- [x] Retrieve接口

- [ ] 补齐其他接口
- [ ] form改成json
- [ ] 跨表查询
    - [ ] 左联右联，条件放在别的表上
- [ ] 多环境切换
- [ ] 异步机制 与消息队列
- [ ] 定时周期任务 gocron
    这俩可以参考 crawlab
- [ ] 权限rbac
- [ ] 缓存机制
- [ ] 自动生成表
  - [ ] 一键初始化数据

写到表里的字段 为可读


1. 如何创建多对多的关系？
   1. 先create当前
   2. 当前.Association("外键").append([]Tag)
   > d.engine.Model(&article).Association("Tag").Append(ts)
2. 修改的时候传入的是map那怎么改关联？
   > d.engine.Model(&article).Association("Tag").Replace(ts)