class Solution:
    def firstUniqChar(self, s: str) -> int:
        cnt: dict[str, int] = {}

        for idx, char in enumerate(s):
            cnt[char] = idx

        print(cnt)
        for idx, char in enumerate(s):
            if idx == cnt[char]:
                return idx
            else:
                cnt[char] = -1

        return -1