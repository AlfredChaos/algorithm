//
// Created by alfred on 2022/7/19.
// https://leetcode.cn/problems/maximum-subarray/
//

#include "stdio.h"

int maxSubArray(int* nums, int numsSize){
    int max = 0;
    int index = 0;
    while(index < numsSize) {
        int sum = nums[index];
        for (int i=index+1; i<numsSize; i++) {
            sum += nums[i];
        }
        if (sum > max) {
            max = sum;
        }
        index++;
    }
    return max;
}

int main() {
    int nums[9] = {-2, 1, -3, 4, -1, 2, 1, -5,4};
    printf("%d\n", maxSubArray(nums, 9));
    return 0;
}