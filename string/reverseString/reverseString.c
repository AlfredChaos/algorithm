void swap(char *a, char *b)
{
    char temp = *a;
    *a = *b;
    *b = temp;
}

void reverseString(char *s, int sSize)
{
    for (int left = 0, right = sSize - 1; left < right; left++)
    {
        swap(s[left], s[right]);
        right--;
    }
}