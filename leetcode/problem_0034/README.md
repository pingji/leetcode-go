# Solutions

## Solution 1：
二分查找中，寻找 leftIdx 即为在数组中寻找第一个大于等于 target 的下标，寻找rightIdx 即为在数组中寻找第一个大于 target 的下标，然后将下标减一。

## Solution 2:
特别处理 num[mid] == target 的时left rigth变换，让它等于mid

注意几个地方
1. 循环的条件必须时 left < right ，如果是 = ，当num[mid] == target时，由于还是mid，导致一直死循环
2. 取last时，mid的计算方法是让它靠近right。原因，如果left和right相连（比如4，5），4的位置是target，5不是。由于要往右搜索，left一直停在了4，必须让5往右移动
3. 这种保留mid，不做mid+1 或者mid-1的操作，就会导致mid位置不会移动，处理时要考虑很多细节
 # References
