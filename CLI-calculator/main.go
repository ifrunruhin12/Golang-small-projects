package main

import (
    "fmt"
    "math"
    "strconv"
    "strings"
)


var precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}


func welcome() {
    fmt.Println("Hello there! Welcome to my Basic CLI Calculator.")
    fmt.Println("Nothing fancy yet, But soon more feature  will be added")
    fmt.Println("You can perform basic operations for now")
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, bool) {
	if b != 0 {
		return a / b, true
	}
	return 0, false
}

func calculate(num1 float64, operator string, num2 float64) (float64, bool) {
	switch operator {
	case "+":
		return add(num1, num2), true
	case "-":
		return subtract(num1, num2), true
	case "*":
		return multiply(num1, num2), true
	case "/":
		return divide(num1, num2)
	default:
		fmt.Println("Error: Invalid operator.")
		return 0, false
	}
}

func formatOutput(result float64) string {
	if math.Mod(result, 1) == 0 {
		return fmt.Sprintf("%d", int(result))
	}
	return fmt.Sprintf("%f", result)
}


func manual() {
	fmt.Println("Manual for Development Version:")
	fmt.Println("Operators:")
	fmt.Println("  +  for Addition")
	fmt.Println("  -  for Subtraction")
	fmt.Println("  *  for Multiplication")
	fmt.Println("  /  for Division")
	fmt.Println("Example: 10+45*2-5/5")
	fmt.Println("(Follows BODMAS rule)")
}

// Convert infix expression to postfix (RPN)
func infixToPostfix(expr string) ([]string, error) {
	var output []string
	var operators []string
	tokens := tokenize(expr)

	for _, token := range tokens {
		if isNumber(token) {
			output = append(output, token)
		} else if isOperator(token) {
			for len(operators) > 0 && operators[len(operators)-1] != "(" && precedence[operators[len(operators)-1]] >= precedence[token] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		} else if token == "(" {
			operators = append(operators, token)
		} else if token == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return nil, fmt.Errorf("mismatched parentheses")
			}
			operators = operators[:len(operators)-1] // Remove "("
		} else {
			return nil, fmt.Errorf("invalid token: %s", token)
		}
	}

	// Append remaining operators
	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output, nil
}

// Evaluate RPN expression
func evaluatePostfix(tokens []string) (float64, error) {
	var stack []float64

	for _, token := range tokens {
		if isNumber(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				result = a / b
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return stack[0], nil
}

// Tokenize input string
func tokenize(expr string) []string {
	var tokens []string
	var current string

	for _, char := range expr {
		if strings.ContainsRune("+-*/()", char) {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			tokens = append(tokens, string(char))
		} else {
			current += string(char)
		}
	}

	if current != "" {
		tokens = append(tokens, current)
	}
	return tokens
}

// Check if a string is a number
func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// Check if a string is an operator
func isOperator(s string) bool {
	_, exists := precedence[s]
	return exists
}

// The main eval function
func eval(expr string) (float64, error) {
	postfix, err := infixToPostfix(expr)
	if err != nil {
		return 0, err
	}
	return evaluatePostfix(postfix)
}

func evaluateExpression(expression string) (float64, bool) {
	expr := strings.ReplaceAll(expression, " ", "")
	result, err := eval(expr)
	if err != nil {
		fmt.Println("Error: Invalid expression.")
		return 0, false
	}
	return result, true
}

func main() {
    welcome()
	var mode string
	fmt.Println("Choose mode:")
	fmt.Println("1 - User-Friendly Version")
	fmt.Println("2 - Development Version")
	fmt.Print("Enter choice: ")
	fmt.Scanln(&mode)

	if mode == "1" {
		var num1, num2 float64
		var operator string

		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)

		fmt.Print("Enter operator (+, -, *, /): ")
		fmt.Scanln(&operator)

		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)

		if result, valid := calculate(num1, operator, num2); valid {
			fmt.Printf("Result: %s\n", formatOutput(result))
		} else {
			fmt.Println("Calculation failed.")
		}
	} else if mode == "2" {
		var readManual string
		fmt.Print("Do you want to read the manual first? (yes/no): ")
		fmt.Scanln(&readManual)
		if readManual == "yes" {
			manual()
		}

		var expression string
		fmt.Print("Enter the mathematical expression: ")
		fmt.Scanln(&expression)
		if result, valid := evaluateExpression(expression); valid {
			fmt.Printf("Result: %s\n", formatOutput(result))
		}
	} else {
		fmt.Println("Invalid choice. Exiting program.")
	}
}

