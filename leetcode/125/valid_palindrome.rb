def is_palindrome(s)
    s.gsub(/[^0-9a-zA-Z]/, '').casecmp(s.gsub(/[^0-9a-zA-Z]/, '').reverse) == 0
end
