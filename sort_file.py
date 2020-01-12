#!/usr/bin/python2.7
from random import randint
# Random array to file
f = open("to_sort.txt", "w")
for _ in range(10): 
  f.write(str(randint(0,100))+ " ")
f.close  

# Read array
f = open("to_sort.txt", "r")
array = f.read().split()
f.close

# Convert to int
array = map(int, array)

def mysort(arr, low, high):
  print(array)
  if (low < high):
    p = partition(arr, low, high)
    mysort(arr, low, p - 1)
    mysort(arr, p + 1, high)

def swap(arr, i, j):
  temp = arr[i]
  arr[i] = arr[j]
  arr[j] = temp
  
def partition(arr, low, high):
  # Take middle element
  p = arr[high]
  i = low - 1
  # loop array 
  for j in range(low, high+1, 1):
    if ( arr[j] <= p ):  
      swap(arr, i + 1, j)
      i = i + 1
  #swap(arr, i , high)
  return i 

mysort(array, 0, len(array) - 1)
print array

