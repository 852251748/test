package main

import (
	"fmt"
	"strconv"
	"strings"
)

//逆波兰式算数计算步骤：
//1.获取表达式(获取普通表达式)
//2.解析表达式(将普通表达式转换成逆波兰式表达式)
//3.计算逆波兰式表达式的值

//1.获取表达式(获取普通表达式)
type Node struct {
	data map[int]interface{}
	pos  int
}

func (this *Node) push(op interface{}) {
	this.data[this.pos] = op
	this.pos++
}

func (this *Node) pop() interface{} {
	if this.pos != 0 {
		temdata := this.data[this.pos-1]
		delete(this.data, this.pos-1)
		this.pos--
		return temdata
	}
	return nil
}

func (this *Node) popHead() {
	if this.pos != 0 {
		for i := 0; i < this.pos; i++ {
			expression = append(expression, this.data[i])
		}
	}
}

type any interface{}

var (
	stack      []int64
	Express    string
	Priority   map[string]int
	expression []any
)

func getExpression() {
	fmt.Println("Please input expression")
	//var expre string
	for {
		fmt.Scan(&Express)
		if Express[len(Express)-1] != '#' {
			fmt.Println("not ending with # .Please input again")
			continue
		}
		break
	}
}

//2.解析表达式(将普通表达式转换成逆波兰式表达式)
//将一个普通的中序表达式转换为逆波兰表达式的一般算法是：
//首先需要分配2个栈，一个作为临时存储运算符的栈S1（含一个结束符号），一个作为输入逆波兰式的栈S2（空栈），S1栈可先放入优先级最低的运算符#，注意，中缀式应以此最低优先级的运算符结束。可指定其他字符，不一定非#不可。从中缀式的左端开始取字符，逐序进行如下步骤：
//（1）若取出的字符是操作数，则分析出完整的运算数，该操作数直接送入S2栈
//（2）若取出的字符是运算符，则将该运算符与S1栈栈顶元素比较，如果该运算符优先级(不包括括号运算符)大于S1栈栈顶运算符优先级，则将该运算符进S1栈，否则，将S1栈的栈顶运算符弹出，送入S2栈中，直至S1栈栈顶运算符低于（不包括等于）该运算符优先级，最后将该运算符送入S1栈。
//（3）若取出的字符是“（”，则直接送入S1栈顶。
//（4）若取出的字符是“）”，则将距离S1栈栈顶最近的“（”之间的运算符，逐个出栈，依次送入S2栈，此时抛弃“（”。
//（5）重复上面的1~4步，直至处理完所有的输入字符
//（6）若取出的字符是“#”，则将S1栈内所有运算符（不包括“#”），逐个出栈，依次送入S2栈。
//完成以上步骤，S2栈便为逆波兰式输出结果。
func praseExpress() {
	var stack1, stack2 Node
	stack1 = struct {
		data map[int]interface{}
		pos  int
	}{data: make(map[int]interface{}, 10), pos: 0}
	stack2 = struct {
		data map[int]interface{}
		pos  int
	}{data: make(map[int]interface{}, 10), pos: 0}

	tempstr := strings.Split(Express, "")

	for _, v := range tempstr {
		num, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			stack2.push(num)
		} else if v == "#" {
			for {
				if tmp := stack1.pop(); tmp != nil {
					stack2.push(tmp.(string))
				} else {
					break
				}
			}
		} else if v == "(" {
			stack1.push(v)
		} else if v == ")" {
			for {
				tmp := stack1.pop().(string)
				if tmp != "(" {
					stack2.push(tmp)
				} else {
					break
				}
			}
		} else {
			for {
				tmp := stack1.pop()
				if tmp != nil {
					if tmp.(string) != "(" {
						if Priority[v] > Priority[tmp.(string)] {
							stack1.push(tmp)
							stack1.push(v)
							break
						} else {
							stack2.push(tmp)
						}
					} else {
						stack1.push(tmp)
						stack1.push(v)
						break
					}
				} else {
					stack1.push(v)
					break
				}
			}
		}
	}
	stack2.popHead()
	fmt.Println(expression)

}

func main() {
	stack = make([]int64, 10)
	Priority = make(map[string]int, 10)
	Priority["+"] = 1
	Priority["-"] = 1
	Priority["*"] = 2
	Priority["/"] = 2
	//var expression []any
	//inputreader := bufio.NewReader(os.Stdin)
	//fmt.Println("Please input expression")
	//for {
	//	input, err := inputreader.ReadString('\n')
	//	if err != nil {
	//		fmt.Println("read string fail!", err)
	//	}
	//	if input == "q\r\n" {
	//		break
	//	}
	//	variable := input[:len(input)-2]
	//	num, ok := strconv.ParseInt(variable, 10, 64)
	//	if ok != nil {
	//		expression = append(expression, variable)
	//	} else {
	//		expression = append(expression, num)
	//	}
	//}
	//fmt.Println(expression)
	//Caculate(expression)
	getExpression()
	praseExpress()
	Caculate(expression)
	//ab+c*ab+e/-  12+3*12+4/-  (1+2)*3-(5+3)/4
}

//3.计算逆波兰式表达式的值
//新建一个表达式,如果当前字符为变量或者为数字，则压栈，如果是运算符，则将栈顶两个元素弹出作相应运算，结果再入栈，最后当表达式扫描完后，栈里的就是结果。
func Caculate(expression []any) {

	var i int
	for _, v := range expression {
		switch v.(type) {
		case int64:
			stack[i] = v.(int64)
			i++
		case string:
			switch v.(string) {
			case "+":
				stack[i-2] = stack[i-2] + stack[i-1]
				i--
			case "-":
				stack[i-2] = stack[i-2] - stack[i-1]
				i--
			case "*":
				stack[i-2] = stack[i-2] * stack[i-1]
				i--
			case "/":
				stack[i-2] = stack[i-2] / stack[i-1]
				i--
			}
		}
	}
	fmt.Println(stack[i-1])
}
