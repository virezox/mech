from gpsoauth import google
import gpsoauth
import sys

sig = google.construct_signature(
   'srpen6@gmail.com', sys.argv[1], gpsoauth.ANDROID_KEY_7_3_29
)

print(sig)
