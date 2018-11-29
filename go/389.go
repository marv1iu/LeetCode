/*
Use a map count the frenquency of each char.
Better solution, use xor to iterate both strings.
*/
func findTheDifference(s string, t string) byte {
    if len(s) < len(t) {
        temp := t
        t = s
        s = temp
    }
    counter1, counter2 := counter(s), counter(t)
    for key, value := range counter1 {
        if _value, ok := counter2[key]; !(ok && _value == value)  {
            return byte(key)
        }
    }
    return s[0]
}

func counter(str string) map[rune]int {
    _counter := make(map[rune]int)
    for _, val := range str {
        _counter[val], _ = _counter[val]
        _counter[val]++
    }
    return _counter
}
