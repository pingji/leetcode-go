## Solutions
### solution 1
回溯,

```mermaid
graph TB
    root[ ] --> Node11[1]
    root --> Node12[2]
    root --> Node13[3]
    Node11 --> Node21[2]
    Node11 --> Node22[3]
    Node11 --> Node23[4]
    Node12 --> Node24[3]
    Node12 --> Node25[4]
    Node13 --> Node26[4]
```