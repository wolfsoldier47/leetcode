package roman_to_integer

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

//     I can be placed before V (5) and X (10) to make 4 and 9.
//     X can be placed before L (50) and C (100) to make 40 and 90.
//     C can be placed before D (500) and M (1000) to make 400 and 900.

// Given a roman numeral, convert it to an integer.


func romanToInt(s string) int {
    r := []rune(s)

    var count int = 0

    for i:=0 ;i<len(r);i++{
        switch string(r[i]) {
            case "I":
                    if i+1 < len(r) && (string(r[i+1]) == "V" || string(r[i+1]) == "X") {
                        if string(r[i+1]) == "V" {
                            count += 4
                            i++
                        } else if string(r[i+1]) == "X" {
                            count += 9
                            i++ // skip next character
                        }
                    } else {
                        count += 1
                    }
            case "V": count += 5
            case "X":
                if  i+1 < len(r) && (string(r[i+1]) == "L" || string(r[i+1]) == "C") {
                        if string(r[i+1]) == "L" {
                            count += 40
                            i++
                        } else if string(r[i+1]) == "C" {
                            count += 90
                            i++
                        }
                } else {
                    count += 10
                }


            case "L": count += 50
            case "C":
            if  i+1 < len(r) && (string(r[i+1]) == "D" || string(r[i+1]) == "M") {
                if string(r[i+1]) == "D" {
                            count += 400
                            i++
                } else if string(r[i+1]) == "M"{
                            count += 900
                            i++
                }
            } else {
                count += 100
            }

            case "D": count += 500
            case "M": count += 1000
        }
    }

    return count
}
