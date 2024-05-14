x = int(input("Please enter x : "))

for i in range(x):
    print("*" * (i+1))
for i in range(x-1):
    print("*" * (x-i-1))