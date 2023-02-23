import torch

pthfile = r'D:/Coding/Python/Python_Program/Pth_Test/spynet_20210409-c6c1bd09.pth'

net = torch.load(pthfile,map_location=torch.device('cpu'))

print(type(net))

print(len(net))

for k in net.keys():
    print(k)  # 查看键

