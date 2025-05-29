package main

import (
	"bufio"
	"fmt"
	"os"
	"sections_algorithm/internal/function"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Укажите имя файла как аргумент командной строки")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не удалось открыть файл: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	readFloat := func() (float64, error) {
		if !scanner.Scan() {
			return 0, fmt.Errorf("ожидалось число, но данные закончились")
		}
		var f float64
		_, err := fmt.Sscanf(scanner.Text(), "%f", &f)
		return f, err
	}

	nf, err := readFloat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения количества отрезков: %v\n", err)
		return
	}
	n := int(nf)
	segments := make([]function.Section, n)

	for i := 0; i < n; i++ {
		l, err := readFloat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения левой границы отрезка %d: %v\n", i+1, err)
			return
		}
		r, err := readFloat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения правой границы отрезка %d: %v\n", i+1, err)
			return
		}
		segments[i] = function.Section{Start: l, End: r}
	}

	result := function.MinPointsCover(segments)
	fmt.Println(result)
}
