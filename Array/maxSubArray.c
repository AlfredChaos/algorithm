//
// Created by alfred on 2022/7/19.
// https://leetcode.cn/problems/maximum-subarray/
//

#include "stdio.h"

// 这道题名为数组，实际上是动态规划
// 找出特征方程，就变得非常容易
// 特征方程就是nums[n] + f(n-1) > nums[n], f(n-1)代表前n-1项之和
// 可以通过数学方法论证：
// 1.假设存在这么一个最大子序列和，nums[0] + nums[1] < nums[1]
// 那么子序列完全应该以f(1)开始，与预设条件不符

// 2.假设nums内数值全都是负数，一定存在一个最大的负数，单独成列

// 3.假设f(n-1)<0，nums[n]>0，
// 如果sum>0，则sum<nums[n]，nums[n]应该才是最大子序列；
// 如果sum<0，则sum<nums[n]，nums[n]应该才是最大子序列；

// 4.假设f(n-1)>0，nums[n]<0，
// 如果sum>0，则sum>nums[n]
// 如果sum<0，则sum>nums[n]

// 5.假设两者都大于0，则无争议

// 综上，满足sum>nums[n]即可
int maxSubArray(int *nums, int numsSize)
{
    int max = 0;
    for (int i = 1; i < numsSize; i++)
    {
        if (nums[i] + nums[i - 1] > nums[i])
        {
            nums[i] += nums[i - 1];
        }
        if (nums[i] > max)
        {
            max = nums[i];
        }
    }
    return max;
}

int main()
{
    int nums[9] = {-2, 1, -3, 4, -1, 2, 1, -5, 4};
    printf("%d\n", maxSubArray(nums, 9));
    return 0;
}