class Solution:
    def numUniqueEmails(self, emails: List[str]) -> int:
        actual = {}

        for email in emails:
            split_email = []

            seenAtmark = False
            seenPlus = False
            for s in email:
                if s == "@":
                    seenAtmark= True

                if not seenAtmark:
                    if s == "+":
                        seenPlus = True

                    if s == ".":
                        continue

                    if not seenPlus:
                        split_email.append(s)

                else:
                    split_email.append(s)

            actual_email = "".join(split_email)

            if actual_email not in actual:
                actual[actual_email] = True


        return len(actual)