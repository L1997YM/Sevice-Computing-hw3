# Sevice-Computing-hw3
## 测试
首先创建input.txt,output.txt,error.txt三个文件，其中在input.txt中输入数据如图

  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/10.png)<br/>
***
1. $ selpg -s 1 -e 1 input.txt <br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/1.png)<br/>
***
2. $ hello | selpg -s 1 -e 3 -f <br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/2.png)<br/>
***
3. $ selpg -s 2 -e 5 input.txt > output.txt <br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/3.png)<br/>
***
4. $ selpg -s 2 -e 10 input.txt 2> error.txt <br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/4.png)<br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/5.png)<br/>
***
5. $ selpg -s 3 -e 5 -l 3 input.txt <br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/6.png)<br/>
***
6. $ selpg -s 1 -e 3 input.txt > output.txt 2> error.txt & <br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/7.png)<br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/8.png)<br/>
  ![](https://github.com/L1997YM/Sevice-Computing-hw3/blob/master/selpg_test/9.png)<br/>
***
