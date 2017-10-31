# This CI workflow executes five steps. 

  * Checkout code given a repo and commit
  * Build the checked out code
  * Run five tests in parallel 
  * Send notification for approval and wait for the specified timeout
  * Once approved, release artifacts
  
