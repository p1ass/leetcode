import heapq

class Heap:
    def __init__(self):
        self.nums = []


    def push(self, val:int):
        # print(self.nums)
        self.nums.append(val)
        self.up(len(self.nums)-1)

    def pop(self)->int:
        val = self.nums[0]

        n = len(self.nums)
        self.nums[0], self.nums[n-1] =self.nums[n-1],self.nums[0]
        self.nums= self.nums[:n-1]
        self.down(0)
        return val

    def up(self,idx:int):
        if idx == 0:
            return

        parent = (idx - (2 - idx % 2)) // 2

        if self.nums[idx] < self.nums[parent]:
            self.nums[idx],self.nums[parent] = self.nums[parent],self.nums[idx]
            self.up(parent)

    def down(self,idx:int):
        left = 2 *idx + 1
        right = 2*idx + 2

        if left >= len(self.nums):
            return

        if right >= len(self.nums):
            if self.nums[idx] > self.nums[left]:
                self.nums[idx], self.nums[left] = self.nums[left], self.nums[idx]
            return

        smaller_child = right
        if self.nums[left] < self.nums[right]:
            smaller_child = left


        if self.nums[idx] > self.nums[smaller_child]:
            self.nums[idx], self.nums[smaller_child] = self.nums[smaller_child], self.nums[idx]
            self.down(smaller_child)





class KthLargest:
    def __init__(self, k: int, nums: List[int]):
        self.k = k
        self.heap =Heap()
        for i in nums:
            if len(self.heap.nums) < k:
                self.heap.push(i)
            else:
                self.heap.push(i)
                self.heap.pop()

    def add(self, val: int) -> int:
        if len(self.heap.nums) < self.k:
            self.heap.push(val)
        else:
            self.heap.push(val)
            self.heap.pop()
        return self.heap.nums[0]

# Your KthLargest object will be instantiated and called as such:
# obj = KthLargest(k, nums)
# param_1 = obj.add(val)
