int guess(int num);

class Solution {
public:
    int guessNumber(int n) {
        long long now = 0, step = n;
        while (step) {
            if (now + step > n) {
                step >>= 1;
                continue;
            }
            int re = guess(now + step);
            if (re >= 0) {
                now = now + step;
            } else {
                step >>= 1;
            }
        }
        return now;
    }
};