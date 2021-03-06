package repl

import (
	"bufio"
	"fmt"
	"io"
	"flag"
	
	"monkey/lexer"
	"monkey/parser"
	"monkey/object"
	"monkey/evaluator"
	"monkey/compiler"
	"monkey/vm"
)

const PROMPT = ">> "
const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

var engine = flag.String("engine", "vm", "use 'vm' or 'eval'")

func Start(in io.Reader, out io.Writer) {
	flag.Parse()
	scanner := bufio.NewScanner(in)

	if *engine == "vm" {
		constants := []object.Object{}
		globals := make([]object.Object, vm.GlobalsSize)
		symbolTable := compiler.NewSymbolTable()
		for i, v := range object.Builtins {
			symbolTable.DefineBuiltin(i, v.Name)
		}

		for {
			fmt.Printf(PROMPT)
			scanned := scanner.Scan()
			if !scanned {
				return
			}
	
			line := scanner.Text()
			l := lexer.New(line)
			p := parser.New(l)
			program := p.ParseProgram()
			if len(p.Errors()) != 0 {
				printParserErrors(out, p.Errors())
				continue
			}
	
			comp := compiler.NewWithState(symbolTable, constants)
			err := comp.Compile(program)
			if err != nil {
				fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
				continue
			}
	
			code := comp.Bytecode()
			constants = code.Constants
	
			machine := vm.NewWithGlobalsStore(code, globals)
			err 	 = machine.Run()
			if err != nil {
				fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
				continue
			}
	
			stackTop := machine.LastPoppedStackElm()
			io.WriteString(out, stackTop.Inspect())
			io.WriteString(out, "\n")
		}

	} else {
		env := object.NewEnvironment()

		for {
			fmt.Printf(PROMPT)
			scanned := scanner.Scan()
			if !scanned {
				return
			}
	
			line := scanner.Text()
			l := lexer.New(line)
			p := parser.New(l)
			program := p.ParseProgram()
			if len(p.Errors()) != 0 {
				printParserErrors(out, p.Errors())
				continue
			}
	
			evaulated := evaulator.Eval(program, env)
			if evaulated != nil {
				io.WriteString(out, evaulated.Inspect())
				io.WriteString(out, "\n")
			}
		}
	}
}

func printParserErrors(out io.Writer, errors[] string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	}
}
