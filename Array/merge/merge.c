void append(int *origin, int *add, int index, int addSize)
{
    int originSize = sizeof(origin) / sizeof(origin[0]);
    for (int i = index; i < addSize; i++)
    {
        origin[originSize] = add[index];
        originSize++;
    }
}

void merge(int *nums1, int nums1Size, int m, int *nums2, int nums2Size, int n)
{
    int p1 = 0, p2 = 0;
    int sorted[m + n];
    for (;;)
    {
        if (p1 == m)
        {
            append(sorted, nums2, p2, n);
            break;
        }
        if (p2 == n)
        {
            append(sorted, nums1, p1, m);
            break;
        }
        if (nums1[p1] < nums2[p2])
        {
            int sortedSize = sizeof(sorted) / sizeof(sorted[0]);
            sorted[sortedSize] = nums1[p1];
            p1++;
        }
        else
        {
            int sortedSize = sizeof(sorted) / sizeof(sorted[0]);
            sorted[sortedSize] = nums2[p2];
            p2++;
        }
    }
    for (int i = 0; i < m + n; i++)
    {
        nums1[i] = sorted[i];
    }
}