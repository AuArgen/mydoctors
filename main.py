s=input()
t=0
for i in s:
    if t==0:
        print("$",end="")
    t+=1    
    if t==60:
        print()
        t=0
    else:
        print(i,end="")