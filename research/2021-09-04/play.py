import gpsoauth
import sys

res = gpsoauth.perform_master_login(
   'srpen6@gmail.com', sys.argv[1], '38B5418D8683ADBB',
)

print(res)
