# prometheus-server

- dal: 数据访问层，负责对数据库的操作(与controller合并)  https://github.com/mongodb/mongo-go-driver
- utils: 工具函数
- logs: 日志存放位置
- middlewares: 中间件，主要用于配置Cors
- controller: 业务逻辑层
- models: 模型层，定义结构