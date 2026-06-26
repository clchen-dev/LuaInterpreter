local vectorMeta = {}

local function vector(x, y)
  local value = {x = x, y = y}
  setmetatable(value, vectorMeta)
  return value
end

vectorMeta.__add = function(left, right)
  return vector(left.x + right.x, left.y + right.y)
end

local result = vector(1, 2) + vector(3, 4)
print(result.x, result.y)
