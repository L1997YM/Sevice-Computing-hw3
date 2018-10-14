# Sevice-Computing-hw3
## 测试
首先创建input.txt,output.txt,error.txt三个文件，其中在input.txt中输入数据如图

  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/10.png)<br/>
***
1. $ selpg -s 1 -e 1 input.txt <br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/1.png)<br/>
***
2. $ hello | selpg -s 1 -e 3 -f <br/>
“hello”的标准输出被 shell／内核重定向至 selpg 的标准输入，同时测试在参数-f作用下，根据换行符换行的情况。 <br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/2.png)<br/>
***
3. $ selpg -s 2 -e 5 input.txt > output.txt <br/>
标准输出被 shell／内核重定向至“output.txt” <br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/3.png)<br/>
***
4. $ selpg -s 2 -e 10 input.txt 2> error.txt <br/>
所有的错误消息被 shell／内核重定向至“error.txt” <br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/4.png)<br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/5.png)<br/>
***
5. $ selpg -s 3 -e 5 -l 3 input.txt <br/>
该命令通过可选参数-l将每页的行数设置为3 <br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/6.png)<br/>
***
6. $ selpg -s 1 -e 3 input.txt > output.txt 2> error.txt & <br/>
selpg 进程在后台运行，并且标准输出和标准错误都被重定向至文件。 <br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/7.png)<br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/8.png)<br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/9.png)<br/>
***
剩余部分之后再补齐。。。<br/>
主要参考学习的博客：https://blog.csdn.net/wyxwyx469410930/article/details/82952728#25__154
