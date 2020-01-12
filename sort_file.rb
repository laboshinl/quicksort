#!/usr/bin/ruby

array = 10.times.map{ rand(100) }

File.open('to_sort.txt', 'w') {
  |file| file.write(array.join(" "))
}

File.open('to_sort.txt', 'r') {
  |file| array = file.read()
}

array = array.split().map(&:to_i)

def mysort(arr, low, high)
  if ( low < high ) 
    p = partition(arr, low, high)
    mysort(arr, low, p - 1)
    mysort(arr, p + 1, high) 
  end
end

def partition(arr, low, high)
  p = arr[high]
  i = low - 1
  for j in Range.new(low,high)
    if ( arr[j] <= p )
      if ( i < j )  
        swap(arr, i+1, j)
      end
      i = i + 1
    end
  end   
  return i 
end

def swap(arr, i, j)
  temp = arr[i]
  arr[i] = arr[j]
  arr[j] = temp
end

mysort(array, 0, array.length - 1 )
p array
