from androguard.misc import AnalyzeAPK

a,d,dx= AnalyzeAPK('2.1.4.apk')
f = open('2.1.4.java', 'w')

for dd in d:
   for clas in dd.get_classes():
      name = clas.get_name()
      if 'bandcamp' in name:
         print(name, file=f)
         src = dd.get_class(name).get_source()
         print(src, file=f)

