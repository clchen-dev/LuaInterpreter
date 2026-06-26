local function counter(start)
  return function(step)
    start = start + step
    return start
  end
end

local nextValue = counter(10)
print(nextValue(2), nextValue(3))
