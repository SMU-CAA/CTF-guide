### [File](./file/cpp_dd1c054d73e92af0e5ad9c7ee1371bd5.zip)

#### function: sub_40111A  

第一次加密

加密算法: `input[i] = ((input[i]<<2|input[i]>>6)^i)`

逆算法: `crypto[i] = (crypto[i] ^ i) << 6 | (crypto[i] ^ i) >> 2;`

#### fucntion: sub_401332

第二次加密

加密算法: `input[i]=(input[i]|input[i -1])&~(input[i]&input[i-1])` <!--实质为xor运算-->

逆算法: `crypto[i] ^= crypto[i - 1];`





解密程序:

```c
#include <stdio.h>

unsigned char crypto[] = {0x99,  0x0B0, 0x87, 0x9E, 0x70,  0x0E8, 0x41,  0x44,
                          0x5,   0x4,   0x8B, 0x9A, 0x74,  0x0BC, 0x55,  0x58,
                          0x0B5, 0x61,  0x8E, 0x36, 0x0AC, 0x9,   0x59,  0x0E5,
                          0x61,  0x0DD, 0x3E, 0x3F, 0x0B9, 0x15,  0x0ED, 0x0D5};

int main() {
  for (int j = 0; j < 4; j++) {

    // 原始加密: input[i]=(input[i]|input[i -1])&~(input[i]&input[i-1])
    // Simplification input[i] ^= input[i - 1];
    // 逆算法:
    for (int i = sizeof(crypto) - 1; i > 0; i--) {
      crypto[i] ^= crypto[i - 1];
    }
  }

  for (int i = 0; i < sizeof(crypto); i++) {

    crypto[i] = (crypto[i] ^ i) << 6 | (crypto[i] ^ i) >> 2;
    printf("%c", crypto[i]);
  }

  return 0;
}
```

