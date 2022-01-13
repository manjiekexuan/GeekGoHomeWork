# GeekGoHomeWork
极客时间 Go基础训练营 作业
## 第一次作业
实现体脂率和计算器

## 第二次作业
内容：为体脂计算器编写单元测试并完善体脂计算器的验证逻辑。
课后作业 1
内容：使用 github 上的 lib： github.com/armstrongli/go-bmi 完成体脂计算器。
本地添加 module 的 replace，并在本地项目扩展 github.com/armstrongli/go-bmi 以支持 BMI、FatRate 的计算。
使用 vendor 保证代码的完整性与可运行。
#### homework_2_module_replace
### 课后作业 2
内容：为体脂计算器编写单元测试并完善体脂计算器的验证逻辑。
BMI 计算：
录入正常身高、体重，确保计算结果符合预期；
录入 0 或负数身高，返回错误；
录入 0 或负数体重，返回错误。
体脂率计算：
录入正常 BMI、年龄、性别，确保计算结果符合预期；
录入非法 BMI、年龄、性别（0、负数、超过 150 的年龄、非男女的性别输入），返回错误；
录入完整的性别、年龄、身高、体重，确保最终获得的健康建议符合预期
#### homework_2_fatRateTesting
