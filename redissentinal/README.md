redis sentinal 实验

3 个从节点，1 个主节点，2 个 sentinel

go进行测试

使用 makefile 编译项目

使用方式：

```bash
cd src && make build && make run
```
如果成功，执行 `docker logs redissentinal_test_app` 可以看到持续输出 `get test22: jaychen`。

