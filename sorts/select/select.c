void swap(int arr[], int left, int right)
{
    int temp = arr[right];
    arr[right] = arr[left];
    arr[left] = temp;
}

void selection(int arr[], int size)
{
    int min = 0;
    for (int i = 0; i < size - 1; i++)
    {
        for (int j = i + 1; j < size - 1; j++)
        {
            if (arr[i] > arr[j])
            {
                min = j;
            }
        }
        swap(arr, i, min);
    }
}