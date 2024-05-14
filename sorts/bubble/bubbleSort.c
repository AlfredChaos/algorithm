void swap(int arr[], int left, int right)
{
    int temp = arr[right];
    arr[right] = arr[left];
    arr[left] = temp;
}

void bubble(int arr[])
{
    int len = sizeof(arr) / sizeof(arr[0]);
    if (len > 1)
    {
        for (int i = 0; i < len; i++)
        {
            for (int j = len - 1; j > i; j--)
            {
                if (arr[i] > arr[j])
                {
                    swap(arr, i, j);
                }
            }
        }
    }
}