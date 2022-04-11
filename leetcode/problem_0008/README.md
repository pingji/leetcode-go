# Solutions

## Solution 1：确定有限状态机（deterministic finite automaton, DFA）

字符串分为以下几个类型

- space
- +/-
- number
- other

状态分为以下几个状态

- start
- signed
- in_number
- end

|           | space | + / -  | number    | other |
| --------: | ----- | ------ | --------- | ----- |
|     start | start | signed | in_number | end   |
|    signed | end   | end    | in_number | end   |
| in_number | end   | end    | in_number | end   |
|       end | end   | end    | end       | end   |

# References
