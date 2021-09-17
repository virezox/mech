from androguard.misc import AnalyzeAPK

a,d,dx= AnalyzeAPK('Bandcamp_v2.4.11_apkpure.com.apk')
f = open('2.4.11.java', 'w')

for dd in d:
   for clas in dd.get_classes():
      name = clas.get_name()
      if 'bandcamp' in name:
         print(name, file=f)
         src = dd.get_class(name).get_source()
         print(src, file=f)

