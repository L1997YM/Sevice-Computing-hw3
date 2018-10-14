package main

import(
    "fmt"
    "os"
    "bufio"
    "github.com/spf13/pflag"
    "os/exec"
    "io"
)

type selpg_args struct{
    s_page int           //起始页
    e_page int           //结束页
    file_name string     //文件名
    page_line int        //每页行数
    is_f bool            //是否用换页符换页
    print_dest string    //打印地址
}

//初始化
func Parser(p *selpg_args){
    pflag.IntVarP(&p.s_page,"start","s",0,"首页")
    pflag.IntVarP(&p.e_page,"end","e",0,"尾页")
    pflag.IntVarP(&p.page_line,"linenum","l",4,"打印的每页行数")
    pflag.BoolVarP(&p.is_f,"arg_f","f",false,"是否用换页符")
    pflag.StringVarP(&p.print_dest,"dest","d","","打印目的地")
    pflag.Parse()
}

//参数处理
//selarg   selpg_args的一个结构体
func processArgs(selarg *selpg_args){
  var is_err bool = false
  //处理第1个参数不是-s的情况
  //Args的类型是[]string，获取运行时给出的参数
	if os.Args[1][0] != '-' || os.Args[1][1] != 's' {
		fmt.Fprintf(os.Stderr,"error：第一个参数必须为-s\n")
		is_err = true
		os.Exit(1)
	}

  //处理起始页参数出错的情况
	if selarg.s_page < 1 {
		fmt.Fprintf(os.Stderr,"error：起始页错误\n")
		is_err = true
		os.Exit(1)
	}

  //处理第3个参数不是-e的情况
	if os.Args[3][0] != '-' || os.Args[3][1] != 'e' {
		fmt.Fprintf(os.Stderr,"error：第三个参数必须是-e\n")
		is_err = true
		os.Exit(1)
	}

  //处理终止页出错的情况
	if selarg.e_page < 1 || selarg.e_page < selarg.s_page {
		fmt.Fprintf(os.Stderr,"error：终止页错误\n")
		is_err = true
		os.Exit(1)
	}

  //处理每页行数大小出错的情况
	if selarg.page_line < 1 {
		fmt.Fprintf(os.Stderr,"error：页行数错误\n")
		is_err = true
		os.Exit(1)
	}

  if is_err == true {
    fmt.Fprintf(os.Stderr,"正确的指令格式：selpg -s Number -e Number [options] [filename]\n")
  }

  if pflag.NArg() > 0 {
    //选择读取的文件名为非pflag命令行第一个参数
	  selarg.file_name = pflag.Arg(0)
    //打开文件，返回文件描述符
	  file, err := os.Open(selarg.file_name)
	  if err != nil {
		  fmt.Fprintf(os.Stderr,"error:文件\"%s\"不存在\n",selarg.file_name)
		  os.Exit(1)
	  }
    //以只读方式打开文件
	  file, err = os.OpenFile(selarg.file_name,os.O_RDONLY,0666)
	  if err != nil {
		  if os.IsPermission(err) {
			  fmt.Fprintf(os.Stderr,"error:文件\"%s\"无法读取\n",selarg.file_name)
			  os.Exit(1)
		  }
	  }
  //关闭文件，使其不能够再进行I/O操作
	file.Close()
  }
}

//处理文件内容
func processInput(selarg * selpg_args) {
	fin := os.Stdin
	fout := os.Stdout
	var (
		page_num int    //当前页数
		line_num int    //当前行数
		err1 error
		err2 error
		err3 error
                err4 error
		line string     //当前行数据
                char1 byte      //当前字符
		cmd *exec.Cmd   //Cmd结构表示一个正在准备或者正在运行的外部命令
		stdin io.WriteCloser  //Write方法和Closer方法的结合
	)

	if selarg.file_name != "" {
		fin, err1 = os.Open(selarg.file_name)
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "error:无法打开文件\"%s\"\n",selarg.file_name)
			os.Exit(1)
		}
	}

        //print_dest 打印地
	if selarg.print_dest != "" {
        //exec.Command返回cmd结构来执行带有相关参数的命令（Path参数和Args参数）
		cmd = exec.Command("cat","-n")
        //返回一个连接到command标准输入的管道pipe
		stdin, err2 = cmd.StdinPipe()
		if err2 != nil {
        //采用默认格式将其参数格式化并写入标准输出。总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符。
			fmt.Println(err2)
		}
	} else {
		stdin = nil
	}

  //bufio包  实现缓存IO
  //rd 底层的io.Reader
  //NewReader相当于NewReaderSize(rd, 4096)
  //NewReaderSize将rd封装成一个带缓存的 bufio.Reader对象，缓存大小由size指定（如果小于16则会被设置为16）。
	rd := bufio.NewReader(fin)
        page_num = 1
        line_num = 0
        //不用换页符换页
	if selarg.is_f == false {
		for true {
			line, err3 = rd.ReadString('\n')
			if err3 != nil {
				break
			}
			line_num ++
			if line_num > selarg.page_line {
				page_num ++
				line_num = 1
			}
			if page_num >= selarg.s_page && page_num <= selarg.e_page {
        //当前行数据缓存到fout中
				fmt.Fprintf(fout, "%s", line)
			}
		}
	} else {
		for true {
			char1, err4 = rd.ReadByte()
			if err4 != nil {
				break
			}
      //换页符'\f'
			if char1 == '\f' {
				page_num ++
			}
			if page_num >= selarg.s_page && page_num <= selarg.e_page {
				fmt.Fprintf(fout, "%c", char1)
			}
		}
		fmt.Print("\n")
	}

  if page_num < selarg.s_page {
    fmt.Fprintf(os.Stderr,"error:起始页%d大于该文件的总页数%d\n",selarg.s_page,page_num)
  } else if page_num < selarg.e_page {
    fmt.Fprintf(os.Stderr,"error:终止页%d大于该文件的总页数%d\n",selarg.e_page,page_num)
  }

  if selarg.print_dest != "" {
    stdin.Close()
    cmd.Stdout = fout
    cmd.Run()
  }
  fin.Close()
  fout.Close()
}

func main() {
  sa := selpg_args{0,0,"",5,false,""}
  Parser(&sa)
  processArgs(&sa)
  processInput(&sa)
}
