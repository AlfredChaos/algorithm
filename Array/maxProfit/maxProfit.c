int maxProfit(int *prices, int pricesSize)
{
    int max = 0;
    int minprice = 1e9;
    int profit = 0;
    for (int i = 0; i < pricesSize; i++)
    {
        profit = prices[i] - minprice;
        if (profit > max)
        {
            max = profit;
        }
        if (minprice > prices[i])
        {
            minprice = prices[i];
        }
    }
    return max;
}