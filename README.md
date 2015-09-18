#Supu CLI
Command line tool to manage your supu.io

###next
It will show you a list of in todo issues

###start [issue_id]
Sends an issue to in progress. This command will:

- [ ] Create a new feature branch.
- [ ] Move the issue to in_progress on your issue tracker.
- [ ] Assign to me the issue on your issue tracker.

### doing [ --user <user> ]
It will show a list of in todo issues

### details [issue_id]
Show issue details for the given issue

### comment [issue_id] <body>
Will add a comment on the issue.

### review [issue_id] [ --assign <user> ]
Sends an issue to be review. This command will:

- [ ] Move to REVIEW the issue on the issue tracker
- [ ] Push your local branch to remote repository
- [ ] Will open a pull request against default branch
- [ ] Assign the pull request to a reviewer

### uat [issue_id] [ --assign <user> ]
Sends an issue to be uated. This command will:

- [ ] Move to UAT the issue on the issue tracker
- [ ] Assign to somebody on the uat team

### done [issue_id] [ --assign <user> ]  [ --not-merge ]
Sends the given issue to done. This command will:

- [ ] Move to done column on the issue tracker
- [ ] Assign to original developer on the issue tracker
- [ ] Merge the pull request





