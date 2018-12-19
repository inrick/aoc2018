#include <stdio.h>
#include <stdint.h>

/* First automatic dump of input.txt. See main.go for solution. */

int main(void) {
  int64_t a = 0, b = 0, c = 0, d = 0, e = 0, f = 0;
  a = 1;
  for (;;) {
    switch (c) {
    case  0: c += 16;               break;
    case  1: d  = 1;                break;
    case  2: b  = 1;                break;
    case  3: f  = d * b;            break;
    case  4: f  = (f == e) ? 1 : 0; break;
    case  5: c += f;                break;
    case  6: c += 1;                break;
    case  7: a += d;                break;
    case  8: b += 1;                break;
    case  9: f  = (b > e) ? 1 : 0;  break;
    case 10: c += f;                break;
    case 11: c  = 2;                break;
    case 12: d += 1;                break;
    case 13: f  = (d > e) ? 1 : 0;  break;
    case 14: c += f;                break;
    case 15: c  = 1;                break;
    case 16: c *= c;                break;
    case 17: e += 2;                break;
    case 18: e *= e;                break;
    case 19: e *= c;                break;
    case 20: e *= 11;               break;
    case 21: f += 7;                break;
    case 22: f *= c;                break;
    case 23: f += 4;                break;
    case 24: e += f;                break;
    case 25: c += a;                break;
    case 26: c  = 0;                break;
    case 27: f  = c;                break;
    case 28: f *= c;                break;
    case 29: f += c;                break;
    case 30: f *= c;                break;
    case 31: f *= 14;               break;
    case 32: f *= c;                break;
    case 33: e += f;                break;
    case 34: a  = 0;                break;
    case 35: c  = 0;                break;
    default: goto done;
    }
    c++;
  }
done:
  printf("%ld\n", a);
  return 0;
}
