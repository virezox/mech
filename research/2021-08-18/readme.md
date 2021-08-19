# August 18 2021

I think I found a better option. Instead of using the Parser, I can just use
the Lexer. That way, the input stays as is. As long as what I said holds true,
about the values being valid JSON, it seems to do the trick. I probably need to
add more cases, but the code here works with basic inputs.
