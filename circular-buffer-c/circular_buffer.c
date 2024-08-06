#include <stdio.h>

enum { N = 10 };  // size of circular buffer

int buffer[N]; // note: only (N - 1) elements can be stored at a given time
int writeIndx = 0;
int readIndx  = 0;

int put (int item) 
{
  if ((writeIndx + 1) % N == readIndx)
  {
     // buffer is full, avoid overflow
     return 0;
  }
  buffer[writeIndx] = item;
  writeIndx = (writeIndx + 1) % N;
  return 1;
}

int get (int *value) 
{
  if (readIndx == writeIndx)
  {
     // buffer is empty
     return 0;
  }

  *value = buffer[readIndx];
  readIndx = (readIndx + 1) % N;
  return 1;
}

int main ()
{
    // printf("%i\n", 10 % 10);
  // test circular buffer
  int value = 1001;
  while (put (value ++));
  while (get (& value))
     printf ("read %d\n", value);
  return 0;
}