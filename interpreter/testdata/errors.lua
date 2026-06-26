local ok, message = pcall(function()
  error("boom")
end)

print(ok, message)
