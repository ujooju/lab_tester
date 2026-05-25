## Repository tester for Gitea.
__Web interface__ displays repositories which are available for testing, score for each branch and the report from the __test runner__.
__Web interface__ has an inbuilt API to manage the test queue.
Every test request on the __web interface__ adds the branch to the queue.
__Test runner__ is polling the __web interface__ to get the next branch to test, then executes the test script, computes the branch score and sends the report back to the server.

### Dependencies
- curl
- OPENSSL
- git

You can make the system compatible with ГОСТ cyphering by reconfiguring OPENSSL, curl and git.

Http requests are made with the curl.
