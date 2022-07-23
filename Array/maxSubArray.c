//
// Created by alfred on 2022/7/19.
// https://leetcode.cn/problems/maximum-subarray/
//

#include "stdio.h"

int maxSubArray(int* nums, int numsSize){
    int max = 0;
    for (int i=1; i<numsSize; i++) {
        if (nums[i] + nums[i-1] > nums[i]) {
            nums[i] += nums[i-1];
        }
        if (nums[i] > max) {
            max = nums[i];
        }
    }
    return max;
}

int main() {
    int nums[9] = {-2, 1, -3, 4, -1, 2, 1, -5,4};
    printf("%d\n", maxSubArray(nums, 9));
    return 0;
}