package fuckup

import (
  "strconv"
)

func Fuckup() {
  println("fuckup")
}


func add(exists map[int]int, num string) (string, error) {
    int_num, err := strconv.Atoi(num)
    if err != nil {
        return "false", err
    }
    exists[int_num] += 1
    return "", nil
}

func remove(exists map[int]int, num string) (string, error) {
    
    int_num, err := strconv.Atoi(num)
    if err != nil {
        return "false", err
    }
    
    if count, exist := exists[int_num]; exist == false || count <= 0 {
        return "false", nil
    }
    exists[int_num] -= 1
    return "true", nil
}

func exist(exists map[int]int, num string) (string, error)  {
    
    int_num, err := strconv.Atoi(num)
    if err != nil {
        return "false", err
    }
    
    count, is := exists[int_num]
    if is && count > 0 {
        return "true", nil;
    } else {
        return "false", nil;
    }
}

func get_next(exists map[int]int, num string) (string, error) {
    
    int_num, err := strconv.Atoi(num)
    
    if err != nil {
        return "false", err
    }
    
    for next := int_num + 1 ; next <= 100 ; next++ {
        if count, exist := exists[next]; exist && count > 0 {
            return strconv.Itoa(next), nil
        }
    }
    
    return "", nil
}


func Solution(queries [][]string) []string {
    exists := make(map[int]int)
    ans := make([]string, 0)
    for _, cmd  := range queries {
        op, num := cmd[0], cmd[1]
        switch op {
            case "ADD" :
                ret, _ := add(exists, num)
                ans = append(ans, ret)
            case "REMOVE":
                ret, _ := remove(exists, num)
                ans = append(ans, ret)
            case "EXISTS" :
                ret, _ := exist(exists, num)
                ans = append(ans, ret)
            case "GET_NEXT":
                ret, _ := get_next(exists, num)
                ans = append(ans, ret)
            default: 
                
        }
    }
    return ans
}

