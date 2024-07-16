package ascii

func Converter(input string, lines []string) [][]string {
    converted := [][]string{}
    symb := []string{}
    for _, char := range input {
        for i := 1; i < 9; i++ {
            symb = append(symb, lines[int((char-32)*9)+i])
        }
        converted = append(converted, symb)
        symb = []string{}
    }
    return converted
}
