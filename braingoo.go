package braingoo

type Interpreter struct {
	source    string
	i         int // Current index of source
	tape      [30000]int8
	pointer   int
	loopStack []int  // Indices of loop brackets from source
	input     []int8 // Input values FIFO
	output    string
}

func NewInterpreter(source string, input []int8) Interpreter {
	return Interpreter{source, 0, [30000]int8{}, 0, make([]int, 0), input, ""}
}

func (ip *Interpreter) increment() {
	ip.tape[ip.pointer]++
}

func (ip *Interpreter) nextInput() {
	ip.tape[ip.pointer], ip.input = ip.input[0], ip.input[1:]
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
	for ip.source[ip.i] != 93 || level != 0 {
		if ip.source[ip.i] == 91 {
			level++
		} else if ip.source[ip.i] == 93 {
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
	for ip.i = 0; ip.i < len(ip.source); ip.i++ {
		switch ip.source[ip.i] {
		case 43:
			ip.increment()
		case 44:
			ip.nextInput()
		case 45:
			ip.decrement()
		case 46:
			ip.output += string(ip.tape[ip.pointer])
		case 60:
			ip.shiftLeft()
		case 62:
			ip.shiftRight()
		case 91:
			ip.startLoop()
		case 93:
			ip.endLoop()
		}
	}
}
