# [File](../file/easy_py_7857d04e6ba0e11d8f5cda2bcae66b1b.zip)

<img src="easy_py.assets/2020-07-18_14-23.png" alt="反编译报错" style="zoom: 50%;" />

加载常量，范围溢出，修改成0试试

<img src="easy_py.assets/qwe.png" style="zoom:30%;" />



反编译结果



<img src="easy_py.assets/2020-07-18_14-29.png" alt="2020-07-18_14-29" style="zoom:38%;" />

跳转到6 ,9 标签根本没有6，9 判断为混淆指令，手动删掉这三条指令

这个字段记录着操作码的长度，删除三条指令长度为9个字节，然后将178-9

 ![qwe](easy_py.assets/qwe-1595053891298.png)



最终反编译结果:

![2020-07-18_14-32](easy_py.assets/2020-07-18_14-32.png)



简化一下，直接爆破省事

![wqe](easy_py.assets/wqe.png)