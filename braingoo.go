package braingoo

type Interpreter struct {
	input     string
	i         int // Current index of input
	tape      [30000]int8
	pointer   int
	loopStack []int // Holds indices of loop brackets from input
	output    string
}

func NewInterpreter(input string) Interpreter {
	return Interpreter{input, 0, [30000]int8{}, 0, make([]int, 0), ""}
}

func (ip *Interpreter) increment() {
	ip.tape[ip.pointer]++
}

func (ip *Interpreter) decrement() {
	ip.tape[ip.pointer]--
}

func (ip *Interpreter) shiftLeft() {
	if ip.pointer == 0 {
		ip.pointer = 29999
	} else {
		ip.pointer--
	}
}

func (ip *Interpreter) shiftRight() {
	if ip.pointer == 29999 {
		ip.pointer = 0
	} else {
		ip.pointer++
	}
}

func (ip *Interpreter) startLoop() {
	if ip.tape[ip.pointer] != 0 {
		ip.loopStack = append(ip.loopStack, ip.i)
	} else {
		ip.jump()
	}
}

func (ip *Interpreter) jump() {
	level := -1
	for ip.input[ip.i] != 93 || level != 0 {
		if ip.input[ip.i] == 91 {
			level++
		} else if ip.input[ip.i] == 93 {
			level--
		}
		ip.i++
	}
}

func (ip *Interpreter) endLoop() {
	if ip.tape[ip.pointer] != 0 {
		ip.i = ip.loopStack[len(ip.loopStack)-1]
	} else if len(ip.loopStack) > 0 {
		ip.loopStack = ip.loopStack[:len(ip.loopStack)-1]
	}
}

func (ip *Interpreter) parse() {
	for ip.i = 0; ip.i < len(ip.input); ip.i++ {
		switch ip.input[ip.i] {
		case 43:
			ip.increment()
			break
		case 45:
			ip.decrement()
			break
		case 46:
			ip.output += string(ip.tape[ip.pointer])
			break
		case 60:
			ip.shiftLeft()
			break
		case 62:
			ip.shiftRight()
			break
		case 91:
			ip.startLoop()
			break
		case 93:
			ip.endLoop()
			break
		}
	}
}
