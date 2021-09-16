from androguard.misc import AnalyzeAPK
a,d,dx= AnalyzeAPK('Bandcamp_v2.4.11_apkpure.com.apk')
src = d[1].get_class('Lcom/bandcamp/android/network/d;').get_source()
f = open('two.java', 'w')
print(src, file=f)
